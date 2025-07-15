// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *CreateModelResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *CreateModelResponseBody
	GetDescription() *string
	SetGroupId(v string) *CreateModelResponseBody
	GetGroupId() *string
	SetModelId(v string) *CreateModelResponseBody
	GetModelId() *string
	SetModelName(v string) *CreateModelResponseBody
	GetModelName() *string
	SetModelRef(v string) *CreateModelResponseBody
	GetModelRef() *string
	SetModifiedTime(v string) *CreateModelResponseBody
	GetModifiedTime() *string
	SetRegionId(v string) *CreateModelResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateModelResponseBody
	GetRequestId() *string
	SetSchema(v string) *CreateModelResponseBody
	GetSchema() *string
}

type CreateModelResponseBody struct {
	// The time when the model was created.
	//
	// example:
	//
	// 2019-01-29T09:33:01Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the created model.
	//
	// example:
	//
	// Model Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group to which the created model belongs.
	//
	// example:
	//
	// 30e792398d6c4569b04c0e53a3494381
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the created model.
	//
	// example:
	//
	// 766c0b9538a04bdf974953b5576783ba
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the created model.
	//
	// example:
	//
	// Test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The URI of the created model.
	//
	// example:
	//
	// https://apigateway.aliyun.com/models/30e792398d6c4569b04c0e53a3494381/766c0b9538a04bdf974953b5576783ba
	ModelRef *string `json:"ModelRef,omitempty" xml:"ModelRef,omitempty"`
	// The time when the model is last modified.
	//
	// example:
	//
	// 2019-01-29T09:33:01Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region to which the created model belongs.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4173F95B-360C-460C-9F6C-4A960B904411
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The definition of the created model.
	//
	// example:
	//
	// {\\"type\\":\\"object\\",\\"properties\\":{\\"id\\":{\\"format\\":\\"int64\\",\\"maximum\\":100,\\"exclusiveMaximum\\":true,\\"type\\":\\"integer\\"},\\"name\\":{\\"maxLength\\":10,\\"type\\":\\"string\\"}}}
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *CreateModelResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateModelResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateModelResponseBody) GetModelId() *string {
	return s.ModelId
}

func (s *CreateModelResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelResponseBody) GetModelRef() *string {
	return s.ModelRef
}

func (s *CreateModelResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CreateModelResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelResponseBody) GetSchema() *string {
	return s.Schema
}

func (s *CreateModelResponseBody) SetCreatedTime(v string) *CreateModelResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateModelResponseBody) SetDescription(v string) *CreateModelResponseBody {
	s.Description = &v
	return s
}

func (s *CreateModelResponseBody) SetGroupId(v string) *CreateModelResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateModelResponseBody) SetModelId(v string) *CreateModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *CreateModelResponseBody) SetModelName(v string) *CreateModelResponseBody {
	s.ModelName = &v
	return s
}

func (s *CreateModelResponseBody) SetModelRef(v string) *CreateModelResponseBody {
	s.ModelRef = &v
	return s
}

func (s *CreateModelResponseBody) SetModifiedTime(v string) *CreateModelResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateModelResponseBody) SetRegionId(v string) *CreateModelResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetSchema(v string) *CreateModelResponseBody {
	s.Schema = &v
	return s
}

func (s *CreateModelResponseBody) Validate() error {
	return dara.Validate(s)
}
