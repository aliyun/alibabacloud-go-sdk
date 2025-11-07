// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudauthstSceneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *DescribeCloudauthstSceneListRequest
	GetProductCode() *string
}

type DescribeCloudauthstSceneListRequest struct {
	// Product Code
	//
	// example:
	//
	// SMART_CARD
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeCloudauthstSceneListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudauthstSceneListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudauthstSceneListRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeCloudauthstSceneListRequest) SetProductCode(v string) *DescribeCloudauthstSceneListRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeCloudauthstSceneListRequest) Validate() error {
	return dara.Validate(s)
}
