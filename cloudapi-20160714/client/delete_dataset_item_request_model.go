// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *DeleteDatasetItemRequest
	GetDatasetId() *string
	SetDatasetItemId(v string) *DeleteDatasetItemRequest
	GetDatasetItemId() *string
	SetSecurityToken(v string) *DeleteDatasetItemRequest
	GetSecurityToken() *string
}

type DeleteDatasetItemRequest struct {
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
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDatasetItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetItemRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeleteDatasetItemRequest) GetDatasetItemId() *string {
	return s.DatasetItemId
}

func (s *DeleteDatasetItemRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDatasetItemRequest) SetDatasetId(v string) *DeleteDatasetItemRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteDatasetItemRequest) SetDatasetItemId(v string) *DeleteDatasetItemRequest {
	s.DatasetItemId = &v
	return s
}

func (s *DeleteDatasetItemRequest) SetSecurityToken(v string) *DeleteDatasetItemRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDatasetItemRequest) Validate() error {
	return dara.Validate(s)
}
