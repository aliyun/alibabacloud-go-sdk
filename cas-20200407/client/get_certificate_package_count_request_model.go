// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificatePackageCountRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCertificatePackageCountRequest struct {
}

func (s GetCertificatePackageCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCertificatePackageCountRequest) GoString() string {
	return s.String()
}

func (s *GetCertificatePackageCountRequest) Validate() error {
	return dara.Validate(s)
}
