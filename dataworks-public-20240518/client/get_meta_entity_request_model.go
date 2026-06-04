// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetMetaEntityRequest
	GetId() *string
}

type GetMetaEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api:api_001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMetaEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaEntityRequest) GoString() string {
	return s.String()
}

func (s *GetMetaEntityRequest) GetId() *string {
	return s.Id
}

func (s *GetMetaEntityRequest) SetId(v string) *GetMetaEntityRequest {
	s.Id = &v
	return s
}

func (s *GetMetaEntityRequest) Validate() error {
	return dara.Validate(s)
}
