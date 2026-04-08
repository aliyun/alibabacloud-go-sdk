// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployTrafficControlTaskCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeployTrafficControlTaskCodeResponseBody
	GetRequestId() *string
}

type DeployTrafficControlTaskCodeResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployTrafficControlTaskCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployTrafficControlTaskCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeployTrafficControlTaskCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployTrafficControlTaskCodeResponseBody) SetRequestId(v string) *DeployTrafficControlTaskCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployTrafficControlTaskCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
