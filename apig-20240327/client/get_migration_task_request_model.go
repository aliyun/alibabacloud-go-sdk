// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMigrationTaskRequest struct {
}

func (s GetMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
