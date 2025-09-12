// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetInstanceSecurityGroupsResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *GetInstanceSecurityGroupsResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetInstanceSecurityGroupsResponseBody
	GetRequestId() *string
	SetSecurityGroups(v []*string) *GetInstanceSecurityGroupsResponseBody
	GetSecurityGroups() []*string
}

type GetInstanceSecurityGroupsResponseBody struct {
	AccessDeniedDetail *string   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	InstanceId         *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroups     []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
}

func (s GetInstanceSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSecurityGroupsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetInstanceSecurityGroupsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceSecurityGroupsResponseBody) GetSecurityGroups() []*string {
	return s.SecurityGroups
}

func (s *GetInstanceSecurityGroupsResponseBody) SetAccessDeniedDetail(v string) *GetInstanceSecurityGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponseBody) SetInstanceId(v string) *GetInstanceSecurityGroupsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponseBody) SetRequestId(v string) *GetInstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponseBody) SetSecurityGroups(v []*string) *GetInstanceSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *GetInstanceSecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
