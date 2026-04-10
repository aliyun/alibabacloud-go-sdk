// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizTraceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteBizTraceRequest struct {
}

func (s DeleteBizTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizTraceRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizTraceRequest) Validate() error {
	return dara.Validate(s)
}
