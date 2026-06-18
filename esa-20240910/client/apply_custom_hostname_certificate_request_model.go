// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCustomHostnameCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnameId(v int64) *ApplyCustomHostnameCertificateRequest
	GetHostnameId() *int64
}

type ApplyCustomHostnameCertificateRequest struct {
	// The ID of the SaaS domain name. Call the [ListCustomHostnames](https://help.aliyun.com/document_detail/3018667.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	HostnameId *int64 `json:"HostnameId,omitempty" xml:"HostnameId,omitempty"`
}

func (s ApplyCustomHostnameCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCustomHostnameCertificateRequest) GoString() string {
	return s.String()
}

func (s *ApplyCustomHostnameCertificateRequest) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *ApplyCustomHostnameCertificateRequest) SetHostnameId(v int64) *ApplyCustomHostnameCertificateRequest {
	s.HostnameId = &v
	return s
}

func (s *ApplyCustomHostnameCertificateRequest) Validate() error {
	return dara.Validate(s)
}
