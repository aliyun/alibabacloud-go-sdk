// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDataAgentKnowledgeBaseResponseBodyData) *CreateDataAgentKnowledgeBaseResponseBody
	GetData() *CreateDataAgentKnowledgeBaseResponseBodyData
	SetErrorCode(v string) *CreateDataAgentKnowledgeBaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataAgentKnowledgeBaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataAgentKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataAgentKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type CreateDataAgentKnowledgeBaseResponseBody struct {
	Data *CreateDataAgentKnowledgeBaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataAgentKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) GetData() *CreateDataAgentKnowledgeBaseResponseBodyData {
	return s.Data
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) SetData(v *CreateDataAgentKnowledgeBaseResponseBodyData) *CreateDataAgentKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) SetErrorCode(v string) *CreateDataAgentKnowledgeBaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) SetErrorMessage(v string) *CreateDataAgentKnowledgeBaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) SetRequestId(v string) *CreateDataAgentKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) SetSuccess(v bool) *CreateDataAgentKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentKnowledgeBaseResponseBodyData struct {
	// example:
	//
	// kb-HZ-ra99akg0t*********1bku
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s CreateDataAgentKnowledgeBaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentKnowledgeBaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataAgentKnowledgeBaseResponseBodyData) GetKbUuid() *string {
	return s.KbUuid
}

func (s *CreateDataAgentKnowledgeBaseResponseBodyData) SetKbUuid(v string) *CreateDataAgentKnowledgeBaseResponseBodyData {
	s.KbUuid = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
