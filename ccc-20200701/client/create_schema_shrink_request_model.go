// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSchemaShrinkRequest
	GetDescription() *string
	SetId(v string) *CreateSchemaShrinkRequest
	GetId() *string
	SetInstanceId(v string) *CreateSchemaShrinkRequest
	GetInstanceId() *string
	SetPropertiesShrink(v string) *CreateSchemaShrinkRequest
	GetPropertiesShrink() *string
	SetRequestId(v string) *CreateSchemaShrinkRequest
	GetRequestId() *string
}

type CreateSchemaShrinkRequest struct {
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// schema id
	//
	// example:
	//
	// profile
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1c450
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSchemaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemaShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSchemaShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSchemaShrinkRequest) GetId() *string {
	return s.Id
}

func (s *CreateSchemaShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSchemaShrinkRequest) GetPropertiesShrink() *string {
	return s.PropertiesShrink
}

func (s *CreateSchemaShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchemaShrinkRequest) SetDescription(v string) *CreateSchemaShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSchemaShrinkRequest) SetId(v string) *CreateSchemaShrinkRequest {
	s.Id = &v
	return s
}

func (s *CreateSchemaShrinkRequest) SetInstanceId(v string) *CreateSchemaShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSchemaShrinkRequest) SetPropertiesShrink(v string) *CreateSchemaShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *CreateSchemaShrinkRequest) SetRequestId(v string) *CreateSchemaShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *CreateSchemaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
