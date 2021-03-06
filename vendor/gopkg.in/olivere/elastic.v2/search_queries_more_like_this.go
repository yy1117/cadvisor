// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"fmt"
	"math"
)

// More like this query find documents that are “like” provided text
// by running it against one or more fields. For more details, see
// http://www.elasticsearch.org/guide/reference/query-dsl/mlt-query/
type MoreLikeThisQuery struct {
	fields                 []string
	likeText               string
	ids                    []string
	docs                   []*MoreLikeThisQueryItem
	include                *bool
	minimumShouldMatch     string
	minTermFreq            *int
	maxQueryTerms          *int
	stopWords              []string
	minDocFreq             *int
	maxDocFreq             *int
	minWordLen             *int
	maxWordLen             *int
	boostTerms             *float64
	boost                  *float64
	analyzer               string
	failOnUnsupportedField *bool
	queryName              string
}

// NewMoreLikeThisQuery creates a new more-like-this query.
func NewMoreLikeThisQuery(likeText string) MoreLikeThisQuery {
	return MoreLikeThisQuery{
		likeText:  likeText,
		fields:    make([]string, 0),
		ids:       make([]string, 0),
		docs:      make([]*MoreLikeThisQueryItem, 0),
		stopWords: make([]string, 0),
	}
}

// Field adds one or more field names to the query.
func (q MoreLikeThisQuery) Field(fields ...string) MoreLikeThisQuery {
	q.fields = append(q.fields, fields...)
	return q
}

// Fields adds one or more field names to the query.
// Deprecated: Use Field for compatibility with elastic.v3.
func (q MoreLikeThisQuery) Fields(fields ...string) MoreLikeThisQuery {
	q.fields = append(q.fields, fields...)
	return q
}

// StopWord sets the stopwords. Any word in this set is considered
// "uninteresting" and ignored. Even if your Analyzer allows stopwords,
// you might want to tell the MoreLikeThis code to ignore them, as for
// the purposes of document similarity it seems reasonable to assume that
// "a stop word is never interesting".
func (q MoreLikeThisQuery) StopWord(stopWords ...string) MoreLikeThisQuery {
	q.stopWords = append(q.stopWords, stopWords...)
	return q
}

// StopWords is an alias for StopWord.
// Deprecated: Use StopWord for compatibility with elastic.v3.
func (q MoreLikeThisQuery) StopWords(stopWords ...string) MoreLikeThisQuery {
	q.stopWords = append(q.stopWords, stopWords...)
	return q
}

// LikeText sets the text to use in order to find documents that are "like" this.
func (q MoreLikeThisQuery) LikeText(likeText string) MoreLikeThisQuery {
	q.likeText = likeText
	return q
}

// Docs sets the documents to use in order to find documents that are "like" this.
func (q MoreLikeThisQuery) Docs(docs ...*MoreLikeThisQueryItem) MoreLikeThisQuery {
	q.docs = append(q.docs, docs...)
	return q
}

// Ids sets the document ids to use in order to find documents that are "like" this.
func (q MoreLikeThisQuery) Ids(ids ...string) MoreLikeThisQuery {
	q.ids = append(q.ids, ids...)
	return q
}

// Include specifies whether the input documents should also be included
// in the results returned. Defaults to false.
func (q MoreLikeThisQuery) Include(include bool) MoreLikeThisQuery {
	q.include = &include
	return q
}

// PercentTermsToMatch will be changed to MinimumShouldMatch.
func (q MoreLikeThisQuery) PercentTermsToMatch(percentTermsToMatch float64) MoreLikeThisQuery {
	q.minimumShouldMatch = fmt.Sprintf("%d%%", int(math.Floor(percentTermsToMatch*100)))
	return q
}

// MinimumShouldMatch sets the number of terms that must match the generated
// query expressed in the common syntax for minimum should match.
// The default value is "30%".
//
// This used to be "PercentTermsToMatch".
func (q MoreLikeThisQuery) MinimumShouldMatch(minimumShouldMatch string) MoreLikeThisQuery {
	q.minimumShouldMatch = minimumShouldMatch
	return q
}

// MinTermFreq is the frequency below which terms will be ignored in the
// source doc. The default frequency is 2.
func (q MoreLikeThisQuery) MinTermFreq(minTermFreq int) MoreLikeThisQuery {
	q.minTermFreq = &minTermFreq
	return q
}

// MaxQueryTerms sets the maximum number of query terms that will be included
// in any generated query. It defaults to 25.
func (q MoreLikeThisQuery) MaxQueryTerms(maxQueryTerms int) MoreLikeThisQuery {
	q.maxQueryTerms = &maxQueryTerms
	return q
}

// MinDocFreq sets the frequency at which words will be ignored which do
// not occur in at least this many docs. The default is 5.
func (q MoreLikeThisQuery) MinDocFreq(minDocFreq int) MoreLikeThisQuery {
	q.minDocFreq = &minDocFreq
	return q
}

// MaxDocFreq sets the maximum frequency for which words may still appear.
// Words that appear in more than this many docs will be ignored.
// It defaults to unbounded.
func (q MoreLikeThisQuery) MaxDocFreq(maxDocFreq int) MoreLikeThisQuery {
	q.maxDocFreq = &maxDocFreq
	return q
}

// MinWordLength sets the minimum word length below which words will be
// ignored. It defaults to 0.
func (q MoreLikeThisQuery) MinWordLen(minWordLen int) MoreLikeThisQuery {
	q.minWordLen = &minWordLen
	return q
}

// MaxWordLen sets the maximum word length above which words will be ignored.
// Defaults to unbounded (0).
func (q MoreLikeThisQuery) MaxWordLen(maxWordLen int) MoreLikeThisQuery {
	q.maxWordLen = &maxWordLen
	return q
}

// BoostTerms sets the boost factor to use when boosting terms.
// It defaults to 1.
func (q MoreLikeThisQuery) BoostTerms(boostTerms float64) MoreLikeThisQuery {
	q.boostTerms = &boostTerms
	return q
}

// Analyzer specifies the analyzer that will be use to analyze the text.
// Defaults to the analyzer associated with the field.
func (q MoreLikeThisQuery) Analyzer(analyzer string) MoreLikeThisQuery {
	q.analyzer = analyzer
	return q
}

// Boost sets the boost for this query.
func (q MoreLikeThisQuery) Boost(boost float64) MoreLikeThisQuery {
	q.boost = &boost
	return q
}

// FailOnUnsupportedField indicates whether to fail or return no result
// when this query is run against a field which is not supported such as
// a binary/numeric field.
func (q MoreLikeThisQuery) FailOnUnsupportedField(fail bool) MoreLikeThisQuery {
	q.failOnUnsupportedField = &fail
	return q
}

// QueryName sets the query name for the filter that can be used when
// searching for matched_filters per hit.
func (q MoreLikeThisQuery) QueryName(queryName string) MoreLikeThisQuery {
	q.queryName = queryName
	return q
}

// Creates the query source for the mlt query.
func (q MoreLikeThisQuery) Source() interface{} {
	// {
	//   "match_all" : { ... }
	// }
	params := make(map[string]interface{})
	source := make(map[string]interface{})
	source["mlt"] = params

	if q.likeText == "" && len(q.docs) == 0 && len(q.ids) == 0 {
		// We have no form of returning errors for invalid queries as of Elastic v2.
		// We also don't have access to the client here, so we can't log anything.
		// All we can do is to return an empty query, I suppose.
		// TODO Is there a better approach here?
		//return nil, errors.New(`more_like_this requires some documents to be "liked"`)
		return source
	}

	if len(q.fields) > 0 {
		params["fields"] = q.fields
	}
	if q.likeText != "" {
		params["like_text"] = q.likeText
	}
	if q.minimumShouldMatch != "" {
		params["minimum_should_match"] = q.minimumShouldMatch
	}
	if q.minTermFreq != nil {
		params["min_term_freq"] = *q.minTermFreq
	}
	if q.maxQueryTerms != nil {
		params["max_query_terms"] = *q.maxQueryTerms
	}
	if len(q.stopWords) > 0 {
		params["stop_words"] = q.stopWords
	}
	if q.minDocFreq != nil {
		params["min_doc_freq"] = *q.minDocFreq
	}
	if q.maxDocFreq != nil {
		params["max_doc_freq"] = *q.maxDocFreq
	}
	if q.minWordLen != nil {
		params["min_word_len"] = *q.minWordLen
	}
	if q.maxWordLen != nil {
		params["max_word_len"] = *q.maxWordLen
	}
	if q.boostTerms != nil {
		params["boost_terms"] = *q.boostTerms
	}
	if q.boost != nil {
		params["boost"] = *q.boost
	}
	if q.analyzer != "" {
		params["analyzer"] = q.analyzer
	}
	if q.failOnUnsupportedField != nil {
		params["fail_on_unsupported_field"] = *q.failOnUnsupportedField
	}
	if q.queryName != "" {
		params["_name"] = q.queryName
	}
	if len(q.ids) > 0 {
		params["ids"] = q.ids
	}
	if len(q.docs) > 0 {
		docs := make([]interface{}, 0)
		for _, doc := range q.docs {
			docs = append(docs, doc.Source())
		}
		params["docs"] = docs
	}
	if q.include != nil {
		params["exclude"] = !(*q.include) // ES 1.x only has exclude
	}

	return source
}

// -- MoreLikeThisQueryItem --

// MoreLikeThisQueryItem represents a single item of a MoreLikeThisQuery
// to be "liked" or "unliked".
type MoreLikeThisQueryItem struct {
	likeText string

	index       string
	typ         string
	id          string
	doc         interface{}
	fields      []string
	routing     string
	fsc         *FetchSourceContext
	version     int64
	versionType string
}

// NewMoreLikeThisQueryItem creates and initializes a MoreLikeThisQueryItem.
func NewMoreLikeThisQueryItem() *MoreLikeThisQueryItem {
	return &MoreLikeThisQueryItem{
		version: -1,
	}
}

// LikeText represents a text to be "liked".
func (item *MoreLikeThisQueryItem) LikeText(likeText string) *MoreLikeThisQueryItem {
	item.likeText = likeText
	return item
}

// Index represents the index of the item.
func (item *MoreLikeThisQueryItem) Index(index string) *MoreLikeThisQueryItem {
	item.index = index
	return item
}

// Type represents the document type of the item.
func (item *MoreLikeThisQueryItem) Type(typ string) *MoreLikeThisQueryItem {
	item.typ = typ
	return item
}

// Id represents the document id of the item.
func (item *MoreLikeThisQueryItem) Id(id string) *MoreLikeThisQueryItem {
	item.id = id
	return item
}

// Doc represents a raw document template for the item.
func (item *MoreLikeThisQueryItem) Doc(doc interface{}) *MoreLikeThisQueryItem {
	item.doc = doc
	return item
}

// Fields represents the list of fields of the item.
func (item *MoreLikeThisQueryItem) Fields(fields ...string) *MoreLikeThisQueryItem {
	item.fields = append(item.fields, fields...)
	return item
}

// Routing sets the routing associated with the item.
func (item *MoreLikeThisQueryItem) Routing(routing string) *MoreLikeThisQueryItem {
	item.routing = routing
	return item
}

// FetchSourceContext represents the fetch source of the item which controls
// if and how _source should be returned.
func (item *MoreLikeThisQueryItem) FetchSourceContext(fsc *FetchSourceContext) *MoreLikeThisQueryItem {
	item.fsc = fsc
	return item
}

// Version specifies the version of the item.
func (item *MoreLikeThisQueryItem) Version(version int64) *MoreLikeThisQueryItem {
	item.version = version
	return item
}

// VersionType represents the version type of the item.
func (item *MoreLikeThisQueryItem) VersionType(versionType string) *MoreLikeThisQueryItem {
	item.versionType = versionType
	return item
}

// Source returns the JSON-serializable fragment of the entity.
func (item *MoreLikeThisQueryItem) Source() interface{} {
	if item.likeText != "" {
		return item.likeText
	}

	source := make(map[string]interface{})

	if item.index != "" {
		source["_index"] = item.index
	}
	if item.typ != "" {
		source["_type"] = item.typ
	}
	if item.id != "" {
		source["_id"] = item.id
	}
	if item.doc != nil {
		source["doc"] = item.doc
	}
	if len(item.fields) > 0 {
		source["fields"] = item.fields
	}
	if item.routing != "" {
		source["_routing"] = item.routing
	}
	if item.fsc != nil {
		source["_source"] = item.fsc.Source()
	}
	if item.version >= 0 {
		source["_version"] = item.version
	}
	if item.versionType != "" {
		source["_version_type"] = item.versionType
	}

	return source
}
