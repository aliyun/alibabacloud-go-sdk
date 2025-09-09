// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeployServiceInstanceResponseBody
	GetRequestId() *string
}

type DeployServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B8A6AEA6-0D8F-589A-A7FF-B44FD842456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployServiceInstanceResponseBody) SetRequestId(v string) *DeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
