// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTarget(v string) *GetAgentlessTaskCountRequest
	GetTarget() *string
	SetTargetType(v int32) *GetAgentlessTaskCountRequest
	GetTargetType() *int32
}

type GetAgentlessTaskCountRequest struct {
	// The instance ID of the asset.
	//
	// example:
	//
	// s-m5edddcwq7d57d******
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Specifies the type of the object being inspected. Valid values:
	//
	// 	- **3**: User Snapshot.
	//
	// 	- **4**: User Image.
	//
	// example:
	//
	// 3
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetAgentlessTaskCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskCountRequest) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskCountRequest) GetTarget() *string {
	return s.Target
}

func (s *GetAgentlessTaskCountRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *GetAgentlessTaskCountRequest) SetTarget(v string) *GetAgentlessTaskCountRequest {
	s.Target = &v
	return s
}

func (s *GetAgentlessTaskCountRequest) SetTargetType(v int32) *GetAgentlessTaskCountRequest {
	s.TargetType = &v
	return s
}

func (s *GetAgentlessTaskCountRequest) Validate() error {
	return dara.Validate(s)
}
