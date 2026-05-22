// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCustomHostnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnameId(v int64) *VerifyCustomHostnameRequest
	GetHostnameId() *int64
}

type VerifyCustomHostnameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	HostnameId *int64 `json:"HostnameId,omitempty" xml:"HostnameId,omitempty"`
}

func (s VerifyCustomHostnameRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCustomHostnameRequest) GoString() string {
	return s.String()
}

func (s *VerifyCustomHostnameRequest) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *VerifyCustomHostnameRequest) SetHostnameId(v int64) *VerifyCustomHostnameRequest {
	s.HostnameId = &v
	return s
}

func (s *VerifyCustomHostnameRequest) Validate() error {
	return dara.Validate(s)
}
