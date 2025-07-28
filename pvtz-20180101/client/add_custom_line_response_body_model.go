// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLineId(v string) *AddCustomLineResponseBody
	GetLineId() *string
	SetName(v string) *AddCustomLineResponseBody
	GetName() *string
	SetRequestId(v string) *AddCustomLineResponseBody
	GetRequestId() *string
}

type AddCustomLineResponseBody struct {
	// The unique ID of the custom line.
	//
	// example:
	//
	// 1065001
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The name of the custom line.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCustomLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomLineResponseBody) GetLineId() *string {
	return s.LineId
}

func (s *AddCustomLineResponseBody) GetName() *string {
	return s.Name
}

func (s *AddCustomLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomLineResponseBody) SetLineId(v string) *AddCustomLineResponseBody {
	s.LineId = &v
	return s
}

func (s *AddCustomLineResponseBody) SetName(v string) *AddCustomLineResponseBody {
	s.Name = &v
	return s
}

func (s *AddCustomLineResponseBody) SetRequestId(v string) *AddCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomLineResponseBody) Validate() error {
	return dara.Validate(s)
}
