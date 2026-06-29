// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationMemberSeatStatsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetOrganizationMemberSeatStatsRequest struct {
}

func (s GetOrganizationMemberSeatStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationMemberSeatStatsRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberSeatStatsRequest) Validate() error {
	return dara.Validate(s)
}
