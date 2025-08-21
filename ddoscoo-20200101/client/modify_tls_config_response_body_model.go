// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTlsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTlsConfigResponseBody
	GetRequestId() *string
}

type ModifyTlsConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTlsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTlsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTlsConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTlsConfigResponseBody) SetRequestId(v string) *ModifyTlsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTlsConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
