// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEnterpriseCodeRequest
	GetClientToken() *string
	SetEnterpriseCode(v string) *CreateEnterpriseCodeRequest
	GetEnterpriseCode() *string
	SetRegionId(v string) *CreateEnterpriseCodeRequest
	GetRegionId() *string
}

type CreateEnterpriseCodeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically set **ClientToken*	- to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enterprise code.
	//
	// The enterprise code must be five characters in length and must contain letters and digits. Each enterprise code must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12P**
	EnterpriseCode *string `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	// The ID of the region.
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

func (s CreateEnterpriseCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseCodeRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseCodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEnterpriseCodeRequest) GetEnterpriseCode() *string {
	return s.EnterpriseCode
}

func (s *CreateEnterpriseCodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnterpriseCodeRequest) SetClientToken(v string) *CreateEnterpriseCodeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEnterpriseCodeRequest) SetEnterpriseCode(v string) *CreateEnterpriseCodeRequest {
	s.EnterpriseCode = &v
	return s
}

func (s *CreateEnterpriseCodeRequest) SetRegionId(v string) *CreateEnterpriseCodeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnterpriseCodeRequest) Validate() error {
	return dara.Validate(s)
}
