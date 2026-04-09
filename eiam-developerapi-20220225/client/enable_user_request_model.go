// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableUserRequest interface {
  dara.Model
  String() string
  GoString() string
}

type EnableUserRequest struct {
}

func (s EnableUserRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableUserRequest) GoString() string {
  return s.String()
}

func (s *EnableUserRequest) Validate() error {
  return dara.Validate(s)
}

