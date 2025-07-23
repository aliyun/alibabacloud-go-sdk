// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppFailOverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AppFailOverRequest
	GetApplicationId() *string
	SetFailZone(v string) *AppFailOverRequest
	GetFailZone() *string
}

type AppFailOverRequest struct {
	// The application ID.
	//
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The destination zone to which you want to switch the disaster recovery application.
	//
	// example:
	//
	// cn-hangzhou-g
	FailZone *string `json:"FailZone,omitempty" xml:"FailZone,omitempty"`
}

func (s AppFailOverRequest) String() string {
	return dara.Prettify(s)
}

func (s AppFailOverRequest) GoString() string {
	return s.String()
}

func (s *AppFailOverRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AppFailOverRequest) GetFailZone() *string {
	return s.FailZone
}

func (s *AppFailOverRequest) SetApplicationId(v string) *AppFailOverRequest {
	s.ApplicationId = &v
	return s
}

func (s *AppFailOverRequest) SetFailZone(v string) *AppFailOverRequest {
	s.FailZone = &v
	return s
}

func (s *AppFailOverRequest) Validate() error {
	return dara.Validate(s)
}
