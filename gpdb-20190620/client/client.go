// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeDBInstanceForDmsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port         *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeDBInstanceForDmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceForDmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceForDmsRequest) SetDBInstanceId(v string) *DescribeDBInstanceForDmsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceForDmsRequest) SetHost(v string) *DescribeDBInstanceForDmsRequest {
	s.Host = &v
	return s
}

func (s *DescribeDBInstanceForDmsRequest) SetOwnerId(v int64) *DescribeDBInstanceForDmsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceForDmsRequest) SetPort(v int64) *DescribeDBInstanceForDmsRequest {
	s.Port = &v
	return s
}

type DescribeDBInstanceForDmsResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Count          *int64                                        `json:"Count,omitempty" xml:"Count,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Instance       *DescribeDBInstanceForDmsResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBInstanceForDmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceForDmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceForDmsResponseBody) SetCode(v string) *DescribeDBInstanceForDmsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBody) SetCount(v int64) *DescribeDBInstanceForDmsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBody) SetHttpStatusCode(v int32) *DescribeDBInstanceForDmsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBody) SetInstance(v *DescribeDBInstanceForDmsResponseBodyInstance) *DescribeDBInstanceForDmsResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBody) SetMessage(v string) *DescribeDBInstanceForDmsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBody) SetRequestId(v string) *DescribeDBInstanceForDmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBody) SetSuccess(v bool) *DescribeDBInstanceForDmsResponseBody {
	s.Success = &v
	return s
}

type DescribeDBInstanceForDmsResponseBodyInstance struct {
	AliUid              *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid                 *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	ConnectionString    *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DbInstanceName      *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	DbType              *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	Port                *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Region              *string `json:"Region,omitempty" xml:"Region,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcCloudInstanceId  *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcIp               *string `json:"VpcIp,omitempty" xml:"VpcIp,omitempty"`
}

func (s DescribeDBInstanceForDmsResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceForDmsResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetAliUid(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetBid(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.Bid = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetConnectionString(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetDbInstanceName(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetDbType(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.DbType = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetDescription(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetInstanceNetworkType(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetPort(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetRegion(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.Region = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetVSwitchId(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetVersion(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.Version = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetVpcCloudInstanceId(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetVpcId(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponseBodyInstance) SetVpcIp(v string) *DescribeDBInstanceForDmsResponseBodyInstance {
	s.VpcIp = &v
	return s
}

type DescribeDBInstanceForDmsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceForDmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceForDmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceForDmsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceForDmsResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceForDmsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceForDmsResponse) SetStatusCode(v int32) *DescribeDBInstanceForDmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceForDmsResponse) SetBody(v *DescribeDBInstanceForDmsResponseBody) *DescribeDBInstanceForDmsResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceSecurityIpsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDBInstanceSecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSecurityIpsRequest) SetInstanceId(v string) *DescribeDBInstanceSecurityIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsRequest) SetOwnerId(v int64) *DescribeDBInstanceSecurityIpsRequest {
	s.OwnerId = &v
	return s
}

type DescribeDBInstanceSecurityIpsResponseBody struct {
	Code           *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Count          *int64                                             `json:"Count,omitempty" xml:"Count,omitempty"`
	HttpStatusCode *int32                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result         []*DescribeDBInstanceSecurityIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success        *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBInstanceSecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSecurityIpsResponseBody) SetCode(v string) *DescribeDBInstanceSecurityIpsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponseBody) SetCount(v int64) *DescribeDBInstanceSecurityIpsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponseBody) SetHttpStatusCode(v int32) *DescribeDBInstanceSecurityIpsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponseBody) SetMessage(v string) *DescribeDBInstanceSecurityIpsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponseBody) SetRequestId(v string) *DescribeDBInstanceSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponseBody) SetResult(v []*DescribeDBInstanceSecurityIpsResponseBodyResult) *DescribeDBInstanceSecurityIpsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponseBody) SetSuccess(v bool) *DescribeDBInstanceSecurityIpsResponseBody {
	s.Success = &v
	return s
}

type DescribeDBInstanceSecurityIpsResponseBodyResult struct {
	GroupName *string   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceSecurityIpsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSecurityIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSecurityIpsResponseBodyResult) SetGroupName(v string) *DescribeDBInstanceSecurityIpsResponseBodyResult {
	s.GroupName = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponseBodyResult) SetWhiteList(v []*string) *DescribeDBInstanceSecurityIpsResponseBodyResult {
	s.WhiteList = v
	return s
}

type DescribeDBInstanceSecurityIpsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceSecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSecurityIpsResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponse) SetStatusCode(v int32) *DescribeDBInstanceSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSecurityIpsResponse) SetBody(v *DescribeDBInstanceSecurityIpsResponseBody) *DescribeDBInstanceSecurityIpsResponse {
	s.Body = v
	return s
}

type DescribeDBInstancesForDmsRequest struct {
	AliUid  *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDBInstancesForDmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesForDmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForDmsRequest) SetAliUid(v int64) *DescribeDBInstancesForDmsRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstancesForDmsRequest) SetOwnerId(v int64) *DescribeDBInstancesForDmsRequest {
	s.OwnerId = &v
	return s
}

type DescribeDBInstancesForDmsResponseBody struct {
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Count          *int64                                            `json:"Count,omitempty" xml:"Count,omitempty"`
	HttpStatusCode *int32                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Instances      []*DescribeDBInstancesForDmsResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBInstancesForDmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesForDmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForDmsResponseBody) SetCode(v string) *DescribeDBInstancesForDmsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBody) SetCount(v int64) *DescribeDBInstancesForDmsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBody) SetHttpStatusCode(v int32) *DescribeDBInstancesForDmsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBody) SetInstances(v []*DescribeDBInstancesForDmsResponseBodyInstances) *DescribeDBInstancesForDmsResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBody) SetMessage(v string) *DescribeDBInstancesForDmsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBody) SetRequestId(v string) *DescribeDBInstancesForDmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBody) SetSuccess(v bool) *DescribeDBInstancesForDmsResponseBody {
	s.Success = &v
	return s
}

type DescribeDBInstancesForDmsResponseBodyInstances struct {
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// BID。
	Bid                 *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	ConnectionString    *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DbInstanceName      *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	DbType              *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	Port                *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Region              *string `json:"Region,omitempty" xml:"Region,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcCloudInstanceId  *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcIp               *string `json:"VpcIp,omitempty" xml:"VpcIp,omitempty"`
}

func (s DescribeDBInstancesForDmsResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesForDmsResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetAliUid(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetBid(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.Bid = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetConnectionString(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetDbInstanceName(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetDbType(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.DbType = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetDescription(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetInstanceNetworkType(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetPort(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.Port = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetRegion(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.Region = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetVSwitchId(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetVersion(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.Version = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetVpcCloudInstanceId(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetVpcId(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponseBodyInstances) SetVpcIp(v string) *DescribeDBInstancesForDmsResponseBodyInstances {
	s.VpcIp = &v
	return s
}

type DescribeDBInstancesForDmsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesForDmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancesForDmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesForDmsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForDmsResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesForDmsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesForDmsResponse) SetStatusCode(v int32) *DescribeDBInstancesForDmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesForDmsResponse) SetBody(v *DescribeDBInstancesForDmsResponseBody) *DescribeDBInstancesForDmsResponse {
	s.Body = v
	return s
}

type ModifyDBInstanceSecurityIpsRequest struct {
	AliUid     *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	WhileList  *string `json:"WhileList,omitempty" xml:"WhileList,omitempty"`
}

func (s ModifyDBInstanceSecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSecurityIpsRequest) SetAliUid(v int64) *ModifyDBInstanceSecurityIpsRequest {
	s.AliUid = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsRequest) SetGroupName(v string) *ModifyDBInstanceSecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsRequest) SetInstanceId(v string) *ModifyDBInstanceSecurityIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsRequest) SetOwnerId(v int64) *ModifyDBInstanceSecurityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsRequest) SetWhileList(v string) *ModifyDBInstanceSecurityIpsRequest {
	s.WhileList = &v
	return s
}

type ModifyDBInstanceSecurityIpsResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDBInstanceSecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSecurityIpsResponseBody) SetCode(v string) *ModifyDBInstanceSecurityIpsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsResponseBody) SetHttpStatusCode(v int32) *ModifyDBInstanceSecurityIpsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsResponseBody) SetMessage(v string) *ModifyDBInstanceSecurityIpsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsResponseBody) SetRequestId(v string) *ModifyDBInstanceSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsResponseBody) SetSuccess(v bool) *ModifyDBInstanceSecurityIpsResponseBody {
	s.Success = &v
	return s
}

type ModifyDBInstanceSecurityIpsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceSecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSecurityIpsResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSecurityIpsResponse) SetStatusCode(v int32) *ModifyDBInstanceSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceSecurityIpsResponse) SetBody(v *ModifyDBInstanceSecurityIpsResponseBody) *ModifyDBInstanceSecurityIpsResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":            tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou":           tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai":           tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen":           tea.String("gpdb.aliyuncs.com"),
		"cn-hongkong":           tea.String("gpdb.aliyuncs.com"),
		"ap-southeast-1":        tea.String("gpdb.aliyuncs.com"),
		"us-west-1":             tea.String("gpdb.aliyuncs.com"),
		"us-east-1":             tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-qingdao":            tea.String("gpdb.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("gpdb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("gpdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceForDmsWithOptions(request *DescribeDBInstanceForDmsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceForDmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceForDms"),
		Version:     tea.String("2019-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceForDmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceForDms(request *DescribeDBInstanceForDmsRequest) (_result *DescribeDBInstanceForDmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceForDmsResponse{}
	_body, _err := client.DescribeDBInstanceForDmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceSecurityIpsWithOptions(request *DescribeDBInstanceSecurityIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceSecurityIps"),
		Version:     tea.String("2019-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceSecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceSecurityIps(request *DescribeDBInstanceSecurityIpsRequest) (_result *DescribeDBInstanceSecurityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceSecurityIpsResponse{}
	_body, _err := client.DescribeDBInstanceSecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancesForDmsWithOptions(request *DescribeDBInstancesForDmsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesForDmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstancesForDms"),
		Version:     tea.String("2019-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstancesForDmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstancesForDms(request *DescribeDBInstancesForDmsRequest) (_result *DescribeDBInstancesForDmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancesForDmsResponse{}
	_body, _err := client.DescribeDBInstancesForDmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceSecurityIpsWithOptions(request *ModifyDBInstanceSecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceSecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.WhileList)) {
		query["WhileList"] = request.WhileList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBInstanceSecurityIps"),
		Version:     tea.String("2019-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBInstanceSecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceSecurityIps(request *ModifyDBInstanceSecurityIpsRequest) (_result *ModifyDBInstanceSecurityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceSecurityIpsResponse{}
	_body, _err := client.ModifyDBInstanceSecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
