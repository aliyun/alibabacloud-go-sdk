// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrivilegePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Privilege) *PrivilegePageResult
	GetData() []*Privilege
	SetRequestId(v string) *PrivilegePageResult
	GetRequestId() *string
	SetTotal(v int64) *PrivilegePageResult
	GetTotal() *int64
}

type PrivilegePageResult struct {
	Data      []*Privilege `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PrivilegePageResult) String() string {
	return dara.Prettify(s)
}

func (s PrivilegePageResult) GoString() string {
	return s.String()
}

func (s *PrivilegePageResult) GetData() []*Privilege {
	return s.Data
}

func (s *PrivilegePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PrivilegePageResult) GetTotal() *int64 {
	return s.Total
}

func (s *PrivilegePageResult) SetData(v []*Privilege) *PrivilegePageResult {
	s.Data = v
	return s
}

func (s *PrivilegePageResult) SetRequestId(v string) *PrivilegePageResult {
	s.RequestId = &v
	return s
}

func (s *PrivilegePageResult) SetTotal(v int64) *PrivilegePageResult {
	s.Total = &v
	return s
}

func (s *PrivilegePageResult) Validate() error {
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
