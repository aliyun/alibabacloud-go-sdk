// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShardRecoveriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveOnly(v bool) *ListShardRecoveriesRequest
	GetActiveOnly() *bool
}

type ListShardRecoveriesRequest struct {
	// Specifies whether to return information about data restoration of shards. Valid values:
	//
	// 	- true: returns information about data restoration of shards that are being restored.
	//
	// 	- false: returns information about data restoration of all shards.
	//
	// example:
	//
	// true
	ActiveOnly *bool `json:"activeOnly,omitempty" xml:"activeOnly,omitempty"`
}

func (s ListShardRecoveriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListShardRecoveriesRequest) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesRequest) GetActiveOnly() *bool {
	return s.ActiveOnly
}

func (s *ListShardRecoveriesRequest) SetActiveOnly(v bool) *ListShardRecoveriesRequest {
	s.ActiveOnly = &v
	return s
}

func (s *ListShardRecoveriesRequest) Validate() error {
	return dara.Validate(s)
}
