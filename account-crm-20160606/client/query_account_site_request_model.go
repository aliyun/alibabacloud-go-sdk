// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPk(v string) *QueryAccountSiteRequest
	GetPk() *string
}

type QueryAccountSiteRequest struct {
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s QueryAccountSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSiteRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountSiteRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryAccountSiteRequest) SetPk(v string) *QueryAccountSiteRequest {
	s.Pk = &v
	return s
}

func (s *QueryAccountSiteRequest) Validate() error {
	return dara.Validate(s)
}
