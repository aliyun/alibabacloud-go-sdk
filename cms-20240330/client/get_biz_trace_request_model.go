// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizTraceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetBizTraceRequest struct {
}

func (s GetBizTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBizTraceRequest) GoString() string {
	return s.String()
}

func (s *GetBizTraceRequest) Validate() error {
	return dara.Validate(s)
}
