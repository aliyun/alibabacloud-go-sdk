// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableExecuteStatementRequest interface {
  dara.Model
  String() string
  GoString() string
}

type EnableExecuteStatementRequest struct {
}

func (s EnableExecuteStatementRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableExecuteStatementRequest) GoString() string {
  return s.String()
}

func (s *EnableExecuteStatementRequest) Validate() error {
  return dara.Validate(s)
}

