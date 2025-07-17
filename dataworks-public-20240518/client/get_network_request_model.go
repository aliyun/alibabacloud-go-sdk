// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetNetworkRequest
	GetId() *int64
}

type GetNetworkRequest struct {
	// The network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkRequest) GetId() *int64 {
	return s.Id
}

func (s *GetNetworkRequest) SetId(v int64) *GetNetworkRequest {
	s.Id = &v
	return s
}

func (s *GetNetworkRequest) Validate() error {
	return dara.Validate(s)
}
