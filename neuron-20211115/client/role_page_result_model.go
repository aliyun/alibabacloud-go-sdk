// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRolePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Role) *RolePageResult
	GetData() []*Role
	SetRequestId(v string) *RolePageResult
	GetRequestId() *string
	SetTotal(v int64) *RolePageResult
	GetTotal() *int64
}

type RolePageResult struct {
	Data      []*Role `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s RolePageResult) String() string {
	return dara.Prettify(s)
}

func (s RolePageResult) GoString() string {
	return s.String()
}

func (s *RolePageResult) GetData() []*Role {
	return s.Data
}

func (s *RolePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *RolePageResult) GetTotal() *int64 {
	return s.Total
}

func (s *RolePageResult) SetData(v []*Role) *RolePageResult {
	s.Data = v
	return s
}

func (s *RolePageResult) SetRequestId(v string) *RolePageResult {
	s.RequestId = &v
	return s
}

func (s *RolePageResult) SetTotal(v int64) *RolePageResult {
	s.Total = &v
	return s
}

func (s *RolePageResult) Validate() error {
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
