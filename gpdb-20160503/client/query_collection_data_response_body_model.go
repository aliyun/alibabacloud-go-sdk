// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCollectionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMatches(v *QueryCollectionDataResponseBodyMatches) *QueryCollectionDataResponseBody
	GetMatches() *QueryCollectionDataResponseBodyMatches
	SetMessage(v string) *QueryCollectionDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCollectionDataResponseBody
	GetRequestId() *string
	SetStatus(v string) *QueryCollectionDataResponseBody
	GetStatus() *string
	SetTotal(v int32) *QueryCollectionDataResponseBody
	GetTotal() *int32
}

type QueryCollectionDataResponseBody struct {
	// Data list.
	Matches *QueryCollectionDataResponseBodyMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Struct"`
	// Detailed information when the request fails.
	//
	// example:
	//
	// 0.1234
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status, with the following values:
	//
	// - **success**: Success.
	//
	// - **fail**: Failure.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Only returned when the Offset is not 0, this value represents the total number of hits for the search criteria.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryCollectionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBody) GetMatches() *QueryCollectionDataResponseBodyMatches {
	return s.Matches
}

func (s *QueryCollectionDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCollectionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCollectionDataResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryCollectionDataResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *QueryCollectionDataResponseBody) SetMatches(v *QueryCollectionDataResponseBodyMatches) *QueryCollectionDataResponseBody {
	s.Matches = v
	return s
}

func (s *QueryCollectionDataResponseBody) SetMessage(v string) *QueryCollectionDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCollectionDataResponseBody) SetRequestId(v string) *QueryCollectionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCollectionDataResponseBody) SetStatus(v string) *QueryCollectionDataResponseBody {
	s.Status = &v
	return s
}

func (s *QueryCollectionDataResponseBody) SetTotal(v int32) *QueryCollectionDataResponseBody {
	s.Total = &v
	return s
}

func (s *QueryCollectionDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataResponseBodyMatches struct {
	Match []*QueryCollectionDataResponseBodyMatchesMatch `json:"match,omitempty" xml:"match,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataResponseBodyMatches) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatches) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatches) GetMatch() []*QueryCollectionDataResponseBodyMatchesMatch {
	return s.Match
}

func (s *QueryCollectionDataResponseBodyMatches) SetMatch(v []*QueryCollectionDataResponseBodyMatchesMatch) *QueryCollectionDataResponseBodyMatches {
	s.Match = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatches) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataResponseBodyMatchesMatch struct {
	// The unique ID of the vector data.
	//
	// example:
	//
	// doca-1234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Metadata.
	Metadata   map[string]*string     `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	MetadataV2 map[string]interface{} `json:"MetadataV2,omitempty" xml:"MetadataV2,omitempty"`
	// The similarity score of this data, which is related to the algorithm `(l2/ip/cosine)` specified when creating the index.
	//
	// example:
	//
	// 0.12345
	Score        *float64                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	SparseValues *QueryCollectionDataResponseBodyMatchesMatchSparseValues `json:"SparseValues,omitempty" xml:"SparseValues,omitempty" type:"Struct"`
	// List of vector data.
	Values *QueryCollectionDataResponseBodyMatchesMatchValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s QueryCollectionDataResponseBodyMatchesMatch) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatchesMatch) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) GetId() *string {
	return s.Id
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) GetMetadataV2() map[string]interface{} {
	return s.MetadataV2
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) GetScore() *float64 {
	return s.Score
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) GetSparseValues() *QueryCollectionDataResponseBodyMatchesMatchSparseValues {
	return s.SparseValues
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) GetValues() *QueryCollectionDataResponseBodyMatchesMatchValues {
	return s.Values
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetId(v string) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Id = &v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetMetadata(v map[string]*string) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Metadata = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetMetadataV2(v map[string]interface{}) *QueryCollectionDataResponseBodyMatchesMatch {
	s.MetadataV2 = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetScore(v float64) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Score = &v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetSparseValues(v *QueryCollectionDataResponseBodyMatchesMatchSparseValues) *QueryCollectionDataResponseBodyMatchesMatch {
	s.SparseValues = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) SetValues(v *QueryCollectionDataResponseBodyMatchesMatchValues) *QueryCollectionDataResponseBodyMatchesMatch {
	s.Values = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatch) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataResponseBodyMatchesMatchSparseValues struct {
	Indices *QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Struct"`
	Values  *QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues  `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s QueryCollectionDataResponseBodyMatchesMatchSparseValues) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatchesMatchSparseValues) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValues) GetIndices() *QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices {
	return s.Indices
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValues) GetValues() *QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues {
	return s.Values
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValues) SetIndices(v *QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices) *QueryCollectionDataResponseBodyMatchesMatchSparseValues {
	s.Indices = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValues) SetValues(v *QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues) *QueryCollectionDataResponseBodyMatchesMatchSparseValues {
	s.Values = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValues) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices struct {
	Indice []*int32 `json:"Indice,omitempty" xml:"Indice,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices) GetIndice() []*int32 {
	return s.Indice
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices) SetIndice(v []*int32) *QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices {
	s.Indice = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValuesIndices) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues struct {
	Value []*float32 `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues) GetValue() []*float32 {
	return s.Value
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues) SetValue(v []*float32) *QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues {
	s.Value = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatchSparseValuesValues) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataResponseBodyMatchesMatchValues struct {
	Value []*float64 `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataResponseBodyMatchesMatchValues) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponseBodyMatchesMatchValues) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponseBodyMatchesMatchValues) GetValue() []*float64 {
	return s.Value
}

func (s *QueryCollectionDataResponseBodyMatchesMatchValues) SetValue(v []*float64) *QueryCollectionDataResponseBodyMatchesMatchValues {
	s.Value = v
	return s
}

func (s *QueryCollectionDataResponseBodyMatchesMatchValues) Validate() error {
	return dara.Validate(s)
}
