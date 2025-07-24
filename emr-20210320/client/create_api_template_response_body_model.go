// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateApiTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateApiTemplateResponseBody
	GetSuccess() *string
	SetTemplateId(v string) *CreateApiTemplateResponseBody
	GetTemplateId() *string
}

type CreateApiTemplateResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Template ID (to be deprecated).
	//
	// example:
	//
	// at-41b4c6a0fc63****
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// Template ID (it is recommended to use the parameter TemplateId).
	//
	// example:
	//
	// at-41b4c6a0fc63****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateApiTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiTemplateResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateApiTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateApiTemplateResponseBody) SetRequestId(v string) *CreateApiTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiTemplateResponseBody) SetSuccess(v string) *CreateApiTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApiTemplateResponseBody) SetTemplateId(v string) *CreateApiTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateApiTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
