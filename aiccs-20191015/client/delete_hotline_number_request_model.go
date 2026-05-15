// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotlineNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotlineNumber(v string) *DeleteHotlineNumberRequest
	GetHotlineNumber() *string
	SetInstanceId(v string) *DeleteHotlineNumberRequest
	GetInstanceId() *string
}

type DeleteHotlineNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 05710000****
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteHotlineNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotlineNumberRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *DeleteHotlineNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHotlineNumberRequest) SetHotlineNumber(v string) *DeleteHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *DeleteHotlineNumberRequest) SetInstanceId(v string) *DeleteHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHotlineNumberRequest) Validate() error {
	return dara.Validate(s)
}
