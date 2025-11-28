// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPrivateRAGServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeployPrivateRAGServiceResponseBody
	GetRequestId() *string
}

type DeployPrivateRAGServiceResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployPrivateRAGServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployPrivateRAGServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployPrivateRAGServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployPrivateRAGServiceResponseBody) SetRequestId(v string) *DeployPrivateRAGServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployPrivateRAGServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
