// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeveloperPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Developer) *DeveloperPageResult
	GetData() []*Developer
	SetRequestId(v string) *DeveloperPageResult
	GetRequestId() *string
	SetTotal(v int64) *DeveloperPageResult
	GetTotal() *int64
}

type DeveloperPageResult struct {
	Data      []*Developer `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s DeveloperPageResult) String() string {
	return dara.Prettify(s)
}

func (s DeveloperPageResult) GoString() string {
	return s.String()
}

func (s *DeveloperPageResult) GetData() []*Developer {
	return s.Data
}

func (s *DeveloperPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeveloperPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *DeveloperPageResult) SetData(v []*Developer) *DeveloperPageResult {
	s.Data = v
	return s
}

func (s *DeveloperPageResult) SetRequestId(v string) *DeveloperPageResult {
	s.RequestId = &v
	return s
}

func (s *DeveloperPageResult) SetTotal(v int64) *DeveloperPageResult {
	s.Total = &v
	return s
}

func (s *DeveloperPageResult) Validate() error {
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
