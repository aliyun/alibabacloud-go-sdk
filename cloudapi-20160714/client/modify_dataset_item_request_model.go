// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatasetItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *ModifyDatasetItemRequest
	GetDatasetId() *string
	SetDatasetItemId(v string) *ModifyDatasetItemRequest
	GetDatasetItemId() *string
	SetDescription(v string) *ModifyDatasetItemRequest
	GetDescription() *string
	SetExpiredTime(v string) *ModifyDatasetItemRequest
	GetExpiredTime() *string
	SetSecurityToken(v string) *ModifyDatasetItemRequest
	GetSecurityToken() *string
}

type ModifyDatasetItemRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// a25a6589b2584ff490e891cc********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The ID of the data entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5045****
	DatasetItemId *string `json:"DatasetItemId,omitempty" xml:"DatasetItemId,omitempty"`
	// The description of the data entry. The description cannot exceed 180 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time in UTC when the data entry expires. The time is in the **yyyy-MM-ddTHH:mm:ssZ*	- format.
	//
	// example:
	//
	// 2022-09-22T12:00:00Z
	ExpiredTime   *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyDatasetItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatasetItemRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatasetItemRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ModifyDatasetItemRequest) GetDatasetItemId() *string {
	return s.DatasetItemId
}

func (s *ModifyDatasetItemRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDatasetItemRequest) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ModifyDatasetItemRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDatasetItemRequest) SetDatasetId(v string) *ModifyDatasetItemRequest {
	s.DatasetId = &v
	return s
}

func (s *ModifyDatasetItemRequest) SetDatasetItemId(v string) *ModifyDatasetItemRequest {
	s.DatasetItemId = &v
	return s
}

func (s *ModifyDatasetItemRequest) SetDescription(v string) *ModifyDatasetItemRequest {
	s.Description = &v
	return s
}

func (s *ModifyDatasetItemRequest) SetExpiredTime(v string) *ModifyDatasetItemRequest {
	s.ExpiredTime = &v
	return s
}

func (s *ModifyDatasetItemRequest) SetSecurityToken(v string) *ModifyDatasetItemRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDatasetItemRequest) Validate() error {
	return dara.Validate(s)
}
