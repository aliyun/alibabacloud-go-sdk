// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnabledPrivilegesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeEnabledPrivilegesResponseBodyData) *DescribeEnabledPrivilegesResponseBody
	GetData() []*DescribeEnabledPrivilegesResponseBodyData
	SetRequestId(v string) *DescribeEnabledPrivilegesResponseBody
	GetRequestId() *string
}

type DescribeEnabledPrivilegesResponseBody struct {
	// The queried permission level and permissions.
	Data []*DescribeEnabledPrivilegesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 246F42E0-A475-15FF-96D2-8DC47FC2F289
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnabledPrivilegesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponseBody) GetData() []*DescribeEnabledPrivilegesResponseBodyData {
	return s.Data
}

func (s *DescribeEnabledPrivilegesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnabledPrivilegesResponseBody) SetData(v []*DescribeEnabledPrivilegesResponseBodyData) *DescribeEnabledPrivilegesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBody) SetRequestId(v string) *DescribeEnabledPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBody) Validate() error {
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

type DescribeEnabledPrivilegesResponseBodyData struct {
	// The description of the permission level.
	//
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The queried permissions.
	//
	// This parameter is required.
	Privileges []*DescribeEnabledPrivilegesResponseBodyDataPrivileges `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
	// The permission level.
	//
	// This parameter is required.
	//
	// example:
	//
	// Global
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribeEnabledPrivilegesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnabledPrivilegesResponseBodyData) GetPrivileges() []*DescribeEnabledPrivilegesResponseBodyDataPrivileges {
	return s.Privileges
}

func (s *DescribeEnabledPrivilegesResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *DescribeEnabledPrivilegesResponseBodyData) SetDescription(v string) *DescribeEnabledPrivilegesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyData) SetPrivileges(v []*DescribeEnabledPrivilegesResponseBodyDataPrivileges) *DescribeEnabledPrivilegesResponseBodyData {
	s.Privileges = v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyData) SetScope(v string) *DescribeEnabledPrivilegesResponseBodyData {
	s.Scope = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyData) Validate() error {
	if s.Privileges != nil {
		for _, item := range s.Privileges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnabledPrivilegesResponseBodyDataPrivileges struct {
	// The description of the permission.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// select
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeEnabledPrivilegesResponseBodyDataPrivileges) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponseBodyDataPrivileges) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponseBodyDataPrivileges) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnabledPrivilegesResponseBodyDataPrivileges) GetKey() *string {
	return s.Key
}

func (s *DescribeEnabledPrivilegesResponseBodyDataPrivileges) SetDescription(v string) *DescribeEnabledPrivilegesResponseBodyDataPrivileges {
	s.Description = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyDataPrivileges) SetKey(v string) *DescribeEnabledPrivilegesResponseBodyDataPrivileges {
	s.Key = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyDataPrivileges) Validate() error {
	return dara.Validate(s)
}
