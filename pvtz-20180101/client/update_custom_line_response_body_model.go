// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLineId(v string) *UpdateCustomLineResponseBody
	GetLineId() *string
	SetRequestId(v string) *UpdateCustomLineResponseBody
	GetRequestId() *string
}

type UpdateCustomLineResponseBody struct {
	// The unique ID of the custom line.
	//
	// example:
	//
	// 765001
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineResponseBody) GetLineId() *string {
	return s.LineId
}

func (s *UpdateCustomLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomLineResponseBody) SetLineId(v string) *UpdateCustomLineResponseBody {
	s.LineId = &v
	return s
}

func (s *UpdateCustomLineResponseBody) SetRequestId(v string) *UpdateCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomLineResponseBody) Validate() error {
	return dara.Validate(s)
}
