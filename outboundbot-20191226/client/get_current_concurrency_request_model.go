// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentConcurrencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetCurrentConcurrencyRequest
	GetInstanceId() *string
}

type GetCurrentConcurrencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCurrentConcurrencyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentConcurrencyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCurrentConcurrencyRequest) SetInstanceId(v string) *GetCurrentConcurrencyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCurrentConcurrencyRequest) Validate() error {
	return dara.Validate(s)
}
