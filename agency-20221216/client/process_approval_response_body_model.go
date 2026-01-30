// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProcessApprovalResponseBody
	GetCode() *string
	SetData(v bool) *ProcessApprovalResponseBody
	GetData() *bool
	SetMessage(v string) *ProcessApprovalResponseBody
	GetMessage() *string
	SetRequestId(v string) *ProcessApprovalResponseBody
	GetRequestId() *string
}

type ProcessApprovalResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// id of request
	//
	// example:
	//
	// 23309219-4A34-589D-A3E0-9B2A3BFFD24F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProcessApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProcessApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessApprovalResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProcessApprovalResponseBody) GetData() *bool {
	return s.Data
}

func (s *ProcessApprovalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProcessApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProcessApprovalResponseBody) SetCode(v string) *ProcessApprovalResponseBody {
	s.Code = &v
	return s
}

func (s *ProcessApprovalResponseBody) SetData(v bool) *ProcessApprovalResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessApprovalResponseBody) SetMessage(v string) *ProcessApprovalResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessApprovalResponseBody) SetRequestId(v string) *ProcessApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessApprovalResponseBody) Validate() error {
	return dara.Validate(s)
}
