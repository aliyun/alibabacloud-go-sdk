// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpProtectionResponseBody
	GetRequestId() *string
}

type UpdateIpProtectionResponseBody struct {
	// Request ID
	//
	// example:
	//
	// B653A6FC-D1AD-5936-A262-F50994ED2574
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpProtectionResponseBody) SetRequestId(v string) *UpdateIpProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
