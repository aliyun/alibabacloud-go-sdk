// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApiKeyRequest interface {
  dara.Model
  String() string
  GoString() string
}

type EnableApiKeyRequest struct {
}

func (s EnableApiKeyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApiKeyRequest) GoString() string {
  return s.String()
}

func (s *EnableApiKeyRequest) Validate() error {
  return dara.Validate(s)
}

