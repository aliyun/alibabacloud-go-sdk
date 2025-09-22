// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRefDocItem interface {
	dara.Model
	String() string
	GoString() string
	SetDocInfo(v *ChatRefDocInfo) *ChatRefDocItem
	GetDocInfo() *ChatRefDocInfo
	SetDocName(v string) *ChatRefDocItem
	GetDocName() *string
	SetDocUrl(v string) *ChatRefDocItem
	GetDocUrl() *string
	SetOriginDocName(v string) *ChatRefDocItem
	GetOriginDocName() *string
	SetOriginDocUrl(v string) *ChatRefDocItem
	GetOriginDocUrl() *string
	SetPageNum(v []*ChatDocumentPageNum) *ChatRefDocItem
	GetPageNum() []*ChatDocumentPageNum
	SetSourceType(v string) *ChatRefDocItem
	GetSourceType() *string
}

type ChatRefDocItem struct {
	DocInfo       *ChatRefDocInfo        `json:"docInfo,omitempty" xml:"docInfo,omitempty"`
	DocName       *string                `json:"docName,omitempty" xml:"docName,omitempty"`
	DocUrl        *string                `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	OriginDocName *string                `json:"originDocName,omitempty" xml:"originDocName,omitempty"`
	OriginDocUrl  *string                `json:"originDocUrl,omitempty" xml:"originDocUrl,omitempty"`
	PageNum       []*ChatDocumentPageNum `json:"pageNum,omitempty" xml:"pageNum,omitempty" type:"Repeated"`
	SourceType    *string                `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ChatRefDocItem) String() string {
	return dara.Prettify(s)
}

func (s ChatRefDocItem) GoString() string {
	return s.String()
}

func (s *ChatRefDocItem) GetDocInfo() *ChatRefDocInfo {
	return s.DocInfo
}

func (s *ChatRefDocItem) GetDocName() *string {
	return s.DocName
}

func (s *ChatRefDocItem) GetDocUrl() *string {
	return s.DocUrl
}

func (s *ChatRefDocItem) GetOriginDocName() *string {
	return s.OriginDocName
}

func (s *ChatRefDocItem) GetOriginDocUrl() *string {
	return s.OriginDocUrl
}

func (s *ChatRefDocItem) GetPageNum() []*ChatDocumentPageNum {
	return s.PageNum
}

func (s *ChatRefDocItem) GetSourceType() *string {
	return s.SourceType
}

func (s *ChatRefDocItem) SetDocInfo(v *ChatRefDocInfo) *ChatRefDocItem {
	s.DocInfo = v
	return s
}

func (s *ChatRefDocItem) SetDocName(v string) *ChatRefDocItem {
	s.DocName = &v
	return s
}

func (s *ChatRefDocItem) SetDocUrl(v string) *ChatRefDocItem {
	s.DocUrl = &v
	return s
}

func (s *ChatRefDocItem) SetOriginDocName(v string) *ChatRefDocItem {
	s.OriginDocName = &v
	return s
}

func (s *ChatRefDocItem) SetOriginDocUrl(v string) *ChatRefDocItem {
	s.OriginDocUrl = &v
	return s
}

func (s *ChatRefDocItem) SetPageNum(v []*ChatDocumentPageNum) *ChatRefDocItem {
	s.PageNum = v
	return s
}

func (s *ChatRefDocItem) SetSourceType(v string) *ChatRefDocItem {
	s.SourceType = &v
	return s
}

func (s *ChatRefDocItem) Validate() error {
	return dara.Validate(s)
}
