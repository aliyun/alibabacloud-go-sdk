// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBasesResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeBasesResponseBodyData) *ListKnowledgeBasesResponseBody
	GetData() *ListKnowledgeBasesResponseBodyData
	SetMaxResults(v int32) *ListKnowledgeBasesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListKnowledgeBasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeBasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeBasesResponseBody
	GetSuccess() *bool
}

type ListKnowledgeBasesResponseBody struct {
	// The business error code. This field is omitted when the request is successful. If the request fails, the common error response protocol is followed.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of knowledge bases on the current page and the cursor for the next page.
	Data *ListKnowledgeBasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of knowledge bases to return per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The readable description of the business error. This field is omitted when the request is successful. Provide the RequestId along with this message when troubleshooting failures.
	//
	// example:
	//
	// The request parameters are invalid.
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

func (s ListKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBasesResponseBody) GetData() *ListKnowledgeBasesResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeBasesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeBasesResponseBody) SetCode(v string) *ListKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetData(v *ListKnowledgeBasesResponseBodyData) *ListKnowledgeBasesResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetMaxResults(v int32) *ListKnowledgeBasesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetMessage(v string) *ListKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetRequestId(v string) *ListKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetSuccess(v bool) *ListKnowledgeBasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBasesResponseBodyData struct {
	// The list of knowledge bases that the current caller has permission to view in the current region.
	KnowledgeBases []*ListKnowledgeBasesResponseBodyDataKnowledgeBases `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty" type:"Repeated"`
	// The pagination cursor. Do not specify this parameter when you request the first page. For subsequent pages, use the NextToken value returned in the previous response as is. An empty response indicates that no more pages are available. If the current page is empty but the cursor is not empty, continue pagination. Keep the same MaxResults value during pagination.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgeBasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponseBodyData) GetKnowledgeBases() []*ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	return s.KnowledgeBases
}

func (s *ListKnowledgeBasesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesResponseBodyData) SetKnowledgeBases(v []*ListKnowledgeBasesResponseBodyDataKnowledgeBases) *ListKnowledgeBasesResponseBodyData {
	s.KnowledgeBases = v
	return s
}

func (s *ListKnowledgeBasesResponseBodyData) SetNextToken(v string) *ListKnowledgeBasesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyData) Validate() error {
	if s.KnowledgeBases != nil {
		for _, item := range s.KnowledgeBases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeBasesResponseBodyDataKnowledgeBases struct {
	// The read-only access role of the current caller for the knowledge base. Valid values:
	//
	// - OWNER: owner
	//
	// - ADMIN: administrator
	//
	// - VIEWER: viewer
	//
	// example:
	//
	// OWNER
	AccessRole *string `json:"AccessRole,omitempty" xml:"AccessRole,omitempty"`
	// The creation time of the knowledge base. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1794369600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the knowledge base. This field can be empty if no description is set.
	//
	// example:
	//
	// Knowledge about sales metrics, orders, and business policies
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The stable ID of the knowledge base, which can be used to limit the search scope.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The last modification time of the knowledge base. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1794373200000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the knowledge base.
	//
	// example:
	//
	// Sales knowledge base
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks user ID of the knowledge base owner.
	//
	// example:
	//
	// 1000000000000000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The public status of the knowledge base. This list returns only normally visible knowledge bases, and the value is ACTIVE.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the knowledge base. Valid values:
	//
	// - PERSONAL: personal knowledge base
	//
	// - ENTERPRISE: enterprise knowledge base
	//
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgeBasesResponseBodyDataKnowledgeBases) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponseBodyDataKnowledgeBases) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetAccessRole() *string {
	return s.AccessRole
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetOwner() *string {
	return s.Owner
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetStatus() *string {
	return s.Status
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetAccessRole(v string) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.AccessRole = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetCreateTime(v int64) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.CreateTime = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetDescription(v string) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.Description = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetKnowledgeBaseId(v string) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetModifyTime(v int64) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.ModifyTime = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetName(v string) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.Name = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetOwner(v string) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.Owner = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetStatus(v string) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.Status = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) SetType(v string) *ListKnowledgeBasesResponseBodyDataKnowledgeBases {
	s.Type = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyDataKnowledgeBases) Validate() error {
	return dara.Validate(s)
}
