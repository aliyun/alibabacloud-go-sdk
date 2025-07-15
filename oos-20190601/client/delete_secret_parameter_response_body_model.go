// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecretParameterResponseBody
	GetRequestId() *string
}

type DeleteSecretParameterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0D02BDF-77F6-49F2-95C9-8E87121D1944
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecretParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretParameterResponseBody) SetRequestId(v string) *DeleteSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
