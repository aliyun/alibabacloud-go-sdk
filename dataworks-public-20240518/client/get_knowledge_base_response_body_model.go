// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeBaseResponseBody
	GetCode() *string
	SetData(v *GetKnowledgeBaseResponseBodyData) *GetKnowledgeBaseResponseBody
	GetData() *GetKnowledgeBaseResponseBodyData
	SetMessage(v string) *GetKnowledgeBaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgeBaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgeBaseResponseBody
	GetSuccess() *bool
}

type GetKnowledgeBaseResponseBody struct {
	// The business error code. This field may be omitted if the request succeeds. If the request fails, the common error response protocol applies.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The public configuration information of the target knowledge base.
	Data *GetKnowledgeBaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// A human-readable description of the business error. This field may be omitted if the request succeeds. Provide the RequestId when troubleshooting failures.
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

func (s GetKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseResponseBody) GetData() *GetKnowledgeBaseResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeBaseResponseBody) SetCode(v string) *GetKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetData(v *GetKnowledgeBaseResponseBodyData) *GetKnowledgeBaseResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetMessage(v string) *GetKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetRequestId(v string) *GetKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetSuccess(v bool) *GetKnowledgeBaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeBaseResponseBodyData struct {
	// The read-only access role of the current caller for the knowledge base. Valid values: OWNER (owner), ADMIN (administrator), and VIEWER (viewer).
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
	// The description of the knowledge base. This field can be empty if it is not set.
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
	// The public status of the knowledge base. This list returns only normal and visible knowledge bases. The value is ACTIVE.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the knowledge base. Valid values: PERSONAL (personal knowledge base) and ENTERPRISE (enterprise knowledge base).
	//
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyData) GetAccessRole() *string {
	return s.AccessRole
}

func (s *GetKnowledgeBaseResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetKnowledgeBaseResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeBaseResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetKnowledgeBaseResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeBaseResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetKnowledgeBaseResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetKnowledgeBaseResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetKnowledgeBaseResponseBodyData) SetAccessRole(v string) *GetKnowledgeBaseResponseBodyData {
	s.AccessRole = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetCreateTime(v int64) *GetKnowledgeBaseResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetDescription(v string) *GetKnowledgeBaseResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetKnowledgeBaseId(v string) *GetKnowledgeBaseResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetModifyTime(v int64) *GetKnowledgeBaseResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetName(v string) *GetKnowledgeBaseResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetOwner(v string) *GetKnowledgeBaseResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetStatus(v string) *GetKnowledgeBaseResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) SetType(v string) *GetKnowledgeBaseResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
