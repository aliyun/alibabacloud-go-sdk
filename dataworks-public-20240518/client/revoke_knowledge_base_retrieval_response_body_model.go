// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeKnowledgeBaseRetrievalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevokeKnowledgeBaseRetrievalResponseBody
	GetCode() *string
	SetData(v *RevokeKnowledgeBaseRetrievalResponseBodyData) *RevokeKnowledgeBaseRetrievalResponseBody
	GetData() *RevokeKnowledgeBaseRetrievalResponseBodyData
	SetMessage(v string) *RevokeKnowledgeBaseRetrievalResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeKnowledgeBaseRetrievalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeKnowledgeBaseRetrievalResponseBody
	GetSuccess() *bool
}

type RevokeKnowledgeBaseRetrievalResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *RevokeKnowledgeBaseRetrievalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The security error message returned when the request fails.
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

func (s RevokeKnowledgeBaseRetrievalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeKnowledgeBaseRetrievalResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) GetData() *RevokeKnowledgeBaseRetrievalResponseBodyData {
	return s.Data
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) SetCode(v string) *RevokeKnowledgeBaseRetrievalResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) SetData(v *RevokeKnowledgeBaseRetrievalResponseBodyData) *RevokeKnowledgeBaseRetrievalResponseBody {
	s.Data = v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) SetMessage(v string) *RevokeKnowledgeBaseRetrievalResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) SetRequestId(v string) *RevokeKnowledgeBaseRetrievalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) SetSuccess(v bool) *RevokeKnowledgeBaseRetrievalResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RevokeKnowledgeBaseRetrievalResponseBodyData struct {
	// The stable ID of the associated knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The DataWorks user ID that receives the share. When a new user is added, the system verifies that the user is a member of the same tenant.
	//
	// example:
	//
	// 1000000000000000
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of principal that receives the share. This field is optional in the request. The default and the only response value is USER.
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s RevokeKnowledgeBaseRetrievalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RevokeKnowledgeBaseRetrievalResponseBodyData) GoString() string {
	return s.String()
}

func (s *RevokeKnowledgeBaseRetrievalResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RevokeKnowledgeBaseRetrievalResponseBodyData) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *RevokeKnowledgeBaseRetrievalResponseBodyData) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *RevokeKnowledgeBaseRetrievalResponseBodyData) SetKnowledgeBaseId(v string) *RevokeKnowledgeBaseRetrievalResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBodyData) SetPrincipalId(v string) *RevokeKnowledgeBaseRetrievalResponseBodyData {
	s.PrincipalId = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBodyData) SetPrincipalType(v string) *RevokeKnowledgeBaseRetrievalResponseBodyData {
	s.PrincipalType = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
