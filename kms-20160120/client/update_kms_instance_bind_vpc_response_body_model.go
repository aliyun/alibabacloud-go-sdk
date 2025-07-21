// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKmsInstanceBindVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKmsInstanceBindVpcResponseBody
	GetRequestId() *string
}

type UpdateKmsInstanceBindVpcResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// d3eca5c8-a856-4347-8eb6-e1898c3fda2e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKmsInstanceBindVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKmsInstanceBindVpcResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKmsInstanceBindVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKmsInstanceBindVpcResponseBody) SetRequestId(v string) *UpdateKmsInstanceBindVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKmsInstanceBindVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
