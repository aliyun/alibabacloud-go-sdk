// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceGroupPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PdpServiceGroup) *PdpServiceGroupPageResult
	GetData() []*PdpServiceGroup
	SetRequestId(v string) *PdpServiceGroupPageResult
	GetRequestId() *string
	SetTotal(v int32) *PdpServiceGroupPageResult
	GetTotal() *int32
}

type PdpServiceGroupPageResult struct {
	Data []*PdpServiceGroup `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpServiceGroupPageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceGroupPageResult) GoString() string {
	return s.String()
}

func (s *PdpServiceGroupPageResult) GetData() []*PdpServiceGroup {
	return s.Data
}

func (s *PdpServiceGroupPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpServiceGroupPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *PdpServiceGroupPageResult) SetData(v []*PdpServiceGroup) *PdpServiceGroupPageResult {
	s.Data = v
	return s
}

func (s *PdpServiceGroupPageResult) SetRequestId(v string) *PdpServiceGroupPageResult {
	s.RequestId = &v
	return s
}

func (s *PdpServiceGroupPageResult) SetTotal(v int32) *PdpServiceGroupPageResult {
	s.Total = &v
	return s
}

func (s *PdpServiceGroupPageResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
