// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantKnowledgeBaseRetrievalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GrantKnowledgeBaseRetrievalResponseBody
	GetCode() *string
	SetData(v *GrantKnowledgeBaseRetrievalResponseBodyData) *GrantKnowledgeBaseRetrievalResponseBody
	GetData() *GrantKnowledgeBaseRetrievalResponseBodyData
	SetMessage(v string) *GrantKnowledgeBaseRetrievalResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantKnowledgeBaseRetrievalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GrantKnowledgeBaseRetrievalResponseBody
	GetSuccess() *bool
}

type GrantKnowledgeBaseRetrievalResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *GrantKnowledgeBaseRetrievalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GrantKnowledgeBaseRetrievalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantKnowledgeBaseRetrievalResponseBody) GoString() string {
	return s.String()
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) GetCode() *string {
	return s.Code
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) GetData() *GrantKnowledgeBaseRetrievalResponseBodyData {
	return s.Data
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) SetCode(v string) *GrantKnowledgeBaseRetrievalResponseBody {
	s.Code = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) SetData(v *GrantKnowledgeBaseRetrievalResponseBodyData) *GrantKnowledgeBaseRetrievalResponseBody {
	s.Data = v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) SetMessage(v string) *GrantKnowledgeBaseRetrievalResponseBody {
	s.Message = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) SetRequestId(v string) *GrantKnowledgeBaseRetrievalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) SetSuccess(v bool) *GrantKnowledgeBaseRetrievalResponseBody {
	s.Success = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GrantKnowledgeBaseRetrievalResponseBodyData struct {
	// The stable ID of the associated knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The DataWorks user ID of the principal that receives the share. The system verifies that the user belongs to the same tenant when the principal is added.
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

func (s GrantKnowledgeBaseRetrievalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GrantKnowledgeBaseRetrievalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GrantKnowledgeBaseRetrievalResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GrantKnowledgeBaseRetrievalResponseBodyData) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GrantKnowledgeBaseRetrievalResponseBodyData) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GrantKnowledgeBaseRetrievalResponseBodyData) SetKnowledgeBaseId(v string) *GrantKnowledgeBaseRetrievalResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBodyData) SetPrincipalId(v string) *GrantKnowledgeBaseRetrievalResponseBodyData {
	s.PrincipalId = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBodyData) SetPrincipalType(v string) *GrantKnowledgeBaseRetrievalResponseBodyData {
	s.PrincipalType = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
