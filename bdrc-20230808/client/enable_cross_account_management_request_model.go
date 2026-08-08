// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCrossAccountManagementRequest interface {
  dara.Model
  String() string
  GoString() string
}

type EnableCrossAccountManagementRequest struct {
}

func (s EnableCrossAccountManagementRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCrossAccountManagementRequest) GoString() string {
  return s.String()
}

func (s *EnableCrossAccountManagementRequest) Validate() error {
  return dara.Validate(s)
}

