// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhitePortalMenuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataportalId(v string) *ListWhitePortalMenuRequest
	GetDataportalId() *string
}

type ListWhitePortalMenuRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// asdagad***213425
	DataportalId *string `json:"DataportalId,omitempty" xml:"DataportalId,omitempty"`
}

func (s ListWhitePortalMenuRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWhitePortalMenuRequest) GoString() string {
	return s.String()
}

func (s *ListWhitePortalMenuRequest) GetDataportalId() *string {
	return s.DataportalId
}

func (s *ListWhitePortalMenuRequest) SetDataportalId(v string) *ListWhitePortalMenuRequest {
	s.DataportalId = &v
	return s
}

func (s *ListWhitePortalMenuRequest) Validate() error {
	return dara.Validate(s)
}
