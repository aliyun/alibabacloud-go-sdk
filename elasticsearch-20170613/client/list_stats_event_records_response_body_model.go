// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatsEventRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListStatsEventRecordsResponseBody
	GetRequestId() *string
	SetResult(v *ListStatsEventRecordsResponseBodyResult) *ListStatsEventRecordsResponseBody
	GetResult() *ListStatsEventRecordsResponseBodyResult
}

type ListStatsEventRecordsResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListStatsEventRecordsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListStatsEventRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStatsEventRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatsEventRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStatsEventRecordsResponseBody) GetResult() *ListStatsEventRecordsResponseBodyResult {
	return s.Result
}

func (s *ListStatsEventRecordsResponseBody) SetRequestId(v string) *ListStatsEventRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatsEventRecordsResponseBody) SetResult(v *ListStatsEventRecordsResponseBodyResult) *ListStatsEventRecordsResponseBody {
	s.Result = v
	return s
}

func (s *ListStatsEventRecordsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStatsEventRecordsResponseBodyResult struct {
	Result []*ListStatsEventRecordsResponseBodyResultResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	Total *string `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListStatsEventRecordsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListStatsEventRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListStatsEventRecordsResponseBodyResult) GetResult() []*ListStatsEventRecordsResponseBodyResultResult {
	return s.Result
}

func (s *ListStatsEventRecordsResponseBodyResult) GetTotal() *string {
	return s.Total
}

func (s *ListStatsEventRecordsResponseBodyResult) SetResult(v []*ListStatsEventRecordsResponseBodyResultResult) *ListStatsEventRecordsResponseBodyResult {
	s.Result = v
	return s
}

func (s *ListStatsEventRecordsResponseBodyResult) SetTotal(v string) *ListStatsEventRecordsResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListStatsEventRecordsResponseBodyResult) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStatsEventRecordsResponseBodyResultResult struct {
	// example:
	//
	// 4
	Cnt *string `json:"cnt,omitempty" xml:"cnt,omitempty"`
	// example:
	//
	// Info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// Executed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// UserOperator
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListStatsEventRecordsResponseBodyResultResult) String() string {
	return dara.Prettify(s)
}

func (s ListStatsEventRecordsResponseBodyResultResult) GoString() string {
	return s.String()
}

func (s *ListStatsEventRecordsResponseBodyResultResult) GetCnt() *string {
	return s.Cnt
}

func (s *ListStatsEventRecordsResponseBodyResultResult) GetLevel() *string {
	return s.Level
}

func (s *ListStatsEventRecordsResponseBodyResultResult) GetStatus() *string {
	return s.Status
}

func (s *ListStatsEventRecordsResponseBodyResultResult) GetType() *string {
	return s.Type
}

func (s *ListStatsEventRecordsResponseBodyResultResult) SetCnt(v string) *ListStatsEventRecordsResponseBodyResultResult {
	s.Cnt = &v
	return s
}

func (s *ListStatsEventRecordsResponseBodyResultResult) SetLevel(v string) *ListStatsEventRecordsResponseBodyResultResult {
	s.Level = &v
	return s
}

func (s *ListStatsEventRecordsResponseBodyResultResult) SetStatus(v string) *ListStatsEventRecordsResponseBodyResultResult {
	s.Status = &v
	return s
}

func (s *ListStatsEventRecordsResponseBodyResultResult) SetType(v string) *ListStatsEventRecordsResponseBodyResultResult {
	s.Type = &v
	return s
}

func (s *ListStatsEventRecordsResponseBodyResultResult) Validate() error {
	return dara.Validate(s)
}
