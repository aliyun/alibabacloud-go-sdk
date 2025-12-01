// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupGatewayListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeBackupGatewayListRequest
	GetClientToken() *string
	SetIdentifier(v string) *DescribeBackupGatewayListRequest
	GetIdentifier() *string
	SetOwnerId(v string) *DescribeBackupGatewayListRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeBackupGatewayListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeBackupGatewayListRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeBackupGatewayListRequest
	GetRegion() *string
}

type DescribeBackupGatewayListRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique identifier of the backup gateway. You can query multiple backup gateways. Separate multiple identifiers with commas (,).
	//
	// example:
	//
	// 7213527653217
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be a positive integer. Default value: 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// > Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which Database Backup (DBS) is activated. Valid values:
	//
	// 	- **cn-hangzhou**: China (Hangzhou)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-qingdao**: China (Qingdao)
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// 	- **ap-southeast-1**: Singapore (Singapore)
	//
	// 	- **cn-hangzhou-finance**: China East 1 Finance
	//
	// 	- **cn-shanghai-finance**: China East 2 Finance
	//
	// 	- **cn-shenzhen-finance**: China South 1 Finance
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeBackupGatewayListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupGatewayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeBackupGatewayListRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeBackupGatewayListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeBackupGatewayListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeBackupGatewayListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupGatewayListRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeBackupGatewayListRequest) SetClientToken(v string) *DescribeBackupGatewayListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetIdentifier(v string) *DescribeBackupGatewayListRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetOwnerId(v string) *DescribeBackupGatewayListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetPageNum(v int32) *DescribeBackupGatewayListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetPageSize(v int32) *DescribeBackupGatewayListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetRegion(v string) *DescribeBackupGatewayListRequest {
	s.Region = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) Validate() error {
	return dara.Validate(s)
}
