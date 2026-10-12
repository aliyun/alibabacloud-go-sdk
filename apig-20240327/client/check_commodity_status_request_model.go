// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCommodityStatusRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CheckCommodityStatusRequest struct {
}

func (s CheckCommodityStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCommodityStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckCommodityStatusRequest) Validate() error {
	return dara.Validate(s)
}
