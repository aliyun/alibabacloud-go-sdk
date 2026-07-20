// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogKmsGrantsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCatalogKmsGrantsRequest struct {
}

func (s GetCatalogKmsGrantsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogKmsGrantsRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogKmsGrantsRequest) Validate() error {
	return dara.Validate(s)
}
