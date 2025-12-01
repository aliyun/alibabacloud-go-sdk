// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupGatewayListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeBackupGatewayListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupGatewayListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeBackupGatewayListResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribeBackupGatewayListResponseBodyItems) *DescribeBackupGatewayListResponseBody
	GetItems() *DescribeBackupGatewayListResponseBodyItems
	SetPageNum(v int32) *DescribeBackupGatewayListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeBackupGatewayListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupGatewayListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupGatewayListResponseBody
	GetSuccess() *bool
	SetTotalElements(v int32) *DescribeBackupGatewayListResponseBody
	GetTotalElements() *int32
	SetTotalPages(v int32) *DescribeBackupGatewayListResponseBody
	GetTotalPages() *int32
}

type DescribeBackupGatewayListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of backup gateways.
	Items *DescribeBackupGatewayListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 46361705-8531-492F-807E-A97E482DD4A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of backup gateways.
	//
	// example:
	//
	// 0
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeBackupGatewayListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupGatewayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupGatewayListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupGatewayListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBackupGatewayListResponseBody) GetItems() *DescribeBackupGatewayListResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupGatewayListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeBackupGatewayListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupGatewayListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupGatewayListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupGatewayListResponseBody) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *DescribeBackupGatewayListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeBackupGatewayListResponseBody) SetErrCode(v string) *DescribeBackupGatewayListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetErrMessage(v string) *DescribeBackupGatewayListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetHttpStatusCode(v int32) *DescribeBackupGatewayListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetItems(v *DescribeBackupGatewayListResponseBodyItems) *DescribeBackupGatewayListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetPageNum(v int32) *DescribeBackupGatewayListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetPageSize(v int32) *DescribeBackupGatewayListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetRequestId(v string) *DescribeBackupGatewayListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetSuccess(v bool) *DescribeBackupGatewayListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetTotalElements(v int32) *DescribeBackupGatewayListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetTotalPages(v int32) *DescribeBackupGatewayListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupGatewayListResponseBodyItems struct {
	BackupGateway []*DescribeBackupGatewayListResponseBodyItemsBackupGateway `json:"BackupGateway,omitempty" xml:"BackupGateway,omitempty" type:"Repeated"`
}

func (s DescribeBackupGatewayListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupGatewayListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponseBodyItems) GetBackupGateway() []*DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	return s.BackupGateway
}

func (s *DescribeBackupGatewayListResponseBodyItems) SetBackupGateway(v []*DescribeBackupGatewayListResponseBodyItemsBackupGateway) *DescribeBackupGatewayListResponseBodyItems {
	s.BackupGateway = v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItems) Validate() error {
	if s.BackupGateway != nil {
		for _, item := range s.BackupGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupGatewayListResponseBodyItemsBackupGateway struct {
	// The time when the backup gateway was created, such as 1554560477000.
	//
	// example:
	//
	// 1554560477000
	BackupGatewayCreateTime *int64 `json:"BackupGatewayCreateTime,omitempty" xml:"BackupGatewayCreateTime,omitempty"`
	// The ID of the backup gateway.
	//
	// example:
	//
	// 2321313123
	BackupGatewayId *string `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The status of the backup gateway. Valid values:
	//
	// 	- ONLINE: The backup gateway is online.
	//
	// 	- OFFLINE: The backup gateway is offline.
	//
	// 	- STOPPED: The backup gateway is stopped.
	//
	// 	- UPGRADING: The backup gateway is being upgraded.
	//
	// example:
	//
	// ONLINE
	BackupGatewayStatus *string `json:"BackupGatewayStatus,omitempty" xml:"BackupGatewayStatus,omitempty"`
	// The display name of the backup gateway.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique identifier of the backup gateway.
	//
	// example:
	//
	// sgdsajhdgu
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The last time when a heartbeat message was sent, such as 1554560477000.
	//
	// example:
	//
	// 1554560477000
	LastHeartbeatTime *int64 `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the host on which the backup gateway is installed.
	//
	// example:
	//
	// test
	SourceEndpointHostname *string `json:"SourceEndpointHostname,omitempty" xml:"SourceEndpointHostname,omitempty"`
	// The public IP address of the host on which the backup gateway is installed.
	//
	// example:
	//
	// XX.XX.XX.XX
	SourceEndpointInternetIP *string `json:"SourceEndpointInternetIP,omitempty" xml:"SourceEndpointInternetIP,omitempty"`
	// The private IP address of the host on which the backup gateway is installed.
	//
	// example:
	//
	// XX.XX.XX.XX
	SourceEndpointIntranetIP *string `json:"SourceEndpointIntranetIP,omitempty" xml:"SourceEndpointIntranetIP,omitempty"`
}

func (s DescribeBackupGatewayListResponseBodyItemsBackupGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupGatewayListResponseBodyItemsBackupGateway) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetBackupGatewayCreateTime() *int64 {
	return s.BackupGatewayCreateTime
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetBackupGatewayId() *string {
	return s.BackupGatewayId
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetBackupGatewayStatus() *string {
	return s.BackupGatewayStatus
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetLastHeartbeatTime() *int64 {
	return s.LastHeartbeatTime
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetRegion() *string {
	return s.Region
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetSourceEndpointHostname() *string {
	return s.SourceEndpointHostname
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetSourceEndpointInternetIP() *string {
	return s.SourceEndpointInternetIP
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) GetSourceEndpointIntranetIP() *string {
	return s.SourceEndpointIntranetIP
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetBackupGatewayCreateTime(v int64) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.BackupGatewayCreateTime = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetBackupGatewayId(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetBackupGatewayStatus(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.BackupGatewayStatus = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetDisplayName(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.DisplayName = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetIdentifier(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.Identifier = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetLastHeartbeatTime(v int64) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.LastHeartbeatTime = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetRegion(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.Region = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetSourceEndpointHostname(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.SourceEndpointHostname = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetSourceEndpointInternetIP(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.SourceEndpointInternetIP = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetSourceEndpointIntranetIP(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.SourceEndpointIntranetIP = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) Validate() error {
	return dara.Validate(s)
}
