// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityTemplateView(v *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) *CreateAuthorityTemplateResponseBody
	GetAuthorityTemplateView() *CreateAuthorityTemplateResponseBodyAuthorityTemplateView
	SetErrorCode(v string) *CreateAuthorityTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateAuthorityTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateAuthorityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAuthorityTemplateResponseBody
	GetSuccess() *bool
	SetTid(v int64) *CreateAuthorityTemplateResponseBody
	GetTid() *int64
}

type CreateAuthorityTemplateResponseBody struct {
	// The details of the permission template.
	AuthorityTemplateView *CreateAuthorityTemplateResponseBodyAuthorityTemplateView `json:"AuthorityTemplateView,omitempty" xml:"AuthorityTemplateView,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateAuthorityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthorityTemplateResponseBody) GetAuthorityTemplateView() *CreateAuthorityTemplateResponseBodyAuthorityTemplateView {
	return s.AuthorityTemplateView
}

func (s *CreateAuthorityTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAuthorityTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateAuthorityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAuthorityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAuthorityTemplateResponseBody) GetTid() *int64 {
	return s.Tid
}

func (s *CreateAuthorityTemplateResponseBody) SetAuthorityTemplateView(v *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) *CreateAuthorityTemplateResponseBody {
	s.AuthorityTemplateView = v
	return s
}

func (s *CreateAuthorityTemplateResponseBody) SetErrorCode(v string) *CreateAuthorityTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBody) SetErrorMessage(v string) *CreateAuthorityTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBody) SetRequestId(v string) *CreateAuthorityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBody) SetSuccess(v bool) *CreateAuthorityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBody) SetTid(v int64) *CreateAuthorityTemplateResponseBody {
	s.Tid = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBody) Validate() error {
	if s.AuthorityTemplateView != nil {
		if err := s.AuthorityTemplateView.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAuthorityTemplateResponseBodyAuthorityTemplateView struct {
	// The time when the permission template was created. The time is in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-01-11 14:17:33
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the permission template.
	//
	// example:
	//
	// 12***
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the permission template.
	//
	// example:
	//
	// This template is used for business testing.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission template.
	//
	// example:
	//
	// Test template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the permission template.
	//
	// example:
	//
	// 1563
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateAuthorityTemplateResponseBodyAuthorityTemplateView) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorityTemplateResponseBodyAuthorityTemplateView) GoString() string {
	return s.String()
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) GetDescription() *string {
	return s.Description
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) GetName() *string {
	return s.Name
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) SetCreateTime(v string) *CreateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.CreateTime = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) SetCreatorId(v int64) *CreateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.CreatorId = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) SetDescription(v string) *CreateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.Description = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) SetName(v string) *CreateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.Name = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) SetTemplateId(v int64) *CreateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.TemplateId = &v
	return s
}

func (s *CreateAuthorityTemplateResponseBodyAuthorityTemplateView) Validate() error {
	return dara.Validate(s)
}
