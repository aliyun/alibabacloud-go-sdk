// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteNumbersRequest
	GetInstanceId() *string
	SetNumbers(v []*string) *DeleteNumbersRequest
	GetNumbers() []*string
	SetTenantId(v string) *DeleteNumbersRequest
	GetTenantId() *string
}

type DeleteNumbersRequest struct {
	// The instance ID.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of numbers to be deleted.
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNumbersRequest) GoString() string {
	return s.String()
}

func (s *DeleteNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNumbersRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *DeleteNumbersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteNumbersRequest) SetInstanceId(v string) *DeleteNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNumbersRequest) SetNumbers(v []*string) *DeleteNumbersRequest {
	s.Numbers = v
	return s
}

func (s *DeleteNumbersRequest) SetTenantId(v string) *DeleteNumbersRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteNumbersRequest) Validate() error {
	return dara.Validate(s)
}
