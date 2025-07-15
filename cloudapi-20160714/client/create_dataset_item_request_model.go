// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *CreateDatasetItemRequest
	GetDatasetId() *string
	SetDescription(v string) *CreateDatasetItemRequest
	GetDescription() *string
	SetExpiredTime(v string) *CreateDatasetItemRequest
	GetExpiredTime() *string
	SetSecurityToken(v string) *CreateDatasetItemRequest
	GetSecurityToken() *string
	SetValue(v string) *CreateDatasetItemRequest
	GetValue() *string
}

type CreateDatasetItemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a25a6589b2584ff490e891cc********
	DatasetId   *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2022-09-22T12:00:00Z
	ExpiredTime   *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 106.43.XXX.XXX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDatasetItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetItemRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetItemRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateDatasetItemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetItemRequest) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *CreateDatasetItemRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateDatasetItemRequest) GetValue() *string {
	return s.Value
}

func (s *CreateDatasetItemRequest) SetDatasetId(v string) *CreateDatasetItemRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetItemRequest) SetDescription(v string) *CreateDatasetItemRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetItemRequest) SetExpiredTime(v string) *CreateDatasetItemRequest {
	s.ExpiredTime = &v
	return s
}

func (s *CreateDatasetItemRequest) SetSecurityToken(v string) *CreateDatasetItemRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateDatasetItemRequest) SetValue(v string) *CreateDatasetItemRequest {
	s.Value = &v
	return s
}

func (s *CreateDatasetItemRequest) Validate() error {
	return dara.Validate(s)
}
