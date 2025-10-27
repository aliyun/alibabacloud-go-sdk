// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckItemWarningStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetClusterCheckItemWarningStatisticsResponseBodyData) *GetClusterCheckItemWarningStatisticsResponseBody
	GetData() *GetClusterCheckItemWarningStatisticsResponseBodyData
	SetRequestId(v string) *GetClusterCheckItemWarningStatisticsResponseBody
	GetRequestId() *string
}

type GetClusterCheckItemWarningStatisticsResponseBody struct {
	// The statistics on risk items by risk level.
	Data *GetClusterCheckItemWarningStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C2B285A3-3493-5C5F-A224-4CCE4BFC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterCheckItemWarningStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckItemWarningStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemWarningStatisticsResponseBody) GetData() *GetClusterCheckItemWarningStatisticsResponseBodyData {
	return s.Data
}

func (s *GetClusterCheckItemWarningStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterCheckItemWarningStatisticsResponseBody) SetData(v *GetClusterCheckItemWarningStatisticsResponseBodyData) *GetClusterCheckItemWarningStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponseBody) SetRequestId(v string) *GetClusterCheckItemWarningStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterCheckItemWarningStatisticsResponseBodyData struct {
	// The number of high-risk items.
	//
	// example:
	//
	// 3
	HighWarningCount *int32 `json:"HighWarningCount,omitempty" xml:"HighWarningCount,omitempty"`
	// The number of low-risk items.
	//
	// example:
	//
	// 1
	LowWarningCount *int32 `json:"LowWarningCount,omitempty" xml:"LowWarningCount,omitempty"`
	// The number of medium-risk items.
	//
	// example:
	//
	// 2
	MediumWarningCount *int32 `json:"MediumWarningCount,omitempty" xml:"MediumWarningCount,omitempty"`
}

func (s GetClusterCheckItemWarningStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckItemWarningStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemWarningStatisticsResponseBodyData) GetHighWarningCount() *int32 {
	return s.HighWarningCount
}

func (s *GetClusterCheckItemWarningStatisticsResponseBodyData) GetLowWarningCount() *int32 {
	return s.LowWarningCount
}

func (s *GetClusterCheckItemWarningStatisticsResponseBodyData) GetMediumWarningCount() *int32 {
	return s.MediumWarningCount
}

func (s *GetClusterCheckItemWarningStatisticsResponseBodyData) SetHighWarningCount(v int32) *GetClusterCheckItemWarningStatisticsResponseBodyData {
	s.HighWarningCount = &v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponseBodyData) SetLowWarningCount(v int32) *GetClusterCheckItemWarningStatisticsResponseBodyData {
	s.LowWarningCount = &v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponseBodyData) SetMediumWarningCount(v int32) *GetClusterCheckItemWarningStatisticsResponseBodyData {
	s.MediumWarningCount = &v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
