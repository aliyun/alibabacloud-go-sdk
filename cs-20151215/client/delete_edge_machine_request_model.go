// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v string) *DeleteEdgeMachineRequest
	GetForce() *string
}

type DeleteEdgeMachineRequest struct {
	// Specifies whether to forcefully delete the cloud-native box. Valid values:
	//
	// 	- `true`: forcefully deletes the cloud-native box.
	//
	// 	- `false`: does not forcefully delete the cloud-native box.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true/false
	Force *string `json:"force,omitempty" xml:"force,omitempty"`
}

func (s DeleteEdgeMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeMachineRequest) GoString() string {
	return s.String()
}

func (s *DeleteEdgeMachineRequest) GetForce() *string {
	return s.Force
}

func (s *DeleteEdgeMachineRequest) SetForce(v string) *DeleteEdgeMachineRequest {
	s.Force = &v
	return s
}

func (s *DeleteEdgeMachineRequest) Validate() error {
	return dara.Validate(s)
}
