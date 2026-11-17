// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseAdministratorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBaseAdministratorsResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeBaseAdministratorsResponseBodyData) *ListKnowledgeBaseAdministratorsResponseBody
	GetData() *ListKnowledgeBaseAdministratorsResponseBodyData
	SetMessage(v string) *ListKnowledgeBaseAdministratorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeBaseAdministratorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeBaseAdministratorsResponseBody
	GetSuccess() *bool
}

type ListKnowledgeBaseAdministratorsResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *ListKnowledgeBaseAdministratorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListKnowledgeBaseAdministratorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseAdministratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) GetData() *ListKnowledgeBaseAdministratorsResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) SetCode(v string) *ListKnowledgeBaseAdministratorsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) SetData(v *ListKnowledgeBaseAdministratorsResponseBodyData) *ListKnowledgeBaseAdministratorsResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) SetMessage(v string) *ListKnowledgeBaseAdministratorsResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) SetRequestId(v string) *ListKnowledgeBaseAdministratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) SetSuccess(v bool) *ListKnowledgeBaseAdministratorsResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBaseAdministratorsResponseBodyData struct {
	// The administrators on the current page, excluding the owner.
	KnowledgeBaseAdministrators []*ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators `json:"KnowledgeBaseAdministrators,omitempty" xml:"KnowledgeBaseAdministrators,omitempty" type:"Repeated"`
	// The maximum number of entries per page used for the current request. If the input parameter is omitted, the value is 20. This is not the actual number of entries returned on the current page.
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

func (s ListKnowledgeBaseAdministratorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseAdministratorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyData) GetKnowledgeBaseAdministrators() []*ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators {
	return s.KnowledgeBaseAdministrators
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyData) SetKnowledgeBaseAdministrators(v []*ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) *ListKnowledgeBaseAdministratorsResponseBodyData {
	s.KnowledgeBaseAdministrators = v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyData) SetMaxResults(v int32) *ListKnowledgeBaseAdministratorsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyData) SetNextToken(v string) *ListKnowledgeBaseAdministratorsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyData) Validate() error {
	if s.KnowledgeBaseAdministrators != nil {
		for _, item := range s.KnowledgeBaseAdministrators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators struct {
	// The creation time of the record, represented as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The display name of the user. This field is optional and is not used to verify the user’s identity.
	//
	// example:
	//
	// User A
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the knowledge base.
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

func (s ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) GetUserId() *string {
	return s.UserId
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) SetCreateTime(v int64) *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators {
	s.CreateTime = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) SetDisplayName(v string) *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators {
	s.DisplayName = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) SetKnowledgeBaseId(v string) *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) SetUserId(v string) *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators {
	s.UserId = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponseBodyDataKnowledgeBaseAdministrators) Validate() error {
	return dara.Validate(s)
}
