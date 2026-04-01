// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCDeploymentSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCDeploymentSetResponseBody
	GetRequestId() *string
}

type DeleteRCDeploymentSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRCDeploymentSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCDeploymentSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCDeploymentSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCDeploymentSetResponseBody) SetRequestId(v string) *DeleteRCDeploymentSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCDeploymentSetResponseBody) Validate() error {
	return dara.Validate(s)
}
