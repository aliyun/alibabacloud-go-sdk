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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enterprise code.
	//
	// The enterprise code must be 5 characters in length and contain both letters and digits. The letters can be uppercase or lowercase. The enterprise code must be globally unique and cannot be the same as that of another enterprise.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12P**
	EnterpriseCode *string `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	// The region ID.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) to query the regions supported by Smart Access Gateway and the corresponding region IDs.
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
