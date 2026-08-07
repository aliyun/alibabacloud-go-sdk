// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFlashSmsProvidersRequest
	GetInstanceId() *string
}

type ListFlashSmsProvidersRequest struct {
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListFlashSmsProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListFlashSmsProvidersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsProvidersRequest) SetInstanceId(v string) *ListFlashSmsProvidersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsProvidersRequest) Validate() error {
	return dara.Validate(s)
}
