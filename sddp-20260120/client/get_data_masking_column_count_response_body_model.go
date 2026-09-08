// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataMaskingColumnCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnCount(v *GetDataMaskingColumnCountResponseBodyColumnCount) *GetDataMaskingColumnCountResponseBody
	GetColumnCount() *GetDataMaskingColumnCountResponseBodyColumnCount
	SetRequestId(v string) *GetDataMaskingColumnCountResponseBody
	GetRequestId() *string
}

type GetDataMaskingColumnCountResponseBody struct {
	ColumnCount *GetDataMaskingColumnCountResponseBodyColumnCount `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty" type:"Struct"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataMaskingColumnCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingColumnCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataMaskingColumnCountResponseBody) GetColumnCount() *GetDataMaskingColumnCountResponseBodyColumnCount {
	return s.ColumnCount
}

func (s *GetDataMaskingColumnCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataMaskingColumnCountResponseBody) SetColumnCount(v *GetDataMaskingColumnCountResponseBodyColumnCount) *GetDataMaskingColumnCountResponseBody {
	s.ColumnCount = v
	return s
}

func (s *GetDataMaskingColumnCountResponseBody) SetRequestId(v string) *GetDataMaskingColumnCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataMaskingColumnCountResponseBody) Validate() error {
	if s.ColumnCount != nil {
		if err := s.ColumnCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataMaskingColumnCountResponseBodyColumnCount struct {
	// example:
	//
	// 10
	MaskedCount *int64 `json:"MaskedCount,omitempty" xml:"MaskedCount,omitempty"`
	// example:
	//
	// 1
	MaskingFailedCount *int64 `json:"MaskingFailedCount,omitempty" xml:"MaskingFailedCount,omitempty"`
	// example:
	//
	// 20
	SensitiveCount *int64 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDataMaskingColumnCountResponseBodyColumnCount) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingColumnCountResponseBodyColumnCount) GoString() string {
	return s.String()
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) GetMaskedCount() *int64 {
	return s.MaskedCount
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) GetMaskingFailedCount() *int64 {
	return s.MaskingFailedCount
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) GetSensitiveCount() *int64 {
	return s.SensitiveCount
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) SetMaskedCount(v int64) *GetDataMaskingColumnCountResponseBodyColumnCount {
	s.MaskedCount = &v
	return s
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) SetMaskingFailedCount(v int64) *GetDataMaskingColumnCountResponseBodyColumnCount {
	s.MaskingFailedCount = &v
	return s
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) SetSensitiveCount(v int64) *GetDataMaskingColumnCountResponseBodyColumnCount {
	s.SensitiveCount = &v
	return s
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) SetTotalCount(v int64) *GetDataMaskingColumnCountResponseBodyColumnCount {
	s.TotalCount = &v
	return s
}

func (s *GetDataMaskingColumnCountResponseBodyColumnCount) Validate() error {
	return dara.Validate(s)
}
