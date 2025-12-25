// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertPrepayInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConvertPrepayInstanceRequest(v *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) *QueryConvertPrepayInstancePriceRequest
	GetConvertPrepayInstanceRequest() *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest
}

type QueryConvertPrepayInstancePriceRequest struct {
	// This parameter is required.
	ConvertPrepayInstanceRequest *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest `json:"ConvertPrepayInstanceRequest,omitempty" xml:"ConvertPrepayInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryConvertPrepayInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceRequest) GetConvertPrepayInstanceRequest() *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest {
	return s.ConvertPrepayInstanceRequest
}

func (s *QueryConvertPrepayInstancePriceRequest) SetConvertPrepayInstanceRequest(v *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) *QueryConvertPrepayInstancePriceRequest {
	s.ConvertPrepayInstanceRequest = v
	return s
}

func (s *QueryConvertPrepayInstancePriceRequest) Validate() error {
	if s.ConvertPrepayInstanceRequest != nil {
		if err := s.ConvertPrepayInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) SetInstanceId(v string) *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) SetRegion(v string) *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest {
	s.Region = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) Validate() error {
	return dara.Validate(s)
}
