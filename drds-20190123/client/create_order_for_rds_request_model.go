// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForRdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *CreateOrderForRdsRequest
	GetParams() *string
	SetRegionId(v string) *CreateOrderForRdsRequest
	GetRegionId() *string
}

type CreateOrderForRdsRequest struct {
	// The JSON string that contains the order details. For more information, see [CreateDBInstance](https://help.aliyun.com/document_detail/26228.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"key":"value"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateOrderForRdsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForRdsRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsRequest) GetParams() *string {
	return s.Params
}

func (s *CreateOrderForRdsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrderForRdsRequest) SetParams(v string) *CreateOrderForRdsRequest {
	s.Params = &v
	return s
}

func (s *CreateOrderForRdsRequest) SetRegionId(v string) *CreateOrderForRdsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrderForRdsRequest) Validate() error {
	return dara.Validate(s)
}
