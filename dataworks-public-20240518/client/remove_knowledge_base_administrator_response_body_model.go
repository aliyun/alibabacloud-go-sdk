// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeBaseAdministratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveKnowledgeBaseAdministratorResponseBody
	GetCode() *string
	SetData(v *RemoveKnowledgeBaseAdministratorResponseBodyData) *RemoveKnowledgeBaseAdministratorResponseBody
	GetData() *RemoveKnowledgeBaseAdministratorResponseBodyData
	SetMessage(v string) *RemoveKnowledgeBaseAdministratorResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveKnowledgeBaseAdministratorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveKnowledgeBaseAdministratorResponseBody
	GetSuccess() *bool
}

type RemoveKnowledgeBaseAdministratorResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *RemoveKnowledgeBaseAdministratorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveKnowledgeBaseAdministratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeBaseAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) GetData() *RemoveKnowledgeBaseAdministratorResponseBodyData {
	return s.Data
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) SetCode(v string) *RemoveKnowledgeBaseAdministratorResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) SetData(v *RemoveKnowledgeBaseAdministratorResponseBodyData) *RemoveKnowledgeBaseAdministratorResponseBody {
	s.Data = v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) SetMessage(v string) *RemoveKnowledgeBaseAdministratorResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) SetRequestId(v string) *RemoveKnowledgeBaseAdministratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) SetSuccess(v bool) *RemoveKnowledgeBaseAdministratorResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveKnowledgeBaseAdministratorResponseBodyData struct {
	// The stable ID of the associated knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The DataWorks user ID of the administrator.
	//
	// example:
	//
	// 1000000000000000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveKnowledgeBaseAdministratorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeBaseAdministratorResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeBaseAdministratorResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RemoveKnowledgeBaseAdministratorResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *RemoveKnowledgeBaseAdministratorResponseBodyData) SetKnowledgeBaseId(v string) *RemoveKnowledgeBaseAdministratorResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponseBodyData) SetUserId(v string) *RemoveKnowledgeBaseAdministratorResponseBodyData {
	s.UserId = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
