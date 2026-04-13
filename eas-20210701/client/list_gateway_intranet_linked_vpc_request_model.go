// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetLinkedVpcRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListGatewayIntranetLinkedVpcRequest struct {
}

func (s ListGatewayIntranetLinkedVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcRequest) Validate() error {
	return dara.Validate(s)
}
