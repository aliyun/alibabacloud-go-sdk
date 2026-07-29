// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAccountCreditRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetYikeAccountCreditRequest struct {
}

func (s GetYikeAccountCreditRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAccountCreditRequest) GoString() string {
	return s.String()
}

func (s *GetYikeAccountCreditRequest) Validate() error {
	return dara.Validate(s)
}
