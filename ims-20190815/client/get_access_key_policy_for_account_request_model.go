// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyPolicyForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAccessKeyPolicyForAccountRequest struct {
}

func (s GetAccessKeyPolicyForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyPolicyForAccountRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyPolicyForAccountRequest) Validate() error {
	return dara.Validate(s)
}
