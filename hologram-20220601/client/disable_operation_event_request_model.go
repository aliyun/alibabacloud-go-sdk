// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOperationEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DisableOperationEventRequest
	GetId() *string
}

type DisableOperationEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1980869072412614657
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DisableOperationEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableOperationEventRequest) GoString() string {
	return s.String()
}

func (s *DisableOperationEventRequest) GetId() *string {
	return s.Id
}

func (s *DisableOperationEventRequest) SetId(v string) *DisableOperationEventRequest {
	s.Id = &v
	return s
}

func (s *DisableOperationEventRequest) Validate() error {
	return dara.Validate(s)
}
