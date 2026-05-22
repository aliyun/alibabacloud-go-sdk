// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomHostnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnameId(v int64) *DeleteCustomHostnameRequest
	GetHostnameId() *int64
}

type DeleteCustomHostnameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	HostnameId *int64 `json:"HostnameId,omitempty" xml:"HostnameId,omitempty"`
}

func (s DeleteCustomHostnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomHostnameRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomHostnameRequest) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *DeleteCustomHostnameRequest) SetHostnameId(v int64) *DeleteCustomHostnameRequest {
	s.HostnameId = &v
	return s
}

func (s *DeleteCustomHostnameRequest) Validate() error {
	return dara.Validate(s)
}
