// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaEntityDefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *CreateMetaEntityDefResponseBody
	GetEntityType() *string
	SetRequestId(v string) *CreateMetaEntityDefResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMetaEntityDefResponseBody
	GetSuccess() *bool
}

type CreateMetaEntityDefResponseBody struct {
	// example:
	//
	// custom_entity-biz_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0A04C673-BEFA-5803-94E5-89E2D9F8C567
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMetaEntityDefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaEntityDefResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetaEntityDefResponseBody) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateMetaEntityDefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetaEntityDefResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMetaEntityDefResponseBody) SetEntityType(v string) *CreateMetaEntityDefResponseBody {
	s.EntityType = &v
	return s
}

func (s *CreateMetaEntityDefResponseBody) SetRequestId(v string) *CreateMetaEntityDefResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetaEntityDefResponseBody) SetSuccess(v bool) *CreateMetaEntityDefResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMetaEntityDefResponseBody) Validate() error {
	return dara.Validate(s)
}
