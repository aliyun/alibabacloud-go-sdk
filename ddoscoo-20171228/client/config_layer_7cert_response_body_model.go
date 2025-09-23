// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7CertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigLayer7CertResponseBody
	GetRequestId() *string
}

type ConfigLayer7CertResponseBody struct {
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer7CertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7CertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigLayer7CertResponseBody) SetRequestId(v string) *ConfigLayer7CertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigLayer7CertResponseBody) Validate() error {
	return dara.Validate(s)
}
