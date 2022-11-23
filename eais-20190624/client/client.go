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

type AttachEaiRequest struct {
	ClientInstanceId             *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachEaiRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachEaiRequest) GoString() string {
	return s.String()
}

func (s *AttachEaiRequest) SetClientInstanceId(v string) *AttachEaiRequest {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaiRequest) SetElasticAcceleratedInstanceId(v string) *AttachEaiRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *AttachEaiRequest) SetRegionId(v string) *AttachEaiRequest {
	s.RegionId = &v
	return s
}

type AttachEaiResponseBody struct {
	ClientInstanceId             *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	RequestId                    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachEaiResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEaiResponseBody) SetClientInstanceId(v string) *AttachEaiResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaiResponseBody) SetElasticAcceleratedInstanceId(v string) *AttachEaiResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *AttachEaiResponseBody) SetRequestId(v string) *AttachEaiResponseBody {
	s.RequestId = &v
	return s
}

type AttachEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachEaiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachEaiResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachEaiResponse) GoString() string {
	return s.String()
}

func (s *AttachEaiResponse) SetHeaders(v map[string]*string) *AttachEaiResponse {
	s.Headers = v
	return s
}

func (s *AttachEaiResponse) SetStatusCode(v int32) *AttachEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEaiResponse) SetBody(v *AttachEaiResponseBody) *AttachEaiResponse {
	s.Body = v
	return s
}

type CreateEaiRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiRequest) SetClientToken(v string) *CreateEaiRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiRequest) SetInstanceName(v string) *CreateEaiRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateEaiRequest) SetInstanceType(v string) *CreateEaiRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateEaiRequest) SetRegionId(v string) *CreateEaiRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiRequest) SetSecurityGroupId(v string) *CreateEaiRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiRequest) SetVSwitchId(v string) *CreateEaiRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiResponseBody struct {
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	RequestId                    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiResponseBody) SetRequestId(v string) *CreateEaiResponseBody {
	s.RequestId = &v
	return s
}

type CreateEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEaiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEaiResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiResponse) SetHeaders(v map[string]*string) *CreateEaiResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiResponse) SetStatusCode(v int32) *CreateEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiResponse) SetBody(v *CreateEaiResponseBody) *CreateEaiResponse {
	s.Body = v
	return s
}

type CreateEaiAllRequest struct {
	ClientImageId                 *string `json:"ClientImageId,omitempty" xml:"ClientImageId,omitempty"`
	ClientInstanceName            *string `json:"ClientInstanceName,omitempty" xml:"ClientInstanceName,omitempty"`
	ClientInstanceType            *string `json:"ClientInstanceType,omitempty" xml:"ClientInstanceType,omitempty"`
	ClientInternetMaxBandwidthIn  *int32  `json:"ClientInternetMaxBandwidthIn,omitempty" xml:"ClientInternetMaxBandwidthIn,omitempty"`
	ClientInternetMaxBandwidthOut *int32  `json:"ClientInternetMaxBandwidthOut,omitempty" xml:"ClientInternetMaxBandwidthOut,omitempty"`
	ClientPassword                *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	ClientSecurityGroupId         *string `json:"ClientSecurityGroupId,omitempty" xml:"ClientSecurityGroupId,omitempty"`
	ClientSystemDiskCategory      *string `json:"ClientSystemDiskCategory,omitempty" xml:"ClientSystemDiskCategory,omitempty"`
	ClientSystemDiskSize          *int32  `json:"ClientSystemDiskSize,omitempty" xml:"ClientSystemDiskSize,omitempty"`
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClientVSwitchId               *string `json:"ClientVSwitchId,omitempty" xml:"ClientVSwitchId,omitempty"`
	ClientZoneId                  *string `json:"ClientZoneId,omitempty" xml:"ClientZoneId,omitempty"`
	EaiInstanceType               *string `json:"EaiInstanceType,omitempty" xml:"EaiInstanceType,omitempty"`
	InstanceName                  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId                      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEaiAllRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiAllRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiAllRequest) SetClientImageId(v string) *CreateEaiAllRequest {
	s.ClientImageId = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientInstanceName(v string) *CreateEaiAllRequest {
	s.ClientInstanceName = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientInstanceType(v string) *CreateEaiAllRequest {
	s.ClientInstanceType = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientInternetMaxBandwidthIn(v int32) *CreateEaiAllRequest {
	s.ClientInternetMaxBandwidthIn = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientInternetMaxBandwidthOut(v int32) *CreateEaiAllRequest {
	s.ClientInternetMaxBandwidthOut = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientPassword(v string) *CreateEaiAllRequest {
	s.ClientPassword = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientSecurityGroupId(v string) *CreateEaiAllRequest {
	s.ClientSecurityGroupId = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientSystemDiskCategory(v string) *CreateEaiAllRequest {
	s.ClientSystemDiskCategory = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientSystemDiskSize(v int32) *CreateEaiAllRequest {
	s.ClientSystemDiskSize = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientToken(v string) *CreateEaiAllRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientVSwitchId(v string) *CreateEaiAllRequest {
	s.ClientVSwitchId = &v
	return s
}

func (s *CreateEaiAllRequest) SetClientZoneId(v string) *CreateEaiAllRequest {
	s.ClientZoneId = &v
	return s
}

func (s *CreateEaiAllRequest) SetEaiInstanceType(v string) *CreateEaiAllRequest {
	s.EaiInstanceType = &v
	return s
}

func (s *CreateEaiAllRequest) SetInstanceName(v string) *CreateEaiAllRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateEaiAllRequest) SetRegionId(v string) *CreateEaiAllRequest {
	s.RegionId = &v
	return s
}

type CreateEaiAllResponseBody struct {
	ClientInstanceId             *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	RequestId                    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiAllResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiAllResponseBody) SetClientInstanceId(v string) *CreateEaiAllResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *CreateEaiAllResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiAllResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiAllResponseBody) SetRequestId(v string) *CreateEaiAllResponseBody {
	s.RequestId = &v
	return s
}

type CreateEaiAllResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEaiAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEaiAllResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiAllResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiAllResponse) SetHeaders(v map[string]*string) *CreateEaiAllResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiAllResponse) SetStatusCode(v int32) *CreateEaiAllResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiAllResponse) SetBody(v *CreateEaiAllResponseBody) *CreateEaiAllResponse {
	s.Body = v
	return s
}

type CreateEaiJupyterRequest struct {
	ClientToken     *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EaisType        *string                                  `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EnvironmentVar  []*CreateEaiJupyterRequestEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	RegionId        *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiJupyterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiJupyterRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterRequest) SetClientToken(v string) *CreateEaiJupyterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetEaisType(v string) *CreateEaiJupyterRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetEnvironmentVar(v []*CreateEaiJupyterRequestEnvironmentVar) *CreateEaiJupyterRequest {
	s.EnvironmentVar = v
	return s
}

func (s *CreateEaiJupyterRequest) SetRegionId(v string) *CreateEaiJupyterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetSecurityGroupId(v string) *CreateEaiJupyterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetVSwitchId(v string) *CreateEaiJupyterRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiJupyterRequestEnvironmentVar struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiJupyterRequestEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiJupyterRequestEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterRequestEnvironmentVar) SetKey(v string) *CreateEaiJupyterRequestEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateEaiJupyterRequestEnvironmentVar) SetValue(v string) *CreateEaiJupyterRequestEnvironmentVar {
	s.Value = &v
	return s
}

type CreateEaiJupyterShrinkRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EaisType             *string `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EnvironmentVarShrink *string `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiJupyterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiJupyterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterShrinkRequest) SetClientToken(v string) *CreateEaiJupyterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetEaisType(v string) *CreateEaiJupyterShrinkRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetEnvironmentVarShrink(v string) *CreateEaiJupyterShrinkRequest {
	s.EnvironmentVarShrink = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetRegionId(v string) *CreateEaiJupyterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetSecurityGroupId(v string) *CreateEaiJupyterShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetVSwitchId(v string) *CreateEaiJupyterShrinkRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiJupyterResponseBody struct {
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	RequestId                    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiJupyterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiJupyterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiJupyterResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiJupyterResponseBody) SetRequestId(v string) *CreateEaiJupyterResponseBody {
	s.RequestId = &v
	return s
}

type CreateEaiJupyterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEaiJupyterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEaiJupyterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiJupyterResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterResponse) SetHeaders(v map[string]*string) *CreateEaiJupyterResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiJupyterResponse) SetStatusCode(v int32) *CreateEaiJupyterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiJupyterResponse) SetBody(v *CreateEaiJupyterResponseBody) *CreateEaiJupyterResponse {
	s.Body = v
	return s
}

type DeleteEaiRequest struct {
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	Force                        *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEaiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaiRequest) GoString() string {
	return s.String()
}

func (s *DeleteEaiRequest) SetElasticAcceleratedInstanceId(v string) *DeleteEaiRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DeleteEaiRequest) SetForce(v bool) *DeleteEaiRequest {
	s.Force = &v
	return s
}

func (s *DeleteEaiRequest) SetRegionId(v string) *DeleteEaiRequest {
	s.RegionId = &v
	return s
}

type DeleteEaiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEaiResponseBody) SetRequestId(v string) *DeleteEaiResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEaiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEaiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaiResponse) GoString() string {
	return s.String()
}

func (s *DeleteEaiResponse) SetHeaders(v map[string]*string) *DeleteEaiResponse {
	s.Headers = v
	return s
}

func (s *DeleteEaiResponse) SetStatusCode(v int32) *DeleteEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEaiResponse) SetBody(v *DeleteEaiResponseBody) *DeleteEaiResponse {
	s.Body = v
	return s
}

type DeleteEaiAllRequest struct {
	ClientInstanceId             *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEaiAllRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaiAllRequest) GoString() string {
	return s.String()
}

func (s *DeleteEaiAllRequest) SetClientInstanceId(v string) *DeleteEaiAllRequest {
	s.ClientInstanceId = &v
	return s
}

func (s *DeleteEaiAllRequest) SetElasticAcceleratedInstanceId(v string) *DeleteEaiAllRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DeleteEaiAllRequest) SetRegionId(v string) *DeleteEaiAllRequest {
	s.RegionId = &v
	return s
}

type DeleteEaiAllResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEaiAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaiAllResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEaiAllResponseBody) SetRequestId(v string) *DeleteEaiAllResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEaiAllResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEaiAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEaiAllResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaiAllResponse) GoString() string {
	return s.String()
}

func (s *DeleteEaiAllResponse) SetHeaders(v map[string]*string) *DeleteEaiAllResponse {
	s.Headers = v
	return s
}

func (s *DeleteEaiAllResponse) SetStatusCode(v int32) *DeleteEaiAllResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEaiAllResponse) SetBody(v *DeleteEaiAllResponseBody) *DeleteEaiAllResponse {
	s.Body = v
	return s
}

type DescribeEaisRequest struct {
	ElasticAcceleratedInstanceIds *string `json:"ElasticAcceleratedInstanceIds,omitempty" xml:"ElasticAcceleratedInstanceIds,omitempty"`
	InstanceName                  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType                  *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PageNumber                    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId                      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEaisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisRequest) GoString() string {
	return s.String()
}

func (s *DescribeEaisRequest) SetElasticAcceleratedInstanceIds(v string) *DescribeEaisRequest {
	s.ElasticAcceleratedInstanceIds = &v
	return s
}

func (s *DescribeEaisRequest) SetInstanceName(v string) *DescribeEaisRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeEaisRequest) SetInstanceType(v string) *DescribeEaisRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeEaisRequest) SetPageNumber(v int32) *DescribeEaisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEaisRequest) SetPageSize(v int32) *DescribeEaisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEaisRequest) SetRegionId(v string) *DescribeEaisRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEaisRequest) SetStatus(v string) *DescribeEaisRequest {
	s.Status = &v
	return s
}

type DescribeEaisResponseBody struct {
	Instances  *DescribeEaisResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEaisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBody) SetInstances(v *DescribeEaisResponseBodyInstances) *DescribeEaisResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeEaisResponseBody) SetPageNumber(v int32) *DescribeEaisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEaisResponseBody) SetPageSize(v int32) *DescribeEaisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEaisResponseBody) SetRequestId(v string) *DescribeEaisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEaisResponseBody) SetTotalCount(v int32) *DescribeEaisResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEaisResponseBodyInstances struct {
	Instance []*DescribeEaisResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeEaisResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstances) SetInstance(v []*DescribeEaisResponseBodyInstancesInstance) *DescribeEaisResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeEaisResponseBodyInstancesInstance struct {
	Category                     *string                                        `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientInstanceId             *string                                        `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	ClientInstanceName           *string                                        `json:"ClientInstanceName,omitempty" xml:"ClientInstanceName,omitempty"`
	ClientInstanceType           *string                                        `json:"ClientInstanceType,omitempty" xml:"ClientInstanceType,omitempty"`
	CreationTime                 *string                                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                  *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ElasticAcceleratedInstanceId *string                                        `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	InstanceName                 *string                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType                 *string                                        `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	JupyterUrl                   *string                                        `json:"JupyterUrl,omitempty" xml:"JupyterUrl,omitempty"`
	RegionId                     *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId              *string                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	StartTime                    *string                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                       *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                         *DescribeEaisResponseBodyInstancesInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId                    *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                       *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeEaisResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetCategory(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.Category = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetClientInstanceId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ClientInstanceId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetClientInstanceName(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ClientInstanceName = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetClientInstanceType(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ClientInstanceType = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetCreationTime(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetDescription(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetElasticAcceleratedInstanceId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetInstanceType(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetJupyterUrl(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.JupyterUrl = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetRegionId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetSecurityGroupId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetStartTime(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetStatus(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetTags(v *DescribeEaisResponseBodyInstancesInstanceTags) *DescribeEaisResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetVSwitchId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetZoneId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

type DescribeEaisResponseBodyInstancesInstanceTags struct {
	Tag []*DescribeEaisResponseBodyInstancesInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEaisResponseBodyInstancesInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstancesInstanceTags) SetTag(v []*DescribeEaisResponseBodyInstancesInstanceTagsTag) *DescribeEaisResponseBodyInstancesInstanceTags {
	s.Tag = v
	return s
}

type DescribeEaisResponseBodyInstancesInstanceTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeEaisResponseBodyInstancesInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisResponseBodyInstancesInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstancesInstanceTagsTag) SetTagKey(v string) *DescribeEaisResponseBodyInstancesInstanceTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstanceTagsTag) SetTagValue(v string) *DescribeEaisResponseBodyInstancesInstanceTagsTag {
	s.TagValue = &v
	return s
}

type DescribeEaisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEaisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEaisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisResponse) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponse) SetHeaders(v map[string]*string) *DescribeEaisResponse {
	s.Headers = v
	return s
}

func (s *DescribeEaisResponse) SetStatusCode(v int32) *DescribeEaisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEaisResponse) SetBody(v *DescribeEaisResponseBody) *DescribeEaisResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DetachEaiRequest struct {
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachEaiRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachEaiRequest) GoString() string {
	return s.String()
}

func (s *DetachEaiRequest) SetElasticAcceleratedInstanceId(v string) *DetachEaiRequest {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DetachEaiRequest) SetRegionId(v string) *DetachEaiRequest {
	s.RegionId = &v
	return s
}

type DetachEaiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachEaiResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEaiResponseBody) SetRequestId(v string) *DetachEaiResponseBody {
	s.RequestId = &v
	return s
}

type DetachEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachEaiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachEaiResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachEaiResponse) GoString() string {
	return s.String()
}

func (s *DetachEaiResponse) SetHeaders(v map[string]*string) *DetachEaiResponse {
	s.Headers = v
	return s
}

func (s *DetachEaiResponse) SetStatusCode(v int32) *DetachEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEaiResponse) SetBody(v *DetachEaiResponseBody) *DetachEaiResponse {
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
		"ap-northeast-1":              tea.String("eais.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("eais.aliyuncs.com"),
		"ap-south-1":                  tea.String("eais.aliyuncs.com"),
		"ap-southeast-1":              tea.String("eais.aliyuncs.com"),
		"ap-southeast-2":              tea.String("eais.aliyuncs.com"),
		"ap-southeast-3":              tea.String("eais.aliyuncs.com"),
		"ap-southeast-5":              tea.String("eais.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("eais.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("eais.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("eais.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("eais.aliyuncs.com"),
		"cn-edge-1":                   tea.String("eais.aliyuncs.com"),
		"cn-fujian":                   tea.String("eais.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("eais.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("eais.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("eais.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("eais.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("eais.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("eais.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("eais.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("eais.aliyuncs.com"),
		"cn-hongkong":                 tea.String("eais.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("eais.aliyuncs.com"),
		"cn-huhehaote":                tea.String("eais.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("eais.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("eais.aliyuncs.com"),
		"cn-qingdao":                  tea.String("eais.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("eais.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("eais.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("eais.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("eais.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("eais.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("eais.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("eais.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("eais.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("eais.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("eais.aliyuncs.com"),
		"cn-wuhan":                    tea.String("eais.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("eais.aliyuncs.com"),
		"cn-yushanfang":               tea.String("eais.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("eais.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("eais.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("eais.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("eais.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("eais.aliyuncs.com"),
		"eu-central-1":                tea.String("eais.aliyuncs.com"),
		"eu-west-1":                   tea.String("eais.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("eais.aliyuncs.com"),
		"me-east-1":                   tea.String("eais.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("eais.aliyuncs.com"),
		"us-east-1":                   tea.String("eais.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eais"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AttachEaiWithOptions(request *AttachEaiRequest, runtime *util.RuntimeOptions) (_result *AttachEaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInstanceId)) {
		query["ClientInstanceId"] = request.ClientInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticAcceleratedInstanceId)) {
		query["ElasticAcceleratedInstanceId"] = request.ElasticAcceleratedInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachEai"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachEaiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachEai(request *AttachEaiRequest) (_result *AttachEaiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachEaiResponse{}
	_body, _err := client.AttachEaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEaiWithOptions(request *CreateEaiRequest, runtime *util.RuntimeOptions) (_result *CreateEaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEai"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEaiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEai(request *CreateEaiRequest) (_result *CreateEaiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEaiResponse{}
	_body, _err := client.CreateEaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEaiAllWithOptions(request *CreateEaiAllRequest, runtime *util.RuntimeOptions) (_result *CreateEaiAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientImageId)) {
		query["ClientImageId"] = request.ClientImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInstanceName)) {
		query["ClientInstanceName"] = request.ClientInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInstanceType)) {
		query["ClientInstanceType"] = request.ClientInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInternetMaxBandwidthIn)) {
		query["ClientInternetMaxBandwidthIn"] = request.ClientInternetMaxBandwidthIn
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInternetMaxBandwidthOut)) {
		query["ClientInternetMaxBandwidthOut"] = request.ClientInternetMaxBandwidthOut
	}

	if !tea.BoolValue(util.IsUnset(request.ClientPassword)) {
		query["ClientPassword"] = request.ClientPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecurityGroupId)) {
		query["ClientSecurityGroupId"] = request.ClientSecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSystemDiskCategory)) {
		query["ClientSystemDiskCategory"] = request.ClientSystemDiskCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSystemDiskSize)) {
		query["ClientSystemDiskSize"] = request.ClientSystemDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVSwitchId)) {
		query["ClientVSwitchId"] = request.ClientVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientZoneId)) {
		query["ClientZoneId"] = request.ClientZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.EaiInstanceType)) {
		query["EaiInstanceType"] = request.EaiInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEaiAll"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEaiAllResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEaiAll(request *CreateEaiAllRequest) (_result *CreateEaiAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEaiAllResponse{}
	_body, _err := client.CreateEaiAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEaiJupyterWithOptions(tmpReq *CreateEaiJupyterRequest, runtime *util.RuntimeOptions) (_result *CreateEaiJupyterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEaiJupyterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EnvironmentVar)) {
		request.EnvironmentVarShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvironmentVar, tea.String("EnvironmentVar"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EaisType)) {
		query["EaisType"] = request.EaisType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentVarShrink)) {
		query["EnvironmentVar"] = request.EnvironmentVarShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEaiJupyter"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEaiJupyterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEaiJupyter(request *CreateEaiJupyterRequest) (_result *CreateEaiJupyterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEaiJupyterResponse{}
	_body, _err := client.CreateEaiJupyterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEaiWithOptions(request *DeleteEaiRequest, runtime *util.RuntimeOptions) (_result *DeleteEaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticAcceleratedInstanceId)) {
		query["ElasticAcceleratedInstanceId"] = request.ElasticAcceleratedInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEai"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEaiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEai(request *DeleteEaiRequest) (_result *DeleteEaiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEaiResponse{}
	_body, _err := client.DeleteEaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEaiAllWithOptions(request *DeleteEaiAllRequest, runtime *util.RuntimeOptions) (_result *DeleteEaiAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInstanceId)) {
		query["ClientInstanceId"] = request.ClientInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticAcceleratedInstanceId)) {
		query["ElasticAcceleratedInstanceId"] = request.ElasticAcceleratedInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEaiAll"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEaiAllResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEaiAll(request *DeleteEaiAllRequest) (_result *DeleteEaiAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEaiAllResponse{}
	_body, _err := client.DeleteEaiAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEaisWithOptions(request *DescribeEaisRequest, runtime *util.RuntimeOptions) (_result *DescribeEaisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticAcceleratedInstanceIds)) {
		query["ElasticAcceleratedInstanceIds"] = request.ElasticAcceleratedInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEais"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEaisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEais(request *DescribeEaisRequest) (_result *DescribeEaisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEaisResponse{}
	_body, _err := client.DescribeEaisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DetachEaiWithOptions(request *DetachEaiRequest, runtime *util.RuntimeOptions) (_result *DetachEaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticAcceleratedInstanceId)) {
		query["ElasticAcceleratedInstanceId"] = request.ElasticAcceleratedInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachEai"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachEaiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachEai(request *DetachEaiRequest) (_result *DetachEaiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachEaiResponse{}
	_body, _err := client.DetachEaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
