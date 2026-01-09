// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeDataResponseBody
	GetCode() *string
	SetData(v []*GetKnowledgeDataResponseBodyData) *GetKnowledgeDataResponseBody
	GetData() []*GetKnowledgeDataResponseBodyData
	SetMsg(v string) *GetKnowledgeDataResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetKnowledgeDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgeDataResponseBody
	GetSuccess() *bool
}

type GetKnowledgeDataResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetKnowledgeDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                             `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKnowledgeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeDataResponseBody) GetData() []*GetKnowledgeDataResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeDataResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetKnowledgeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeDataResponseBody) SetCode(v string) *GetKnowledgeDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeDataResponseBody) SetData(v []*GetKnowledgeDataResponseBodyData) *GetKnowledgeDataResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeDataResponseBody) SetMsg(v string) *GetKnowledgeDataResponseBody {
	s.Msg = &v
	return s
}

func (s *GetKnowledgeDataResponseBody) SetRequestId(v string) *GetKnowledgeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeDataResponseBody) SetSuccess(v bool) *GetKnowledgeDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeDataResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKnowledgeDataResponseBodyData struct {
	AppId               *string `json:"appId,omitempty" xml:"appId,omitempty"`
	InternalKnowledgeId *string `json:"internalKnowledgeId,omitempty" xml:"internalKnowledgeId,omitempty"`
	KnowledgeName       *string `json:"knowledgeName,omitempty" xml:"knowledgeName,omitempty"`
	Message             *string `json:"message,omitempty" xml:"message,omitempty"`
	Status              *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetKnowledgeDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeDataResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetKnowledgeDataResponseBodyData) GetInternalKnowledgeId() *string {
	return s.InternalKnowledgeId
}

func (s *GetKnowledgeDataResponseBodyData) GetKnowledgeName() *string {
	return s.KnowledgeName
}

func (s *GetKnowledgeDataResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeDataResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetKnowledgeDataResponseBodyData) SetAppId(v string) *GetKnowledgeDataResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetKnowledgeDataResponseBodyData) SetInternalKnowledgeId(v string) *GetKnowledgeDataResponseBodyData {
	s.InternalKnowledgeId = &v
	return s
}

func (s *GetKnowledgeDataResponseBodyData) SetKnowledgeName(v string) *GetKnowledgeDataResponseBodyData {
	s.KnowledgeName = &v
	return s
}

func (s *GetKnowledgeDataResponseBodyData) SetMessage(v string) *GetKnowledgeDataResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetKnowledgeDataResponseBodyData) SetStatus(v string) *GetKnowledgeDataResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetKnowledgeDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
