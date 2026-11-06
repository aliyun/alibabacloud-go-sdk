// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayDfpExtProcServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetGatewayDfpExtProcServiceRequest struct {
}

func (s GetGatewayDfpExtProcServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDfpExtProcServiceRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayDfpExtProcServiceRequest) Validate() error {
	return dara.Validate(s)
}
