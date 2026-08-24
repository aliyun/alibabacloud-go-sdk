// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanGlobalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetVirusScanGlobalConfigRequest struct {
}

func (s GetVirusScanGlobalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *GetVirusScanGlobalConfigRequest) Validate() error {
	return dara.Validate(s)
}
