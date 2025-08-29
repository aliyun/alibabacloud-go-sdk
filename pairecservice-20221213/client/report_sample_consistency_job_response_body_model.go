// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportSampleConsistencyJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeaturesDifference(v []*ReportSampleConsistencyJobResponseBodyFeaturesDifference) *ReportSampleConsistencyJobResponseBody
	GetFeaturesDifference() []*ReportSampleConsistencyJobResponseBodyFeaturesDifference
	SetReplyTableFeatures(v int64) *ReportSampleConsistencyJobResponseBody
	GetReplyTableFeatures() *int64
	SetReplyTableLostFeatures(v int64) *ReportSampleConsistencyJobResponseBody
	GetReplyTableLostFeatures() *int64
	SetRequestId(v int64) *ReportSampleConsistencyJobResponseBody
	GetRequestId() *int64
	SetSampleTableFeatures(v int64) *ReportSampleConsistencyJobResponseBody
	GetSampleTableFeatures() *int64
	SetSampleTableLostFeatures(v int64) *ReportSampleConsistencyJobResponseBody
	GetSampleTableLostFeatures() *int64
}

type ReportSampleConsistencyJobResponseBody struct {
	FeaturesDifference      []*ReportSampleConsistencyJobResponseBodyFeaturesDifference `json:"FeaturesDifference,omitempty" xml:"FeaturesDifference,omitempty" type:"Repeated"`
	ReplyTableFeatures      *int64                                                      `json:"ReplyTableFeatures,omitempty" xml:"ReplyTableFeatures,omitempty"`
	ReplyTableLostFeatures  *int64                                                      `json:"ReplyTableLostFeatures,omitempty" xml:"ReplyTableLostFeatures,omitempty"`
	RequestId               *int64                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SampleTableFeatures     *int64                                                      `json:"SampleTableFeatures,omitempty" xml:"SampleTableFeatures,omitempty"`
	SampleTableLostFeatures *int64                                                      `json:"SampleTableLostFeatures,omitempty" xml:"SampleTableLostFeatures,omitempty"`
}

func (s ReportSampleConsistencyJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportSampleConsistencyJobResponseBody) GoString() string {
	return s.String()
}

func (s *ReportSampleConsistencyJobResponseBody) GetFeaturesDifference() []*ReportSampleConsistencyJobResponseBodyFeaturesDifference {
	return s.FeaturesDifference
}

func (s *ReportSampleConsistencyJobResponseBody) GetReplyTableFeatures() *int64 {
	return s.ReplyTableFeatures
}

func (s *ReportSampleConsistencyJobResponseBody) GetReplyTableLostFeatures() *int64 {
	return s.ReplyTableLostFeatures
}

func (s *ReportSampleConsistencyJobResponseBody) GetRequestId() *int64 {
	return s.RequestId
}

func (s *ReportSampleConsistencyJobResponseBody) GetSampleTableFeatures() *int64 {
	return s.SampleTableFeatures
}

func (s *ReportSampleConsistencyJobResponseBody) GetSampleTableLostFeatures() *int64 {
	return s.SampleTableLostFeatures
}

func (s *ReportSampleConsistencyJobResponseBody) SetFeaturesDifference(v []*ReportSampleConsistencyJobResponseBodyFeaturesDifference) *ReportSampleConsistencyJobResponseBody {
	s.FeaturesDifference = v
	return s
}

func (s *ReportSampleConsistencyJobResponseBody) SetReplyTableFeatures(v int64) *ReportSampleConsistencyJobResponseBody {
	s.ReplyTableFeatures = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBody) SetReplyTableLostFeatures(v int64) *ReportSampleConsistencyJobResponseBody {
	s.ReplyTableLostFeatures = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBody) SetRequestId(v int64) *ReportSampleConsistencyJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBody) SetSampleTableFeatures(v int64) *ReportSampleConsistencyJobResponseBody {
	s.SampleTableFeatures = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBody) SetSampleTableLostFeatures(v int64) *ReportSampleConsistencyJobResponseBody {
	s.SampleTableLostFeatures = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ReportSampleConsistencyJobResponseBodyFeaturesDifference struct {
	Count       *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	Ratio       *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s ReportSampleConsistencyJobResponseBodyFeaturesDifference) String() string {
	return dara.Prettify(s)
}

func (s ReportSampleConsistencyJobResponseBodyFeaturesDifference) GoString() string {
	return s.String()
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) GetCount() *int64 {
	return s.Count
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) GetFeatureType() *string {
	return s.FeatureType
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) GetRatio() *string {
	return s.Ratio
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) SetCount(v int64) *ReportSampleConsistencyJobResponseBodyFeaturesDifference {
	s.Count = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) SetFeatureName(v string) *ReportSampleConsistencyJobResponseBodyFeaturesDifference {
	s.FeatureName = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) SetFeatureType(v string) *ReportSampleConsistencyJobResponseBodyFeaturesDifference {
	s.FeatureType = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) SetRatio(v string) *ReportSampleConsistencyJobResponseBodyFeaturesDifference {
	s.Ratio = &v
	return s
}

func (s *ReportSampleConsistencyJobResponseBodyFeaturesDifference) Validate() error {
	return dara.Validate(s)
}
