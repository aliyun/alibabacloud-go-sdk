// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorityTemplateViewList(v *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList) *ListAuthorityTemplateResponseBody
	GetAuthorityTemplateViewList() *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList
	SetErrorCode(v string) *ListAuthorityTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAuthorityTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListAuthorityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAuthorityTemplateResponseBody
	GetSuccess() *bool
	SetTid(v int64) *ListAuthorityTemplateResponseBody
	GetTid() *int64
	SetTotalCount(v int64) *ListAuthorityTemplateResponseBody
	GetTotalCount() *int64
}

type ListAuthorityTemplateResponseBody struct {
	AuthorityTemplateViewList *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList `json:"AuthorityTemplateViewList,omitempty" xml:"AuthorityTemplateViewList,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 31853A2B-DC9D-5B39-8492-D2AC8BCF550E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorityTemplateResponseBody) GetAuthorityTemplateViewList() *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList {
	return s.AuthorityTemplateViewList
}

func (s *ListAuthorityTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAuthorityTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAuthorityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAuthorityTemplateResponseBody) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorityTemplateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorityTemplateResponseBody) SetAuthorityTemplateViewList(v *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList) *ListAuthorityTemplateResponseBody {
	s.AuthorityTemplateViewList = v
	return s
}

func (s *ListAuthorityTemplateResponseBody) SetErrorCode(v string) *ListAuthorityTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAuthorityTemplateResponseBody) SetErrorMessage(v string) *ListAuthorityTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAuthorityTemplateResponseBody) SetRequestId(v string) *ListAuthorityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorityTemplateResponseBody) SetSuccess(v bool) *ListAuthorityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ListAuthorityTemplateResponseBody) SetTid(v int64) *ListAuthorityTemplateResponseBody {
	s.Tid = &v
	return s
}

func (s *ListAuthorityTemplateResponseBody) SetTotalCount(v int64) *ListAuthorityTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAuthorityTemplateResponseBodyAuthorityTemplateViewList struct {
	AuthorityTemplateView []*ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView `json:"AuthorityTemplateView,omitempty" xml:"AuthorityTemplateView,omitempty" type:"Repeated"`
}

func (s ListAuthorityTemplateResponseBodyAuthorityTemplateViewList) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityTemplateResponseBodyAuthorityTemplateViewList) GoString() string {
	return s.String()
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList) GetAuthorityTemplateView() []*ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView {
	return s.AuthorityTemplateView
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList) SetAuthorityTemplateView(v []*ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList {
	s.AuthorityTemplateView = v
	return s
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewList) Validate() error {
	return dara.Validate(s)
}

type ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView struct {
	// example:
	//
	// 2023-10-26 11:37:47
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 522****
	CreatorId   *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2592
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) GoString() string {
	return s.String()
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) GetDescription() *string {
	return s.Description
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) GetName() *string {
	return s.Name
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) SetCreateTime(v string) *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView {
	s.CreateTime = &v
	return s
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) SetCreatorId(v int64) *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView {
	s.CreatorId = &v
	return s
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) SetDescription(v string) *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView {
	s.Description = &v
	return s
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) SetName(v string) *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView {
	s.Name = &v
	return s
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) SetTemplateId(v int64) *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView {
	s.TemplateId = &v
	return s
}

func (s *ListAuthorityTemplateResponseBodyAuthorityTemplateViewListAuthorityTemplateView) Validate() error {
	return dara.Validate(s)
}
