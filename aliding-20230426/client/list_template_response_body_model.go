// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *ListTemplateResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *ListTemplateResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTemplateResponseBody
	GetRequestId() *string
	SetTemplateList(v []*ListTemplateResponseBodyTemplateList) *ListTemplateResponseBody
	GetTemplateList() []*ListTemplateResponseBodyTemplateList
	SetVendorRequestId(v string) *ListTemplateResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListTemplateResponseBody
	GetVendorType() *string
}

type ListTemplateResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// []
	TemplateList []*ListTemplateResponseBodyTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListTemplateResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateResponseBody) GetTemplateList() []*ListTemplateResponseBodyTemplateList {
	return s.TemplateList
}

func (s *ListTemplateResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListTemplateResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListTemplateResponseBody) SetHasMore(v bool) *ListTemplateResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListTemplateResponseBody) SetNextToken(v string) *ListTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateResponseBody) SetRequestId(v string) *ListTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateResponseBody) SetTemplateList(v []*ListTemplateResponseBodyTemplateList) *ListTemplateResponseBody {
	s.TemplateList = v
	return s
}

func (s *ListTemplateResponseBody) SetVendorRequestId(v string) *ListTemplateResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListTemplateResponseBody) SetVendorType(v string) *ListTemplateResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListTemplateResponseBody) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTemplateResponseBodyTemplateList struct {
	// example:
	//
	// URL
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// example:
	//
	// 1596506100000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// WORKBOOK
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// user_template
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1596506100000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// workspaceId
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTemplateResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBodyTemplateList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListTemplateResponseBodyTemplateList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTemplateResponseBodyTemplateList) GetDocType() *string {
	return s.DocType
}

func (s *ListTemplateResponseBodyTemplateList) GetId() *string {
	return s.Id
}

func (s *ListTemplateResponseBodyTemplateList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplateResponseBodyTemplateList) GetTitle() *string {
	return s.Title
}

func (s *ListTemplateResponseBodyTemplateList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListTemplateResponseBodyTemplateList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTemplateResponseBodyTemplateList) SetCoverUrl(v string) *ListTemplateResponseBodyTemplateList {
	s.CoverUrl = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetCreateTime(v int64) *ListTemplateResponseBodyTemplateList {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetDocType(v string) *ListTemplateResponseBodyTemplateList {
	s.DocType = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetId(v string) *ListTemplateResponseBodyTemplateList {
	s.Id = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetTemplateType(v string) *ListTemplateResponseBodyTemplateList {
	s.TemplateType = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetTitle(v string) *ListTemplateResponseBodyTemplateList {
	s.Title = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetUpdateTime(v int64) *ListTemplateResponseBodyTemplateList {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetWorkspaceId(v string) *ListTemplateResponseBodyTemplateList {
	s.WorkspaceId = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) Validate() error {
	return dara.Validate(s)
}
