// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySampleConsistencyJobDifferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDifferenceHistogram(v []*QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) *QuerySampleConsistencyJobDifferenceResponseBody
	GetDifferenceHistogram() []*QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram
	SetNumberFeatureDifferences(v []*QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) *QuerySampleConsistencyJobDifferenceResponseBody
	GetNumberFeatureDifferences() []*QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences
	SetRequestId(v string) *QuerySampleConsistencyJobDifferenceResponseBody
	GetRequestId() *string
	SetStringFeatureDifferences(v []*QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) *QuerySampleConsistencyJobDifferenceResponseBody
	GetStringFeatureDifferences() []*QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences
}

type QuerySampleConsistencyJobDifferenceResponseBody struct {
	DifferenceHistogram      []*QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram      `json:"DifferenceHistogram,omitempty" xml:"DifferenceHistogram,omitempty" type:"Repeated"`
	NumberFeatureDifferences []*QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences `json:"NumberFeatureDifferences,omitempty" xml:"NumberFeatureDifferences,omitempty" type:"Repeated"`
	RequestId                *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StringFeatureDifferences []*QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences `json:"StringFeatureDifferences,omitempty" xml:"StringFeatureDifferences,omitempty" type:"Repeated"`
}

func (s QuerySampleConsistencyJobDifferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySampleConsistencyJobDifferenceResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) GetDifferenceHistogram() []*QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram {
	return s.DifferenceHistogram
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) GetNumberFeatureDifferences() []*QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences {
	return s.NumberFeatureDifferences
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) GetStringFeatureDifferences() []*QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences {
	return s.StringFeatureDifferences
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) SetDifferenceHistogram(v []*QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) *QuerySampleConsistencyJobDifferenceResponseBody {
	s.DifferenceHistogram = v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) SetNumberFeatureDifferences(v []*QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) *QuerySampleConsistencyJobDifferenceResponseBody {
	s.NumberFeatureDifferences = v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) SetRequestId(v string) *QuerySampleConsistencyJobDifferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) SetStringFeatureDifferences(v []*QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) *QuerySampleConsistencyJobDifferenceResponseBody {
	s.StringFeatureDifferences = v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram struct {
	Abscissa *string `json:"Abscissa,omitempty" xml:"Abscissa,omitempty"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) String() string {
	return dara.Prettify(s)
}

func (s QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) GoString() string {
	return s.String()
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) GetAbscissa() *string {
	return s.Abscissa
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) GetValue() *int64 {
	return s.Value
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) SetAbscissa(v string) *QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram {
	s.Abscissa = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) SetValue(v int64) *QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram {
	s.Value = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyDifferenceHistogram) Validate() error {
	return dara.Validate(s)
}

type QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences struct {
	DiffValue               *float64 `json:"DiffValue,omitempty" xml:"DiffValue,omitempty"`
	ItemId                  *string  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ReplyTableFeatureValue  *float64 `json:"ReplyTableFeatureValue,omitempty" xml:"ReplyTableFeatureValue,omitempty"`
	RequestId               *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SampleTableFeatureValue *float64 `json:"SampleTableFeatureValue,omitempty" xml:"SampleTableFeatureValue,omitempty"`
	UserId                  *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) String() string {
	return dara.Prettify(s)
}

func (s QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) GoString() string {
	return s.String()
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) GetDiffValue() *float64 {
	return s.DiffValue
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) GetItemId() *string {
	return s.ItemId
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) GetReplyTableFeatureValue() *float64 {
	return s.ReplyTableFeatureValue
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) GetSampleTableFeatureValue() *float64 {
	return s.SampleTableFeatureValue
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) GetUserId() *string {
	return s.UserId
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) SetDiffValue(v float64) *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences {
	s.DiffValue = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) SetItemId(v string) *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences {
	s.ItemId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) SetReplyTableFeatureValue(v float64) *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences {
	s.ReplyTableFeatureValue = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) SetRequestId(v string) *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences {
	s.RequestId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) SetSampleTableFeatureValue(v float64) *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences {
	s.SampleTableFeatureValue = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) SetUserId(v string) *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences {
	s.UserId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyNumberFeatureDifferences) Validate() error {
	return dara.Validate(s)
}

type QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences struct {
	ItemId                  *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ReplyTableFeatureValue  *string `json:"ReplyTableFeatureValue,omitempty" xml:"ReplyTableFeatureValue,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SampleTableFeatureValue *string `json:"SampleTableFeatureValue,omitempty" xml:"SampleTableFeatureValue,omitempty"`
	UserId                  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) String() string {
	return dara.Prettify(s)
}

func (s QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) GoString() string {
	return s.String()
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) GetItemId() *string {
	return s.ItemId
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) GetReplyTableFeatureValue() *string {
	return s.ReplyTableFeatureValue
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) GetSampleTableFeatureValue() *string {
	return s.SampleTableFeatureValue
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) GetUserId() *string {
	return s.UserId
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) SetItemId(v string) *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences {
	s.ItemId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) SetReplyTableFeatureValue(v string) *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences {
	s.ReplyTableFeatureValue = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) SetRequestId(v string) *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences {
	s.RequestId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) SetSampleTableFeatureValue(v string) *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences {
	s.SampleTableFeatureValue = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) SetUserId(v string) *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences {
	s.UserId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponseBodyStringFeatureDifferences) Validate() error {
	return dara.Validate(s)
}
