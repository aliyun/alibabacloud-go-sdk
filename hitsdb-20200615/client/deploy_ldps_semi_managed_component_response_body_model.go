// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployLdpsSemiManagedComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeployLdpsSemiManagedComponentResponseBody
	GetRequestId() *string
}

type DeployLdpsSemiManagedComponentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployLdpsSemiManagedComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployLdpsSemiManagedComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DeployLdpsSemiManagedComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployLdpsSemiManagedComponentResponseBody) SetRequestId(v string) *DeployLdpsSemiManagedComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
