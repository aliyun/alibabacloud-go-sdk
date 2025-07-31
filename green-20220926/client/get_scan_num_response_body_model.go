// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLimitNumber(v int64) *GetScanNumResponseBody
	GetLimitNumber() *int64
	SetRequestId(v string) *GetScanNumResponseBody
	GetRequestId() *string
	SetScanNumber(v int64) *GetScanNumResponseBody
	GetScanNumber() *int64
	SetSumNumber(v int64) *GetScanNumResponseBody
	GetSumNumber() *int64
	SetTag(v bool) *GetScanNumResponseBody
	GetTag() *bool
}

type GetScanNumResponseBody struct {
	// example:
	//
	// 10
	LimitNumber *int64 `json:"LimitNumber,omitempty" xml:"LimitNumber,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	ScanNumber *int64 `json:"ScanNumber,omitempty" xml:"ScanNumber,omitempty"`
	// example:
	//
	// 10
	SumNumber *int64 `json:"SumNumber,omitempty" xml:"SumNumber,omitempty"`
	// example:
	//
	// True
	Tag *bool `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetScanNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScanNumResponseBody) GoString() string {
	return s.String()
}

func (s *GetScanNumResponseBody) GetLimitNumber() *int64 {
	return s.LimitNumber
}

func (s *GetScanNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScanNumResponseBody) GetScanNumber() *int64 {
	return s.ScanNumber
}

func (s *GetScanNumResponseBody) GetSumNumber() *int64 {
	return s.SumNumber
}

func (s *GetScanNumResponseBody) GetTag() *bool {
	return s.Tag
}

func (s *GetScanNumResponseBody) SetLimitNumber(v int64) *GetScanNumResponseBody {
	s.LimitNumber = &v
	return s
}

func (s *GetScanNumResponseBody) SetRequestId(v string) *GetScanNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScanNumResponseBody) SetScanNumber(v int64) *GetScanNumResponseBody {
	s.ScanNumber = &v
	return s
}

func (s *GetScanNumResponseBody) SetSumNumber(v int64) *GetScanNumResponseBody {
	s.SumNumber = &v
	return s
}

func (s *GetScanNumResponseBody) SetTag(v bool) *GetScanNumResponseBody {
	s.Tag = &v
	return s
}

func (s *GetScanNumResponseBody) Validate() error {
	return dara.Validate(s)
}
