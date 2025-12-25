// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteInstanceRequest(v *DeleteInstanceRequestDeleteInstanceRequest) *DeleteInstanceRequest
	GetDeleteInstanceRequest() *DeleteInstanceRequestDeleteInstanceRequest
}

type DeleteInstanceRequest struct {
	// This parameter is required.
	DeleteInstanceRequest *DeleteInstanceRequestDeleteInstanceRequest `json:"DeleteInstanceRequest,omitempty" xml:"DeleteInstanceRequest,omitempty" type:"Struct"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetDeleteInstanceRequest() *DeleteInstanceRequestDeleteInstanceRequest {
	return s.DeleteInstanceRequest
}

func (s *DeleteInstanceRequest) SetDeleteInstanceRequest(v *DeleteInstanceRequestDeleteInstanceRequest) *DeleteInstanceRequest {
	s.DeleteInstanceRequest = v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	if s.DeleteInstanceRequest != nil {
		if err := s.DeleteInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteInstanceRequestDeleteInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223493C7-FCA9-13D4-B75B-AF8B32F4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteInstanceRequestDeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequestDeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequestDeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequestDeleteInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteInstanceRequestDeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequestDeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequestDeleteInstanceRequest) SetRegion(v string) *DeleteInstanceRequestDeleteInstanceRequest {
	s.Region = &v
	return s
}

func (s *DeleteInstanceRequestDeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
