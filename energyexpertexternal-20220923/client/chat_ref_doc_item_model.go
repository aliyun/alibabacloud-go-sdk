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
	// Related document information.
	DocInfo *ChatRefDocInfo `json:"docInfo,omitempty" xml:"docInfo,omitempty"`
	// Document name. If the original file is of types like doc, it will be converted to PDF for easier document location. This file should be used for document location information.
	//
	// example:
	//
	// a.pdf
	DocName *string `json:"docName,omitempty" xml:"docName,omitempty"`
	// Document link.
	//
	// example:
	//
	// https://carbon-aidoc.oss-accelerate.aliyuncs.com/jobs/42dbc7-3a9e-4e18-8939-3fd2d247bd3c/document/97178d94c75941d3b932883c810c5e.pdf
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// Original document name.
	//
	// example:
	//
	// a.doc
	OriginDocName *string `json:"originDocName,omitempty" xml:"originDocName,omitempty"`
	// Original document link.
	//
	// example:
	//
	// https://carbon-aidoc.oss-accelerate.aliyuncs.com/jobs/42dbc7-3a9e-4e18-8939-3fd2d247bd3c/document/97178d94c75941d3b932883c810c5e.doc
	OriginDocUrl *string `json:"originDocUrl,omitempty" xml:"originDocUrl,omitempty"`
	// Page number information.
	PageNum []*ChatDocumentPageNum `json:"pageNum,omitempty" xml:"pageNum,omitempty" type:"Repeated"`
	// Return file type: \\"doc\\" for document files, \\"web\\" for internet information.
	//
	// example:
	//
	// doc
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
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
	if s.DocInfo != nil {
		if err := s.DocInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PageNum != nil {
		for _, item := range s.PageNum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
