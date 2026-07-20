// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackIdRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCompliancePackIdRequest struct {
}

func (s GetCompliancePackIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackIdRequest) GoString() string {
	return s.String()
}

func (s *GetCompliancePackIdRequest) Validate() error {
	return dara.Validate(s)
}
