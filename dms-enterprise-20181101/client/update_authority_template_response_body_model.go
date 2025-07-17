// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityTemplateView(v *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) *UpdateAuthorityTemplateResponseBody
	GetAuthorityTemplateView() *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView
	SetErrorCode(v string) *UpdateAuthorityTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateAuthorityTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateAuthorityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAuthorityTemplateResponseBody
	GetSuccess() *bool
	SetTid(v int64) *UpdateAuthorityTemplateResponseBody
	GetTid() *int64
}

type UpdateAuthorityTemplateResponseBody struct {
	// The details of the permission template.
	AuthorityTemplateView *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView `json:"AuthorityTemplateView,omitempty" xml:"AuthorityTemplateView,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
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
	// The ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateAuthorityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorityTemplateResponseBody) GetAuthorityTemplateView() *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView {
	return s.AuthorityTemplateView
}

func (s *UpdateAuthorityTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAuthorityTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAuthorityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAuthorityTemplateResponseBody) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateAuthorityTemplateResponseBody) SetAuthorityTemplateView(v *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) *UpdateAuthorityTemplateResponseBody {
	s.AuthorityTemplateView = v
	return s
}

func (s *UpdateAuthorityTemplateResponseBody) SetErrorCode(v string) *UpdateAuthorityTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBody) SetErrorMessage(v string) *UpdateAuthorityTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBody) SetRequestId(v string) *UpdateAuthorityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBody) SetSuccess(v bool) *UpdateAuthorityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBody) SetTid(v int64) *UpdateAuthorityTemplateResponseBody {
	s.Tid = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateAuthorityTemplateResponseBodyAuthorityTemplateView struct {
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

func (s UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) GoString() string {
	return s.String()
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) GetDescription() *string {
	return s.Description
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) GetName() *string {
	return s.Name
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) SetCreatorId(v int64) *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.CreatorId = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) SetDescription(v string) *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.Description = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) SetName(v string) *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.Name = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) SetTemplateId(v int64) *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView {
	s.TemplateId = &v
	return s
}

func (s *UpdateAuthorityTemplateResponseBodyAuthorityTemplateView) Validate() error {
	return dara.Validate(s)
}
