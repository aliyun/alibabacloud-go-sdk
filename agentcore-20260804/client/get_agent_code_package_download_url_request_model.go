// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentCodePackageDownloadUrlRequest struct {
}

func (s GetAgentCodePackageDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
