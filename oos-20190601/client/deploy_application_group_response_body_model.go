// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeployApplicationGroupResponseBody
	GetRequestId() *string
}

type DeployApplicationGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8AF4800A-A316-589A-90C4-313B1FEEB084
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApplicationGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployApplicationGroupResponseBody) SetRequestId(v string) *DeployApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployApplicationGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
