// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteApplicationRequest
	GetApplicationId() *string
}

type DeleteApplicationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationRequest) SetApplicationId(v string) *DeleteApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationRequest) Validate() error {
	return dara.Validate(s)
}
