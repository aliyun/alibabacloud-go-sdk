// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceGroupInfo interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ServiceGroupInfo
	GetId() *int64
	SetName(v string) *ServiceGroupInfo
	GetName() *string
}

type ServiceGroupInfo struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 分组1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ServiceGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s ServiceGroupInfo) GoString() string {
	return s.String()
}

func (s *ServiceGroupInfo) GetId() *int64 {
	return s.Id
}

func (s *ServiceGroupInfo) GetName() *string {
	return s.Name
}

func (s *ServiceGroupInfo) SetId(v int64) *ServiceGroupInfo {
	s.Id = &v
	return s
}

func (s *ServiceGroupInfo) SetName(v string) *ServiceGroupInfo {
	s.Name = &v
	return s
}

func (s *ServiceGroupInfo) Validate() error {
	return dara.Validate(s)
}
