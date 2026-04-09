// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitParentIdsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListOrganizationalUnitParentIdsRequest struct {
}

func (s ListOrganizationalUnitParentIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsRequest) Validate() error {
	return dara.Validate(s)
}
