// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlExecutionResultRequest interface {
	dara.Model
	String() string
	GoString() string
}

type FetchSqlExecutionResultRequest struct {
}

func (s FetchSqlExecutionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlExecutionResultRequest) GoString() string {
	return s.String()
}

func (s *FetchSqlExecutionResultRequest) Validate() error {
	return dara.Validate(s)
}
