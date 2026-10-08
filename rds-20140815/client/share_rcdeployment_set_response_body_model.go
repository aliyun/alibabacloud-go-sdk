// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareRCDeploymentSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ShareRCDeploymentSetResponseBody
	GetRequestId() *string
}

type ShareRCDeploymentSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ShareRCDeploymentSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShareRCDeploymentSetResponseBody) GoString() string {
	return s.String()
}

func (s *ShareRCDeploymentSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShareRCDeploymentSetResponseBody) SetRequestId(v string) *ShareRCDeploymentSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShareRCDeploymentSetResponseBody) Validate() error {
	return dara.Validate(s)
}
