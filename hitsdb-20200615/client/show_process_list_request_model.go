// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShowProcessListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ShowProcessListRequest
	GetInstanceId() *string
	SetRegionId(v string) *ShowProcessListRequest
	GetRegionId() *string
	SetSecurityToken(v string) *ShowProcessListRequest
	GetSecurityToken() *string
}

type ShowProcessListRequest struct {
	// This parameter is required.
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ShowProcessListRequest) String() string {
	return dara.Prettify(s)
}

func (s ShowProcessListRequest) GoString() string {
	return s.String()
}

func (s *ShowProcessListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ShowProcessListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ShowProcessListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ShowProcessListRequest) SetInstanceId(v string) *ShowProcessListRequest {
	s.InstanceId = &v
	return s
}

func (s *ShowProcessListRequest) SetRegionId(v string) *ShowProcessListRequest {
	s.RegionId = &v
	return s
}

func (s *ShowProcessListRequest) SetSecurityToken(v string) *ShowProcessListRequest {
	s.SecurityToken = &v
	return s
}

func (s *ShowProcessListRequest) Validate() error {
	return dara.Validate(s)
}
