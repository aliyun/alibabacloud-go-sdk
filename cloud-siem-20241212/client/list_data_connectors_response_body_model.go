// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataConnectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataConnector(v []*ListDataConnectorsResponseBodyDataConnector) *ListDataConnectorsResponseBody
	GetDataConnector() []*ListDataConnectorsResponseBodyDataConnector
	SetMaxResults(v int32) *ListDataConnectorsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataConnectorsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataConnectorsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataConnectorsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataConnectorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataConnectorsResponseBody
	GetTotalCount() *int32
}

type ListDataConnectorsResponseBody struct {
	// The list of collectors.
	DataConnector []*ListDataConnectorsResponseBodyDataConnector `json:"DataConnector,omitempty" xml:"DataConnector,omitempty" type:"Repeated"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataConnectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataConnectorsResponseBody) GetDataConnector() []*ListDataConnectorsResponseBodyDataConnector {
	return s.DataConnector
}

func (s *ListDataConnectorsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataConnectorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataConnectorsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataConnectorsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataConnectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataConnectorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataConnectorsResponseBody) SetDataConnector(v []*ListDataConnectorsResponseBodyDataConnector) *ListDataConnectorsResponseBody {
	s.DataConnector = v
	return s
}

func (s *ListDataConnectorsResponseBody) SetMaxResults(v int32) *ListDataConnectorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataConnectorsResponseBody) SetNextToken(v string) *ListDataConnectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataConnectorsResponseBody) SetPageNumber(v int32) *ListDataConnectorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataConnectorsResponseBody) SetPageSize(v int32) *ListDataConnectorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataConnectorsResponseBody) SetRequestId(v string) *ListDataConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataConnectorsResponseBody) SetTotalCount(v int32) *ListDataConnectorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataConnectorsResponseBody) Validate() error {
	if s.DataConnector != nil {
		for _, item := range s.DataConnector {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataConnectorsResponseBodyDataConnector struct {
	// The configuration item ID of the collector access object in the multi-cloud configuration.
	//
	// example:
	//
	// cn-smbc-prod-cloudtrail.s3.cn-north-1.amazonaws.com.cn_AKIAX4AUG6AYSJO7FAG4
	AuthConfigId *string `json:"AuthConfigId,omitempty" xml:"AuthConfigId,omitempty"`
	// The cloud product of the authentication configuration.
	//
	// example:
	//
	// salesForceRestAPI
	AuthConfigProduct *string `json:"AuthConfigProduct,omitempty" xml:"AuthConfigProduct,omitempty"`
	// The cloud vendor of the authentication configuration.
	//
	// example:
	//
	// SALESFORCE
	AuthConfigVendor *string `json:"AuthConfigVendor,omitempty" xml:"AuthConfigVendor,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-06-01T02:14:24Z
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The collector configuration information.
	//
	// example:
	//
	// {\\"endpoint\\":\\"csrm-sentinel.s3.cn-north-1.amazonaws.com.cn_AKIARGUYUGSX6A6VUXQJ\\",\\"bucket\\":\\"csrm-sentinel\\",\\"prefix\\":\\"AliTest/\\",\\"format\\":{\\"type\\":\\"JSON\\"},\\"encoding\\":\\"UTF-8\\"}
	DataConnectorConfig *string `json:"DataConnectorConfig,omitempty" xml:"DataConnectorConfig,omitempty"`
	// The collector ID.
	//
	// example:
	//
	// dc-07423146117d77db266f78bc41f4fd80
	DataConnectorId *string `json:"DataConnectorId,omitempty" xml:"DataConnectorId,omitempty"`
	// The data connector name.
	//
	// example:
	//
	// dc-***
	DataConnectorName *string `json:"DataConnectorName,omitempty" xml:"DataConnectorName,omitempty"`
	// The connector status.
	//
	// example:
	//
	// enable
	DataConnectorStatus *string `json:"DataConnectorStatus,omitempty" xml:"DataConnectorStatus,omitempty"`
	// The connector type.
	//
	// example:
	//
	// oss
	DataConnectorType *string `json:"DataConnectorType,omitempty" xml:"DataConnectorType,omitempty"`
	// The destination data source ID. This parameter is required only for synchronization.
	//
	// example:
	//
	// ds-t3ywipile7gctobaunx0
	DestDataSourceId *string `json:"DestDataSourceId,omitempty" xml:"DestDataSourceId,omitempty"`
	// The Simple Log Service project name.
	//
	// example:
	//
	// aliyun-cloudsiem-channel-1371069058301795-cn-shanghai
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// The log storage region ID.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The Simple Log Service Logstore name.
	//
	// example:
	//
	// ls-wecom-crontasklog-prod
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the SLS data import job associated with the collector.
	//
	// example:
	//
	// 1a2b3c4d5e6f7a8b9c0d1e2f3a4b****
	SlsIngestionJobName *string `json:"SlsIngestionJobName,omitempty" xml:"SlsIngestionJobName,omitempty"`
	// The status of the SLS data import job associated with the collector.
	//
	// example:
	//
	// running
	SlsIngestionJobState *string `json:"SlsIngestionJobState,omitempty" xml:"SlsIngestionJobState,omitempty"`
	// The source data type.
	//
	// example:
	//
	// s3
	SrcDataType *string `json:"SrcDataType,omitempty" xml:"SrcDataType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-01-04 22:31:54
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataConnectorsResponseBodyDataConnector) String() string {
	return dara.Prettify(s)
}

func (s ListDataConnectorsResponseBodyDataConnector) GoString() string {
	return s.String()
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetAuthConfigId() *string {
	return s.AuthConfigId
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetAuthConfigProduct() *string {
	return s.AuthConfigProduct
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetAuthConfigVendor() *string {
	return s.AuthConfigVendor
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetDataConnectorConfig() *string {
	return s.DataConnectorConfig
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetDataConnectorId() *string {
	return s.DataConnectorId
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetDataConnectorName() *string {
	return s.DataConnectorName
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetDataConnectorStatus() *string {
	return s.DataConnectorStatus
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetDataConnectorType() *string {
	return s.DataConnectorType
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetDestDataSourceId() *string {
	return s.DestDataSourceId
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetSlsIngestionJobName() *string {
	return s.SlsIngestionJobName
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetSlsIngestionJobState() *string {
	return s.SlsIngestionJobState
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetSrcDataType() *string {
	return s.SrcDataType
}

func (s *ListDataConnectorsResponseBodyDataConnector) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetAuthConfigId(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.AuthConfigId = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetAuthConfigProduct(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.AuthConfigProduct = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetAuthConfigVendor(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.AuthConfigVendor = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetCreationTime(v int64) *ListDataConnectorsResponseBodyDataConnector {
	s.CreationTime = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetDataConnectorConfig(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.DataConnectorConfig = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetDataConnectorId(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.DataConnectorId = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetDataConnectorName(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.DataConnectorName = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetDataConnectorStatus(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.DataConnectorStatus = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetDataConnectorType(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.DataConnectorType = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetDestDataSourceId(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.DestDataSourceId = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetLogProjectName(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.LogProjectName = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetLogRegionId(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.LogRegionId = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetLogStoreName(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.LogStoreName = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetSlsIngestionJobName(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.SlsIngestionJobName = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetSlsIngestionJobState(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.SlsIngestionJobState = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetSrcDataType(v string) *ListDataConnectorsResponseBodyDataConnector {
	s.SrcDataType = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) SetUpdateTime(v int64) *ListDataConnectorsResponseBodyDataConnector {
	s.UpdateTime = &v
	return s
}

func (s *ListDataConnectorsResponseBodyDataConnector) Validate() error {
	return dara.Validate(s)
}
