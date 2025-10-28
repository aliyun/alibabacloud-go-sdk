// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDefineRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDebugEnable(v bool) *ListUserDefineRegionRequest
	GetDebugEnable() *bool
}

type ListUserDefineRegionRequest struct {
	// Specifies whether remote debugging is allowed.
	//
	// example:
	//
	// false
	DebugEnable *bool `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty"`
}

func (s ListUserDefineRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefineRegionRequest) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionRequest) GetDebugEnable() *bool {
	return s.DebugEnable
}

func (s *ListUserDefineRegionRequest) SetDebugEnable(v bool) *ListUserDefineRegionRequest {
	s.DebugEnable = &v
	return s
}

func (s *ListUserDefineRegionRequest) Validate() error {
	return dara.Validate(s)
}
