// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundColors(v [][]*string) *UpdateRangeRequest
	GetBackgroundColors() [][]*string
	SetHyperlinks(v [][]*UpdateRangeRequestHyperlinks) *UpdateRangeRequest
	GetHyperlinks() [][]*UpdateRangeRequestHyperlinks
	SetNumberFormat(v string) *UpdateRangeRequest
	GetNumberFormat() *string
	SetRangeAddress(v string) *UpdateRangeRequest
	GetRangeAddress() *string
	SetSheetId(v string) *UpdateRangeRequest
	GetSheetId() *string
	SetTenantContext(v *UpdateRangeRequestTenantContext) *UpdateRangeRequest
	GetTenantContext() *UpdateRangeRequestTenantContext
	SetValues(v [][]*string) *UpdateRangeRequest
	GetValues() [][]*string
	SetWorkbookId(v string) *UpdateRangeRequest
	GetWorkbookId() *string
}

type UpdateRangeRequest struct {
	// example:
	//
	// [["#ff0000","#ff0000","#ff0000"]]
	BackgroundColors [][]*string `json:"BackgroundColors,omitempty" xml:"BackgroundColors,omitempty" type:"Repeated"`
	// example:
	//
	// [["type": "path","link": "https://www.dingtalk.com","text": "test"]]
	Hyperlinks [][]*UpdateRangeRequestHyperlinks `json:"Hyperlinks,omitempty" xml:"Hyperlinks,omitempty" type:"Repeated"`
	// example:
	//
	// General
	NumberFormat *string `json:"NumberFormat,omitempty" xml:"NumberFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A3:C3
	RangeAddress *string `json:"RangeAddress,omitempty" xml:"RangeAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                          `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *UpdateRangeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// [["1","2","3"]]
	Values [][]*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s UpdateRangeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRangeRequest) GetBackgroundColors() [][]*string {
	return s.BackgroundColors
}

func (s *UpdateRangeRequest) GetHyperlinks() [][]*UpdateRangeRequestHyperlinks {
	return s.Hyperlinks
}

func (s *UpdateRangeRequest) GetNumberFormat() *string {
	return s.NumberFormat
}

func (s *UpdateRangeRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *UpdateRangeRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *UpdateRangeRequest) GetTenantContext() *UpdateRangeRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateRangeRequest) GetValues() [][]*string {
	return s.Values
}

func (s *UpdateRangeRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *UpdateRangeRequest) SetBackgroundColors(v [][]*string) *UpdateRangeRequest {
	s.BackgroundColors = v
	return s
}

func (s *UpdateRangeRequest) SetHyperlinks(v [][]*UpdateRangeRequestHyperlinks) *UpdateRangeRequest {
	s.Hyperlinks = v
	return s
}

func (s *UpdateRangeRequest) SetNumberFormat(v string) *UpdateRangeRequest {
	s.NumberFormat = &v
	return s
}

func (s *UpdateRangeRequest) SetRangeAddress(v string) *UpdateRangeRequest {
	s.RangeAddress = &v
	return s
}

func (s *UpdateRangeRequest) SetSheetId(v string) *UpdateRangeRequest {
	s.SheetId = &v
	return s
}

func (s *UpdateRangeRequest) SetTenantContext(v *UpdateRangeRequestTenantContext) *UpdateRangeRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateRangeRequest) SetValues(v [][]*string) *UpdateRangeRequest {
	s.Values = v
	return s
}

func (s *UpdateRangeRequest) SetWorkbookId(v string) *UpdateRangeRequest {
	s.WorkbookId = &v
	return s
}

func (s *UpdateRangeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRangeRequestHyperlinks struct {
	// example:
	//
	// path
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// example:
	//
	// test
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s UpdateRangeRequestHyperlinks) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeRequestHyperlinks) GoString() string {
	return s.String()
}

func (s *UpdateRangeRequestHyperlinks) GetType() *string {
	return s.Type
}

func (s *UpdateRangeRequestHyperlinks) GetLink() *string {
	return s.Link
}

func (s *UpdateRangeRequestHyperlinks) GetText() *string {
	return s.Text
}

func (s *UpdateRangeRequestHyperlinks) SetType(v string) *UpdateRangeRequestHyperlinks {
	s.Type = &v
	return s
}

func (s *UpdateRangeRequestHyperlinks) SetLink(v string) *UpdateRangeRequestHyperlinks {
	s.Link = &v
	return s
}

func (s *UpdateRangeRequestHyperlinks) SetText(v string) *UpdateRangeRequestHyperlinks {
	s.Text = &v
	return s
}

func (s *UpdateRangeRequestHyperlinks) Validate() error {
	return dara.Validate(s)
}

type UpdateRangeRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateRangeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateRangeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateRangeRequestTenantContext) SetTenantId(v string) *UpdateRangeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateRangeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
