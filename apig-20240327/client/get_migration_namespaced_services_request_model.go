// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationNamespacedServicesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMigrationNamespacedServicesRequest struct {
}

func (s GetMigrationNamespacedServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationNamespacedServicesRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationNamespacedServicesRequest) Validate() error {
	return dara.Validate(s)
}
