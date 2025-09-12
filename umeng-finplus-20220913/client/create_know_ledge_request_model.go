// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowLedgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateKnowLedgeRequestBody) *CreateKnowLedgeRequest
	GetBody() *CreateKnowLedgeRequestBody
}

type CreateKnowLedgeRequest struct {
	Body *CreateKnowLedgeRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateKnowLedgeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowLedgeRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowLedgeRequest) GetBody() *CreateKnowLedgeRequestBody {
	return s.Body
}

func (s *CreateKnowLedgeRequest) SetBody(v *CreateKnowLedgeRequestBody) *CreateKnowLedgeRequest {
	s.Body = v
	return s
}

func (s *CreateKnowLedgeRequest) Validate() error {
	return dara.Validate(s)
}

type CreateKnowLedgeRequestBody struct {
	AppId               *string `json:"appId,omitempty" xml:"appId,omitempty"`
	InternalKnowledgeId *string `json:"internalKnowledgeId,omitempty" xml:"internalKnowledgeId,omitempty"`
	KnowledgeName       *string `json:"knowledgeName,omitempty" xml:"knowledgeName,omitempty"`
}

func (s CreateKnowLedgeRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowLedgeRequestBody) GoString() string {
	return s.String()
}

func (s *CreateKnowLedgeRequestBody) GetAppId() *string {
	return s.AppId
}

func (s *CreateKnowLedgeRequestBody) GetInternalKnowledgeId() *string {
	return s.InternalKnowledgeId
}

func (s *CreateKnowLedgeRequestBody) GetKnowledgeName() *string {
	return s.KnowledgeName
}

func (s *CreateKnowLedgeRequestBody) SetAppId(v string) *CreateKnowLedgeRequestBody {
	s.AppId = &v
	return s
}

func (s *CreateKnowLedgeRequestBody) SetInternalKnowledgeId(v string) *CreateKnowLedgeRequestBody {
	s.InternalKnowledgeId = &v
	return s
}

func (s *CreateKnowLedgeRequestBody) SetKnowledgeName(v string) *CreateKnowLedgeRequestBody {
	s.KnowledgeName = &v
	return s
}

func (s *CreateKnowLedgeRequestBody) Validate() error {
	return dara.Validate(s)
}
