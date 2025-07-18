// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeletePrivateAccessApplicationRequest
	GetApplicationId() *string
}

type DeletePrivateAccessApplicationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s DeletePrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeletePrivateAccessApplicationRequest) SetApplicationId(v string) *DeletePrivateAccessApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeletePrivateAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}
