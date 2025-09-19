// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstallCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInstallCodeResponseBody
	GetRequestId() *string
}

type DeleteInstallCodeResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0DC1F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstallCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstallCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstallCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstallCodeResponseBody) SetRequestId(v string) *DeleteInstallCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstallCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
