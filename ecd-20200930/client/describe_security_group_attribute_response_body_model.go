// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressPermissions(v []*Permission) *DescribeSecurityGroupAttributeResponseBody
	GetEgressPermissions() []*Permission
	SetIngressPermissions(v []*Permission) *DescribeSecurityGroupAttributeResponseBody
	GetIngressPermissions() []*Permission
	SetRequestId(v string) *DescribeSecurityGroupAttributeResponseBody
	GetRequestId() *string
}

type DescribeSecurityGroupAttributeResponseBody struct {
	EgressPermissions  []*Permission `json:"EgressPermissions,omitempty" xml:"EgressPermissions,omitempty" type:"Repeated"`
	IngressPermissions []*Permission `json:"IngressPermissions,omitempty" xml:"IngressPermissions,omitempty" type:"Repeated"`
	RequestId          *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetEgressPermissions() []*Permission {
	return s.EgressPermissions
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetIngressPermissions() []*Permission {
	return s.IngressPermissions
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetEgressPermissions(v []*Permission) *DescribeSecurityGroupAttributeResponseBody {
	s.EgressPermissions = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetIngressPermissions(v []*Permission) *DescribeSecurityGroupAttributeResponseBody {
	s.IngressPermissions = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetRequestId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) Validate() error {
	if s.EgressPermissions != nil {
		for _, item := range s.EgressPermissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IngressPermissions != nil {
		for _, item := range s.IngressPermissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
