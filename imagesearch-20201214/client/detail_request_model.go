// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DetailRequest
	GetInstanceName() *string
}

type DetailRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailRequest) GoString() string {
	return s.String()
}

func (s *DetailRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DetailRequest) SetInstanceName(v string) *DetailRequest {
	s.InstanceName = &v
	return s
}

func (s *DetailRequest) Validate() error {
	return dara.Validate(s)
}
