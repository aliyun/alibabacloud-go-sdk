// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDataAgentKnowledgeBaseResponseBodyData) *DeleteDataAgentKnowledgeBaseResponseBody
	GetData() *DeleteDataAgentKnowledgeBaseResponseBodyData
	SetErrorCode(v string) *DeleteDataAgentKnowledgeBaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataAgentKnowledgeBaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataAgentKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataAgentKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type DeleteDataAgentKnowledgeBaseResponseBody struct {
	Data *DeleteDataAgentKnowledgeBaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DeleteDataAgentKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) GetData() *DeleteDataAgentKnowledgeBaseResponseBodyData {
	return s.Data
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) SetData(v *DeleteDataAgentKnowledgeBaseResponseBodyData) *DeleteDataAgentKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) SetErrorCode(v string) *DeleteDataAgentKnowledgeBaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) SetErrorMessage(v string) *DeleteDataAgentKnowledgeBaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) SetRequestId(v string) *DeleteDataAgentKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) SetSuccess(v bool) *DeleteDataAgentKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDataAgentKnowledgeBaseResponseBodyData struct {
	// example:
	//
	// kb-HZ-rtl5lwx********q32d3ux
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s DeleteDataAgentKnowledgeBaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentKnowledgeBaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentKnowledgeBaseResponseBodyData) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DeleteDataAgentKnowledgeBaseResponseBodyData) SetKbUuid(v string) *DeleteDataAgentKnowledgeBaseResponseBodyData {
	s.KbUuid = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
