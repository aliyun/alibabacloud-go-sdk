// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyCertificateResponseBody
	GetRequestId() *string
}

type ApplyCertificateResponseBody struct {
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyCertificateResponseBody) SetRequestId(v string) *ApplyCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
