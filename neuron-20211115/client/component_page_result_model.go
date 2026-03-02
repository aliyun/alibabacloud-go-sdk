// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*PdpComponent) *ComponentPageResult
	GetList() []*PdpComponent
	SetRequestId(v string) *ComponentPageResult
	GetRequestId() *string
	SetTotal(v int32) *ComponentPageResult
	GetTotal() *int32
}

type ComponentPageResult struct {
	List []*PdpComponent `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ComponentPageResult) String() string {
	return dara.Prettify(s)
}

func (s ComponentPageResult) GoString() string {
	return s.String()
}

func (s *ComponentPageResult) GetList() []*PdpComponent {
	return s.List
}

func (s *ComponentPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ComponentPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *ComponentPageResult) SetList(v []*PdpComponent) *ComponentPageResult {
	s.List = v
	return s
}

func (s *ComponentPageResult) SetRequestId(v string) *ComponentPageResult {
	s.RequestId = &v
	return s
}

func (s *ComponentPageResult) SetTotal(v int32) *ComponentPageResult {
	s.Total = &v
	return s
}

func (s *ComponentPageResult) Validate() error {
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
