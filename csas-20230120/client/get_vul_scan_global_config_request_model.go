// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulScanGlobalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetVulScanGlobalConfigRequest struct {
}

func (s GetVulScanGlobalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *GetVulScanGlobalConfigRequest) Validate() error {
	return dara.Validate(s)
}
