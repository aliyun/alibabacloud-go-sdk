// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEnterpriseCodeRequest
	GetClientToken() *string
	SetEnterpriseCode(v string) *DeleteEnterpriseCodeRequest
	GetEnterpriseCode() *string
	SetRegionId(v string) *DeleteEnterpriseCodeRequest
	GetRegionId() *string
}

type DeleteEnterpriseCodeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enterprise code that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***
	EnterpriseCode *string `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	// The ID of the region to which the enterprise code belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEnterpriseCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseCodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseCodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEnterpriseCodeRequest) GetEnterpriseCode() *string {
	return s.EnterpriseCode
}

func (s *DeleteEnterpriseCodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnterpriseCodeRequest) SetClientToken(v string) *DeleteEnterpriseCodeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEnterpriseCodeRequest) SetEnterpriseCode(v string) *DeleteEnterpriseCodeRequest {
	s.EnterpriseCode = &v
	return s
}

func (s *DeleteEnterpriseCodeRequest) SetRegionId(v string) *DeleteEnterpriseCodeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnterpriseCodeRequest) Validate() error {
	return dara.Validate(s)
}
