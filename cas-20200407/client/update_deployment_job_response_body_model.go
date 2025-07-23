// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDeploymentJobResponseBody
	GetRequestId() *string
}

type UpdateDeploymentJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeploymentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeploymentJobResponseBody) SetRequestId(v string) *UpdateDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentJobResponseBody) Validate() error {
	return dara.Validate(s)
}
