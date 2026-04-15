// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int32) *DeleteAlertDestinationRequest
	GetId() *int32
}

type DeleteAlertDestinationRequest struct {
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteAlertDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertDestinationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertDestinationRequest) GetId() *int32 {
	return s.Id
}

func (s *DeleteAlertDestinationRequest) SetId(v int32) *DeleteAlertDestinationRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlertDestinationRequest) Validate() error {
	return dara.Validate(s)
}
