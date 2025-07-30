// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEveryOneSellsFormListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v string) *GetEveryOneSellsFormListRequest
	GetAuth() *string
}

type GetEveryOneSellsFormListRequest struct {
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
}

func (s GetEveryOneSellsFormListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEveryOneSellsFormListRequest) GoString() string {
	return s.String()
}

func (s *GetEveryOneSellsFormListRequest) GetAuth() *string {
	return s.Auth
}

func (s *GetEveryOneSellsFormListRequest) SetAuth(v string) *GetEveryOneSellsFormListRequest {
	s.Auth = &v
	return s
}

func (s *GetEveryOneSellsFormListRequest) Validate() error {
	return dara.Validate(s)
}
