// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *ModifyDatasetRequest
	GetDatasetId() *string
	SetDatasetName(v string) *ModifyDatasetRequest
	GetDatasetName() *string
	SetDescription(v string) *ModifyDatasetRequest
	GetDescription() *string
	SetSecurityToken(v string) *ModifyDatasetRequest
	GetSecurityToken() *string
}

type ModifyDatasetRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// a25a6589b2584ff490e891cc********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// NewDatasetName
	DatasetName   *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatasetRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ModifyDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ModifyDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDatasetRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDatasetRequest) SetDatasetId(v string) *ModifyDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *ModifyDatasetRequest) SetDatasetName(v string) *ModifyDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *ModifyDatasetRequest) SetDescription(v string) *ModifyDatasetRequest {
	s.Description = &v
	return s
}

func (s *ModifyDatasetRequest) SetSecurityToken(v string) *ModifyDatasetRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDatasetRequest) Validate() error {
	return dara.Validate(s)
}
