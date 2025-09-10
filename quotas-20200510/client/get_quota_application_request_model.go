// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetQuotaApplicationRequest
	GetApplicationId() *string
}

type GetQuotaApplicationRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s GetQuotaApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetQuotaApplicationRequest) SetApplicationId(v string) *GetQuotaApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetQuotaApplicationRequest) Validate() error {
	return dara.Validate(s)
}
