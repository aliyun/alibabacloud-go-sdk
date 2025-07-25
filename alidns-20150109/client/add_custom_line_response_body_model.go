// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLineCode(v string) *AddCustomLineResponseBody
	GetLineCode() *string
	SetLineId(v int64) *AddCustomLineResponseBody
	GetLineId() *int64
	SetRequestId(v string) *AddCustomLineResponseBody
	GetRequestId() *string
}

type AddCustomLineResponseBody struct {
	// The code of the custom line.
	//
	// example:
	//
	// hra0yc-597
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The unique ID of the custom line.
	//
	// example:
	//
	// 597
	LineId *int64 `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCustomLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomLineResponseBody) GetLineCode() *string {
	return s.LineCode
}

func (s *AddCustomLineResponseBody) GetLineId() *int64 {
	return s.LineId
}

func (s *AddCustomLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomLineResponseBody) SetLineCode(v string) *AddCustomLineResponseBody {
	s.LineCode = &v
	return s
}

func (s *AddCustomLineResponseBody) SetLineId(v int64) *AddCustomLineResponseBody {
	s.LineId = &v
	return s
}

func (s *AddCustomLineResponseBody) SetRequestId(v string) *AddCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomLineResponseBody) Validate() error {
	return dara.Validate(s)
}
