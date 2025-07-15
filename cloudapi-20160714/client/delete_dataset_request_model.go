// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *DeleteDatasetRequest
	GetDatasetId() *string
	SetSecurityToken(v string) *DeleteDatasetRequest
	GetSecurityToken() *string
}

type DeleteDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a25a6589b2584ff490e891cc********
	DatasetId     *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeleteDatasetRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDatasetRequest) SetDatasetId(v string) *DeleteDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteDatasetRequest) SetSecurityToken(v string) *DeleteDatasetRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDatasetRequest) Validate() error {
	return dara.Validate(s)
}
