// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportResponseRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetReportResponseRequest struct {
}

func (s GetReportResponseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponseRequest) GoString() string {
	return s.String()
}

func (s *GetReportResponseRequest) Validate() error {
	return dara.Validate(s)
}
