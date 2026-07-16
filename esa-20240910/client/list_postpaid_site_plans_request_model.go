// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostpaidSitePlansRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListPostpaidSitePlansRequest struct {
}

func (s ListPostpaidSitePlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidSitePlansRequest) GoString() string {
	return s.String()
}

func (s *ListPostpaidSitePlansRequest) Validate() error {
	return dara.Validate(s)
}
