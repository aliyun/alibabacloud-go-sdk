// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListMemberSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterListMemberSubscriptionsRequest struct {
}

func (s ModelRouterListMemberSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
