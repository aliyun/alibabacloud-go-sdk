// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteEipRequest
	GetInstanceId() *string
}

type DeleteEipRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5sg1owx0g4ojy66ab2tez77r2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteEipRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEipRequest) GoString() string {
	return s.String()
}

func (s *DeleteEipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteEipRequest) SetInstanceId(v string) *DeleteEipRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteEipRequest) Validate() error {
	return dara.Validate(s)
}
