// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeBgpPackByIpRequest struct {
	// The region ID of the asset to query.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The IP address of the asset to query.
	//
	// >  You can call the [DescribeInstanceIpAddress](https://help.aliyun.com/document_detail/472620.html) operation to query the IDs of Elastic Compute Service (ECS) instances, Server Load Balancer (SLB) instances, and elastic IP addresses (EIPs) within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 118.31.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBgpPackByIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpRequest) SetDdosRegionId(v string) *DescribeBgpPackByIpRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeBgpPackByIpRequest) SetIp(v string) *DescribeBgpPackByIpRequest {
	s.Ip = &v
	return s
}

type DescribeBgpPackByIpResponseBody struct {
	// The HTTP status code of the request.
	//
	// For more information about status codes, see [Common parameters](https://help.aliyun.com/document_detail/118841.html).
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configurations of the instance that is associated with the asset.
	DdosbgpInfo *DescribeBgpPackByIpResponseBodyDdosbgpInfo `json:"DdosbgpInfo,omitempty" xml:"DdosbgpInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E46A08E4-A1CD-5BE9-B580-C4D6E9BC5484
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBgpPackByIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponseBody) SetCode(v int32) *DescribeBgpPackByIpResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) SetDdosbgpInfo(v *DescribeBgpPackByIpResponseBodyDdosbgpInfo) *DescribeBgpPackByIpResponseBody {
	s.DdosbgpInfo = v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) SetRequestId(v string) *DescribeBgpPackByIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) SetSuccess(v bool) *DescribeBgpPackByIpResponseBody {
	s.Success = &v
	return s
}

type DescribeBgpPackByIpResponseBodyDdosbgpInfo struct {
	// The basic protection threshold of the instance. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	BaseThreshold *int32 `json:"BaseThreshold,omitempty" xml:"BaseThreshold,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	DdosbgpInstanceId *string `json:"DdosbgpInstanceId,omitempty" xml:"DdosbgpInstanceId,omitempty"`
	// The burstable protection threshold of the instance. Unit: Gbit/s.
	//
	// example:
	//
	// 301
	ElasticThreshold *int32 `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	// The expiration time of the instance. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640448000000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 118.31.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetBaseThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.BaseThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetDdosbgpInstanceId(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.DdosbgpInstanceId = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetElasticThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetExpireTime(v int64) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetIp(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.Ip = &v
	return s
}

type DescribeBgpPackByIpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBgpPackByIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBgpPackByIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponse) SetHeaders(v map[string]*string) *DescribeBgpPackByIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpPackByIpResponse) SetStatusCode(v int32) *DescribeBgpPackByIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBgpPackByIpResponse) SetBody(v *DescribeBgpPackByIpResponseBody) *DescribeBgpPackByIpResponse {
	s.Body = v
	return s
}

type DescribeCapRequest struct {
	// The start time of the DDoS attack event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > You can call the [DescribeDdosEventList](https://help.aliyun.com/document_detail/354236.html) operation to query the start time of each DDoS attack event that occurred on an asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1637812279000
	BegTime *int64 `json:"BegTime,omitempty" xml:"BegTime,omitempty"`
	// The region ID of the asset that is under DDoS attacks. The asset is assigned a public IP address.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The ID of the asset that is under DDoS attacks.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp10bclrt56fblts****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset that is under DDoS attacks. The asset is assigned a public IP address. Valid values:
	//
	// 	- **ecs**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **slb**: a Server Load Balancer (SLB) instance.
	//
	// 	- **eip**: an elastic IP address (EIP).
	//
	// 	- **ipv6**: an IPv6 gateway.
	//
	// 	- **swas**: a simple application server.
	//
	// 	- **waf**: a Web Application Firewall (WAF) instance of the Exclusive edition.
	//
	// 	- **ga_basic**: a Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public IP address of the asset that is under DDoS attacks.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeCapRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapRequest) SetBegTime(v int64) *DescribeCapRequest {
	s.BegTime = &v
	return s
}

func (s *DescribeCapRequest) SetDdosRegionId(v string) *DescribeCapRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceId(v string) *DescribeCapRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceType(v string) *DescribeCapRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeCapRequest) SetInternetIp(v string) *DescribeCapRequest {
	s.InternetIp = &v
	return s
}

type DescribeCapResponseBody struct {
	// The download link to the traffic data that is captured when a DDoS attack event occurs.
	CapUrl *DescribeCapResponseBodyCapUrl `json:"CapUrl,omitempty" xml:"CapUrl,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C728D7E9-9A39-52E0-966B-5C33118BDBB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapResponseBody) SetCapUrl(v *DescribeCapResponseBodyCapUrl) *DescribeCapResponseBody {
	s.CapUrl = v
	return s
}

func (s *DescribeCapResponseBody) SetRequestId(v string) *DescribeCapResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCapResponseBodyCapUrl struct {
	// The download link to the traffic data.
	//
	// example:
	//
	// http://beaver-pack****.oss-cn-hangzhou.aliyuncs.com/ddos-2021112511-121.89.XX.XX.cap?Expires=1637824408&OSSAccessKeyId=LTAIXu2lJhw3****&Signature=KKSzOMSUajtwcqfqxkU1nK69d4****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeCapResponseBodyCapUrl) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapResponseBodyCapUrl) GoString() string {
	return s.String()
}

func (s *DescribeCapResponseBodyCapUrl) SetUrl(v string) *DescribeCapResponseBodyCapUrl {
	s.Url = &v
	return s
}

type DescribeCapResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCapResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapResponse) SetHeaders(v map[string]*string) *DescribeCapResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapResponse) SetStatusCode(v int32) *DescribeCapResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapResponse) SetBody(v *DescribeCapResponseBody) *DescribeCapResponse {
	s.Body = v
	return s
}

type DescribeDdosCountRequest struct {
	// The region ID of the asset to query.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The type of the asset to query. Valid values:
	//
	// 	- **ecs**: Elastic Compute Service (ECS) instances.
	//
	// 	- **slb**: Server Load Balancer (SLB) instances.
	//
	// 	- **eip**: elastic IP addresses (EIPs).
	//
	// 	- **ipv6**: IPv6 gateways.
	//
	// 	- **swas**: simple application servers.
	//
	// 	- **waf**: Web Application Firewall (WAF) instances of the Exclusive edition.
	//
	// 	- **ga_basic**: Global Accelerator (GA) instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDdosCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountRequest) SetDdosRegionId(v string) *DescribeDdosCountRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosCountRequest) SetInstanceType(v string) *DescribeDdosCountRequest {
	s.InstanceType = &v
	return s
}

type DescribeDdosCountResponseBody struct {
	// The number of assets that are under DDoS attacks.
	DdosCount *DescribeDdosCountResponseBodyDdosCount `json:"DdosCount,omitempty" xml:"DdosCount,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7D66C762-324E-51CE-B461-2007996087B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDdosCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponseBody) SetDdosCount(v *DescribeDdosCountResponseBodyDdosCount) *DescribeDdosCountResponseBody {
	s.DdosCount = v
	return s
}

func (s *DescribeDdosCountResponseBody) SetRequestId(v string) *DescribeDdosCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDdosCountResponseBodyDdosCount struct {
	// The number of assets for which blackhole filtering is triggered.
	//
	// example:
	//
	// 0
	BlackholeCount *int32 `json:"BlackholeCount,omitempty" xml:"BlackholeCount,omitempty"`
	// The number of assets for which traffic scrubbing is triggered.
	//
	// example:
	//
	// 4
	DefenseCount *int32 `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	// The total number of assets.
	//
	// example:
	//
	// 0
	InstacenCount *int32 `json:"InstacenCount,omitempty" xml:"InstacenCount,omitempty"`
}

func (s DescribeDdosCountResponseBodyDdosCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountResponseBodyDdosCount) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetBlackholeCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.BlackholeCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetDefenseCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.DefenseCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetInstacenCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.InstacenCount = &v
	return s
}

type DescribeDdosCountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponse) SetHeaders(v map[string]*string) *DescribeDdosCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosCountResponse) SetStatusCode(v int32) *DescribeDdosCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosCountResponse) SetBody(v *DescribeDdosCountResponseBody) *DescribeDdosCountResponse {
	s.Body = v
	return s
}

type DescribeDdosCreditRequest struct {
	// The ID of the region.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DescribeDdosCreditRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditRequest) SetDdosRegionId(v string) *DescribeDdosCreditRequest {
	s.DdosRegionId = &v
	return s
}

type DescribeDdosCreditResponseBody struct {
	// The details of the security credit score of the current Alibaba Cloud account in the specified region.
	DdosCredit *DescribeDdosCreditResponseBodyDdosCredit `json:"DdosCredit,omitempty" xml:"DdosCredit,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E1F7BD73-8E9D-58D9-8658-CFC97112C641
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDdosCreditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBody) SetDdosCredit(v *DescribeDdosCreditResponseBodyDdosCredit) *DescribeDdosCreditResponseBody {
	s.DdosCredit = v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetRequestId(v string) *DescribeDdosCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetSuccess(v bool) *DescribeDdosCreditResponseBody {
	s.Success = &v
	return s
}

type DescribeDdosCreditResponseBodyDdosCredit struct {
	// The time period after which blackhole filtering is automatically deactivated in the specified region. Unit: minutes.
	//
	// example:
	//
	// 150
	BlackholeTime *int32 `json:"BlackholeTime,omitempty" xml:"BlackholeTime,omitempty"`
	// The security credit score. The full score is **1000**.
	//
	// example:
	//
	// 550
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The security credit level. Valid values:
	//
	// 	- **A**: outstanding
	//
	// 	- **B**: excellent
	//
	// 	- **C**: good
	//
	// 	- **D**: average
	//
	// 	- **E**: poor
	//
	// 	- **F**: poorer
	//
	// example:
	//
	// D
	ScoreLevel *string `json:"ScoreLevel,omitempty" xml:"ScoreLevel,omitempty"`
}

func (s DescribeDdosCreditResponseBodyDdosCredit) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponseBodyDdosCredit) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetBlackholeTime(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.BlackholeTime = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScore(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.Score = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScoreLevel(v string) *DescribeDdosCreditResponseBodyDdosCredit {
	s.ScoreLevel = &v
	return s
}

type DescribeDdosCreditResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosCreditResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponse) SetHeaders(v map[string]*string) *DescribeDdosCreditResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosCreditResponse) SetStatusCode(v int32) *DescribeDdosCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosCreditResponse) SetBody(v *DescribeDdosCreditResponseBody) *DescribeDdosCreditResponse {
	s.Body = v
	return s
}

type DescribeDdosEventListRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The region ID of the asset to query.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The ID of asset to query.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp10bclrt56fblts****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset to query. Valid values:
	//
	// 	- **ecs**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **slb**: a Server Load Balancer (SLB) instance.
	//
	// 	- **eip**: an elastic IP address (EIP).
	//
	// 	- **ipv6**: an IPv6 gateway.
	//
	// 	- **swas**: a simple application server.
	//
	// 	- **waf**: a Web Application Firewall (WAF) instance of the Exclusive edition.
	//
	// 	- **ga_basic**: a Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset to query.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDdosEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListRequest) SetCurrentPage(v int32) *DescribeDdosEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetDdosRegionId(v string) *DescribeDdosEventListRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceId(v string) *DescribeDdosEventListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceType(v string) *DescribeDdosEventListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInternetIp(v string) *DescribeDdosEventListRequest {
	s.InternetIp = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetPageSize(v int32) *DescribeDdosEventListRequest {
	s.PageSize = &v
	return s
}

type DescribeDdosEventListResponseBody struct {
	// The details of the DDoS attack events.
	DdosEventList *DescribeDdosEventListResponseBodyDdosEventList `json:"DdosEventList,omitempty" xml:"DdosEventList,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0907F8-A9F3-5E11-977B-D59CD98C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBody) SetDdosEventList(v *DescribeDdosEventListResponseBodyDdosEventList) *DescribeDdosEventListResponseBody {
	s.DdosEventList = v
	return s
}

func (s *DescribeDdosEventListResponseBody) SetRequestId(v string) *DescribeDdosEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventListResponseBody) SetTotal(v int32) *DescribeDdosEventListResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosEventListResponseBodyDdosEventList struct {
	DdosEvent []*DescribeDdosEventListResponseBodyDdosEventListDdosEvent `json:"DdosEvent,omitempty" xml:"DdosEvent,omitempty" type:"Repeated"`
}

func (s DescribeDdosEventListResponseBodyDdosEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBodyDdosEventList) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetDdosEvent(v []*DescribeDdosEventListResponseBodyDdosEventListDdosEvent) *DescribeDdosEventListResponseBodyDdosEventList {
	s.DdosEvent = v
	return s
}

type DescribeDdosEventListResponseBodyDdosEventListDdosEvent struct {
	// The status of the DDoS attack event. Valid values:
	//
	// 	- **mitigating**: indicates that traffic scrubbing is in progress.
	//
	// 	- **blackholed**: indicates that blackhole filtering is triggered for the asset.
	//
	// 	- **normal**: indicates that the DDoS attack event ends.
	//
	// example:
	//
	// normal
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The type of the DDoS attack event. Valid values:
	//
	// 	- **defense**: an attack event that triggers traffic scrubbing
	//
	// 	- **blackhole**: an attack event that triggers blackhole filtering
	//
	// example:
	//
	// blackhole
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The time of the last attack. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > This parameter is returned only when the asset is attacked multiple times within a DDoS attack event.
	//
	// example:
	//
	// 1637817679000
	DelayTime *int64 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The end time of the DDoS attack event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637817679000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the DDoS attack event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637812279000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when blackhole filtering is deactivated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **blackhole**.
	//
	// example:
	//
	// 1637814079000
	UnBlackholeTime *int64 `json:"UnBlackholeTime,omitempty" xml:"UnBlackholeTime,omitempty"`
}

func (s DescribeDdosEventListResponseBodyDdosEventListDdosEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDdosStatus(v string) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DdosStatus = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDdosType(v string) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDelayTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DelayTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetEndTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetStartTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetUnBlackholeTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.UnBlackholeTime = &v
	return s
}

type DescribeDdosEventListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponse) SetHeaders(v map[string]*string) *DescribeDdosEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventListResponse) SetStatusCode(v int32) *DescribeDdosEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventListResponse) SetBody(v *DescribeDdosEventListResponseBody) *DescribeDdosEventListResponse {
	s.Body = v
	return s
}

type DescribeDdosThresholdRequest struct {
	// The region ID of the asset.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The type of the threshold. Valid values:
	//
	// 	- **defense**: traffic scrubbing threshold
	//
	// 	- **blackhole**: DDoS mitigation threshold
	//
	// This parameter is required.
	//
	// example:
	//
	// defense
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The ID of asset N to query.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The type of the asset that is assigned a public IP address. Valid values:
	//
	// 	- **ecs**: ECS instances.
	//
	// 	- **slb**: SLB instances.
	//
	// 	- **eip**: EIPs.
	//
	// 	- **ipv6**: IPv6 gateways.
	//
	// 	- **swas**: simple application servers.
	//
	// 	- **waf**: Web Application Firewall (WAF) instances of the Exclusive edition.
	//
	// 	- **ga_basic**: Global Accelerator (GA) instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDdosThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdRequest) SetDdosRegionId(v string) *DescribeDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetDdosType(v string) *DescribeDdosThresholdRequest {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceIds(v []*string) *DescribeDdosThresholdRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceType(v string) *DescribeDdosThresholdRequest {
	s.InstanceType = &v
	return s
}

type DescribeDdosThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E9B3C090-55AD-59C6-979E-FCFD81E7D9E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the threshold.
	Thresholds *DescribeDdosThresholdResponseBodyThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
}

func (s DescribeDdosThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBody) SetRequestId(v string) *DescribeDdosThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosThresholdResponseBody) SetThresholds(v *DescribeDdosThresholdResponseBodyThresholds) *DescribeDdosThresholdResponseBody {
	s.Thresholds = v
	return s
}

type DescribeDdosThresholdResponseBodyThresholds struct {
	Threshold []*DescribeDdosThresholdResponseBodyThresholdsThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Repeated"`
}

func (s DescribeDdosThresholdResponseBodyThresholds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponseBodyThresholds) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetThreshold(v []*DescribeDdosThresholdResponseBodyThresholdsThreshold) *DescribeDdosThresholdResponseBodyThresholds {
	s.Threshold = v
	return s
}

type DescribeDdosThresholdResponseBodyThresholdsThreshold struct {
	// If the value of the **DdosType*	- parameter is **defense**, the Bps parameter indicates the current traffic scrubbing threshold. Unit: Mbit/s.
	//
	// If the value of the **DdosType*	- parameter is **blackhole**, the Bps parameter indicates the basic protection threshold. Unit: Mbit/s.
	//
	// example:
	//
	// 500
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The type of the threshold. Valid values:
	//
	// 	- **defense**: traffic scrubbing threshold
	//
	// 	- **blackhole**: DDoS mitigation threshold
	//
	// example:
	//
	// defense
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The burstable protection threshold (the maximum DDoS mitigation threshold). Unit: Mbit/s.
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **blackhole**.
	//
	// example:
	//
	// 12310
	ElasticBps *int32 `json:"ElasticBps,omitempty" xml:"ElasticBps,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp10bclrt56fblts****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Indicates whether the threshold is automatically adjusted. Valid values:
	//
	// 	- **true**: The scrubbing thresholds are automatically adjusted based on the traffic load on the asset.
	//
	// 	- **false**: The scrubbing thresholds are not automatically adjusted. You must manually specify the scrubbing thresholds.
	//
	// example:
	//
	// false
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	// The maximum traffic scrubbing threshold. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	MaxBps *int32 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The maximum packet scrubbing threshold. Unit: pps.
	//
	// example:
	//
	// 150000
	MaxPps *int32 `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
	// The packet scrubbing threshold. Unit: pps.
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **defense**.
	//
	// example:
	//
	// 150000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeDdosThresholdResponseBodyThresholdsThreshold) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponseBodyThresholdsThreshold) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.Bps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetDdosType(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetElasticBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.ElasticBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetInstanceId(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetInternetIp(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.InternetIp = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetIsAuto(v bool) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.IsAuto = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetMaxBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.MaxBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetMaxPps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.MaxPps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetPps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.Pps = &v
	return s
}

type DescribeDdosThresholdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponse) SetHeaders(v map[string]*string) *DescribeDdosThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosThresholdResponse) SetStatusCode(v int32) *DescribeDdosThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosThresholdResponse) SetBody(v *DescribeDdosThresholdResponseBody) *DescribeDdosThresholdResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The region ID of the asset.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **mitigating**: queries assets for which traffic scrubbing is triggered.
	//
	// 	- **blackholed**: queries assets for which blackhole filtering is triggered.
	//
	// 	- **normal**: queries assets that are not under DDoS attacks.
	//
	// example:
	//
	// blackholed
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The ID of the asset. The formats of asset IDs vary based on the value of the **InstanceType**. parameter.
	//
	// 	- If you set **InstanceType*	- to **ecs**, specify the ID of the ECS instance. For example, you can specify i-bp1cb6x80tfgocid\\*\\*\\*\\*.
	//
	// 	- If you set **InstanceType*	- to **slb**, specify the ID of the SLB instance. For example, you can specify alb-vn2dqg3v31y2vd\\*\\*\\*\\*.
	//
	// 	- If you set **InstanceType*	- to **eip**, specify the ID of the EIP. For example, you can specify eip-j6ce6dcx9epi7rs46\\*\\*\\*\\*.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 121.199.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// launch-advisor-2022****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the asset to query. Valid values:
	//
	// 	- **ecs**: ECS instances.
	//
	// 	- **slb**: SLB instances.
	//
	// 	- **eip**: EIPs.
	//
	// 	- **ipv6**: IPv6 gateways.
	//
	// 	- **swas**: simple application servers.
	//
	// 	- **waf**: Web Application Firewall (WAF) instances of the Exclusive edition.
	//
	// 	- **ga_basic**: Global Accelerator (GA) instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetCurrentPage(v int32) *DescribeInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceRequest) SetDdosRegionId(v string) *DescribeInstanceRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeInstanceRequest) SetDdosStatus(v string) *DescribeInstanceRequest {
	s.DdosStatus = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceIp(v string) *DescribeInstanceRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceName(v string) *DescribeInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceType(v string) *DescribeInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceRequest) SetPageSize(v int32) *DescribeInstanceRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceResponseBody struct {
	// The details of the assets.
	InstanceList *DescribeInstanceResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C728D7E9-9A39-52E0-966B-5C33118BDBB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetInstanceList(v *DescribeInstanceResponseBodyInstanceList) *DescribeInstanceResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTotal(v int32) *DescribeInstanceResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceResponseBodyInstanceList struct {
	Instance []*DescribeInstanceResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceList) SetInstance(v []*DescribeInstanceResponseBodyInstanceListInstance) *DescribeInstanceResponseBodyInstanceList {
	s.Instance = v
	return s
}

type DescribeInstanceResponseBodyInstanceListInstance struct {
	// The basic protection threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 5200
	BlackholeThreshold *int32 `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	// The traffic scrubbing threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 300
	DefenseBpsThreshold *int32 `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	// The packet scrubbing threshold for the asset. Unit: packets per second (pps).
	//
	// example:
	//
	// 70000
	DefensePpsThreshold *int32 `json:"DefensePpsThreshold,omitempty" xml:"DefensePpsThreshold,omitempty"`
	// The burstable protection threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 12310
	ElasticThreshold *int32 `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 121.199.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **mitigating**: indicates that traffic scrubbing is triggered for the asset.
	//
	// 	- **blackholed**: indicates that blackhole filtering is triggered for the asset.
	//
	// 	- **normal**: indicates that the instance is normal.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the asset.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP protocol that is supported by the asset. Valid values:
	//
	// 	- **v4**: IPv4
	//
	// 	- **v6**: IPv6
	//
	// example:
	//
	// v4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the asset is associated with an Anti-DDoS Origin Basic instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	IsBgppack *bool `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceListInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetBlackholeThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefenseBpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefensePpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetElasticThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceId(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceIp(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceName(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceStatus(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceType(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIpVersion(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIsBgppack(v bool) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IsBgppack = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceIpAddressRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The region ID of the asset.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **defense**: queries assets for which traffic scrubbing is performed.
	//
	// 	- **blackhole**: queries assets for which blackhole filtering is triggered.
	//
	// example:
	//
	// normal
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The ID of the instance to which the asset is added.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the asset that is assigned a public IP address. Valid values:
	//
	// 	- **ecs**: ECS instances.
	//
	// 	- **slb**: SLB instances.
	//
	// 	- **eip**: EIPs.
	//
	// 	- **ipv6**: IPv6 gateways.
	//
	// 	- **swas**: simple application servers.
	//
	// 	- **waf**: Web Application Firewall (WAF) instances of the Exclusive edition.
	//
	// 	- **ga_basic**: Global Accelerator (GA) instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressRequest) SetCurrentPage(v int32) *DescribeInstanceIpAddressRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetDdosRegionId(v string) *DescribeInstanceIpAddressRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetDdosStatus(v string) *DescribeInstanceIpAddressRequest {
	s.DdosStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceId(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceIp(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceName(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceType(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetPageSize(v int32) *DescribeInstanceIpAddressRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceIpAddressResponseBody struct {
	// An array that consists of details of the instance.
	InstanceList []*DescribeInstanceIpAddressResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0907F8-A9F3-5E11-977B-D59CD98C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBody) SetInstanceList(v []*DescribeInstanceIpAddressResponseBodyInstanceList) *DescribeInstanceIpAddressResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceIpAddressResponseBody) SetRequestId(v string) *DescribeInstanceIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBody) SetTotal(v int32) *DescribeInstanceIpAddressResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceIpAddressResponseBodyInstanceList struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The DDoS mitigation status of the instance. Valid values:
	//
	// 	- **normal**: not under DDoS attacks.
	//
	// 	- **abnormal**: under DDoS attacks.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **ecs**
	//
	// 	- **slb**
	//
	// 	- **eip**
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// An array that consists of the details of the asset.
	IpAddressConfig []*DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig `json:"IpAddressConfig,omitempty" xml:"IpAddressConfig,omitempty" type:"Repeated"`
}

func (s DescribeInstanceIpAddressResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceName(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceStatus(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceType(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetIpAddressConfig(v []*DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.IpAddressConfig = v
	return s
}

type DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig struct {
	// The basic protection threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 5200
	BlackholeThreshold *int32 `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	// The traffic scrubbing threshold for the asset measured in Mbit/s. Unit: Mbit/s.
	//
	// example:
	//
	// 300
	DefenseBpsThreshold *int32 `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	// The traffic scrubbing threshold for the asset measured in packets per second (PPS). Unit: PPS.
	//
	// example:
	//
	// 70000
	DefensePpsThreshold *int32 `json:"DefensePpsThreshold,omitempty" xml:"DefensePpsThreshold,omitempty"`
	// The burstable protection threshold for the asset. Unit: Mbit/s.
	//
	// example:
	//
	// 12310
	ElasticThreshold *int32 `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **mitigating**: indicates that traffic scrubbing is in progress.
	//
	// 	- **blackholed**: indicates that blackhole filtering is triggered for the asset.
	//
	// 	- **normal**: indicates that no DDoS attacks are launched against the asset.
	//
	// example:
	//
	// normal
	IpStatus *string `json:"IpStatus,omitempty" xml:"IpStatus,omitempty"`
	// The IP version of the IP address. Valid values:
	//
	// 	- **v4**: IPv4.
	//
	// 	- **v6**: IPv6.
	//
	// example:
	//
	// v4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the asset is added to the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsBgppack *bool `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
	// Indicates whether best-effort protection is enabled for the asset. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 0
	IsFullProtection *int32 `json:"IsFullProtection,omitempty" xml:"IsFullProtection,omitempty"`
	// The region code of the asset.
	//
	// example:
	//
	// cn-hangzhou-dg-a01
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetBlackholeThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetDefenseBpsThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetDefensePpsThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetElasticThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetInstanceIp(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIpStatus(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IpStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIpVersion(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIsBgppack(v bool) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IsBgppack = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIsFullProtection(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IsFullProtection = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetRegionId(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.RegionId = &v
	return s
}

type DescribeInstanceIpAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponse) SetHeaders(v map[string]*string) *DescribeInstanceIpAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceIpAddressResponse) SetStatusCode(v int32) *DescribeInstanceIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceIpAddressResponse) SetBody(v *DescribeInstanceIpAddressResponseBody) *DescribeInstanceIpAddressResponse {
	s.Body = v
	return s
}

type DescribeIpDdosThresholdRequest struct {
	// The region ID of the asset.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The type of the threshold. Valid values:
	//
	// 	- **defense**: traffic scrubbing threshold
	//
	// 	- **blackhole**: DDoS mitigation threshold
	//
	// This parameter is required.
	//
	// example:
	//
	// defense
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The ID of the asset.
	//
	// > You can call the [DescribeInstanceIpAddress](https://help.aliyun.com/document_detail/429562.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1i88rqjza51s****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset that is assigned a public IP address. Valid values:
	//
	// 	- **ecs**: ECS instances.
	//
	// 	- **slb**: SLB instances.
	//
	// 	- **eip**: EIPs.
	//
	// 	- **ipv6**: IPv6 gateways.
	//
	// 	- **swas**: simple application servers.
	//
	// 	- **waf**: Web Application Firewall (WAF) instances of the Exclusive edition.
	//
	// 	- **ga_basic**: Global Accelerator (GA) instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeIpDdosThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdRequest) SetDdosRegionId(v string) *DescribeIpDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetDdosType(v string) *DescribeIpDdosThresholdRequest {
	s.DdosType = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInstanceId(v string) *DescribeIpDdosThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInstanceType(v string) *DescribeIpDdosThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInternetIp(v string) *DescribeIpDdosThresholdRequest {
	s.InternetIp = &v
	return s
}

type DescribeIpDdosThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 025F1688-680B-551A-9A8E-1A0D796315CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the threshold.
	Threshold *DescribeIpDdosThresholdResponseBodyThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Struct"`
}

func (s DescribeIpDdosThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponseBody) SetRequestId(v string) *DescribeIpDdosThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBody) SetThreshold(v *DescribeIpDdosThresholdResponseBodyThreshold) *DescribeIpDdosThresholdResponseBody {
	s.Threshold = v
	return s
}

type DescribeIpDdosThresholdResponseBodyThreshold struct {
	// If the value of the **DdosType*	- parameter is **defense**, the Bps parameter indicates the current traffic scrubbing threshold. Unit: Mbit/s.
	//
	// If the value of the **DdosType*	- parameter is **blackhole**, the Bps parameter indicates the basic protection threshold. Unit: Mbit/s.
	//
	// example:
	//
	// 7500
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The type of the threshold. Valid values:
	//
	// 	- **defense**: traffic scrubbing threshold
	//
	// 	- **blackhole**: DDoS mitigation threshold
	//
	// example:
	//
	// defense
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The burstable protection threshold (the maximum DDoS mitigation threshold). Unit: Mbit/s.
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **blackhole**.
	//
	// example:
	//
	// 12310
	ElasticBps *int32 `json:"ElasticBps,omitempty" xml:"ElasticBps,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1i88rqjza51s****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Indicates whether the threshold is automatically adjusted. Valid values:
	//
	// 	- **true**: The scrubbing thresholds are automatically adjusted based on the traffic load on the asset.
	//
	// 	- **false**: The scrubbing thresholds are not automatically adjusted. You must manually specify the scrubbing thresholds.
	//
	// example:
	//
	// false
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	// The maximum traffic scrubbing threshold. Unit: Mbit/s.
	//
	// example:
	//
	// 7500
	MaxBps *int32 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The maximum packet scrubbing threshold. Unit: pps.
	//
	// example:
	//
	// 5000000
	MaxPps *int32 `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
	// The packet scrubbing threshold. Unit: packets per second (pps).
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **defense**.
	//
	// example:
	//
	// 5000000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeIpDdosThresholdResponseBodyThreshold) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdResponseBodyThreshold) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.Bps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetDdosType(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.DdosType = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetElasticBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.ElasticBps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetInstanceId(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetInternetIp(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetIsAuto(v bool) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.IsAuto = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetMaxBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.MaxBps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetMaxPps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.MaxPps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetPps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.Pps = &v
	return s
}

type DescribeIpDdosThresholdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpDdosThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponse) SetHeaders(v map[string]*string) *DescribeIpDdosThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpDdosThresholdResponse) SetStatusCode(v int32) *DescribeIpDdosThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpDdosThresholdResponse) SetBody(v *DescribeIpDdosThresholdResponseBody) *DescribeIpDdosThresholdResponse {
	s.Body = v
	return s
}

type DescribeIpLocationServiceRequest struct {
	// The IP address of the asset to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeIpLocationServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceRequest) SetInternetIp(v string) *DescribeIpLocationServiceRequest {
	s.InternetIp = &v
	return s
}

type DescribeIpLocationServiceResponseBody struct {
	// The details of the asset.
	Instance *DescribeIpLocationServiceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C728D7E9-9A39-52E0-966B-5C33118BDBB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpLocationServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponseBody) SetInstance(v *DescribeIpLocationServiceResponseBodyInstance) *DescribeIpLocationServiceResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeIpLocationServiceResponseBody) SetRequestId(v string) *DescribeIpLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIpLocationServiceResponseBodyInstance struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **ecs**: an ECS instance.
	//
	// 	- **slb**: an SLB instance.
	//
	// 	- **eip**: an EIP.
	//
	// 	- **ipv6**: an IPv6 gateway.
	//
	// 	- **swas**: a simple application server.
	//
	// 	- **waf**: a Web Application Firewall (WAF) instance of the Exclusive edition.
	//
	// 	- **ga_basic**: a Global Accelerator (GA) instance.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The region to which the public IP address of the asset belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeIpLocationServiceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceId(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceName(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceType(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInternetIp(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetRegion(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.Region = &v
	return s
}

type DescribeIpLocationServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpLocationServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponse) SetHeaders(v map[string]*string) *DescribeIpLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpLocationServiceResponse) SetStatusCode(v int32) *DescribeIpLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpLocationServiceResponse) SetBody(v *DescribeIpLocationServiceResponseBody) *DescribeIpLocationServiceResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	// An array consisting of regions in which Anti-DDoS Origin Basic is available.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5093C7EE-8E27-5FC9-9B88-40626BA540C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The English name of the region.
	//
	// example:
	//
	// East China 1
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	// The Chinese name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The code of the region.
	//
	// example:
	//
	// cn-hangzhou-dg-a01
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionNoAlias *string `json:"RegionNoAlias,omitempty" xml:"RegionNoAlias,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionNo(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionNo = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionNoAlias(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionNoAlias = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type ModifyDefenseThresholdRequest struct {
	// The traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset. When you modify Bps, Pps is required. Otherwise, Bps does not take effect.
	//
	// You can use the monitoring tool that is provided by the asset to query the Internet traffic of the asset:
	//
	// 	- If the asset is an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// 	- If the asset is an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// example:
	//
	// 100
	Bps         *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the asset for which you want to change the scrubbing thresholds.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The ID of the asset.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6idy3c57psf7vu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **ecs**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **slb**: a Server Load Balancer (SLB) instance.
	//
	// 	- **eip**: an elastic IP address (EIP).
	//
	// 	- **ipv6**: an IPv6 gateway.
	//
	// 	- **swas**: a simple application server.
	//
	// 	- **waf**: a Web Application Firewall (WAF) instance of the Exclusive edition.
	//
	// 	- **ga_basic**: a Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Specifies whether to automatically adjust the scrubbing threshold based on the traffic load on the asset. Valid values:
	//
	// 	- **true**: automatically adjusts the scrubbing thresholds. You do not need to configure the **Bps*	- and **Pps*	- parameters.
	//
	// 	- **false**: The scrubbing threshold is not automatically adjusted. You must configure the **Bps*	- and **Pps*	- parameters.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	// The packet scrubbing threshold. Unit: packets per second (PPS). When you modify Pps, Bps is required. Otherwise, Pps does not take effect.
	//
	// The packet scrubbing threshold cannot exceed the peak number of inbound or outbound packets, whichever is larger, of the asset. You can use the monitoring tool that is provided by the asset to query the number of packets of the asset:
	//
	// 	- If the asset is an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// 	- If the asset is an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// example:
	//
	// 70000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s ModifyDefenseThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseThresholdRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdRequest) SetBps(v int32) *ModifyDefenseThresholdRequest {
	s.Bps = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetClientToken(v string) *ModifyDefenseThresholdRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetDdosRegionId(v string) *ModifyDefenseThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceId(v string) *ModifyDefenseThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceType(v string) *ModifyDefenseThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInternetIp(v string) *ModifyDefenseThresholdRequest {
	s.InternetIp = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetIsAuto(v bool) *ModifyDefenseThresholdRequest {
	s.IsAuto = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetPps(v int32) *ModifyDefenseThresholdRequest {
	s.Pps = &v
	return s
}

type ModifyDefenseThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0607F8-A9F3-5E11-977B-D59CD58C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdResponseBody) SetRequestId(v string) *ModifyDefenseThresholdResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseThresholdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseThresholdResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdResponse) SetHeaders(v map[string]*string) *ModifyDefenseThresholdResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseThresholdResponse) SetStatusCode(v int32) *ModifyDefenseThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseThresholdResponse) SetBody(v *ModifyDefenseThresholdResponseBody) *ModifyDefenseThresholdResponse {
	s.Body = v
	return s
}

type ModifyIpDefenseThresholdRequest struct {
	// The traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset. When you modify Bps, Pps is required. Otherwise, Bps does not take effect.
	//
	// You can use the monitoring tool that is provided by the asset to query the Internet traffic of the asset:
	//
	// 	- If the asset is an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// 	- If the asset is an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// 	- If the asset is an EIP, see [View monitoring data](https://help.aliyun.com/document_detail/85354.html).
	//
	// example:
	//
	// 100
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The region ID of the asset.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The ID of the asset.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6idy3c57psf7vu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **ecs**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **slb**: a Server Load Balancer (SLB) instance.
	//
	// 	- **eip**: an elastic IP address (EIP).
	//
	// 	- **ipv6**: an IPv6 gateway.
	//
	// 	- **swas**: a simple application server.
	//
	// 	- **waf**: a Web Application Firewall (WAF) instance of the Exclusive edition.
	//
	// 	- **ga_basic**: a Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Specifies whether to automatically adjust the scrubbing threshold based on the traffic load on the asset. Valid values:
	//
	// 	- **true**: automatically adjusts the scrubbing threshold. You do not need to configure the **Bps*	- and **Pps*	- parameters.
	//
	// 	- **false**: The scrubbing threshold is not automatically adjusted. You must configure the **Bps*	- and **Pps*	- parameters. This is the default value.
	//
	// example:
	//
	// false
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	// The packet scrubbing threshold. Unit: packets per second (PPS). When you modify Pps, Bps is required. Otherwise, Pps does not take effect.
	//
	// The packet scrubbing threshold cannot exceed the peak number of inbound or outbound packets, whichever is larger, of the asset. You can use the monitoring tool that is provided by the asset to query the number of packets of the asset:
	//
	// 	- If the asset is an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// 	- If the asset is an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// 	- If the asset is an EIP, see [View monitoring data](https://help.aliyun.com/document_detail/85354.html).
	//
	// example:
	//
	// 70000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s ModifyIpDefenseThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpDefenseThresholdRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpDefenseThresholdRequest) SetBps(v int32) *ModifyIpDefenseThresholdRequest {
	s.Bps = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetDdosRegionId(v string) *ModifyIpDefenseThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetInstanceId(v string) *ModifyIpDefenseThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetInstanceType(v string) *ModifyIpDefenseThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetInternetIp(v string) *ModifyIpDefenseThresholdRequest {
	s.InternetIp = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetIsAuto(v bool) *ModifyIpDefenseThresholdRequest {
	s.IsAuto = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetPps(v int32) *ModifyIpDefenseThresholdRequest {
	s.Pps = &v
	return s
}

type ModifyIpDefenseThresholdResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0607F8-A9F3-5E11-977B-D59CD58C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpDefenseThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpDefenseThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpDefenseThresholdResponseBody) SetRequestId(v string) *ModifyIpDefenseThresholdResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpDefenseThresholdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpDefenseThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpDefenseThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpDefenseThresholdResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpDefenseThresholdResponse) SetHeaders(v map[string]*string) *ModifyIpDefenseThresholdResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpDefenseThresholdResponse) SetStatusCode(v int32) *ModifyIpDefenseThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpDefenseThresholdResponse) SetBody(v *ModifyIpDefenseThresholdResponseBody) *ModifyIpDefenseThresholdResponse {
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
		"cn-qingdao":                  tea.String("antiddos.aliyuncs.com"),
		"cn-beijing":                  tea.String("antiddos.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("antiddos-openapi.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":                tea.String("antiddos-openapi.cn-huhehaote.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("antiddos-openapi.cn-wulanchabu.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("antiddos.aliyuncs.com"),
		"cn-shanghai":                 tea.String("antiddos.aliyuncs.com"),
		"cn-nanjing":                  tea.String("antiddos-openapi.cn-hangzhou-cloudstone.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("antiddos.aliyuncs.com"),
		"cn-heyuan":                   tea.String("antiddos-openapi.cn-heyuan.aliyuncs.com"),
		"cn-guangzhou":                tea.String("antiddos-openapi.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":                  tea.String("antiddos-openapi.cn-chengdu.aliyuncs.com"),
		"cn-hongkong":                 tea.String("antiddos.aliyuncs.com"),
		"ap-northeast-1":              tea.String("antiddos-openapi.ap-northeast-1.aliyuncs.com"),
		"ap-northeast-2":              tea.String("antiddos-openapi.ap-northeast-2.aliyuncs.com"),
		"ap-southeast-1":              tea.String("antiddos.aliyuncs.com"),
		"ap-southeast-2":              tea.String("antiddos-openapi.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              tea.String("antiddos-openapi.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":              tea.String("antiddos-openapi.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-6":              tea.String("antiddos-openapi.ap-southeast-6.aliyuncs.com"),
		"us-east-1":                   tea.String("antiddos.aliyuncs.com"),
		"us-west-1":                   tea.String("antiddos.aliyuncs.com"),
		"eu-west-1":                   tea.String("antiddos-openapi.eu-west-1.aliyuncs.com"),
		"eu-central-1":                tea.String("antiddos-openapi.eu-central-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("antiddos-openapi.ap-south-1.aliyuncs.com"),
		"me-east-1":                   tea.String("antiddos-openapi.me-east-1.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("antiddos.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("antiddos.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("antiddos.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("antiddos.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("antiddos.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("antiddos.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("antiddos.aliyuncs.com"),
		"cn-edge-1":                   tea.String("antiddos.aliyuncs.com"),
		"cn-fujian":                   tea.String("antiddos.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("antiddos.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("antiddos.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("antiddos.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("antiddos.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("antiddos.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("antiddos.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("antiddos.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("antiddos.aliyuncs.com"),
		"cn-wuhan":                    tea.String("antiddos.aliyuncs.com"),
		"cn-yushanfang":               tea.String("antiddos.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("antiddos.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("antiddos.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("antiddos-openapi.cn-zhangjiakou.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("antiddos.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("antiddos.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("antiddos.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("antiddos-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// Queries the configurations of the Anti-DDoS Origin instance that is associated with an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeBgpPackByIp operation to query the configurations of the Anti-DDoS Origin instance that is associated with an asset. The configurations include the basic protection threshold, burstable protection threshold, and expiration time.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeBgpPackByIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBgpPackByIpResponse
func (client *Client) DescribeBgpPackByIpWithOptions(request *DescribeBgpPackByIpRequest, runtime *util.RuntimeOptions) (_result *DescribeBgpPackByIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBgpPackByIp"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBgpPackByIpResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBgpPackByIpResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the configurations of the Anti-DDoS Origin instance that is associated with an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeBgpPackByIp operation to query the configurations of the Anti-DDoS Origin instance that is associated with an asset. The configurations include the basic protection threshold, burstable protection threshold, and expiration time.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeBgpPackByIpRequest
//
// @return DescribeBgpPackByIpResponse
func (client *Client) DescribeBgpPackByIp(request *DescribeBgpPackByIpRequest) (_result *DescribeBgpPackByIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBgpPackByIpResponse{}
	_body, _err := client.DescribeBgpPackByIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download link to the traffic data that is captured when a DDoS attack event occurs.
//
// Description:
//
// You can call the DescribeCap operation to query the download link to the traffic data that is captured when a DDoS attack event occurs. You can download the traffic data from the download link and use the data as evidence.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCapResponse
func (client *Client) DescribeCapWithOptions(request *DescribeCapRequest, runtime *util.RuntimeOptions) (_result *DescribeCapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BegTime)) {
		query["BegTime"] = request.BegTime
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCap"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeCapResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeCapResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the download link to the traffic data that is captured when a DDoS attack event occurs.
//
// Description:
//
// You can call the DescribeCap operation to query the download link to the traffic data that is captured when a DDoS attack event occurs. You can download the traffic data from the download link and use the data as evidence.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCapRequest
//
// @return DescribeCapResponse
func (client *Client) DescribeCap(request *DescribeCapRequest) (_result *DescribeCapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCapResponse{}
	_body, _err := client.DescribeCapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are under DDoS attacks in a specific region. The assets are assigned public IP addresses.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosCountResponse
func (client *Client) DescribeDdosCountWithOptions(request *DescribeDdosCountRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosCount"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDdosCountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDdosCountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the number of assets that are under DDoS attacks in a specific region. The assets are assigned public IP addresses.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCountRequest
//
// @return DescribeDdosCountResponse
func (client *Client) DescribeDdosCount(request *DescribeDdosCountRequest) (_result *DescribeDdosCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosCountResponse{}
	_body, _err := client.DescribeDdosCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the security credit score of the current Alibaba Cloud account in a specific region.
//
// Description:
//
// You can call the DescribeDdosCredit operation to query the details of the security credit score of the current Alibaba Cloud account in a specific region. The details include the security credit score, security credit level, and the time period after which blackhole filtering is automatically deactivated.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosCreditResponse
func (client *Client) DescribeDdosCreditWithOptions(request *DescribeDdosCreditRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosCreditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosCredit"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDdosCreditResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDdosCreditResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of the security credit score of the current Alibaba Cloud account in a specific region.
//
// Description:
//
// You can call the DescribeDdosCredit operation to query the details of the security credit score of the current Alibaba Cloud account in a specific region. The details include the security credit score, security credit level, and the time period after which blackhole filtering is automatically deactivated.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCreditRequest
//
// @return DescribeDdosCreditResponse
func (client *Client) DescribeDdosCredit(request *DescribeDdosCreditRequest) (_result *DescribeDdosCreditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosCreditResponse{}
	_body, _err := client.DescribeDdosCreditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS attack events that occur on an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeDdosEventList operation to query the details of the DDoS attack events that occur on an asset by page. The details include the start time, end time, and status of each DDoS attack event.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosEventListResponse
func (client *Client) DescribeDdosEventListWithOptions(request *DescribeDdosEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosEventList"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDdosEventListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDdosEventListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of the DDoS attack events that occur on an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeDdosEventList operation to query the details of the DDoS attack events that occur on an asset by page. The details include the start time, end time, and status of each DDoS attack event.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosEventListRequest
//
// @return DescribeDdosEventListResponse
func (client *Client) DescribeDdosEventList(request *DescribeDdosEventListRequest) (_result *DescribeDdosEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosEventListResponse{}
	_body, _err := client.DescribeDdosEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// You can call the DescribeDdosThreshold operation to query the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The details include the current traffic scrubbing threshold, maximum traffic scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosThresholdResponse
func (client *Client) DescribeDdosThresholdWithOptions(request *DescribeDdosThresholdRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosType)) {
		query["DdosType"] = request.DdosType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosThreshold"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDdosThresholdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDdosThresholdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// You can call the DescribeDdosThreshold operation to query the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The details include the current traffic scrubbing threshold, maximum traffic scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosThresholdRequest
//
// @return DescribeDdosThresholdResponse
func (client *Client) DescribeDdosThreshold(request *DescribeDdosThresholdRequest) (_result *DescribeDdosThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosThresholdResponse{}
	_body, _err := client.DescribeDdosThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the assets within the current Alibaba Cloud account. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses. This operation is phased out. We recommend that you use the DescribeInstanceIpAddress operation.
//
// Description:
//
// You can call the DescribeInstance operation to query the details of the assets that are within the current Alibaba Cloud account by page. The details include the IDs and IP addresses of the assets, the basic protection thresholds and traffic scrubbing thresholds that are configured for the assets in Anti-DDoS Origin, and whether the assets are associated with Anti-DDoS Origin instances.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosStatus)) {
		query["DdosStatus"] = request.DdosStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIp)) {
		query["InstanceIp"] = request.InstanceIp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of the assets within the current Alibaba Cloud account. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses. This operation is phased out. We recommend that you use the DescribeInstanceIpAddress operation.
//
// Description:
//
// You can call the DescribeInstance operation to query the details of the assets that are within the current Alibaba Cloud account by page. The details include the IDs and IP addresses of the assets, the basic protection thresholds and traffic scrubbing thresholds that are configured for the assets in Anti-DDoS Origin, and whether the assets are associated with Anti-DDoS Origin instances.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the assets within the current Alibaba Cloud account and the details of the Anti-DDoS Origin instance to which the assets belong. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeInstanceIpAddress operation to query the DDoS mitigation information and the details of the Anti-DDoS Origin instance. The information and the details include the basic protection threshold and traffic scrubbing threshold for the assets, DDoS mitigation status of the assets, ID of the instance, and the mitigation status of the instance.
//
// ## Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceIpAddressResponse
func (client *Client) DescribeInstanceIpAddressWithOptions(request *DescribeInstanceIpAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosStatus)) {
		query["DdosStatus"] = request.DdosStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIp)) {
		query["InstanceIp"] = request.InstanceIp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceIpAddress"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceIpAddressResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of the assets within the current Alibaba Cloud account and the details of the Anti-DDoS Origin instance to which the assets belong. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeInstanceIpAddress operation to query the DDoS mitigation information and the details of the Anti-DDoS Origin instance. The information and the details include the basic protection threshold and traffic scrubbing threshold for the assets, DDoS mitigation status of the assets, ID of the instance, and the mitigation status of the instance.
//
// ## Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceIpAddressRequest
//
// @return DescribeInstanceIpAddressResponse
func (client *Client) DescribeInstanceIpAddress(request *DescribeInstanceIpAddressRequest) (_result *DescribeInstanceIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceIpAddressResponse{}
	_body, _err := client.DescribeInstanceIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeIpDdosThreshold operation to query the details of the DDoS mitigation threshold or traffic scrubbing threshold for a specific asset. The details include the current traffic scrubbing threshold, maximum scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpDdosThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpDdosThresholdResponse
func (client *Client) DescribeIpDdosThresholdWithOptions(request *DescribeIpDdosThresholdRequest, runtime *util.RuntimeOptions) (_result *DescribeIpDdosThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosType)) {
		query["DdosType"] = request.DdosType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpDdosThreshold"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeIpDdosThresholdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeIpDdosThresholdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeIpDdosThreshold operation to query the details of the DDoS mitigation threshold or traffic scrubbing threshold for a specific asset. The details include the current traffic scrubbing threshold, maximum scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpDdosThresholdRequest
//
// @return DescribeIpDdosThresholdResponse
func (client *Client) DescribeIpDdosThreshold(request *DescribeIpDdosThresholdRequest) (_result *DescribeIpDdosThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpDdosThresholdResponse{}
	_body, _err := client.DescribeIpDdosThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the region to which the public IP address of the asset within the current Alibaba Cloud account belongs. The asset can be an elastic IP address (EIP). The asset can also be an Elastic Compute Service (ECS) instance or Server Load Balancer (SLB) instance that is assigned a public IP address.
//
// Description:
//
// You can call the DescribeIpLocationService operation to query the region of the public IP address for a specified asset that is within the current Alibaba Cloud account. You can also query the details of the Anti-DDoS Origin instance to which the asset is added. The details include the ID and name.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpLocationServiceResponse
func (client *Client) DescribeIpLocationServiceWithOptions(request *DescribeIpLocationServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeIpLocationServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpLocationService"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeIpLocationServiceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeIpLocationServiceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the region to which the public IP address of the asset within the current Alibaba Cloud account belongs. The asset can be an elastic IP address (EIP). The asset can also be an Elastic Compute Service (ECS) instance or Server Load Balancer (SLB) instance that is assigned a public IP address.
//
// Description:
//
// You can call the DescribeIpLocationService operation to query the region of the public IP address for a specified asset that is within the current Alibaba Cloud account. You can also query the details of the Anti-DDoS Origin instance to which the asset is added. The details include the ID and name.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpLocationServiceRequest
//
// @return DescribeIpLocationServiceResponse
func (client *Client) DescribeIpLocationService(request *DescribeIpLocationServiceRequest) (_result *DescribeIpLocationServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpLocationServiceResponse{}
	_body, _err := client.DescribeIpLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions in which Anti-DDoS Origin Basic is available.
//
// Description:
//
// You can call this operation to query information about the regions in which Anti-DDoS Origin Basic is available. The information includes the region ID, region name, and code.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the regions in which Anti-DDoS Origin Basic is available.
//
// Description:
//
// You can call this operation to query information about the regions in which Anti-DDoS Origin Basic is available. The information includes the region ID, region name, and code.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the scrubbing thresholds for an asset that is assigned a public IP address.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDefenseThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseThresholdResponse
func (client *Client) ModifyDefenseThresholdWithOptions(request *ModifyDefenseThresholdRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bps)) {
		query["Bps"] = request.Bps
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	if !tea.BoolValue(util.IsUnset(request.IsAuto)) {
		query["IsAuto"] = request.IsAuto
	}

	if !tea.BoolValue(util.IsUnset(request.Pps)) {
		query["Pps"] = request.Pps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseThreshold"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyDefenseThresholdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyDefenseThresholdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the scrubbing thresholds for an asset that is assigned a public IP address.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDefenseThresholdRequest
//
// @return ModifyDefenseThresholdResponse
func (client *Client) ModifyDefenseThreshold(request *ModifyDefenseThresholdRequest) (_result *ModifyDefenseThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseThresholdResponse{}
	_body, _err := client.ModifyDefenseThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the scrubbing thresholds for an asset that is assigned a public IP address. This operation is a synchronous operation that supports Terraform.
//
// Description:
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyIpDefenseThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpDefenseThresholdResponse
func (client *Client) ModifyIpDefenseThresholdWithOptions(request *ModifyIpDefenseThresholdRequest, runtime *util.RuntimeOptions) (_result *ModifyIpDefenseThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bps)) {
		query["Bps"] = request.Bps
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	if !tea.BoolValue(util.IsUnset(request.IsAuto)) {
		query["IsAuto"] = request.IsAuto
	}

	if !tea.BoolValue(util.IsUnset(request.Pps)) {
		query["Pps"] = request.Pps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIpDefenseThreshold"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyIpDefenseThresholdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyIpDefenseThresholdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the scrubbing thresholds for an asset that is assigned a public IP address. This operation is a synchronous operation that supports Terraform.
//
// Description:
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyIpDefenseThresholdRequest
//
// @return ModifyIpDefenseThresholdResponse
func (client *Client) ModifyIpDefenseThreshold(request *ModifyIpDefenseThresholdRequest) (_result *ModifyIpDefenseThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpDefenseThresholdResponse{}
	_body, _err := client.ModifyIpDefenseThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
