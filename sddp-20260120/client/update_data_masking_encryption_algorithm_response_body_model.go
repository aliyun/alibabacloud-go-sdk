// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingEncryptionAlgorithmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataMaskingEncryptionAlgorithmResponseBody
	GetRequestId() *string
}

type UpdateDataMaskingEncryptionAlgorithmResponseBody struct {
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataMaskingEncryptionAlgorithmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingEncryptionAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponseBody) SetRequestId(v string) *UpdateDataMaskingEncryptionAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmResponseBody) Validate() error {
	return dara.Validate(s)
}
