// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaEntityDefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *GetMetaEntityDefRequest
	GetEntityType() *string
}

type GetMetaEntityDefRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s GetMetaEntityDefRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaEntityDefRequest) GoString() string {
	return s.String()
}

func (s *GetMetaEntityDefRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *GetMetaEntityDefRequest) SetEntityType(v string) *GetMetaEntityDefRequest {
	s.EntityType = &v
	return s
}

func (s *GetMetaEntityDefRequest) Validate() error {
	return dara.Validate(s)
}
