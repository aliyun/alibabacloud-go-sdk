// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteNetworkRequest
	GetId() *int64
}

type DeleteNetworkRequest struct {
	// The ID of the network that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteNetworkRequest) SetId(v int64) *DeleteNetworkRequest {
	s.Id = &v
	return s
}

func (s *DeleteNetworkRequest) Validate() error {
	return dara.Validate(s)
}
