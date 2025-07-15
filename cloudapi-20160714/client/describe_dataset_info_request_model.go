// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *DescribeDatasetInfoRequest
	GetDatasetId() *string
	SetSecurityToken(v string) *DescribeDatasetInfoRequest
	GetSecurityToken() *string
}

type DescribeDatasetInfoRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// a25a6589b2584ff490e891cc********
	DatasetId     *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDatasetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatasetInfoRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeDatasetInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDatasetInfoRequest) SetDatasetId(v string) *DescribeDatasetInfoRequest {
	s.DatasetId = &v
	return s
}

func (s *DescribeDatasetInfoRequest) SetSecurityToken(v string) *DescribeDatasetInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDatasetInfoRequest) Validate() error {
	return dara.Validate(s)
}
