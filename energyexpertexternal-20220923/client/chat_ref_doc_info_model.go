// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRefDocInfo interface {
	dara.Model
	String() string
	GoString() string
	SetPageListInfo(v []*ChatRefDocPageInfo) *ChatRefDocInfo
	GetPageListInfo() []*ChatRefDocPageInfo
	SetPages(v int64) *ChatRefDocInfo
	GetPages() *int64
}

type ChatRefDocInfo struct {
	PageListInfo []*ChatRefDocPageInfo `json:"pageListInfo,omitempty" xml:"pageListInfo,omitempty" type:"Repeated"`
	Pages        *int64                `json:"pages,omitempty" xml:"pages,omitempty"`
}

func (s ChatRefDocInfo) String() string {
	return dara.Prettify(s)
}

func (s ChatRefDocInfo) GoString() string {
	return s.String()
}

func (s *ChatRefDocInfo) GetPageListInfo() []*ChatRefDocPageInfo {
	return s.PageListInfo
}

func (s *ChatRefDocInfo) GetPages() *int64 {
	return s.Pages
}

func (s *ChatRefDocInfo) SetPageListInfo(v []*ChatRefDocPageInfo) *ChatRefDocInfo {
	s.PageListInfo = v
	return s
}

func (s *ChatRefDocInfo) SetPages(v int64) *ChatRefDocInfo {
	s.Pages = &v
	return s
}

func (s *ChatRefDocInfo) Validate() error {
	return dara.Validate(s)
}
