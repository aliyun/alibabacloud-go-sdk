// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSessionDistributionResponseBody
	GetRequestId() *string
	SetTimeList(v []*GetSessionDistributionResponseBodyTimeList) *GetSessionDistributionResponseBody
	GetTimeList() []*GetSessionDistributionResponseBodyTimeList
}

type GetSessionDistributionResponseBody struct {
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetSessionDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetSessionDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSessionDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSessionDistributionResponseBody) GetTimeList() []*GetSessionDistributionResponseBodyTimeList {
	return s.TimeList
}

func (s *GetSessionDistributionResponseBody) SetRequestId(v string) *GetSessionDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionDistributionResponseBody) SetTimeList(v []*GetSessionDistributionResponseBodyTimeList) *GetSessionDistributionResponseBody {
	s.TimeList = v
	return s
}

func (s *GetSessionDistributionResponseBody) Validate() error {
	if s.TimeList != nil {
		for _, item := range s.TimeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSessionDistributionResponseBodyTimeList struct {
	// example:
	//
	// 100
	ActiveSessionCount *int64 `json:"ActiveSessionCount,omitempty" xml:"ActiveSessionCount,omitempty"`
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 2019-06-06 01:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 50
	ErrorSessionCount *int64 `json:"ErrorSessionCount,omitempty" xml:"ErrorSessionCount,omitempty"`
	// example:
	//
	// 1000
	LoginSessionCount *int64 `json:"LoginSessionCount,omitempty" xml:"LoginSessionCount,omitempty"`
}

func (s GetSessionDistributionResponseBodyTimeList) String() string {
	return dara.Prettify(s)
}

func (s GetSessionDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionResponseBodyTimeList) GetActiveSessionCount() *int64 {
	return s.ActiveSessionCount
}

func (s *GetSessionDistributionResponseBodyTimeList) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetSessionDistributionResponseBodyTimeList) GetEndDate() *string {
	return s.EndDate
}

func (s *GetSessionDistributionResponseBodyTimeList) GetErrorSessionCount() *int64 {
	return s.ErrorSessionCount
}

func (s *GetSessionDistributionResponseBodyTimeList) GetLoginSessionCount() *int64 {
	return s.LoginSessionCount
}

func (s *GetSessionDistributionResponseBodyTimeList) SetActiveSessionCount(v int64) *GetSessionDistributionResponseBodyTimeList {
	s.ActiveSessionCount = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetBeginDate(v string) *GetSessionDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetEndDate(v string) *GetSessionDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetErrorSessionCount(v int64) *GetSessionDistributionResponseBodyTimeList {
	s.ErrorSessionCount = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetLoginSessionCount(v int64) *GetSessionDistributionResponseBodyTimeList {
	s.LoginSessionCount = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) Validate() error {
	return dara.Validate(s)
}
