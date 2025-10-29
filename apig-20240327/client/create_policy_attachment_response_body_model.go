// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePolicyAttachmentResponseBody
	GetCode() *string
	SetData(v *CreatePolicyAttachmentResponseBodyData) *CreatePolicyAttachmentResponseBody
	GetData() *CreatePolicyAttachmentResponseBodyData
	SetMessage(v string) *CreatePolicyAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePolicyAttachmentResponseBody
	GetRequestId() *string
}

type CreatePolicyAttachmentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *CreatePolicyAttachmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C64***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePolicyAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePolicyAttachmentResponseBody) GetData() *CreatePolicyAttachmentResponseBodyData {
	return s.Data
}

func (s *CreatePolicyAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolicyAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyAttachmentResponseBody) SetCode(v string) *CreatePolicyAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyAttachmentResponseBody) SetData(v *CreatePolicyAttachmentResponseBodyData) *CreatePolicyAttachmentResponseBody {
	s.Data = v
	return s
}

func (s *CreatePolicyAttachmentResponseBody) SetMessage(v string) *CreatePolicyAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolicyAttachmentResponseBody) SetRequestId(v string) *CreatePolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyAttachmentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolicyAttachmentResponseBodyData struct {
	// Policy Mount ID
	//
	// example:
	//
	// pr-cqooju5lhtgquuj6***
	PolicyAttachmentId *string `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
}

func (s CreatePolicyAttachmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyAttachmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentResponseBodyData) GetPolicyAttachmentId() *string {
	return s.PolicyAttachmentId
}

func (s *CreatePolicyAttachmentResponseBodyData) SetPolicyAttachmentId(v string) *CreatePolicyAttachmentResponseBodyData {
	s.PolicyAttachmentId = &v
	return s
}

func (s *CreatePolicyAttachmentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
