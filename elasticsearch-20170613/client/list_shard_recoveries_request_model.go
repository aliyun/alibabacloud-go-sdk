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
	// Specifies whether to display shard data recovery tracking information. Valid values:
	//
	// - true: Displays only ongoing shard data recovery tracking information.
	//
	// - false: Displays all shard data recovery tracking information.
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
