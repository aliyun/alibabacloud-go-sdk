// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionId(v int64) *DeleteConnectionRequest
	GetConnectionId() *int64
}

type DeleteConnectionRequest struct {
	// The data source ID. You can call the [ListConnection](https://help.aliyun.com/document_detail/173911.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConnectionId *int64 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s DeleteConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectionRequest) GetConnectionId() *int64 {
	return s.ConnectionId
}

func (s *DeleteConnectionRequest) SetConnectionId(v int64) *DeleteConnectionRequest {
	s.ConnectionId = &v
	return s
}

func (s *DeleteConnectionRequest) Validate() error {
	return dara.Validate(s)
}
