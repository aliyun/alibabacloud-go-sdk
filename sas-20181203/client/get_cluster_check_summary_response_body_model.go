// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetClusterCheckSummaryResponseBodyData) *GetClusterCheckSummaryResponseBody
	GetData() *GetClusterCheckSummaryResponseBodyData
	SetRequestId(v string) *GetClusterCheckSummaryResponseBody
	GetRequestId() *string
}

type GetClusterCheckSummaryResponseBody struct {
	// Return data.
	Data *GetClusterCheckSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterCheckSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterCheckSummaryResponseBody) GetData() *GetClusterCheckSummaryResponseBodyData {
	return s.Data
}

func (s *GetClusterCheckSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterCheckSummaryResponseBody) SetData(v *GetClusterCheckSummaryResponseBodyData) *GetClusterCheckSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterCheckSummaryResponseBody) SetRequestId(v string) *GetClusterCheckSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterCheckSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterCheckSummaryResponseBodyData struct {
	// Total number of items that failed the check.
	//
	// example:
	//
	// 6
	NotPassCount *int32 `json:"NotPassCount,omitempty" xml:"NotPassCount,omitempty"`
	// Number of high-risk inspection items that have not passed.
	//
	// example:
	//
	// 1
	NotPassHighCount *int32 `json:"NotPassHighCount,omitempty" xml:"NotPassHighCount,omitempty"`
	// Number of low-risk inspection items that have not passed.
	//
	// example:
	//
	// 3
	NotPassLowCount *int32 `json:"NotPassLowCount,omitempty" xml:"NotPassLowCount,omitempty"`
	// Number of medium-risk failed inspection items.
	//
	// example:
	//
	// 2
	NotPassMediumCount *int32 `json:"NotPassMediumCount,omitempty" xml:"NotPassMediumCount,omitempty"`
}

func (s GetClusterCheckSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterCheckSummaryResponseBodyData) GetNotPassCount() *int32 {
	return s.NotPassCount
}

func (s *GetClusterCheckSummaryResponseBodyData) GetNotPassHighCount() *int32 {
	return s.NotPassHighCount
}

func (s *GetClusterCheckSummaryResponseBodyData) GetNotPassLowCount() *int32 {
	return s.NotPassLowCount
}

func (s *GetClusterCheckSummaryResponseBodyData) GetNotPassMediumCount() *int32 {
	return s.NotPassMediumCount
}

func (s *GetClusterCheckSummaryResponseBodyData) SetNotPassCount(v int32) *GetClusterCheckSummaryResponseBodyData {
	s.NotPassCount = &v
	return s
}

func (s *GetClusterCheckSummaryResponseBodyData) SetNotPassHighCount(v int32) *GetClusterCheckSummaryResponseBodyData {
	s.NotPassHighCount = &v
	return s
}

func (s *GetClusterCheckSummaryResponseBodyData) SetNotPassLowCount(v int32) *GetClusterCheckSummaryResponseBodyData {
	s.NotPassLowCount = &v
	return s
}

func (s *GetClusterCheckSummaryResponseBodyData) SetNotPassMediumCount(v int32) *GetClusterCheckSummaryResponseBodyData {
	s.NotPassMediumCount = &v
	return s
}

func (s *GetClusterCheckSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
