// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlFileRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSqlFileRequest struct {
}

func (s GetSqlFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlFileRequest) GoString() string {
	return s.String()
}

func (s *GetSqlFileRequest) Validate() error {
	return dara.Validate(s)
}
