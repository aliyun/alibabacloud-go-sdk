// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndAttachPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAndAttachPolicyResponseBody
	GetCode() *string
	SetData(v *CreateAndAttachPolicyResponseBodyData) *CreateAndAttachPolicyResponseBody
	GetData() *CreateAndAttachPolicyResponseBodyData
	SetMessage(v string) *CreateAndAttachPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAndAttachPolicyResponseBody
	GetRequestId() *string
}

type CreateAndAttachPolicyResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateAndAttachPolicyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAndAttachPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndAttachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndAttachPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAndAttachPolicyResponseBody) GetData() *CreateAndAttachPolicyResponseBodyData {
	return s.Data
}

func (s *CreateAndAttachPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAndAttachPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndAttachPolicyResponseBody) SetCode(v string) *CreateAndAttachPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAndAttachPolicyResponseBody) SetData(v *CreateAndAttachPolicyResponseBodyData) *CreateAndAttachPolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreateAndAttachPolicyResponseBody) SetMessage(v string) *CreateAndAttachPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAndAttachPolicyResponseBody) SetRequestId(v string) *CreateAndAttachPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndAttachPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAndAttachPolicyResponseBodyData struct {
	Attachment *Attachment `json:"attachment,omitempty" xml:"attachment,omitempty"`
	// example:
	//
	// p-cq7l5s5lhtgi6qasrdc0
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s CreateAndAttachPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAndAttachPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAndAttachPolicyResponseBodyData) GetAttachment() *Attachment {
	return s.Attachment
}

func (s *CreateAndAttachPolicyResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateAndAttachPolicyResponseBodyData) SetAttachment(v *Attachment) *CreateAndAttachPolicyResponseBodyData {
	s.Attachment = v
	return s
}

func (s *CreateAndAttachPolicyResponseBodyData) SetPolicyId(v string) *CreateAndAttachPolicyResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *CreateAndAttachPolicyResponseBodyData) Validate() error {
	if s.Attachment != nil {
		if err := s.Attachment.Validate(); err != nil {
			return err
		}
	}
	return nil
}
