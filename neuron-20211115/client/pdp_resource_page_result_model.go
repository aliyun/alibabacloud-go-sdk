// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpResourcePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*PdpResource) *PdpResourcePageResult
	GetList() []*PdpResource
	SetRequestId(v string) *PdpResourcePageResult
	GetRequestId() *string
	SetTotal(v int32) *PdpResourcePageResult
	GetTotal() *int32
}

type PdpResourcePageResult struct {
	List []*PdpResource `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpResourcePageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpResourcePageResult) GoString() string {
	return s.String()
}

func (s *PdpResourcePageResult) GetList() []*PdpResource {
	return s.List
}

func (s *PdpResourcePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpResourcePageResult) GetTotal() *int32 {
	return s.Total
}

func (s *PdpResourcePageResult) SetList(v []*PdpResource) *PdpResourcePageResult {
	s.List = v
	return s
}

func (s *PdpResourcePageResult) SetRequestId(v string) *PdpResourcePageResult {
	s.RequestId = &v
	return s
}

func (s *PdpResourcePageResult) SetTotal(v int32) *PdpResourcePageResult {
	s.Total = &v
	return s
}

func (s *PdpResourcePageResult) Validate() error {
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
