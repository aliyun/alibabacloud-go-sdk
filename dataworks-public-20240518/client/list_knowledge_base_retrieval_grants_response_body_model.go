// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseRetrievalGrantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBaseRetrievalGrantsResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeBaseRetrievalGrantsResponseBodyData) *ListKnowledgeBaseRetrievalGrantsResponseBody
	GetData() *ListKnowledgeBaseRetrievalGrantsResponseBodyData
	SetMessage(v string) *ListKnowledgeBaseRetrievalGrantsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeBaseRetrievalGrantsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeBaseRetrievalGrantsResponseBody
	GetSuccess() *bool
}

type ListKnowledgeBaseRetrievalGrantsResponseBody struct {
	// The business error code. You can use this code to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *ListKnowledgeBaseRetrievalGrantsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKnowledgeBaseRetrievalGrantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseRetrievalGrantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) GetData() *ListKnowledgeBaseRetrievalGrantsResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) SetCode(v string) *ListKnowledgeBaseRetrievalGrantsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) SetData(v *ListKnowledgeBaseRetrievalGrantsResponseBodyData) *ListKnowledgeBaseRetrievalGrantsResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) SetMessage(v string) *ListKnowledgeBaseRetrievalGrantsResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) SetRequestId(v string) *ListKnowledgeBaseRetrievalGrantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) SetSuccess(v bool) *ListKnowledgeBaseRetrievalGrantsResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBaseRetrievalGrantsResponseBodyData struct {
	// The user grants on the current page.
	KnowledgeBaseRetrievalGrants []*ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants `json:"KnowledgeBaseRetrievalGrants,omitempty" xml:"KnowledgeBaseRetrievalGrants,omitempty" type:"Repeated"`
	// The maximum number of entries per page used in the current request. If you do not specify this parameter in the request, the value is 20. This value is not the actual number of entries returned on the current page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The cursor for the next page. If no subsequent pages exist, null is returned.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgeBaseRetrievalGrantsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseRetrievalGrantsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyData) GetKnowledgeBaseRetrievalGrants() []*ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants {
	return s.KnowledgeBaseRetrievalGrants
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyData) SetKnowledgeBaseRetrievalGrants(v []*ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) *ListKnowledgeBaseRetrievalGrantsResponseBodyData {
	s.KnowledgeBaseRetrievalGrants = v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyData) SetMaxResults(v int32) *ListKnowledgeBaseRetrievalGrantsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyData) SetNextToken(v string) *ListKnowledgeBaseRetrievalGrantsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyData) Validate() error {
	if s.KnowledgeBaseRetrievalGrants != nil {
		for _, item := range s.KnowledgeBaseRetrievalGrants {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants struct {
	// The knowledge base ID.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The DataWorks user ID of the principal that receives the share. When a principal is added, the system verifies whether the user is a member of the same tenant.
	//
	// example:
	//
	// 1000000000000000
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of principal that receives the share. Default value: USER. Only USER is supported in the initial release.
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) SetKnowledgeBaseId(v string) *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) SetPrincipalId(v string) *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants {
	s.PrincipalId = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) SetPrincipalType(v string) *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants {
	s.PrincipalType = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponseBodyDataKnowledgeBaseRetrievalGrants) Validate() error {
	return dara.Validate(s)
}
