// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAgentTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateAppAgentTemplateResponseBody
	GetId() *string
	SetRequestId(v string) *CreateAppAgentTemplateResponseBody
	GetRequestId() *string
}

type CreateAppAgentTemplateResponseBody struct {
	// example:
	//
	// 1223131213231313213
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppAgentTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateAppAgentTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppAgentTemplateResponseBody) SetId(v string) *CreateAppAgentTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateAppAgentTemplateResponseBody) SetRequestId(v string) *CreateAppAgentTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppAgentTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
