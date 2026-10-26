// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionReportRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetInstanceInspectionReportRequest struct {
}

func (s GetInstanceInspectionReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportRequest) Validate() error {
	return dara.Validate(s)
}
