// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppFailBackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AppFailBackRequest
	GetApplicationId() *string
}

type AppFailBackRequest struct {
	// The application ID.
	//
	// example:
	//
	// 61ZW1DY5Y3FSAOO2
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s AppFailBackRequest) String() string {
	return dara.Prettify(s)
}

func (s AppFailBackRequest) GoString() string {
	return s.String()
}

func (s *AppFailBackRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AppFailBackRequest) SetApplicationId(v string) *AppFailBackRequest {
	s.ApplicationId = &v
	return s
}

func (s *AppFailBackRequest) Validate() error {
	return dara.Validate(s)
}
