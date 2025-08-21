// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *CheckCertificateRequest
	GetAppKey() *int64
}

type CheckCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s CheckCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCertificateRequest) GoString() string {
	return s.String()
}

func (s *CheckCertificateRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *CheckCertificateRequest) SetAppKey(v int64) *CheckCertificateRequest {
	s.AppKey = &v
	return s
}

func (s *CheckCertificateRequest) Validate() error {
	return dara.Validate(s)
}
