// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeploymentJobResponseBody
	GetRequestId() *string
}

type DeleteDeploymentJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeploymentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeploymentJobResponseBody) SetRequestId(v string) *DeleteDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentJobResponseBody) Validate() error {
	return dara.Validate(s)
}
