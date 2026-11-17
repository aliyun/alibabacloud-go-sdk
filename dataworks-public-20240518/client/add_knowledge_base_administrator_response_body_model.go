// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeBaseAdministratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddKnowledgeBaseAdministratorResponseBody
	GetCode() *string
	SetData(v *AddKnowledgeBaseAdministratorResponseBodyData) *AddKnowledgeBaseAdministratorResponseBody
	GetData() *AddKnowledgeBaseAdministratorResponseBodyData
	SetMessage(v string) *AddKnowledgeBaseAdministratorResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddKnowledgeBaseAdministratorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddKnowledgeBaseAdministratorResponseBody
	GetSuccess() *bool
}

type AddKnowledgeBaseAdministratorResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *AddKnowledgeBaseAdministratorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s AddKnowledgeBaseAdministratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeBaseAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *AddKnowledgeBaseAdministratorResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddKnowledgeBaseAdministratorResponseBody) GetData() *AddKnowledgeBaseAdministratorResponseBodyData {
	return s.Data
}

func (s *AddKnowledgeBaseAdministratorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddKnowledgeBaseAdministratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKnowledgeBaseAdministratorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddKnowledgeBaseAdministratorResponseBody) SetCode(v string) *AddKnowledgeBaseAdministratorResponseBody {
	s.Code = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBody) SetData(v *AddKnowledgeBaseAdministratorResponseBodyData) *AddKnowledgeBaseAdministratorResponseBody {
	s.Data = v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBody) SetMessage(v string) *AddKnowledgeBaseAdministratorResponseBody {
	s.Message = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBody) SetRequestId(v string) *AddKnowledgeBaseAdministratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBody) SetSuccess(v bool) *AddKnowledgeBaseAdministratorResponseBody {
	s.Success = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddKnowledgeBaseAdministratorResponseBodyData struct {
	// The creation time of the record, represented as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The display name of the user. This parameter is optional and is not used as an identity basis.
	//
	// example:
	//
	// User A
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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

func (s AddKnowledgeBaseAdministratorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeBaseAdministratorResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) SetCreateTime(v int64) *AddKnowledgeBaseAdministratorResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) SetDisplayName(v string) *AddKnowledgeBaseAdministratorResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) SetKnowledgeBaseId(v string) *AddKnowledgeBaseAdministratorResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) SetUserId(v string) *AddKnowledgeBaseAdministratorResponseBodyData {
	s.UserId = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
