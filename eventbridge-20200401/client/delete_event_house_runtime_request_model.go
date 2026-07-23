// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventHouseRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteEventHouseRuntimeRequest
	GetName() *string
}

type DeleteEventHouseRuntimeRequest struct {
	// The name of the EventHouse Runtime. If this parameter is not specified, the default Runtime is used. In most cases, you do not need to specify this parameter.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteEventHouseRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventHouseRuntimeRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventHouseRuntimeRequest) GetName() *string {
	return s.Name
}

func (s *DeleteEventHouseRuntimeRequest) SetName(v string) *DeleteEventHouseRuntimeRequest {
	s.Name = &v
	return s
}

func (s *DeleteEventHouseRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
