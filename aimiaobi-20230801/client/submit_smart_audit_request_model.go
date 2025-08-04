// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrlList(v []*SubmitSmartAuditRequestImageUrlList) *SubmitSmartAuditRequest
	GetImageUrlList() []*SubmitSmartAuditRequestImageUrlList
	SetSubCodes(v []*string) *SubmitSmartAuditRequest
	GetSubCodes() []*string
	SetText(v string) *SubmitSmartAuditRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitSmartAuditRequest
	GetWorkspaceId() *string
	SetImageUrls(v []*SubmitSmartAuditRequestImageUrls) *SubmitSmartAuditRequest
	GetImageUrls() []*SubmitSmartAuditRequestImageUrls
}

type SubmitSmartAuditRequest struct {
	ImageUrlList []*SubmitSmartAuditRequestImageUrlList `json:"ImageUrlList,omitempty" xml:"ImageUrlList,omitempty" type:"Repeated"`
	SubCodes     []*string                              `json:"SubCodes,omitempty" xml:"SubCodes,omitempty" type:"Repeated"`
	Text         *string                                `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// xxxx
	WorkspaceId *string                             `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ImageUrls   []*SubmitSmartAuditRequestImageUrls `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
}

func (s SubmitSmartAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditRequest) GetImageUrlList() []*SubmitSmartAuditRequestImageUrlList {
	return s.ImageUrlList
}

func (s *SubmitSmartAuditRequest) GetSubCodes() []*string {
	return s.SubCodes
}

func (s *SubmitSmartAuditRequest) GetText() *string {
	return s.Text
}

func (s *SubmitSmartAuditRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitSmartAuditRequest) GetImageUrls() []*SubmitSmartAuditRequestImageUrls {
	return s.ImageUrls
}

func (s *SubmitSmartAuditRequest) SetImageUrlList(v []*SubmitSmartAuditRequestImageUrlList) *SubmitSmartAuditRequest {
	s.ImageUrlList = v
	return s
}

func (s *SubmitSmartAuditRequest) SetSubCodes(v []*string) *SubmitSmartAuditRequest {
	s.SubCodes = v
	return s
}

func (s *SubmitSmartAuditRequest) SetText(v string) *SubmitSmartAuditRequest {
	s.Text = &v
	return s
}

func (s *SubmitSmartAuditRequest) SetWorkspaceId(v string) *SubmitSmartAuditRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitSmartAuditRequest) SetImageUrls(v []*SubmitSmartAuditRequestImageUrls) *SubmitSmartAuditRequest {
	s.ImageUrls = v
	return s
}

func (s *SubmitSmartAuditRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartAuditRequestImageUrlList struct {
	Id  *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitSmartAuditRequestImageUrlList) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditRequestImageUrlList) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditRequestImageUrlList) GetId() *string {
	return s.Id
}

func (s *SubmitSmartAuditRequestImageUrlList) GetUrl() *string {
	return s.Url
}

func (s *SubmitSmartAuditRequestImageUrlList) SetId(v string) *SubmitSmartAuditRequestImageUrlList {
	s.Id = &v
	return s
}

func (s *SubmitSmartAuditRequestImageUrlList) SetUrl(v string) *SubmitSmartAuditRequestImageUrlList {
	s.Url = &v
	return s
}

func (s *SubmitSmartAuditRequestImageUrlList) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartAuditRequestImageUrls struct {
	Id  *string `json:"id,omitempty" xml:"id,omitempty"`
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitSmartAuditRequestImageUrls) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditRequestImageUrls) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditRequestImageUrls) GetId() *string {
	return s.Id
}

func (s *SubmitSmartAuditRequestImageUrls) GetUrl() *string {
	return s.Url
}

func (s *SubmitSmartAuditRequestImageUrls) SetId(v string) *SubmitSmartAuditRequestImageUrls {
	s.Id = &v
	return s
}

func (s *SubmitSmartAuditRequestImageUrls) SetUrl(v string) *SubmitSmartAuditRequestImageUrls {
	s.Url = &v
	return s
}

func (s *SubmitSmartAuditRequestImageUrls) Validate() error {
	return dara.Validate(s)
}
