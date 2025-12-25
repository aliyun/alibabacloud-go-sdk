// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPrepayInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConvertPrepayInstanceRequest(v *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) *ConvertPrepayInstanceRequest
	GetConvertPrepayInstanceRequest() *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest
}

type ConvertPrepayInstanceRequest struct {
	// This parameter is required.
	ConvertPrepayInstanceRequest *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest `json:"ConvertPrepayInstanceRequest,omitempty" xml:"ConvertPrepayInstanceRequest,omitempty" type:"Struct"`
}

func (s ConvertPrepayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertPrepayInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceRequest) GetConvertPrepayInstanceRequest() *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest {
	return s.ConvertPrepayInstanceRequest
}

func (s *ConvertPrepayInstanceRequest) SetConvertPrepayInstanceRequest(v *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) *ConvertPrepayInstanceRequest {
	s.ConvertPrepayInstanceRequest = v
	return s
}

func (s *ConvertPrepayInstanceRequest) Validate() error {
	if s.ConvertPrepayInstanceRequest != nil {
		if err := s.ConvertPrepayInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertPrepayInstanceRequestConvertPrepayInstanceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) SetInstanceId(v string) *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) SetRegion(v string) *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest {
	s.Region = &v
	return s
}

func (s *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) Validate() error {
	return dara.Validate(s)
}
