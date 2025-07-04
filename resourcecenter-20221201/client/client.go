// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssociateDefaultFilterRequest struct {
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// My Filters
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s AssociateDefaultFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateDefaultFilterRequest) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterRequest) SetFilterName(v string) *AssociateDefaultFilterRequest {
	s.FilterName = &v
	return s
}

type AssociateDefaultFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54673B22-2001-556A-B394-B8697AA9670B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateDefaultFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateDefaultFilterResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterResponseBody) SetRequestId(v string) *AssociateDefaultFilterResponseBody {
	s.RequestId = &v
	return s
}

type AssociateDefaultFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateDefaultFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateDefaultFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateDefaultFilterResponse) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterResponse) SetHeaders(v map[string]*string) *AssociateDefaultFilterResponse {
	s.Headers = v
	return s
}

func (s *AssociateDefaultFilterResponse) SetStatusCode(v int32) *AssociateDefaultFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateDefaultFilterResponse) SetBody(v *AssociateDefaultFilterResponseBody) *AssociateDefaultFilterResponse {
	s.Body = v
	return s
}

type CreateDeliveryChannelRequest struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	//
	// This parameter is required.
	DeliveryChannelFilter *CreateDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The name of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-delivery
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *CreateDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *CreateDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s CreateDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *CreateDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *CreateDeliveryChannelRequest) SetDeliveryChannelFilter(v *CreateDeliveryChannelRequestDeliveryChannelFilter) *CreateDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *CreateDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateDeliveryChannelRequest) SetResourceChangeDelivery(v *CreateDeliveryChannelRequestResourceChangeDelivery) *CreateDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *CreateDeliveryChannelRequest) SetResourceSnapshotDelivery(v *CreateDeliveryChannelRequestResourceSnapshotDelivery) *CreateDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

type CreateDeliveryChannelRequestDeliveryChannelFilter struct {
	// An array of effective resource types for the delivery channel.
	//
	// 	- Example: ["ACS::VPC::VPC", "ACS::ECS::Instance"].
	//
	// 	- If you want to deliver items of all resource types supported by Resource Center, set this parameter to ["ALL"].
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s CreateDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *CreateDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

type CreateDeliveryChannelRequestResourceChangeDelivery struct {
	// The Simple Log Service configurations.
	SlsProperties *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with resourcecenter-.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with resourcecenter-.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Valid values:
	//
	// 	- `SLS`
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceChangeDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) *CreateDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *CreateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *CreateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

type CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// 	- If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object.
	//
	// 	- You need to set this parameter to the ARN of a bucket whose name is prefixed with resourcecenter-.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type CreateDeliveryChannelRequestResourceSnapshotDelivery struct {
	// The custom expression.
	//
	// example:
	//
	// select 	- from resources limit 100;
	CustomExpression *string `json:"CustomExpression,omitempty" xml:"CustomExpression,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 09:00Z
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination.
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with resourcecenter-.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with resourcecenter-.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Valid values:
	//
	// 	- `OSS` for standard delivery
	//
	// 	- `OSS` or `SLS` for custom delivery
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

type CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// 	- If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object.
	//
	// 	- You need to set this parameter to the ARN of a bucket whose name is prefixed with resourcecenter-.
	//
	// 	- This parameter takes effect only if you use custom delivery for scheduled resource snapshots. You do not need to configure this parameter if you use standard delivery for scheduled resource snapshots.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type CreateDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-0bzhsqpnk***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 42A89312-0616-591E-B614-07BC87D3D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateDeliveryChannelResponseBody) SetRequestId(v string) *CreateDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeliveryChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelResponse) SetHeaders(v map[string]*string) *CreateDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateDeliveryChannelResponse) SetStatusCode(v int32) *CreateDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeliveryChannelResponse) SetBody(v *CreateDeliveryChannelResponseBody) *CreateDeliveryChannelResponse {
	s.Body = v
	return s
}

type CreateFilterRequest struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s CreateFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFilterRequest) GoString() string {
	return s.String()
}

func (s *CreateFilterRequest) SetFilterConfiguration(v string) *CreateFilterRequest {
	s.FilterConfiguration = &v
	return s
}

func (s *CreateFilterRequest) SetFilterName(v string) *CreateFilterRequest {
	s.FilterName = &v
	return s
}

type CreateFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EEF1EE1F-50F6-5494-B3DA-8F597DEB31BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFilterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFilterResponseBody) SetRequestId(v string) *CreateFilterResponseBody {
	s.RequestId = &v
	return s
}

type CreateFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFilterResponse) GoString() string {
	return s.String()
}

func (s *CreateFilterResponse) SetHeaders(v map[string]*string) *CreateFilterResponse {
	s.Headers = v
	return s
}

func (s *CreateFilterResponse) SetStatusCode(v int32) *CreateFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFilterResponse) SetBody(v *CreateFilterResponseBody) *CreateFilterResponse {
	s.Body = v
	return s
}

type CreateMultiAccountDeliveryChannelRequest struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	//
	// This parameter is required.
	DeliveryChannelFilter *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The name of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-multi-account-delivery
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s CreateMultiAccountDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *CreateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetDeliveryChannelFilter(v *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) *CreateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetResourceChangeDelivery(v *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) *CreateMultiAccountDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetResourceSnapshotDelivery(v *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) *CreateMultiAccountDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

type CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter struct {
	// An array of effective account scopes for the delivery channel.
	//
	// This parameter is required.
	AccountScopes []*string `json:"AccountScopes,omitempty" xml:"AccountScopes,omitempty" type:"Repeated"`
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetAccountScopes(v []*string) *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.AccountScopes = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

type CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery struct {
	// The Simple Log Service configurations.
	SlsProperties *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 1911422487776***:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Set the value to `SLS`.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

type CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery struct {
	// The custom expression.
	//
	// example:
	//
	// select 	- from resources limit 100;
	CustomExpression *string `json:"CustomExpression,omitempty" xml:"CustomExpression,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 09:00Z
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`. Example: `acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`. Example: `acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination. Valid values:
	//
	// 	- `OSS` for standard delivery
	//
	// 	- `OSS` or `SLS` for custom delivery
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

type CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// >  This parameter takes effect only if you use custom delivery for scheduled resource snapshots. You do not need to configure this parameter if you use standard delivery for scheduled resource snapshots.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type CreateMultiAccountDeliveryChannelResponseBody struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 994BFEFE-4BB5-5A27-8917-4583DEEF2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *CreateMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type CreateMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *CreateMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *CreateMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponse) SetBody(v *CreateMultiAccountDeliveryChannelResponseBody) *CreateMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

type CreateSavedQueryRequest struct {
	// The description of the template.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryRequest) SetDescription(v string) *CreateSavedQueryRequest {
	s.Description = &v
	return s
}

func (s *CreateSavedQueryRequest) SetExpression(v string) *CreateSavedQueryRequest {
	s.Expression = &v
	return s
}

func (s *CreateSavedQueryRequest) SetName(v string) *CreateSavedQueryRequest {
	s.Name = &v
	return s
}

type CreateSavedQueryResponseBody struct {
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryResponseBody) SetQueryId(v string) *CreateSavedQueryResponseBody {
	s.QueryId = &v
	return s
}

func (s *CreateSavedQueryResponseBody) SetRequestId(v string) *CreateSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

type CreateSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryResponse) SetHeaders(v map[string]*string) *CreateSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *CreateSavedQueryResponse) SetStatusCode(v int32) *CreateSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSavedQueryResponse) SetBody(v *CreateSavedQueryResponseBody) *CreateSavedQueryResponse {
	s.Body = v
	return s
}

type DeleteDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s DeleteDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryChannelRequest) SetDeliveryChannelId(v string) *DeleteDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

type DeleteDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AD5F848D-CCDC-5464-93E1-4BA50A482***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryChannelResponseBody) SetRequestId(v string) *DeleteDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeliveryChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryChannelResponse) SetHeaders(v map[string]*string) *DeleteDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeliveryChannelResponse) SetStatusCode(v int32) *DeleteDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeliveryChannelResponse) SetBody(v *DeleteDeliveryChannelResponseBody) *DeleteDeliveryChannelResponse {
	s.Body = v
	return s
}

type DeleteFilterRequest struct {
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s DeleteFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilterRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilterRequest) SetFilterName(v string) *DeleteFilterRequest {
	s.FilterName = &v
	return s
}

type DeleteFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A4A63E3C-89EC-51F9-9934-C9AF1BCBAAA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilterResponseBody) SetRequestId(v string) *DeleteFilterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilterResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilterResponse) SetHeaders(v map[string]*string) *DeleteFilterResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilterResponse) SetStatusCode(v int32) *DeleteFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilterResponse) SetBody(v *DeleteFilterResponseBody) *DeleteFilterResponse {
	s.Body = v
	return s
}

type DeleteMultiAccountDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-0bzhsqpnkxxx
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s DeleteMultiAccountDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiAccountDeliveryChannelRequest) SetDeliveryChannelId(v string) *DeleteMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

type DeleteMultiAccountDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3C5CDBF6-4DB7-53E9-ADDC-5919E3FAC***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMultiAccountDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *DeleteMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultiAccountDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *DeleteMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *DeleteMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultiAccountDeliveryChannelResponse) SetBody(v *DeleteMultiAccountDeliveryChannelResponseBody) *DeleteMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

type DeleteSavedQueryRequest struct {
	// The ID of the template.
	//
	// You can call the [ListSavedQueries](~~ListSavedQueries~~) operation to obtain the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DeleteSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryRequest) SetQueryId(v string) *DeleteSavedQueryRequest {
	s.QueryId = &v
	return s
}

type DeleteSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D170D58E-6256-5344-8F5E-922EC9ECB7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryResponseBody) SetRequestId(v string) *DeleteSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryResponse) SetHeaders(v map[string]*string) *DeleteSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *DeleteSavedQueryResponse) SetStatusCode(v int32) *DeleteSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSavedQueryResponse) SetBody(v *DeleteSavedQueryResponseBody) *DeleteSavedQueryResponse {
	s.Body = v
	return s
}

type DisableMultiAccountResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4951F920-48DB-5731-96AA-3A7C8AE617D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableMultiAccountResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *DisableMultiAccountResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

type DisableMultiAccountResourceCenterResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableMultiAccountResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *DisableMultiAccountResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *DisableMultiAccountResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetBody(v *DisableMultiAccountResourceCenterResponseBody) *DisableMultiAccountResourceCenterResponse {
	s.Body = v
	return s
}

type DisableResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D170D58E-6256-5344-8F5E-922EC9ECB7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponseBody) SetRequestId(v string) *DisableResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

type DisableResourceCenterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponse) SetHeaders(v map[string]*string) *DisableResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableResourceCenterResponse) SetStatusCode(v int32) *DisableResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableResourceCenterResponse) SetBody(v *DisableResourceCenterResponseBody) *DisableResourceCenterResponse {
	s.Body = v
	return s
}

type DisassociateDefaultFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BCAB07BA-82FA-5DC0-9322-FB7ED726481D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateDefaultFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociateDefaultFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateDefaultFilterResponseBody) SetRequestId(v string) *DisassociateDefaultFilterResponseBody {
	s.RequestId = &v
	return s
}

type DisassociateDefaultFilterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateDefaultFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateDefaultFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateDefaultFilterResponse) GoString() string {
	return s.String()
}

func (s *DisassociateDefaultFilterResponse) SetHeaders(v map[string]*string) *DisassociateDefaultFilterResponse {
	s.Headers = v
	return s
}

func (s *DisassociateDefaultFilterResponse) SetStatusCode(v int32) *DisassociateDefaultFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateDefaultFilterResponse) SetBody(v *DisassociateDefaultFilterResponseBody) *DisassociateDefaultFilterResponse {
	s.Body = v
	return s
}

type EnableMultiAccountResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 767038B7-2027-5508-858B-E213232D57D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the feature. Valid values:
	//
	// 	- Pending: The feature is being enabled.
	//
	// 	- Enabled: The feature is enabled.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableMultiAccountResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *EnableMultiAccountResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetStatus(v string) *EnableMultiAccountResourceCenterResponseBody {
	s.Status = &v
	return s
}

type EnableMultiAccountResourceCenterResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableMultiAccountResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *EnableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *EnableMultiAccountResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *EnableMultiAccountResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetBody(v *EnableMultiAccountResourceCenterResponseBody) *EnableMultiAccountResourceCenterResponse {
	s.Body = v
	return s
}

type EnableResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 45357BEF-AB50-5E4D-B05D-5A882A4BE924
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The activation status of the service. Valid values:
	//
	// 	- Pending: The service is being activated.
	//
	// 	- Enabled: The service is activated.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *EnableResourceCenterResponseBody) SetRequestId(v string) *EnableResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableResourceCenterResponseBody) SetStatus(v string) *EnableResourceCenterResponseBody {
	s.Status = &v
	return s
}

type EnableResourceCenterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *EnableResourceCenterResponse) SetHeaders(v map[string]*string) *EnableResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *EnableResourceCenterResponse) SetStatusCode(v int32) *EnableResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableResourceCenterResponse) SetBody(v *EnableResourceCenterResponseBody) *EnableResourceCenterResponse {
	s.Body = v
	return s
}

type ExecuteMultiAccountSQLQueryRequest struct {
	// The SQL statement to be executed.
	//
	// The number of characters in the SQL statement must be less than 2,000.
	//
	// For more information about the SQL syntax, see [Basic SQL syntax](https://help.aliyun.com/document_detail/2539395.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM resources LIMIT 100;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 1000.
	//
	// Default value: 1000.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. The value of this parameter can be one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched.
	//
	// 	- ID of a member: Resources within the member are searched.
	//
	// 	- ID of a member/ID of a Resource group: Resources within the member in the resource group are searched.
	//
	// For more information about how to obtain the ID of a resource directory, the Root folder, a folder, a member, or a resource group, see [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html), [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html), [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html), [ListAccounts](https://help.aliyun.com/document_detail/160016.html), or [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ExecuteMultiAccountSQLQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryRequest) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetExpression(v string) *ExecuteMultiAccountSQLQueryRequest {
	s.Expression = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetMaxResults(v int32) *ExecuteMultiAccountSQLQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetNextToken(v string) *ExecuteMultiAccountSQLQueryRequest {
	s.NextToken = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetScope(v string) *ExecuteMultiAccountSQLQueryRequest {
	s.Scope = &v
	return s
}

type ExecuteMultiAccountSQLQueryResponseBody struct {
	// The columns.
	Columns []*ExecuteMultiAccountSQLQueryResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 44C8A952-D6B0-5BC8-82D5-93BA02E26F2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of search results.
	Rows []interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ExecuteMultiAccountSQLQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetColumns(v []*ExecuteMultiAccountSQLQueryResponseBodyColumns) *ExecuteMultiAccountSQLQueryResponseBody {
	s.Columns = v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetMaxResults(v int32) *ExecuteMultiAccountSQLQueryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetNextToken(v string) *ExecuteMultiAccountSQLQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetRequestId(v string) *ExecuteMultiAccountSQLQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetRows(v []interface{}) *ExecuteMultiAccountSQLQueryResponseBody {
	s.Rows = v
	return s
}

type ExecuteMultiAccountSQLQueryResponseBodyColumns struct {
	// The name of the column.
	//
	// example:
	//
	// resource_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// varchar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExecuteMultiAccountSQLQueryResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) SetName(v string) *ExecuteMultiAccountSQLQueryResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) SetType(v string) *ExecuteMultiAccountSQLQueryResponseBodyColumns {
	s.Type = &v
	return s
}

type ExecuteMultiAccountSQLQueryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteMultiAccountSQLQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteMultiAccountSQLQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponse) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponse) SetHeaders(v map[string]*string) *ExecuteMultiAccountSQLQueryResponse {
	s.Headers = v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponse) SetStatusCode(v int32) *ExecuteMultiAccountSQLQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponse) SetBody(v *ExecuteMultiAccountSQLQueryResponseBody) *ExecuteMultiAccountSQLQueryResponse {
	s.Body = v
	return s
}

type ExecuteSQLQueryRequest struct {
	// The SQL statement to be executed.
	//
	// The number of characters in the SQL statement must be less than 2,000.
	//
	// For more information about the SQL syntax, see [Basic SQL syntax](https://help.aliyun.com/document_detail/2539395.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM resources LIMIT 100;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: 1 to 1000.
	//
	// 	- Default value: 1000.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope.
	//
	// Set this parameter to the ID of a resource group.
	//
	// For information about how to obtain the ID of a resource group, see [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
	//
	// example:
	//
	// rg-acfmzawhxxc****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ExecuteSQLQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryRequest) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryRequest) SetExpression(v string) *ExecuteSQLQueryRequest {
	s.Expression = &v
	return s
}

func (s *ExecuteSQLQueryRequest) SetMaxResults(v int32) *ExecuteSQLQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *ExecuteSQLQueryRequest) SetNextToken(v string) *ExecuteSQLQueryRequest {
	s.NextToken = &v
	return s
}

func (s *ExecuteSQLQueryRequest) SetScope(v string) *ExecuteSQLQueryRequest {
	s.Scope = &v
	return s
}

type ExecuteSQLQueryResponseBody struct {
	// The columns.
	Columns []*ExecuteSQLQueryResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of search results.
	Rows []interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ExecuteSQLQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryResponseBody) SetColumns(v []*ExecuteSQLQueryResponseBodyColumns) *ExecuteSQLQueryResponseBody {
	s.Columns = v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetMaxResults(v int32) *ExecuteSQLQueryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetNextToken(v string) *ExecuteSQLQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetRequestId(v string) *ExecuteSQLQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetRows(v []interface{}) *ExecuteSQLQueryResponseBody {
	s.Rows = v
	return s
}

type ExecuteSQLQueryResponseBodyColumns struct {
	// The name of the column.
	//
	// example:
	//
	// resource_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// varchar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExecuteSQLQueryResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryResponseBodyColumns) SetName(v string) *ExecuteSQLQueryResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *ExecuteSQLQueryResponseBodyColumns) SetType(v string) *ExecuteSQLQueryResponseBodyColumns {
	s.Type = &v
	return s
}

type ExecuteSQLQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteSQLQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSQLQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryResponse) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryResponse) SetHeaders(v map[string]*string) *ExecuteSQLQueryResponse {
	s.Headers = v
	return s
}

func (s *ExecuteSQLQueryResponse) SetStatusCode(v int32) *ExecuteSQLQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteSQLQueryResponse) SetBody(v *ExecuteSQLQueryResponseBody) *ExecuteSQLQueryResponse {
	s.Body = v
	return s
}

type GetDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

type GetDeliveryChannelResponseBody struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *GetDeliveryChannelResponseBodyDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17502A1B-7026-5551-8E1C-CCABD0CBC***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *GetDeliveryChannelResponseBodyResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *GetDeliveryChannelResponseBodyResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s GetDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelDescription(v string) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelFilter(v *GetDeliveryChannelResponseBodyDeliveryChannelFilter) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelFilter = v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelName(v string) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetRequestId(v string) *GetDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetResourceChangeDelivery(v *GetDeliveryChannelResponseBodyResourceChangeDelivery) *GetDeliveryChannelResponseBody {
	s.ResourceChangeDelivery = v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetResourceSnapshotDelivery(v *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) *GetDeliveryChannelResponseBody {
	s.ResourceSnapshotDelivery = v
	return s
}

type GetDeliveryChannelResponseBodyDeliveryChannelFilter struct {
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s GetDeliveryChannelResponseBodyDeliveryChannelFilter) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyDeliveryChannelFilter) SetResourceTypes(v []*string) *GetDeliveryChannelResponseBodyDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

type GetDeliveryChannelResponseBodyResourceChangeDelivery struct {
	// The Simple Log Service configurations.
	SlsProperties *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 1911422487776***:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceChangeDelivery) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) SetSlsProperties(v *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetArn(v string) *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetType(v string) *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetType = &v
	return s
}

type GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties struct {
	// The Alibaba Cloud Resource Name (ARN) of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type GetDeliveryChannelResponseBodyResourceSnapshotDelivery struct {
	// The custom expression.
	//
	// example:
	//
	// select 	- from resources limit 100;
	CustomExpression *string `json:"CustomExpression,omitempty" xml:"CustomExpression,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 09:00Z
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination.
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDelivery) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetCustomExpression(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetDeliveryTime(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetSlsProperties(v *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetArn(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetType(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

type GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type GetDeliveryChannelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetDeliveryChannelResponse) SetStatusCode(v int32) *GetDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeliveryChannelResponse) SetBody(v *GetDeliveryChannelResponseBody) *GetDeliveryChannelResponse {
	s.Body = v
	return s
}

type GetDeliveryChannelStatisticsRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetDeliveryChannelStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsRequest) SetDeliveryChannelId(v string) *GetDeliveryChannelStatisticsRequest {
	s.DeliveryChannelId = &v
	return s
}

type GetDeliveryChannelStatisticsResponseBody struct {
	// The statistics on the delivery channel.
	DeliveryChannelStatistics *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics `json:"DeliveryChannelStatistics,omitempty" xml:"DeliveryChannelStatistics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 80DF0610-504C-56D7-BDCF-7C92FD687***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeliveryChannelStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsResponseBody) SetDeliveryChannelStatistics(v *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) *GetDeliveryChannelStatisticsResponseBody {
	s.DeliveryChannelStatistics = v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBody) SetRequestId(v string) *GetDeliveryChannelStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The last delivery time of resource configuration change events.
	//
	// example:
	//
	// 2025-06-03T16:05:15Z
	LatestChangeDeliveryTime *string `json:"LatestChangeDeliveryTime,omitempty" xml:"LatestChangeDeliveryTime,omitempty"`
	// The last delivery time of scheduled resource snapshots.
	//
	// example:
	//
	// 2025-06-03T16:00:00Z
	LatestSnapshotDeliveryTime *string `json:"LatestSnapshotDeliveryTime,omitempty" xml:"LatestSnapshotDeliveryTime,omitempty"`
}

func (s GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelId(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelName(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestChangeDeliveryTime(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestChangeDeliveryTime = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestSnapshotDeliveryTime(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestSnapshotDeliveryTime = &v
	return s
}

type GetDeliveryChannelStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeliveryChannelStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeliveryChannelStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryChannelStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsResponse) SetHeaders(v map[string]*string) *GetDeliveryChannelStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetDeliveryChannelStatisticsResponse) SetStatusCode(v int32) *GetDeliveryChannelStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponse) SetBody(v *GetDeliveryChannelStatisticsResponseBody) *GetDeliveryChannelStatisticsResponse {
	s.Body = v
	return s
}

type GetExampleQueryRequest struct {
	// The ID of the template.
	//
	// >  You can call the [ListExampleQueries](~~ListExampleQueries~~) operation to obtain the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetExampleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryRequest) GoString() string {
	return s.String()
}

func (s *GetExampleQueryRequest) SetQueryId(v string) *GetExampleQueryRequest {
	s.QueryId = &v
	return s
}

type GetExampleQueryResponseBody struct {
	// The information about the sample query template.
	ExampleQuery *GetExampleQueryResponseBodyExampleQuery `json:"ExampleQuery,omitempty" xml:"ExampleQuery,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65DC00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExampleQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponseBody) SetExampleQuery(v *GetExampleQueryResponseBodyExampleQuery) *GetExampleQueryResponseBody {
	s.ExampleQuery = v
	return s
}

func (s *GetExampleQueryResponseBody) SetRequestId(v string) *GetExampleQueryResponseBody {
	s.RequestId = &v
	return s
}

type GetExampleQueryResponseBodyExampleQuery struct {
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// example:
	//
	// SELECT
	//
	//   resource_id,
	//
	//   resource_name,
	//
	//   region_id,
	//
	//   zone_id,
	//
	//   resource_type,
	//
	//   account_id,
	//
	//   create_time,
	//
	//   resource_group_id,
	//
	//   tags,
	//
	//   ip_addresses,
	//
	//   vpc_id,
	//
	//   v_switch_id
	//
	// FROM
	//
	//   resources
	//
	// ORDER BY
	//
	//   resource_type,
	//
	//   resource_id
	//
	// LIMIT
	//
	//   1000 OFFSET 0;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetExampleQueryResponseBodyExampleQuery) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryResponseBodyExampleQuery) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetDescription(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Description = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetExpression(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Expression = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetName(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Name = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetQueryId(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.QueryId = &v
	return s
}

type GetExampleQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExampleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExampleQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryResponse) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponse) SetHeaders(v map[string]*string) *GetExampleQueryResponse {
	s.Headers = v
	return s
}

func (s *GetExampleQueryResponse) SetStatusCode(v int32) *GetExampleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExampleQueryResponse) SetBody(v *GetExampleQueryResponseBody) *GetExampleQueryResponse {
	s.Body = v
	return s
}

type GetMultiAccountDeliveryChannelRequest struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetMultiAccountDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

type GetMultiAccountDeliveryChannelResponseBody struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE3EAB47-D3A6-5FEA-8353-31C8C0D36***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s GetMultiAccountDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelDescription(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelFilter(v *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelFilter = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelName(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetResourceChangeDelivery(v *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) *GetMultiAccountDeliveryChannelResponseBody {
	s.ResourceChangeDelivery = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetResourceSnapshotDelivery(v *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) *GetMultiAccountDeliveryChannelResponseBody {
	s.ResourceSnapshotDelivery = v
	return s
}

type GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter struct {
	// The effective account scopes of the delivery channel.
	AccountScopes []*string `json:"AccountScopes,omitempty" xml:"AccountScopes,omitempty" type:"Repeated"`
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) SetAccountScopes(v []*string) *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter {
	s.AccountScopes = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) SetResourceTypes(v []*string) *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

type GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery struct {
	// The Simple Log Service configurations.
	SlsProperties *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 1911422487776***:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) SetSlsProperties(v *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetType(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetType = &v
	return s
}

type GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties struct {
	// The Alibaba Cloud Resource Name (ARN) of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery struct {
	// The custom expression.
	//
	// example:
	//
	// select 	- from resources limit 100;
	CustomExpression *string `json:"CustomExpression,omitempty" xml:"CustomExpression,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 09:00Z
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination.
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetCustomExpression(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetDeliveryTime(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetSlsProperties(v *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetType(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

type GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type GetMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *GetMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponse) SetBody(v *GetMultiAccountDeliveryChannelResponseBody) *GetMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

type GetMultiAccountDeliveryChannelStatisticsRequest struct {
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetMultiAccountDeliveryChannelStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsRequest) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelStatisticsRequest {
	s.DeliveryChannelId = &v
	return s
}

type GetMultiAccountDeliveryChannelStatisticsResponseBody struct {
	// The statistics on the delivery channel.
	DeliveryChannelStatistics *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics `json:"DeliveryChannelStatistics,omitempty" xml:"DeliveryChannelStatistics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 80DF0610-504C-56D7-BDCF-7C92FD687***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultiAccountDeliveryChannelStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBody) SetDeliveryChannelStatistics(v *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) *GetMultiAccountDeliveryChannelStatisticsResponseBody {
	s.DeliveryChannelStatistics = v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBody) SetRequestId(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-multi-account-delivery
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The last delivery time of resource configuration change events.
	//
	// example:
	//
	// 2025-06-03T16:05:15Z
	LatestChangeDeliveryTime *string `json:"LatestChangeDeliveryTime,omitempty" xml:"LatestChangeDeliveryTime,omitempty"`
	// The last delivery time of scheduled resource snapshots.
	//
	// example:
	//
	// 2025-06-03T16:00:00Z
	LatestSnapshotDeliveryTime *string `json:"LatestSnapshotDeliveryTime,omitempty" xml:"LatestSnapshotDeliveryTime,omitempty"`
}

func (s GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelName(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestChangeDeliveryTime(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestChangeDeliveryTime = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestSnapshotDeliveryTime(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestSnapshotDeliveryTime = &v
	return s
}

type GetMultiAccountDeliveryChannelStatisticsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountDeliveryChannelStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountDeliveryChannelStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) SetHeaders(v map[string]*string) *GetMultiAccountDeliveryChannelStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) SetStatusCode(v int32) *GetMultiAccountDeliveryChannelStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) SetBody(v *GetMultiAccountDeliveryChannelStatisticsResponseBody) *GetMultiAccountDeliveryChannelStatisticsResponse {
	s.Body = v
	return s
}

type GetMultiAccountResourceCenterServiceStatusResponseBody struct {
	// The initialization status of the feature. Valid values:
	//
	// 	- Pending: The feature is being initialized.
	//
	// 	- Finished: The feature is initialized.
	//
	// example:
	//
	// Pending
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 81671397-1425-51F1-A144-4799E01BEBFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// example:
	//
	// Enabled
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

type GetMultiAccountResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetBody(v *GetMultiAccountResourceCenterServiceStatusResponseBody) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

type GetMultiAccountResourceConfigurationRequest struct {
	// The ID of the management account or member of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1619302****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetMultiAccountResourceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationRequest) SetAccountId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceRegionId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceType(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBody struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 1619302****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The configurations of the resource.
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2023-02-14T03:12:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2023-09-18T07:04:21Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// > Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B2DCC08B-C12A-5705-879C-5A1450016156
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-acfmzy6d****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// test_resource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags []*GetMultiAccountResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID of the resource.
	//
	// example:
	//
	// cn-shanghai-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetAccountId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetMultiAccountResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetCreateTime(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetExpireTime(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetIpAddressAttributes(v []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) *GetMultiAccountResourceConfigurationResponseBody {
	s.IpAddressAttributes = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetMultiAccountResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRegionId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRequestId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceName(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceType(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetTags(v []*GetMultiAccountResourceConfigurationResponseBodyTags) *GetMultiAccountResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetZoneId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 172.27.199.42
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetIpAddress(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetNetworkType(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetVersion(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.Version = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBodyTags struct {
	// The key of tag N.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetKey(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetValue(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

type GetMultiAccountResourceConfigurationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetStatusCode(v int32) *GetMultiAccountResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetBody(v *GetMultiAccountResourceConfigurationResponseBody) *GetMultiAccountResourceConfigurationResponse {
	s.Body = v
	return s
}

type GetMultiAccountResourceCountsRequest struct {
	Filter []*GetMultiAccountResourceCountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s GetMultiAccountResourceCountsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsRequest) SetFilter(v []*GetMultiAccountResourceCountsRequestFilter) *GetMultiAccountResourceCountsRequest {
	s.Filter = v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) SetGroupByKey(v string) *GetMultiAccountResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) SetScope(v string) *GetMultiAccountResourceCountsRequest {
	s.Scope = &v
	return s
}

type GetMultiAccountResourceCountsRequestFilter struct {
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Equals
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Value     []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetMultiAccountResourceCountsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCountsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetKey(v string) *GetMultiAccountResourceCountsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetMatchType(v string) *GetMultiAccountResourceCountsRequestFilter {
	s.MatchType = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetValue(v []*string) *GetMultiAccountResourceCountsRequestFilter {
	s.Value = v
	return s
}

type GetMultiAccountResourceCountsResponseBody struct {
	Filters []*GetMultiAccountResourceCountsResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId      *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceCounts []*GetMultiAccountResourceCountsResponseBodyResourceCounts `json:"ResourceCounts,omitempty" xml:"ResourceCounts,omitempty" type:"Repeated"`
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s GetMultiAccountResourceCountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBody) SetFilters(v []*GetMultiAccountResourceCountsResponseBodyFilters) *GetMultiAccountResourceCountsResponseBody {
	s.Filters = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetGroupByKey(v string) *GetMultiAccountResourceCountsResponseBody {
	s.GroupByKey = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetRequestId(v string) *GetMultiAccountResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetResourceCounts(v []*GetMultiAccountResourceCountsResponseBodyResourceCounts) *GetMultiAccountResourceCountsResponseBody {
	s.ResourceCounts = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetScope(v string) *GetMultiAccountResourceCountsResponseBody {
	s.Scope = &v
	return s
}

type GetMultiAccountResourceCountsResponseBodyFilters struct {
	// example:
	//
	// RegionId
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetMultiAccountResourceCountsResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) SetKey(v string) *GetMultiAccountResourceCountsResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) SetValues(v []*string) *GetMultiAccountResourceCountsResponseBodyFilters {
	s.Values = v
	return s
}

type GetMultiAccountResourceCountsResponseBodyResourceCounts struct {
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// ACS::ECS::NetworkInterface
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetMultiAccountResourceCountsResponseBodyResourceCounts) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBodyResourceCounts) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) SetCount(v int64) *GetMultiAccountResourceCountsResponseBodyResourceCounts {
	s.Count = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) SetGroupName(v string) *GetMultiAccountResourceCountsResponseBodyResourceCounts {
	s.GroupName = &v
	return s
}

type GetMultiAccountResourceCountsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceCountsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceCountsResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceCountsResponse) SetStatusCode(v int32) *GetMultiAccountResourceCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponse) SetBody(v *GetMultiAccountResourceCountsResponseBody) *GetMultiAccountResourceCountsResponse {
	s.Body = v
	return s
}

type GetResourceCenterServiceStatusResponseBody struct {
	// The initialization status of the service. Valid values:
	//
	// 	- Pending: The service is being initialized.
	//
	// 	- Finished: The service is initialized.
	//
	// example:
	//
	// Pending
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD5F848D-CCDC-5464-93E1-4BA50A4826DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Enabled: The service is activated.
	//
	// 	- Disabled: The service is deactivated.
	//
	// example:
	//
	// Enabled
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetResourceCenterServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

type GetResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceCenterServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetBody(v *GetResourceCenterServiceStatusResponseBody) *GetResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

type GetResourceConfigurationRequest struct {
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-bp1kyg72m****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// For more information about the resource types supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationRequest) SetResourceId(v string) *GetResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceRegionId(v string) *GetResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceType(v string) *GetResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

type GetResourceConfigurationResponseBody struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The configurations of the resource.
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2021-07-30T09:20:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*GetResourceConfigurationResponseBodyIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// > Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F1CE0D52-32DA-531A-87A4-B9A5B68D5D8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-acfmv4k****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// vtb-uf6978gdqbi****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::VPC::VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags []*GetResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID of the resource.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetResourceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBody) SetAccountId(v string) *GetResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetCreateTime(v string) *GetResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetExpireTime(v string) *GetResourceConfigurationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetIpAddressAttributes(v []*GetResourceConfigurationResponseBodyIpAddressAttributes) *GetResourceConfigurationResponseBody {
	s.IpAddressAttributes = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRegionId(v string) *GetResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRequestId(v string) *GetResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceName(v string) *GetResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceType(v string) *GetResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetTags(v []*GetResourceConfigurationResponseBodyTags) *GetResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetZoneId(v string) *GetResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

type GetResourceConfigurationResponseBodyIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.1.2
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetResourceConfigurationResponseBodyIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBodyIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetIpAddress(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetNetworkType(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetVersion(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.Version = &v
	return s
}

type GetResourceConfigurationResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceConfigurationResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBodyTags) SetKey(v string) *GetResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyTags) SetValue(v string) *GetResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

type GetResourceConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetResourceConfigurationResponse) SetStatusCode(v int32) *GetResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceConfigurationResponse) SetBody(v *GetResourceConfigurationResponseBody) *GetResourceConfigurationResponse {
	s.Body = v
	return s
}

type GetResourceCountsRequest struct {
	// The filter conditions.
	Filter []*GetResourceCountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The dimension by which resources are queried. Valid values:
	//
	// 	- ResourceType
	//
	// 	- Region
	//
	// 	- ResourceGroupId
	//
	// 	- TagKey
	//
	// 	- TagValue
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
}

func (s GetResourceCountsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequest) SetFilter(v []*GetResourceCountsRequestFilter) *GetResourceCountsRequest {
	s.Filter = v
	return s
}

func (s *GetResourceCountsRequest) SetGroupByKey(v string) *GetResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

type GetResourceCountsRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetResourceCountsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequestFilter) SetKey(v string) *GetResourceCountsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetMatchType(v string) *GetResourceCountsRequestFilter {
	s.MatchType = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetValue(v []*string) *GetResourceCountsRequestFilter {
	s.Value = v
	return s
}

type GetResourceCountsResponseBody struct {
	// The filter conditions.
	Filters []*GetResourceCountsResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The dimension by which resources are queried.
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D98D9B0-318D-56A4-910C-93B5F945AF2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The numbers of resources.
	ResourceCounts []*GetResourceCountsResponseBodyResourceCounts `json:"ResourceCounts,omitempty" xml:"ResourceCounts,omitempty" type:"Repeated"`
}

func (s GetResourceCountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBody) SetFilters(v []*GetResourceCountsResponseBodyFilters) *GetResourceCountsResponseBody {
	s.Filters = v
	return s
}

func (s *GetResourceCountsResponseBody) SetGroupByKey(v string) *GetResourceCountsResponseBody {
	s.GroupByKey = &v
	return s
}

func (s *GetResourceCountsResponseBody) SetRequestId(v string) *GetResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCountsResponseBody) SetResourceCounts(v []*GetResourceCountsResponseBodyResourceCounts) *GetResourceCountsResponseBody {
	s.ResourceCounts = v
	return s
}

type GetResourceCountsResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetResourceCountsResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBodyFilters) SetKey(v string) *GetResourceCountsResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *GetResourceCountsResponseBodyFilters) SetValues(v []*string) *GetResourceCountsResponseBodyFilters {
	s.Values = v
	return s
}

type GetResourceCountsResponseBodyResourceCounts struct {
	// The number of resources.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The group name.
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetResourceCountsResponseBodyResourceCounts) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponseBodyResourceCounts) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBodyResourceCounts) SetCount(v int64) *GetResourceCountsResponseBodyResourceCounts {
	s.Count = &v
	return s
}

func (s *GetResourceCountsResponseBodyResourceCounts) SetGroupName(v string) *GetResourceCountsResponseBodyResourceCounts {
	s.GroupName = &v
	return s
}

type GetResourceCountsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceCountsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponse) SetHeaders(v map[string]*string) *GetResourceCountsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCountsResponse) SetStatusCode(v int32) *GetResourceCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCountsResponse) SetBody(v *GetResourceCountsResponseBody) *GetResourceCountsResponse {
	s.Body = v
	return s
}

type GetSavedQueryRequest struct {
	// The template ID.
	//
	// >  You can call the [ListSavedQueries](~~ListSavedQueries~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *GetSavedQueryRequest) SetQueryId(v string) *GetSavedQueryRequest {
	s.QueryId = &v
	return s
}

type GetSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6D98D9B0-318D-56A4-910C-93B5F945AF2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the template.
	SavedQuery *GetSavedQueryResponseBodySavedQuery `json:"SavedQuery,omitempty" xml:"SavedQuery,omitempty" type:"Struct"`
}

func (s GetSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponseBody) SetRequestId(v string) *GetSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavedQueryResponseBody) SetSavedQuery(v *GetSavedQueryResponseBodySavedQuery) *GetSavedQueryResponseBody {
	s.SavedQuery = v
	return s
}

type GetSavedQueryResponseBodySavedQuery struct {
	// The time when the template was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-30T01:43:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The time when the template was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-30T01:43:16Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSavedQueryResponseBodySavedQuery) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryResponseBodySavedQuery) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponseBodySavedQuery) SetCreateTime(v string) *GetSavedQueryResponseBodySavedQuery {
	s.CreateTime = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetDescription(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Description = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetExpression(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Expression = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetName(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Name = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetQueryId(v string) *GetSavedQueryResponseBodySavedQuery {
	s.QueryId = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetUpdateTime(v string) *GetSavedQueryResponseBodySavedQuery {
	s.UpdateTime = &v
	return s
}

type GetSavedQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponse) SetHeaders(v map[string]*string) *GetSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *GetSavedQueryResponse) SetStatusCode(v int32) *GetSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavedQueryResponse) SetBody(v *GetSavedQueryResponseBody) *GetSavedQueryResponse {
	s.Body = v
	return s
}

type ListDeliveryChannelsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the MaxResults parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDeliveryChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsRequest) SetMaxResults(v int32) *ListDeliveryChannelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeliveryChannelsRequest) SetNextToken(v string) *ListDeliveryChannelsRequest {
	s.NextToken = &v
	return s
}

type ListDeliveryChannelsResponseBody struct {
	// The delivery channels.
	DeliveryChannels []*ListDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17502A1B-7026-5551-8E1C-CCABD0CBC***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeliveryChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListDeliveryChannelsResponseBodyDeliveryChannels) *ListDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListDeliveryChannelsResponseBody) SetMaxResults(v int32) *ListDeliveryChannelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDeliveryChannelsResponseBody) SetNextToken(v string) *ListDeliveryChannelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDeliveryChannelsResponseBody) SetRequestId(v string) *ListDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeliveryChannelsResponseBodyDeliveryChannels struct {
	// The time when the delivery channel was created.
	//
	// example:
	//
	// 2024-06-20T08:46:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-4m6ffpkgu***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
}

func (s ListDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetCreateTime(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.CreateTime = &v
	return s
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelDescription(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

type ListDeliveryChannelsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeliveryChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryChannelsResponse) SetStatusCode(v int32) *ListDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryChannelsResponse) SetBody(v *ListDeliveryChannelsResponseBody) *ListDeliveryChannelsResponse {
	s.Body = v
	return s
}

type ListExampleQueriesRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExampleQueriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesRequest) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesRequest) SetMaxResults(v string) *ListExampleQueriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExampleQueriesRequest) SetNextToken(v string) *ListExampleQueriesRequest {
	s.NextToken = &v
	return s
}

type ListExampleQueriesResponseBody struct {
	// The information about the sample query templates.
	ExampleQueries []*ListExampleQueriesResponseBodyExampleQueries `json:"ExampleQueries,omitempty" xml:"ExampleQueries,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExampleQueriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponseBody) SetExampleQueries(v []*ListExampleQueriesResponseBodyExampleQueries) *ListExampleQueriesResponseBody {
	s.ExampleQueries = v
	return s
}

func (s *ListExampleQueriesResponseBody) SetMaxResults(v string) *ListExampleQueriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExampleQueriesResponseBody) SetNextToken(v string) *ListExampleQueriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExampleQueriesResponseBody) SetRequestId(v string) *ListExampleQueriesResponseBody {
	s.RequestId = &v
	return s
}

type ListExampleQueriesResponseBodyExampleQueries struct {
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s ListExampleQueriesResponseBodyExampleQueries) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesResponseBodyExampleQueries) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetDescription(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.Description = &v
	return s
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetName(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.Name = &v
	return s
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetQueryId(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.QueryId = &v
	return s
}

type ListExampleQueriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExampleQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExampleQueriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponse) SetHeaders(v map[string]*string) *ListExampleQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListExampleQueriesResponse) SetStatusCode(v int32) *ListExampleQueriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExampleQueriesResponse) SetBody(v *ListExampleQueriesResponseBody) *ListExampleQueriesResponse {
	s.Body = v
	return s
}

type ListFiltersResponseBody struct {
	// The name of the default filter.
	//
	// example:
	//
	// My Filters
	DefaultFilterName *string `json:"DefaultFilterName,omitempty" xml:"DefaultFilterName,omitempty"`
	// The configurations of the filter.
	Filters []*ListFiltersResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AA39FB9C-CB74-5E73-8DFE-3A2B096F0759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFiltersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFiltersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFiltersResponseBody) SetDefaultFilterName(v string) *ListFiltersResponseBody {
	s.DefaultFilterName = &v
	return s
}

func (s *ListFiltersResponseBody) SetFilters(v []*ListFiltersResponseBodyFilters) *ListFiltersResponseBody {
	s.Filters = v
	return s
}

func (s *ListFiltersResponseBody) SetRequestId(v string) *ListFiltersResponseBody {
	s.RequestId = &v
	return s
}

type ListFiltersResponseBodyFilters struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	//
	// example:
	//
	// My devices
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s ListFiltersResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s ListFiltersResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *ListFiltersResponseBodyFilters) SetFilterConfiguration(v string) *ListFiltersResponseBodyFilters {
	s.FilterConfiguration = &v
	return s
}

func (s *ListFiltersResponseBodyFilters) SetFilterName(v string) *ListFiltersResponseBodyFilters {
	s.FilterName = &v
	return s
}

type ListFiltersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFiltersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFiltersResponse) GoString() string {
	return s.String()
}

func (s *ListFiltersResponse) SetHeaders(v map[string]*string) *ListFiltersResponse {
	s.Headers = v
	return s
}

func (s *ListFiltersResponse) SetStatusCode(v int32) *ListFiltersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFiltersResponse) SetBody(v *ListFiltersResponseBody) *ListFiltersResponse {
	s.Body = v
	return s
}

type ListMultiAccountDeliveryChannelsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the MaxResults parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsRequest) SetMaxResults(v int32) *ListMultiAccountDeliveryChannelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsRequest) SetNextToken(v string) *ListMultiAccountDeliveryChannelsRequest {
	s.NextToken = &v
	return s
}

type ListMultiAccountDeliveryChannelsResponseBody struct {
	// The delivery channels.
	DeliveryChannels []*ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17502A1B-7026-5551-8E1C-CCABD0CBC***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) *ListMultiAccountDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetMaxResults(v int32) *ListMultiAccountDeliveryChannelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetNextToken(v string) *ListMultiAccountDeliveryChannelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBody) SetRequestId(v string) *ListMultiAccountDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

type ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels struct {
	// The time when the delivery channel was created.
	//
	// example:
	//
	// 2023-08-17T00:23:55Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-0bzhsqpnk***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-multi-account-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetCreateTime(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.CreateTime = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelDescription(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListMultiAccountDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

type ListMultiAccountDeliveryChannelsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListMultiAccountDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponse) SetStatusCode(v int32) *ListMultiAccountDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsResponse) SetBody(v *ListMultiAccountDeliveryChannelsResponseBody) *ListMultiAccountDeliveryChannelsResponse {
	s.Body = v
	return s
}

type ListMultiAccountResourceGroupsRequest struct {
	// The ID of the management account or member of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1394339739****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAS2Nboi3t4xGrdlG5/Ks/Q1xPG9jzviYEuZydevXIkgF
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of resource groups.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsRequest) SetAccountId(v string) *ListMultiAccountResourceGroupsRequest {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetMaxResults(v int32) *ListMultiAccountResourceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetNextToken(v string) *ListMultiAccountResourceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetResourceGroupIds(v []*string) *ListMultiAccountResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

type ListMultiAccountResourceGroupsResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAU5VsT9R1adMTuz9GzginZ3Y+7Y/5JATS+6q5GK9kT75
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FF0A66E-781F-51EE-9531-928F197558F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource groups.
	ResourceGroups []*ListMultiAccountResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetNextToken(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetRequestId(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetResourceGroups(v []*ListMultiAccountResourceGroupsResponseBodyResourceGroups) *ListMultiAccountResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

type ListMultiAccountResourceGroupsResponseBodyResourceGroups struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 1394339739****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	//
	// example:
	//
	// group1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetAccountId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetCreateDate(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.CreateDate = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetDisplayName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.DisplayName = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Id = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Name = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetStatus(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Status = &v
	return s
}

type ListMultiAccountResourceGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponse) SetHeaders(v map[string]*string) *ListMultiAccountResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetStatusCode(v int32) *ListMultiAccountResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetBody(v *ListMultiAccountResourceGroupsResponseBody) *ListMultiAccountResourceGroupsResponse {
	s.Body = v
	return s
}

type ListMultiAccountResourceRelationshipsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The filter conditions for resources associated with the resource.
	RelatedResourceFilter []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter `json:"RelatedResourceFilter,omitempty" xml:"RelatedResourceFilter,omitempty" type:"Repeated"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The search scope. Valid values:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to query the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to query the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to query the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetNextToken(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetRegionId(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.RegionId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetRelatedResourceFilter(v []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) *ListMultiAccountResourceRelationshipsRequest {
	s.RelatedResourceFilter = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetResourceId(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetResourceType(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetScope(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.Scope = &v
	return s
}

type ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RelatedResourceRegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetKey(v string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.Key = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetMatchType(v string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetValue(v []*string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.Value = v
	return s
}

type ListMultiAccountResourceRelationshipsResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BCAB07BA-82FA-5DC0-9322-FB7ED726481D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource relationships.
	ResourceRelationships []*ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships `json:"ResourceRelationships,omitempty" xml:"ResourceRelationships,omitempty" type:"Repeated"`
	// The search scope. Valid values:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched.
	//
	// 	- ID of a member: Resources within the member are searched.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetNextToken(v string) *ListMultiAccountResourceRelationshipsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetRequestId(v string) *ListMultiAccountResourceRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetResourceRelationships(v []*ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) *ListMultiAccountResourceRelationshipsResponseBody {
	s.ResourceRelationships = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetScope(v string) *ListMultiAccountResourceRelationshipsResponseBody {
	s.Scope = &v
	return s
}

type ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships struct {
	// The ID of the management account or member.
	//
	// example:
	//
	// 193396142051****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the associated resource.
	//
	// example:
	//
	// vpc-uf6m5okksddm6c9lh7***
	RelatedResourceId *string `json:"RelatedResourceId,omitempty" xml:"RelatedResourceId,omitempty"`
	// The region ID of the associated resource.
	//
	// example:
	//
	// cn-shanghai
	RelatedResourceRegionId *string `json:"RelatedResourceRegionId,omitempty" xml:"RelatedResourceRegionId,omitempty"`
	// The type of the associated resource.
	//
	// example:
	//
	// ACS::VPC::VPC
	RelatedResourceType *string `json:"RelatedResourceType,omitempty" xml:"RelatedResourceType,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetAccountId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRegionId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RegionId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceRegionId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceRegionId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceType(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetResourceId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetResourceType(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceType = &v
	return s
}

type ListMultiAccountResourceRelationshipsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountResourceRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsResponse) SetHeaders(v map[string]*string) *ListMultiAccountResourceRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponse) SetStatusCode(v int32) *ListMultiAccountResourceRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponse) SetBody(v *ListMultiAccountResourceRelationshipsResponseBody) *ListMultiAccountResourceRelationshipsResponse {
	s.Body = v
	return s
}

type ListMultiAccountTagKeysRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. The value of this parameter can be one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID of the resource directory. The ID is indicated by the `ResourceDirectoryId` parameter in the response of the operation.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID of the Root folder. The ID is indicated by the `RootFolderId` parameter in the response of the operation.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID of the folder. The ID is indicated by the `FolderId` parameter in the response of the operation.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID of the member. The ID is indicated by the `AccountId` parameter in the response of the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListMultiAccountTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysRequest) SetMatchType(v string) *ListMultiAccountTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetMaxResults(v int32) *ListMultiAccountTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetNextToken(v string) *ListMultiAccountTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetScope(v string) *ListMultiAccountTagKeysRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetTagKey(v string) *ListMultiAccountTagKeysRequest {
	s.TagKey = &v
	return s
}

type ListMultiAccountTagKeysResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FA6086F9-6363-51A5-A507-88E3201EBCCB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponseBody) SetNextToken(v string) *ListMultiAccountTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetRequestId(v string) *ListMultiAccountTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetTagKeys(v []*string) *ListMultiAccountTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListMultiAccountTagKeysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetStatusCode(v int32) *ListMultiAccountTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetBody(v *ListMultiAccountTagKeysResponseBody) *ListMultiAccountTagKeysResponse {
	s.Body = v
	return s
}

type ListMultiAccountTagValuesRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. You can set the value to one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListMultiAccountTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesRequest) SetMatchType(v string) *ListMultiAccountTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetMaxResults(v int32) *ListMultiAccountTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetNextToken(v string) *ListMultiAccountTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetScope(v string) *ListMultiAccountTagValuesRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagKey(v string) *ListMultiAccountTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagValue(v string) *ListMultiAccountTagValuesRequest {
	s.TagValue = &v
	return s
}

type ListMultiAccountTagValuesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65DC00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponseBody) SetNextToken(v string) *ListMultiAccountTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetRequestId(v string) *ListMultiAccountTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetTagValues(v []*string) *ListMultiAccountTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListMultiAccountTagValuesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetStatusCode(v int32) *ListMultiAccountTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetBody(v *ListMultiAccountTagValuesResponseBody) *ListMultiAccountTagValuesResponse {
	s.Body = v
	return s
}

type ListResourceRelationshipsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The filter conditions for resources associated with the resource.
	RelatedResourceFilter []*ListResourceRelationshipsRequestRelatedResourceFilter `json:"RelatedResourceFilter,omitempty" xml:"RelatedResourceFilter,omitempty" type:"Repeated"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceRelationshipsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsRequest) SetMaxResults(v int32) *ListResourceRelationshipsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetNextToken(v string) *ListResourceRelationshipsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetRegionId(v string) *ListResourceRelationshipsRequest {
	s.RegionId = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetRelatedResourceFilter(v []*ListResourceRelationshipsRequestRelatedResourceFilter) *ListResourceRelationshipsRequest {
	s.RelatedResourceFilter = v
	return s
}

func (s *ListResourceRelationshipsRequest) SetResourceId(v string) *ListResourceRelationshipsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetResourceType(v string) *ListResourceRelationshipsRequest {
	s.ResourceType = &v
	return s
}

type ListResourceRelationshipsRequestRelatedResourceFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RelatedResourceRegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListResourceRelationshipsRequestRelatedResourceFilter) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsRequestRelatedResourceFilter) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) SetKey(v string) *ListResourceRelationshipsRequestRelatedResourceFilter {
	s.Key = &v
	return s
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) SetMatchType(v string) *ListResourceRelationshipsRequestRelatedResourceFilter {
	s.MatchType = &v
	return s
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) SetValue(v []*string) *ListResourceRelationshipsRequestRelatedResourceFilter {
	s.Value = v
	return s
}

type ListResourceRelationshipsResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 682A3004-38E3-5122-9A11-CCDFAB9C3C4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource relationships.
	ResourceRelationships []*ListResourceRelationshipsResponseBodyResourceRelationships `json:"ResourceRelationships,omitempty" xml:"ResourceRelationships,omitempty" type:"Repeated"`
}

func (s ListResourceRelationshipsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponseBody) SetMaxResults(v int32) *ListResourceRelationshipsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetNextToken(v string) *ListResourceRelationshipsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetRequestId(v string) *ListResourceRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetResourceRelationships(v []*ListResourceRelationshipsResponseBodyResourceRelationships) *ListResourceRelationshipsResponseBody {
	s.ResourceRelationships = v
	return s
}

type ListResourceRelationshipsResponseBodyResourceRelationships struct {
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the associated resource.
	//
	// example:
	//
	// vpc-uf6m5okksddm6c9lh7***
	RelatedResourceId *string `json:"RelatedResourceId,omitempty" xml:"RelatedResourceId,omitempty"`
	// The region ID of the associated resource.
	//
	// example:
	//
	// cn-shanghai
	RelatedResourceRegionId *string `json:"RelatedResourceRegionId,omitempty" xml:"RelatedResourceRegionId,omitempty"`
	// The type of the associated resource.
	//
	// example:
	//
	// ACS::VPC::VPC
	RelatedResourceType *string `json:"RelatedResourceType,omitempty" xml:"RelatedResourceType,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceRelationshipsResponseBodyResourceRelationships) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsResponseBodyResourceRelationships) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRegionId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RegionId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceRegionId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceRegionId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceType(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceType = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetResourceId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetResourceType(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceType = &v
	return s
}

type ListResourceRelationshipsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceRelationshipsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponse) SetHeaders(v map[string]*string) *ListResourceRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceRelationshipsResponse) SetStatusCode(v int32) *ListResourceRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceRelationshipsResponse) SetBody(v *ListResourceRelationshipsResponseBody) *ListResourceRelationshipsResponse {
	s.Body = v
	return s
}

type ListResourceTypesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The query conditions.
	Query []*string `json:"Query,omitempty" xml:"Query,omitempty" type:"Repeated"`
	// The resource type.
	//
	// For more information about the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) SetAcceptLanguage(v string) *ListResourceTypesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListResourceTypesRequest) SetQuery(v []*string) *ListResourceTypesRequest {
	s.Query = v
	return s
}

func (s *ListResourceTypesRequest) SetResourceType(v string) *ListResourceTypesRequest {
	s.ResourceType = &v
	return s
}

type ListResourceTypesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E5556E4C-479A-5BBB-B325-F07563E7E917
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource types.
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

type ListResourceTypesResponseBodyResourceTypes struct {
	// The code mapping of the resource type.
	CodeMapping *ListResourceTypesResponseBodyResourceTypesCodeMapping `json:"CodeMapping,omitempty" xml:"CodeMapping,omitempty" type:"Struct"`
	// The supported filter conditions.
	FilterKeys []*string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty" type:"Repeated"`
	// The name of the Alibaba Cloud service.
	//
	// example:
	//
	// Container Service for Kubernetes
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The name of supported related resource types.
	RelatedResourceTypes []*string `json:"RelatedResourceTypes,omitempty" xml:"RelatedResourceTypes,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the resource type.
	//
	// example:
	//
	// Cluster
	ResourceTypeName *string `json:"ResourceTypeName,omitempty" xml:"ResourceTypeName,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetCodeMapping(v *ListResourceTypesResponseBodyResourceTypesCodeMapping) *ListResourceTypesResponseBodyResourceTypes {
	s.CodeMapping = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetFilterKeys(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.FilterKeys = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProductName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ProductName = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetRelatedResourceTypes(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.RelatedResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceTypeName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceTypeName = &v
	return s
}

type ListResourceTypesResponseBodyResourceTypesCodeMapping struct {
	// The resource group.
	//
	// example:
	//
	// cs.cluster
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The tag.
	//
	// example:
	//
	// cs.cluster
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetResourceGroup(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.ResourceGroup = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetTag(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.Tag = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetStatusCode(v int32) *ListResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type ListSavedQueriesRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSavedQueriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesRequest) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesRequest) SetMaxResults(v string) *ListSavedQueriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSavedQueriesRequest) SetNextToken(v string) *ListSavedQueriesRequest {
	s.NextToken = &v
	return s
}

type ListSavedQueriesResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the custom query templates.
	SavedQueries []*ListSavedQueriesResponseBodySavedQueries `json:"SavedQueries,omitempty" xml:"SavedQueries,omitempty" type:"Repeated"`
}

func (s ListSavedQueriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponseBody) SetMaxResults(v string) *ListSavedQueriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetNextToken(v string) *ListSavedQueriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetRequestId(v string) *ListSavedQueriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetSavedQueries(v []*ListSavedQueriesResponseBodySavedQueries) *ListSavedQueriesResponseBody {
	s.SavedQueries = v
	return s
}

type ListSavedQueriesResponseBodySavedQueries struct {
	// The time when the template was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-13T05:50:35Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template name.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The time when the template was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-14 10:27:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSavedQueriesResponseBodySavedQueries) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesResponseBodySavedQueries) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetCreateTime(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.CreateTime = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetDescription(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.Description = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetName(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.Name = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetQueryId(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.QueryId = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetUpdateTime(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.UpdateTime = &v
	return s
}

type ListSavedQueriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSavedQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSavedQueriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponse) SetHeaders(v map[string]*string) *ListSavedQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListSavedQueriesResponse) SetStatusCode(v int32) *ListSavedQueriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSavedQueriesResponse) SetBody(v *ListSavedQueriesResponseBody) *ListSavedQueriesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// AAAAAUYb00R0gHZBE8FVDeoh2ME93VeeEPUHs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetMatchType(v string) *ListTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetTagKey(v string) *ListTagKeysRequest {
	s.TagKey = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAUDnubHKJbVTCdlIGYUPtsu3EoN3bfdgjDA****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44C8A952-D6B0-5BC8-82D5-93BA02E26F2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*string) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetMatchType(v string) *ListTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) SetTagValue(v string) *ListTagValuesRequest {
	s.TagValue = &v
	return s
}

type ListTagValuesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C1840B83-1193-5E83-AFA6-4B8D303E29F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type SearchMultiAccountResourcesRequest struct {
	// The filter conditions.
	Filter []*SearchMultiAccountResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.``
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. You can set the value to one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The method that is used to sort the entries returned.
	SortCriterion *SearchMultiAccountResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchMultiAccountResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequest) SetFilter(v []*SearchMultiAccountResourcesRequestFilter) *SearchMultiAccountResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetMaxResults(v int32) *SearchMultiAccountResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetNextToken(v string) *SearchMultiAccountResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetScope(v string) *SearchMultiAccountResourcesRequest {
	s.Scope = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetSortCriterion(v *SearchMultiAccountResourcesRequestSortCriterion) *SearchMultiAccountResourcesRequest {
	s.SortCriterion = v
	return s
}

type SearchMultiAccountResourcesRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// ResourceGroupId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchMultiAccountResourcesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestFilter) SetKey(v string) *SearchMultiAccountResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetMatchType(v string) *SearchMultiAccountResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetValue(v []*string) *SearchMultiAccountResourcesRequestFilter {
	s.Value = v
	return s
}

type SearchMultiAccountResourcesRequestSortCriterion struct {
	// The attribute based on which the entries are sorted.
	//
	// The value CreateTime indicates the creation time of resources.
	//
	// example:
	//
	// CreateTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- ASC: The entries are sorted in ascending order. This value is the default value.
	//
	// 	- DESC: The entries are sorted in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s SearchMultiAccountResourcesRequestSortCriterion) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetKey(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetOrder(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

type SearchMultiAccountResourcesResponseBody struct {
	// The filter conditions.
	Filters []*SearchMultiAccountResourcesResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources []*SearchMultiAccountResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The search scope.
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched.
	//
	// 	- ID of a member: Resources within the member are searched.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBody) SetFilters(v []*SearchMultiAccountResourcesResponseBodyFilters) *SearchMultiAccountResourcesResponseBody {
	s.Filters = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetMaxResults(v int32) *SearchMultiAccountResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetNextToken(v string) *SearchMultiAccountResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetRequestId(v string) *SearchMultiAccountResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetResources(v []*SearchMultiAccountResourcesResponseBodyResources) *SearchMultiAccountResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetScope(v string) *SearchMultiAccountResourcesResponseBody {
	s.Scope = &v
	return s
}

type SearchMultiAccountResourcesResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SearchMultiAccountResourcesResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetKey(v string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetMatchType(v string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.MatchType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetValues(v []*string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.Values = v
	return s
}

type SearchMultiAccountResourcesResponseBodyResources struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource was created.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2023-06-14T14:35:45Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vtb-bp11lbh452fr8940s****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*SearchMultiAccountResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetAccountId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetCreateTime(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetExpireTime(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetIpAddressAttributes(v []*SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) *SearchMultiAccountResourcesResponseBodyResources {
	s.IpAddressAttributes = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetIpAddresses(v []*string) *SearchMultiAccountResourcesResponseBodyResources {
	s.IpAddresses = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetRegionId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceGroupId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceName(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceType(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetTags(v []*SearchMultiAccountResourcesResponseBodyResourcesTags) *SearchMultiAccountResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetZoneId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.1.2
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetIpAddress(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetNetworkType(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetVersion(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.Version = &v
	return s
}

type SearchMultiAccountResourcesResponseBodyResourcesTags struct {
	// The key of tag N.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetKey(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetValue(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

type SearchMultiAccountResourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMultiAccountResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMultiAccountResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponse) SetHeaders(v map[string]*string) *SearchMultiAccountResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetStatusCode(v int32) *SearchMultiAccountResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetBody(v *SearchMultiAccountResourcesResponseBody) *SearchMultiAccountResourcesResponse {
	s.Body = v
	return s
}

type SearchResourcesRequest struct {
	// The filter conditions.
	Filter []*SearchResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The method that is used to sort the entries returned.
	SortCriterion *SearchResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequest) SetFilter(v []*SearchResourcesRequestFilter) *SearchResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchResourcesRequest) SetMaxResults(v int32) *SearchResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchResourcesRequest) SetNextToken(v string) *SearchResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchResourcesRequest) SetResourceGroupId(v string) *SearchResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchResourcesRequest) SetSortCriterion(v *SearchResourcesRequestSortCriterion) *SearchResourcesRequest {
	s.SortCriterion = v
	return s
}

type SearchResourcesRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchResourcesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestFilter) SetKey(v string) *SearchResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetMatchType(v string) *SearchResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetValue(v []*string) *SearchResourcesRequestFilter {
	s.Value = v
	return s
}

type SearchResourcesRequestSortCriterion struct {
	// The attribute based on which the entries are sorted.
	//
	// The value CreateTime indicates the creation time of resources.
	//
	// example:
	//
	// CreateTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- ASC: The entries are sorted in ascending order. This value is the default value.
	//
	// 	- DESC: The entries are sorted in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s SearchResourcesRequestSortCriterion) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestSortCriterion) SetKey(v string) *SearchResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestSortCriterion) SetOrder(v string) *SearchResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

type SearchResourcesResponseBody struct {
	// The filter conditions.
	Filters []*SearchResourcesResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources []*SearchResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBody) SetFilters(v []*SearchResourcesResponseBodyFilters) *SearchResourcesResponseBody {
	s.Filters = v
	return s
}

func (s *SearchResourcesResponseBody) SetMaxResults(v int32) *SearchResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchResourcesResponseBody) SetNextToken(v string) *SearchResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchResourcesResponseBody) SetRequestId(v string) *SearchResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchResourcesResponseBody) SetResources(v []*SearchResourcesResponseBodyResources) *SearchResourcesResponseBody {
	s.Resources = v
	return s
}

type SearchResourcesResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyFilters) SetKey(v string) *SearchResourcesResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *SearchResourcesResponseBodyFilters) SetMatchType(v string) *SearchResourcesResponseBodyFilters {
	s.MatchType = &v
	return s
}

func (s *SearchResourcesResponseBodyFilters) SetValues(v []*string) *SearchResourcesResponseBodyFilters {
	s.Values = v
	return s
}

type SearchResourcesResponseBodyResources struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource was created.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2021-07-30T09:20:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*SearchResourcesResponseBodyResourcesIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vtb-bp11lbh452fr8940s****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*SearchResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResources) SetAccountId(v string) *SearchResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetCreateTime(v string) *SearchResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetExpireTime(v string) *SearchResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetIpAddressAttributes(v []*SearchResourcesResponseBodyResourcesIpAddressAttributes) *SearchResourcesResponseBodyResources {
	s.IpAddressAttributes = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetIpAddresses(v []*string) *SearchResourcesResponseBodyResources {
	s.IpAddresses = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetRegionId(v string) *SearchResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceGroupId(v string) *SearchResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceId(v string) *SearchResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceName(v string) *SearchResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceType(v string) *SearchResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetTags(v []*SearchResourcesResponseBodyResourcesTags) *SearchResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetZoneId(v string) *SearchResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type SearchResourcesResponseBodyResourcesIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.1.2
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchResourcesResponseBodyResourcesIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResourcesIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetIpAddress(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetNetworkType(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetVersion(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.Version = &v
	return s
}

type SearchResourcesResponseBodyResourcesTags struct {
	// The key of tag N.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchResourcesResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResourcesTags) SetKey(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesTags) SetValue(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

type SearchResourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponse) SetHeaders(v map[string]*string) *SearchResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchResourcesResponse) SetStatusCode(v int32) *SearchResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResourcesResponse) SetBody(v *SearchResourcesResponseBody) *SearchResourcesResponse {
	s.Body = v
	return s
}

type UpdateDeliveryChannelRequest struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *UpdateDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-4m6ffpkgu***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *UpdateDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *UpdateDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s UpdateDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelFilter(v *UpdateDeliveryChannelRequestDeliveryChannelFilter) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetResourceChangeDelivery(v *UpdateDeliveryChannelRequestResourceChangeDelivery) *UpdateDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetResourceSnapshotDelivery(v *UpdateDeliveryChannelRequestResourceSnapshotDelivery) *UpdateDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

type UpdateDeliveryChannelRequestDeliveryChannelFilter struct {
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s UpdateDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *UpdateDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

type UpdateDeliveryChannelRequestResourceChangeDelivery struct {
	// Specifies whether to enable delivery of resource configuration change events. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 1911422487776***:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Set the value to `SLS`.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateDeliveryChannelRequestResourceChangeDelivery) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetEnabled(v string) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

type UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type UpdateDeliveryChannelRequestResourceSnapshotDelivery struct {
	// The custom expression.
	//
	// example:
	//
	// select 	- from resources limit 100;
	CustomExpression *string `json:"CustomExpression,omitempty" xml:"CustomExpression,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 09:00Z
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// Specifies whether to enable delivery of scheduled resource snapshots. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`. Example: `acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`. Example: `acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination. Valid values:
	//
	// 	- `OSS` for standard delivery
	//
	// 	- `OSS` or `SLS` for custom delivery
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetEnabled(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

type UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// >  This parameter takes effect only if you use custom delivery for scheduled resource snapshots. You do not need to configure this parameter if you use standard delivery for scheduled resource snapshots.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type UpdateDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AD5F848D-CCDC-5464-93E1-4BA50A482***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelResponseBody) SetRequestId(v string) *UpdateDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeliveryChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeliveryChannelResponse) SetStatusCode(v int32) *UpdateDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeliveryChannelResponse) SetBody(v *UpdateDeliveryChannelResponseBody) *UpdateDeliveryChannelResponse {
	s.Body = v
	return s
}

type UpdateFilterRequest struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s UpdateFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFilterRequest) GoString() string {
	return s.String()
}

func (s *UpdateFilterRequest) SetFilterConfiguration(v string) *UpdateFilterRequest {
	s.FilterConfiguration = &v
	return s
}

func (s *UpdateFilterRequest) SetFilterName(v string) *UpdateFilterRequest {
	s.FilterName = &v
	return s
}

type UpdateFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3C5CDBF6-4DB7-53E9-ADDC-5919E3FACF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFilterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFilterResponseBody) SetRequestId(v string) *UpdateFilterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFilterResponse) GoString() string {
	return s.String()
}

func (s *UpdateFilterResponse) SetHeaders(v map[string]*string) *UpdateFilterResponse {
	s.Headers = v
	return s
}

func (s *UpdateFilterResponse) SetStatusCode(v int32) *UpdateFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFilterResponse) SetBody(v *UpdateFilterResponseBody) *UpdateFilterResponse {
	s.Body = v
	return s
}

type UpdateMultiAccountDeliveryChannelRequest struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-4m6ffpkgu***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-multi-account-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s UpdateMultiAccountDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelFilter(v *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetResourceChangeDelivery(v *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) *UpdateMultiAccountDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetResourceSnapshotDelivery(v *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) *UpdateMultiAccountDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

type UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter struct {
	// An array of effective account scopes for the delivery channel.
	AccountScopes []*string `json:"AccountScopes,omitempty" xml:"AccountScopes,omitempty" type:"Repeated"`
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetAccountScopes(v []*string) *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.AccountScopes = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

type UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery struct {
	// Specifies whether to enable delivery of resource configuration change events. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 1911422487776***:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Set the value to `SLS`.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetEnabled(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

type UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery struct {
	// The custom expression.
	//
	// example:
	//
	// select 	- from resources limit 100;
	CustomExpression *string `json:"CustomExpression,omitempty" xml:"CustomExpression,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 09:00Z
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// Specifies whether to enable delivery of scheduled resource snapshots. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`. Example: `acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`. Example: `acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination. Valid values:
	//
	// 	- `OSS` for standard delivery
	//
	// 	- `OSS` or `SLS` for custom delivery
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetEnabled(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

type UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// >  This parameter takes effect only if you use custom delivery for scheduled resource snapshots. You do not need to configure this parameter if you use standard delivery for scheduled resource snapshots.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

type UpdateMultiAccountDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *UpdateMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *UpdateMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelResponse) SetBody(v *UpdateMultiAccountDeliveryChannelResponseBody) *UpdateMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

type UpdateSavedQueryRequest struct {
	// The description of the template.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be unique.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s UpdateSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryRequest) SetDescription(v string) *UpdateSavedQueryRequest {
	s.Description = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetExpression(v string) *UpdateSavedQueryRequest {
	s.Expression = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetName(v string) *UpdateSavedQueryRequest {
	s.Name = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetQueryId(v string) *UpdateSavedQueryRequest {
	s.QueryId = &v
	return s
}

type UpdateSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryResponseBody) SetRequestId(v string) *UpdateSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryResponse) SetHeaders(v map[string]*string) *UpdateSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *UpdateSavedQueryResponse) SetStatusCode(v int32) *UpdateSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSavedQueryResponse) SetBody(v *UpdateSavedQueryResponseBody) *UpdateSavedQueryResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("resourcecenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Sets a default filter.
//
// @param request - AssociateDefaultFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateDefaultFilterResponse
func (client *Client) AssociateDefaultFilterWithOptions(request *AssociateDefaultFilterRequest, runtime *util.RuntimeOptions) (_result *AssociateDefaultFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateDefaultFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateDefaultFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a default filter.
//
// @param request - AssociateDefaultFilterRequest
//
// @return AssociateDefaultFilterResponse
func (client *Client) AssociateDefaultFilter(request *AssociateDefaultFilterRequest) (_result *AssociateDefaultFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateDefaultFilterResponse{}
	_body, _err := client.AssociateDefaultFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a single-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the ResourceSnapshotDelivery.CustomExpression parameter empty.
//
//   - Custom delivery: Set the ResourceSnapshotDelivery.CustomExpression parameter to an appropriate value.
//
// @param request - CreateDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryChannelResponse
func (client *Client) CreateDeliveryChannelWithOptions(request *CreateDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *CreateDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelDescription)) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelFilter)) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceChangeDelivery)) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSnapshotDelivery)) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a single-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the ResourceSnapshotDelivery.CustomExpression parameter empty.
//
//   - Custom delivery: Set the ResourceSnapshotDelivery.CustomExpression parameter to an appropriate value.
//
// @param request - CreateDeliveryChannelRequest
//
// @return CreateDeliveryChannelResponse
func (client *Client) CreateDeliveryChannel(request *CreateDeliveryChannelRequest) (_result *CreateDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeliveryChannelResponse{}
	_body, _err := client.CreateDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a filter.
//
// @param request - CreateFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFilterResponse
func (client *Client) CreateFilterWithOptions(request *CreateFilterRequest, runtime *util.RuntimeOptions) (_result *CreateFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterConfiguration)) {
		query["FilterConfiguration"] = request.FilterConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a filter.
//
// @param request - CreateFilterRequest
//
// @return CreateFilterResponse
func (client *Client) CreateFilter(request *CreateFilterRequest) (_result *CreateFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFilterResponse{}
	_body, _err := client.CreateFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a multi-account delivery channel.
//
// Description:
//
// In Resource Center, you can create multi-account delivery channels by using the management account of your resource directory or the delegated administrator account of Resource Center to deliver resource configuration change events and scheduled resource snapshots within the members in your resource directory to Object Storage Service (OSS) or Simple Log Service. Then, other Alibaba Cloud services consume standardized resource information from OSS or Simple Log Service.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - CreateMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiAccountDeliveryChannelResponse
func (client *Client) CreateMultiAccountDeliveryChannelWithOptions(request *CreateMultiAccountDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *CreateMultiAccountDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelDescription)) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelFilter)) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceChangeDelivery)) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSnapshotDelivery)) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMultiAccountDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a multi-account delivery channel.
//
// Description:
//
// In Resource Center, you can create multi-account delivery channels by using the management account of your resource directory or the delegated administrator account of Resource Center to deliver resource configuration change events and scheduled resource snapshots within the members in your resource directory to Object Storage Service (OSS) or Simple Log Service. Then, other Alibaba Cloud services consume standardized resource information from OSS or Simple Log Service.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - CreateMultiAccountDeliveryChannelRequest
//
// @return CreateMultiAccountDeliveryChannelResponse
func (client *Client) CreateMultiAccountDeliveryChannel(request *CreateMultiAccountDeliveryChannelRequest) (_result *CreateMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CreateMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom query template.
//
// @param request - CreateSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavedQueryResponse
func (client *Client) CreateSavedQueryWithOptions(request *CreateSavedQueryRequest, runtime *util.RuntimeOptions) (_result *CreateSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom query template.
//
// @param request - CreateSavedQueryRequest
//
// @return CreateSavedQueryResponse
func (client *Client) CreateSavedQuery(request *CreateSavedQueryRequest) (_result *CreateSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSavedQueryResponse{}
	_body, _err := client.CreateSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single-account delivery channel.
//
// @param request - DeleteDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeliveryChannelResponse
func (client *Client) DeleteDeliveryChannelWithOptions(request *DeleteDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *DeleteDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single-account delivery channel.
//
// @param request - DeleteDeliveryChannelRequest
//
// @return DeleteDeliveryChannelResponse
func (client *Client) DeleteDeliveryChannel(request *DeleteDeliveryChannelRequest) (_result *DeleteDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeliveryChannelResponse{}
	_body, _err := client.DeleteDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a filter.
//
// @param request - DeleteFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFilterResponse
func (client *Client) DeleteFilterWithOptions(request *DeleteFilterRequest, runtime *util.RuntimeOptions) (_result *DeleteFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a filter.
//
// @param request - DeleteFilterRequest
//
// @return DeleteFilterResponse
func (client *Client) DeleteFilter(request *DeleteFilterRequest) (_result *DeleteFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFilterResponse{}
	_body, _err := client.DeleteFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a multi-account delivery channel.
//
// @param request - DeleteMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultiAccountDeliveryChannelResponse
func (client *Client) DeleteMultiAccountDeliveryChannelWithOptions(request *DeleteMultiAccountDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *DeleteMultiAccountDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMultiAccountDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a multi-account delivery channel.
//
// @param request - DeleteMultiAccountDeliveryChannelRequest
//
// @return DeleteMultiAccountDeliveryChannelResponse
func (client *Client) DeleteMultiAccountDeliveryChannel(request *DeleteMultiAccountDeliveryChannelRequest) (_result *DeleteMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMultiAccountDeliveryChannelResponse{}
	_body, _err := client.DeleteMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom query template.
//
// @param request - DeleteSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSavedQueryResponse
func (client *Client) DeleteSavedQueryWithOptions(request *DeleteSavedQueryRequest, runtime *util.RuntimeOptions) (_result *DeleteSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom query template.
//
// @param request - DeleteSavedQueryRequest
//
// @return DeleteSavedQueryResponse
func (client *Client) DeleteSavedQuery(request *DeleteSavedQueryRequest) (_result *DeleteSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSavedQueryResponse{}
	_body, _err := client.DeleteSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @param request - DisableMultiAccountResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableMultiAccountResourceCenterResponse
func (client *Client) DisableMultiAccountResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableMultiAccountResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @return DisableMultiAccountResourceCenterResponse
func (client *Client) DisableMultiAccountResourceCenter() (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.DisableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deactivates the Resource Center service.
//
// @param request - DisableResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableResourceCenterResponse
func (client *Client) DisableResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *DisableResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates the Resource Center service.
//
// @return DisableResourceCenterResponse
func (client *Client) DisableResourceCenter() (_result *DisableResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.DisableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the default filter.
//
// @param request - DisassociateDefaultFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateDefaultFilterResponse
func (client *Client) DisassociateDefaultFilterWithOptions(runtime *util.RuntimeOptions) (_result *DisassociateDefaultFilterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisassociateDefaultFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociateDefaultFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the default filter.
//
// @return DisassociateDefaultFilterResponse
func (client *Client) DisassociateDefaultFilter() (_result *DisassociateDefaultFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociateDefaultFilterResponse{}
	_body, _err := client.DisassociateDefaultFilterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// Description:
//
// If you have created a resource directory for your enterprise, you can enable the cross-account resource search feature by using the management account of the resource directory or a delegated administrator account of Resource Center to view the resources of members in the resource directory. For more information about a resource directory, see [Resource Directory overview](https://help.aliyun.com/document_detail/200506.html).
//
// @param request - EnableMultiAccountResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMultiAccountResourceCenterResponse
func (client *Client) EnableMultiAccountResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableMultiAccountResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// Description:
//
// If you have created a resource directory for your enterprise, you can enable the cross-account resource search feature by using the management account of the resource directory or a delegated administrator account of Resource Center to view the resources of members in the resource directory. For more information about a resource directory, see [Resource Directory overview](https://help.aliyun.com/document_detail/200506.html).
//
// @return EnableMultiAccountResourceCenterResponse
func (client *Client) EnableMultiAccountResourceCenter() (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.EnableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates the Resource Center service.
//
// @param request - EnableResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceCenterResponse
func (client *Client) EnableResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *EnableResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the Resource Center service.
//
// @return EnableResourceCenterResponse
func (client *Client) EnableResourceCenter() (_result *EnableResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.EnableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query resources across accounts.
//
// @param request - ExecuteMultiAccountSQLQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteMultiAccountSQLQueryResponse
func (client *Client) ExecuteMultiAccountSQLQueryWithOptions(request *ExecuteMultiAccountSQLQueryRequest, runtime *util.RuntimeOptions) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteMultiAccountSQLQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteMultiAccountSQLQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query resources across accounts.
//
// @param request - ExecuteMultiAccountSQLQueryRequest
//
// @return ExecuteMultiAccountSQLQueryResponse
func (client *Client) ExecuteMultiAccountSQLQuery(request *ExecuteMultiAccountSQLQueryRequest) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteMultiAccountSQLQueryResponse{}
	_body, _err := client.ExecuteMultiAccountSQLQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query the resources that can be accessed within the current account.
//
// @param request - ExecuteSQLQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSQLQueryResponse
func (client *Client) ExecuteSQLQueryWithOptions(request *ExecuteSQLQueryRequest, runtime *util.RuntimeOptions) (_result *ExecuteSQLQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteSQLQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteSQLQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query the resources that can be accessed within the current account.
//
// @param request - ExecuteSQLQueryRequest
//
// @return ExecuteSQLQueryResponse
func (client *Client) ExecuteSQLQuery(request *ExecuteSQLQueryRequest) (_result *ExecuteSQLQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteSQLQueryResponse{}
	_body, _err := client.ExecuteSQLQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a single-account delivery channel.
//
// @param request - GetDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryChannelResponse
func (client *Client) GetDeliveryChannelWithOptions(request *GetDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *GetDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a single-account delivery channel.
//
// @param request - GetDeliveryChannelRequest
//
// @return GetDeliveryChannelResponse
func (client *Client) GetDeliveryChannel(request *GetDeliveryChannelRequest) (_result *GetDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeliveryChannelResponse{}
	_body, _err := client.GetDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on a single-account delivery channel.
//
// @param request - GetDeliveryChannelStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryChannelStatisticsResponse
func (client *Client) GetDeliveryChannelStatisticsWithOptions(request *GetDeliveryChannelStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetDeliveryChannelStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeliveryChannelStatistics"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeliveryChannelStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on a single-account delivery channel.
//
// @param request - GetDeliveryChannelStatisticsRequest
//
// @return GetDeliveryChannelStatisticsResponse
func (client *Client) GetDeliveryChannelStatistics(request *GetDeliveryChannelStatisticsRequest) (_result *GetDeliveryChannelStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeliveryChannelStatisticsResponse{}
	_body, _err := client.GetDeliveryChannelStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a sample query template.
//
// @param request - GetExampleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExampleQueryResponse
func (client *Client) GetExampleQueryWithOptions(request *GetExampleQueryRequest, runtime *util.RuntimeOptions) (_result *GetExampleQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExampleQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExampleQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a sample query template.
//
// @param request - GetExampleQueryRequest
//
// @return GetExampleQueryResponse
func (client *Client) GetExampleQuery(request *GetExampleQueryRequest) (_result *GetExampleQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExampleQueryResponse{}
	_body, _err := client.GetExampleQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a multi-account delivery channel.
//
// @param request - GetMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountDeliveryChannelResponse
func (client *Client) GetMultiAccountDeliveryChannelWithOptions(request *GetMultiAccountDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *GetMultiAccountDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a multi-account delivery channel.
//
// @param request - GetMultiAccountDeliveryChannelRequest
//
// @return GetMultiAccountDeliveryChannelResponse
func (client *Client) GetMultiAccountDeliveryChannel(request *GetMultiAccountDeliveryChannelRequest) (_result *GetMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountDeliveryChannelResponse{}
	_body, _err := client.GetMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on a multi-account delivery channel.
//
// @param request - GetMultiAccountDeliveryChannelStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountDeliveryChannelStatisticsResponse
func (client *Client) GetMultiAccountDeliveryChannelStatisticsWithOptions(request *GetMultiAccountDeliveryChannelStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetMultiAccountDeliveryChannelStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountDeliveryChannelStatistics"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountDeliveryChannelStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on a multi-account delivery channel.
//
// @param request - GetMultiAccountDeliveryChannelStatisticsRequest
//
// @return GetMultiAccountDeliveryChannelStatisticsResponse
func (client *Client) GetMultiAccountDeliveryChannelStatistics(request *GetMultiAccountDeliveryChannelStatisticsRequest) (_result *GetMultiAccountDeliveryChannelStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountDeliveryChannelStatisticsResponse{}
	_body, _err := client.GetMultiAccountDeliveryChannelStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @param request - GetMultiAccountResourceCenterServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceCenterServiceStatusResponse
func (client *Client) GetMultiAccountResourceCenterServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountResourceCenterServiceStatus"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @return GetMultiAccountResourceCenterServiceStatusResponse
func (client *Client) GetMultiAccountResourceCenterServiceStatus() (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.GetMultiAccountResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the management account or a member of a resource directory.
//
// @param request - GetMultiAccountResourceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceConfigurationResponse
func (client *Client) GetMultiAccountResourceConfigurationWithOptions(request *GetMultiAccountResourceConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountResourceConfiguration"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the management account or a member of a resource directory.
//
// @param request - GetMultiAccountResourceConfigurationRequest
//
// @return GetMultiAccountResourceConfigurationResponse
func (client *Client) GetMultiAccountResourceConfiguration(request *GetMultiAccountResourceConfigurationRequest) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.GetMultiAccountResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取多账号资源数量
//
// @param request - GetMultiAccountResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceCountsResponse
func (client *Client) GetMultiAccountResourceCountsWithOptions(request *GetMultiAccountResourceCountsRequest, runtime *util.RuntimeOptions) (_result *GetMultiAccountResourceCountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.GroupByKey)) {
		query["GroupByKey"] = request.GroupByKey
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountResourceCounts"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountResourceCountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取多账号资源数量
//
// @param request - GetMultiAccountResourceCountsRequest
//
// @return GetMultiAccountResourceCountsResponse
func (client *Client) GetMultiAccountResourceCounts(request *GetMultiAccountResourceCountsRequest) (_result *GetMultiAccountResourceCountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountResourceCountsResponse{}
	_body, _err := client.GetMultiAccountResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Resource Center service.
//
// @param request - GetResourceCenterServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCenterServiceStatusResponse
func (client *Client) GetResourceCenterServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetResourceCenterServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetResourceCenterServiceStatus"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Resource Center service.
//
// @return GetResourceCenterServiceStatusResponse
func (client *Client) GetResourceCenterServiceStatus() (_result *GetResourceCenterServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.GetResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the current account.
//
// @param request - GetResourceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationResponse
func (client *Client) GetResourceConfigurationWithOptions(request *GetResourceConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetResourceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceConfiguration"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the current account.
//
// @param request - GetResourceConfigurationRequest
//
// @return GetResourceConfigurationResponse
func (client *Client) GetResourceConfiguration(request *GetResourceConfigurationRequest) (_result *GetResourceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.GetResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the numbers of resources on which the current account has access permissions.
//
// @param request - GetResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCountsResponse
func (client *Client) GetResourceCountsWithOptions(request *GetResourceCountsRequest, runtime *util.RuntimeOptions) (_result *GetResourceCountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.GroupByKey)) {
		query["GroupByKey"] = request.GroupByKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceCounts"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceCountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the numbers of resources on which the current account has access permissions.
//
// @param request - GetResourceCountsRequest
//
// @return GetResourceCountsResponse
func (client *Client) GetResourceCounts(request *GetResourceCountsRequest) (_result *GetResourceCountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceCountsResponse{}
	_body, _err := client.GetResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a custom query template.
//
// @param request - GetSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavedQueryResponse
func (client *Client) GetSavedQueryWithOptions(request *GetSavedQueryRequest, runtime *util.RuntimeOptions) (_result *GetSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a custom query template.
//
// @param request - GetSavedQueryRequest
//
// @return GetSavedQueryResponse
func (client *Client) GetSavedQuery(request *GetSavedQueryRequest) (_result *GetSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSavedQueryResponse{}
	_body, _err := client.GetSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of single-account delivery channels.
//
// @param request - ListDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryChannelsResponse
func (client *Client) ListDeliveryChannelsWithOptions(request *ListDeliveryChannelsRequest, runtime *util.RuntimeOptions) (_result *ListDeliveryChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeliveryChannels"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeliveryChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of single-account delivery channels.
//
// @param request - ListDeliveryChannelsRequest
//
// @return ListDeliveryChannelsResponse
func (client *Client) ListDeliveryChannels(request *ListDeliveryChannelsRequest) (_result *ListDeliveryChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeliveryChannelsResponse{}
	_body, _err := client.ListDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all sample query templates.
//
// @param request - ListExampleQueriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExampleQueriesResponse
func (client *Client) ListExampleQueriesWithOptions(request *ListExampleQueriesRequest, runtime *util.RuntimeOptions) (_result *ListExampleQueriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExampleQueries"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExampleQueriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all sample query templates.
//
// @param request - ListExampleQueriesRequest
//
// @return ListExampleQueriesResponse
func (client *Client) ListExampleQueries(request *ListExampleQueriesRequest) (_result *ListExampleQueriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExampleQueriesResponse{}
	_body, _err := client.ListExampleQueriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of filters.
//
// @param request - ListFiltersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFiltersResponse
func (client *Client) ListFiltersWithOptions(runtime *util.RuntimeOptions) (_result *ListFiltersResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListFilters"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFiltersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of filters.
//
// @return ListFiltersResponse
func (client *Client) ListFilters() (_result *ListFiltersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFiltersResponse{}
	_body, _err := client.ListFiltersWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of multi-account delivery channels.
//
// @param request - ListMultiAccountDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountDeliveryChannelsResponse
func (client *Client) ListMultiAccountDeliveryChannelsWithOptions(request *ListMultiAccountDeliveryChannelsRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountDeliveryChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountDeliveryChannels"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountDeliveryChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of multi-account delivery channels.
//
// @param request - ListMultiAccountDeliveryChannelsRequest
//
// @return ListMultiAccountDeliveryChannelsResponse
func (client *Client) ListMultiAccountDeliveryChannels(request *ListMultiAccountDeliveryChannelsRequest) (_result *ListMultiAccountDeliveryChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountDeliveryChannelsResponse{}
	_body, _err := client.ListMultiAccountDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource groups within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountResourceGroupsResponse
func (client *Client) ListMultiAccountResourceGroupsWithOptions(request *ListMultiAccountResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountResourceGroups"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource groups within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountResourceGroupsRequest
//
// @return ListMultiAccountResourceGroupsResponse
func (client *Client) ListMultiAccountResourceGroups(request *ListMultiAccountResourceGroupsRequest) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.ListMultiAccountResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the relationships between resources within the management account or members of your resource directory.
//
// Description:
//
//	  Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
//		- By default, the operation returns up to 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
// @param request - ListMultiAccountResourceRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountResourceRelationshipsResponse
func (client *Client) ListMultiAccountResourceRelationshipsWithOptions(request *ListMultiAccountResourceRelationshipsRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountResourceRelationshipsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedResourceFilter)) {
		query["RelatedResourceFilter"] = request.RelatedResourceFilter
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountResourceRelationships"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountResourceRelationshipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the relationships between resources within the management account or members of your resource directory.
//
// Description:
//
//	  Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
//		- By default, the operation returns up to 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
// @param request - ListMultiAccountResourceRelationshipsRequest
//
// @return ListMultiAccountResourceRelationshipsResponse
func (client *Client) ListMultiAccountResourceRelationships(request *ListMultiAccountResourceRelationshipsRequest) (_result *ListMultiAccountResourceRelationshipsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountResourceRelationshipsResponse{}
	_body, _err := client.ListMultiAccountResourceRelationshipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the management account or a member of your resource directory.
//
// @param request - ListMultiAccountTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagKeysResponse
func (client *Client) ListMultiAccountTagKeysWithOptions(request *ListMultiAccountTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountTagKeys"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the management account or a member of your resource directory.
//
// @param request - ListMultiAccountTagKeysRequest
//
// @return ListMultiAccountTagKeysResponse
func (client *Client) ListMultiAccountTagKeys(request *ListMultiAccountTagKeysRequest) (_result *ListMultiAccountTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.ListMultiAccountTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagValuesResponse
func (client *Client) ListMultiAccountTagValuesWithOptions(request *ListMultiAccountTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagValue)) {
		query["TagValue"] = request.TagValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountTagValues"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountTagValuesRequest
//
// @return ListMultiAccountTagValuesResponse
func (client *Client) ListMultiAccountTagValues(request *ListMultiAccountTagValuesRequest) (_result *ListMultiAccountTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.ListMultiAccountTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of resource relationships on which the current account has access permissions.
//
// Description:
//
//	  You can call this operation to query only the resource relationships on which the current account has access permissions.
//
//		- By default, this operation returns up to 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the query scope. For information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only entries that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Entries that meet any value of the filter condition are returned.
//
// @param request - ListResourceRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRelationshipsResponse
func (client *Client) ListResourceRelationshipsWithOptions(request *ListResourceRelationshipsRequest, runtime *util.RuntimeOptions) (_result *ListResourceRelationshipsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedResourceFilter)) {
		query["RelatedResourceFilter"] = request.RelatedResourceFilter
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceRelationships"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceRelationshipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of resource relationships on which the current account has access permissions.
//
// Description:
//
//	  You can call this operation to query only the resource relationships on which the current account has access permissions.
//
//		- By default, this operation returns up to 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the query scope. For information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only entries that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Entries that meet any value of the filter condition are returned.
//
// @param request - ListResourceRelationshipsRequest
//
// @return ListResourceRelationshipsResponse
func (client *Client) ListResourceRelationships(request *ListResourceRelationshipsRequest) (_result *ListResourceRelationshipsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceRelationshipsResponse{}
	_body, _err := client.ListResourceRelationshipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the metadata of resource types
//
// @param request - ListResourceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithOptions(request *ListResourceTypesRequest, runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypes"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the metadata of resource types
//
// @param request - ListResourceTypesRequest
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all custom query templates.
//
// @param request - ListSavedQueriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSavedQueriesResponse
func (client *Client) ListSavedQueriesWithOptions(request *ListSavedQueriesRequest, runtime *util.RuntimeOptions) (_result *ListSavedQueriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSavedQueries"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSavedQueriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all custom query templates.
//
// @param request - ListSavedQueriesRequest
//
// @return ListSavedQueriesResponse
func (client *Client) ListSavedQueries(request *ListSavedQueriesRequest) (_result *ListSavedQueriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSavedQueriesResponse{}
	_body, _err := client.ListSavedQueriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the current account.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the current account.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the current account.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagValue)) {
		query["TagValue"] = request.TagValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the current account.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for resources within the management account or members of a resource directory.
//
// Description:
//
//	  You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
//		- Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
//		- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
//		- You can visit [Sample Code Center](https://api.alibabacloud.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchMultiAccountResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMultiAccountResourcesResponse
func (client *Client) SearchMultiAccountResourcesWithOptions(request *SearchMultiAccountResourcesRequest, runtime *util.RuntimeOptions) (_result *SearchMultiAccountResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SortCriterion)) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchMultiAccountResources"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for resources within the management account or members of a resource directory.
//
// Description:
//
//	  You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
//		- Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
//		- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
//		- You can visit [Sample Code Center](https://api.alibabacloud.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchMultiAccountResourcesRequest
//
// @return SearchMultiAccountResourcesResponse
func (client *Client) SearchMultiAccountResources(request *SearchMultiAccountResourcesRequest) (_result *SearchMultiAccountResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.SearchMultiAccountResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Search for resources that you can access within the current account.
//
// Description:
//
//	  You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
//		- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
//		- You can visit [Sample Code Center](https://api.aliyun.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResourcesResponse
func (client *Client) SearchResourcesWithOptions(request *SearchResourcesRequest, runtime *util.RuntimeOptions) (_result *SearchResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SortCriterion)) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchResources"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search for resources that you can access within the current account.
//
// Description:
//
//	  You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
//		- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
//		- You can visit [Sample Code Center](https://api.aliyun.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchResourcesRequest
//
// @return SearchResourcesResponse
func (client *Client) SearchResources(request *SearchResourcesRequest) (_result *SearchResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchResourcesResponse{}
	_body, _err := client.SearchResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a single-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - UpdateDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeliveryChannelResponse
func (client *Client) UpdateDeliveryChannelWithOptions(request *UpdateDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *UpdateDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelDescription)) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelFilter)) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceChangeDelivery)) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSnapshotDelivery)) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a single-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - UpdateDeliveryChannelRequest
//
// @return UpdateDeliveryChannelResponse
func (client *Client) UpdateDeliveryChannel(request *UpdateDeliveryChannelRequest) (_result *UpdateDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeliveryChannelResponse{}
	_body, _err := client.UpdateDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a filter.
//
// @param request - UpdateFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFilterResponse
func (client *Client) UpdateFilterWithOptions(request *UpdateFilterRequest, runtime *util.RuntimeOptions) (_result *UpdateFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterConfiguration)) {
		query["FilterConfiguration"] = request.FilterConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a filter.
//
// @param request - UpdateFilterRequest
//
// @return UpdateFilterResponse
func (client *Client) UpdateFilter(request *UpdateFilterRequest) (_result *UpdateFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFilterResponse{}
	_body, _err := client.UpdateFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a multi-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - UpdateMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultiAccountDeliveryChannelResponse
func (client *Client) UpdateMultiAccountDeliveryChannelWithOptions(request *UpdateMultiAccountDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *UpdateMultiAccountDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelDescription)) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelFilter)) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceChangeDelivery)) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSnapshotDelivery)) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMultiAccountDeliveryChannel"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a multi-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - UpdateMultiAccountDeliveryChannelRequest
//
// @return UpdateMultiAccountDeliveryChannelResponse
func (client *Client) UpdateMultiAccountDeliveryChannel(request *UpdateMultiAccountDeliveryChannelRequest) (_result *UpdateMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.UpdateMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom query template.
//
// @param request - UpdateSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSavedQueryResponse
func (client *Client) UpdateSavedQueryWithOptions(request *UpdateSavedQueryRequest, runtime *util.RuntimeOptions) (_result *UpdateSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom query template.
//
// @param request - UpdateSavedQueryRequest
//
// @return UpdateSavedQueryResponse
func (client *Client) UpdateSavedQuery(request *UpdateSavedQueryRequest) (_result *UpdateSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSavedQueryResponse{}
	_body, _err := client.UpdateSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
