// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type VerifyMigrationTaskRequest struct {
}

func (s VerifyMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *VerifyMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
