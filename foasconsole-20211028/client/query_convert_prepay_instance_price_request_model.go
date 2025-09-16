// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertPrepayInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryConvertPrepayInstancePriceRequest
	GetInstanceId() *string
	SetRegion(v string) *QueryConvertPrepayInstancePriceRequest
	GetRegion() *string
}

type QueryConvertPrepayInstancePriceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryConvertPrepayInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryConvertPrepayInstancePriceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryConvertPrepayInstancePriceRequest) SetInstanceId(v string) *QueryConvertPrepayInstancePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceRequest) SetRegion(v string) *QueryConvertPrepayInstancePriceRequest {
	s.Region = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceRequest) Validate() error {
	return dara.Validate(s)
}
