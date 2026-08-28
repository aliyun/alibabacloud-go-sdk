// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMigrationTaskRequest struct {
}

func (s DeleteMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
