// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadPublicKeyResponseBody
	GetRequestId() *string
}

type UploadPublicKeyResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UploadPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadPublicKeyResponseBody) SetRequestId(v string) *UploadPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
