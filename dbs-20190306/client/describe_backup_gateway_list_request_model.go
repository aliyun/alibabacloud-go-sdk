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
	// A client token used to ensure the idempotence of the request. This prevents duplicate requests.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique identifier of the backup gateway. You can query multiple gateways by separating the identifiers with commas (,).
	//
	// example:
	//
	// 7213527653217
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than or equal to 0 and cannot exceed the maximum value of an integer. The default value is 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of records on each page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// > The default value is 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region of the DBS instance. Valid values:
	//
	// - **cn-hangzhou**: China (Hangzhou)
	//
	// - **cn-shanghai**: China (Shanghai)
	//
	// - **cn-qingdao**: China (Qingdao)
	//
	// - **cn-beijing**: China (Beijing)
	//
	// - **cn-shenzhen**: China (Shenzhen)
	//
	// - **cn-hongkong**: China (Hong Kong)
	//
	// - **ap-southeast-1**: Singapore
	//
	// - **cn-hangzhou-finance**: Hangzhou Finance Cloud
	//
	// - **cn-shanghai-finance**: Shanghai Finance Cloud
	//
	// - **cn-shenzhen-finance**: Shenzhen Finance Cloud
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
