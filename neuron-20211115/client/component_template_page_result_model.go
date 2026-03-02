// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentTemplatePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ComponentTemplate) *ComponentTemplatePageResult
	GetList() []*ComponentTemplate
	SetRequestId(v string) *ComponentTemplatePageResult
	GetRequestId() *string
	SetTotal(v int32) *ComponentTemplatePageResult
	GetTotal() *int32
}

type ComponentTemplatePageResult struct {
	List []*ComponentTemplate `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ComponentTemplatePageResult) String() string {
	return dara.Prettify(s)
}

func (s ComponentTemplatePageResult) GoString() string {
	return s.String()
}

func (s *ComponentTemplatePageResult) GetList() []*ComponentTemplate {
	return s.List
}

func (s *ComponentTemplatePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ComponentTemplatePageResult) GetTotal() *int32 {
	return s.Total
}

func (s *ComponentTemplatePageResult) SetList(v []*ComponentTemplate) *ComponentTemplatePageResult {
	s.List = v
	return s
}

func (s *ComponentTemplatePageResult) SetRequestId(v string) *ComponentTemplatePageResult {
	s.RequestId = &v
	return s
}

func (s *ComponentTemplatePageResult) SetTotal(v int32) *ComponentTemplatePageResult {
	s.Total = &v
	return s
}

func (s *ComponentTemplatePageResult) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
