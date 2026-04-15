// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int32) *GetAlertDestinationRequest
	GetId() *int32
}

type GetAlertDestinationRequest struct {
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetAlertDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertDestinationRequest) GoString() string {
	return s.String()
}

func (s *GetAlertDestinationRequest) GetId() *int32 {
	return s.Id
}

func (s *GetAlertDestinationRequest) SetId(v int32) *GetAlertDestinationRequest {
	s.Id = &v
	return s
}

func (s *GetAlertDestinationRequest) Validate() error {
	return dara.Validate(s)
}
