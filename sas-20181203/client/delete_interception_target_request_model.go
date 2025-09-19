// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterceptionTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetIds(v string) *DeleteInterceptionTargetRequest
	GetTargetIds() *string
}

type DeleteInterceptionTargetRequest struct {
	// The IDs of the network objects that you want to remove. You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to query the IDs of the network objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,11,111
	TargetIds *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
}

func (s DeleteInterceptionTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterceptionTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteInterceptionTargetRequest) GetTargetIds() *string {
	return s.TargetIds
}

func (s *DeleteInterceptionTargetRequest) SetTargetIds(v string) *DeleteInterceptionTargetRequest {
	s.TargetIds = &v
	return s
}

func (s *DeleteInterceptionTargetRequest) Validate() error {
	return dara.Validate(s)
}
