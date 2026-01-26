// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomHostnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnameId(v int64) *GetCustomHostnameRequest
	GetHostnameId() *int64
}

type GetCustomHostnameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	HostnameId *int64 `json:"HostnameId,omitempty" xml:"HostnameId,omitempty"`
}

func (s GetCustomHostnameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHostnameRequest) GoString() string {
	return s.String()
}

func (s *GetCustomHostnameRequest) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *GetCustomHostnameRequest) SetHostnameId(v int64) *GetCustomHostnameRequest {
	s.HostnameId = &v
	return s
}

func (s *GetCustomHostnameRequest) Validate() error {
	return dara.Validate(s)
}
