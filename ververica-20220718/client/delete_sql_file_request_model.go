// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSqlFileRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteSqlFileRequest struct {
}

func (s DeleteSqlFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSqlFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteSqlFileRequest) Validate() error {
	return dara.Validate(s)
}
