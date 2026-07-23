// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventHouseRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetEventHouseRuntimeRequest
	GetName() *string
}

type GetEventHouseRuntimeRequest struct {
	// The name of the EventHouse Runtime. If this parameter is not specified, the default Runtime is queried.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetEventHouseRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEventHouseRuntimeRequest) GoString() string {
	return s.String()
}

func (s *GetEventHouseRuntimeRequest) GetName() *string {
	return s.Name
}

func (s *GetEventHouseRuntimeRequest) SetName(v string) *GetEventHouseRuntimeRequest {
	s.Name = &v
	return s
}

func (s *GetEventHouseRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
