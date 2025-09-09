// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigInstanceSecurityGroupsResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *ConfigInstanceSecurityGroupsResponseBody
	GetRequestId() *string
}

type ConfigInstanceSecurityGroupsResponseBody struct {
	// The ID of the bastion host for which security groups were configured.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0ECCC399-4D35-48A7-8379-5C6180E66235
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigInstanceSecurityGroupsResponseBody) SetInstanceId(v string) *ConfigInstanceSecurityGroupsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponseBody) SetRequestId(v string) *ConfigInstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
