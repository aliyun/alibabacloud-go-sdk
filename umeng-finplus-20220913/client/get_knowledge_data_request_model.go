// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *GetKnowledgeDataRequestBody) *GetKnowledgeDataRequest
	GetBody() *GetKnowledgeDataRequestBody
}

type GetKnowledgeDataRequest struct {
	Body *GetKnowledgeDataRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s GetKnowledgeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeDataRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeDataRequest) GetBody() *GetKnowledgeDataRequestBody {
	return s.Body
}

func (s *GetKnowledgeDataRequest) SetBody(v *GetKnowledgeDataRequestBody) *GetKnowledgeDataRequest {
	s.Body = v
	return s
}

func (s *GetKnowledgeDataRequest) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeDataRequestBody struct {
	AppId           *string   `json:"appId,omitempty" xml:"appId,omitempty"`
	KnowledgeIdList []*string `json:"knowledgeIdList,omitempty" xml:"knowledgeIdList,omitempty" type:"Repeated"`
}

func (s GetKnowledgeDataRequestBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeDataRequestBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeDataRequestBody) GetAppId() *string {
	return s.AppId
}

func (s *GetKnowledgeDataRequestBody) GetKnowledgeIdList() []*string {
	return s.KnowledgeIdList
}

func (s *GetKnowledgeDataRequestBody) SetAppId(v string) *GetKnowledgeDataRequestBody {
	s.AppId = &v
	return s
}

func (s *GetKnowledgeDataRequestBody) SetKnowledgeIdList(v []*string) *GetKnowledgeDataRequestBody {
	s.KnowledgeIdList = v
	return s
}

func (s *GetKnowledgeDataRequestBody) Validate() error {
	return dara.Validate(s)
}
