// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStopGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DeleteStopGatewayRequest
	GetGatewayId() *string
}

type DeleteStopGatewayRequest struct {
	// The ID that can uniquely identify the DBGateway. You can obtain the DBGateway ID by calling the [DescribeCloudbenchTask](https://help.aliyun.com/document_detail/230669.html) operation. The DBGateway ID is the value of the **ClientGatewayId*	- field in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22938c83fcfbced4b4869b9695e3****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s DeleteStopGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStopGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteStopGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeleteStopGatewayRequest) SetGatewayId(v string) *DeleteStopGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteStopGatewayRequest) Validate() error {
	return dara.Validate(s)
}
