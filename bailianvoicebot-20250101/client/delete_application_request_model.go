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
	SetBusinessUnitId(v string) *DeleteApplicationRequest
	GetBusinessUnitId() *string
}

type DeleteApplicationRequest struct {
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

func (s DeleteApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *DeleteApplicationRequest) SetApplicationId(v string) *DeleteApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationRequest) SetBusinessUnitId(v string) *DeleteApplicationRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *DeleteApplicationRequest) Validate() error {
	return dara.Validate(s)
}
