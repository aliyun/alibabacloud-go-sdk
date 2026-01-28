// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *GetApplicationRequest
	GetBusinessUnitId() *string
}

type GetApplicationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *GetApplicationRequest) SetApplicationId(v string) *GetApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationRequest) SetBusinessUnitId(v string) *GetApplicationRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *GetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
