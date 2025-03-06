// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AttachEaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-wz93g6pyat2g7t7o****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// i-wz93g6pyat2g7t7o****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// C3BCB7DA-BEB6-4982-A765-6EA61EC84474
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type AttachEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-bp14ws9hbt1oe0u9****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// eais.ei-a6.2xlarge
	EiInstanceType *string `json:"EiInstanceType,omitempty" xml:"EiInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachEaisEiRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachEaisEiRequest) GoString() string {
	return s.String()
}

func (s *AttachEaisEiRequest) SetClientInstanceId(v string) *AttachEaisEiRequest {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaisEiRequest) SetEiInstanceId(v string) *AttachEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *AttachEaisEiRequest) SetEiInstanceType(v string) *AttachEaisEiRequest {
	s.EiInstanceType = &v
	return s
}

func (s *AttachEaisEiRequest) SetRegionId(v string) *AttachEaisEiRequest {
	s.RegionId = &v
	return s
}

type AttachEaisEiResponseBody struct {
	// example:
	//
	// i-bp14ws9hbt1oe0u9****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// C3BCB7DA-BEB6-4982-A765-6EA61EC8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEaisEiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEaisEiResponseBody) SetClientInstanceId(v string) *AttachEaisEiResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaisEiResponseBody) SetEiInstanceId(v string) *AttachEaisEiResponseBody {
	s.EiInstanceId = &v
	return s
}

func (s *AttachEaisEiResponseBody) SetRequestId(v string) *AttachEaisEiResponseBody {
	s.RequestId = &v
	return s
}

type AttachEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEaisEiResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachEaisEiResponse) GoString() string {
	return s.String()
}

func (s *AttachEaisEiResponse) SetHeaders(v map[string]*string) *AttachEaisEiResponse {
	s.Headers = v
	return s
}

func (s *AttachEaisEiResponse) SetStatusCode(v int32) *AttachEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEaisEiResponse) SetBody(v *AttachEaisEiResponseBody) *AttachEaisEiResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais-hzs4h26yyt5xkcke****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceRegionId(v string) *ChangeResourceGroupRequest {
	s.ResourceRegionId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// example:
	//
	// C3BCB7DA-BEB6-4982-A765-6EA61EC8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateEaiRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Image       *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// eais-test01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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

func (s *CreateEaiRequest) SetImage(v string) *CreateEaiRequest {
	s.Image = &v
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

func (s *CreateEaiRequest) SetResourceGroupId(v string) *CreateEaiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiRequest) SetSecurityGroupId(v string) *CreateEaiRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiRequest) SetTag(v []*CreateEaiRequestTag) *CreateEaiRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiRequest) SetVSwitchId(v string) *CreateEaiRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiRequestTag) SetKey(v string) *CreateEaiRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiRequestTag) SetValue(v string) *CreateEaiRequestTag {
	s.Value = &v
	return s
}

type CreateEaiResponseBody struct {
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// A655AB0E-31BB-45AD-9255-FCE93F6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateEaiEciRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// eais-test01
	EaisName *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType *string                 `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	Eci      *CreateEaiEciRequestEci `json:"Eci,omitempty" xml:"Eci,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiEciRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEciRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequest) SetClientToken(v string) *CreateEaiEciRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEciRequest) SetEaisName(v string) *CreateEaiEciRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEciRequest) SetEaisType(v string) *CreateEaiEciRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEciRequest) SetEci(v *CreateEaiEciRequestEci) *CreateEaiEciRequest {
	s.Eci = v
	return s
}

func (s *CreateEaiEciRequest) SetRegionId(v string) *CreateEaiEciRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEciRequest) SetResourceGroupId(v string) *CreateEaiEciRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEciRequest) SetSecurityGroupId(v string) *CreateEaiEciRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEciRequest) SetTag(v []*CreateEaiEciRequestTag) *CreateEaiEciRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEciRequest) SetVSwitchId(v string) *CreateEaiEciRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiEciRequestEci struct {
	Container *CreateEaiEciRequestEciContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// example:
	//
	// eip-uf66jeqopgqa9hdn****
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// example:
	//
	// test-nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ecs.c5.xlarge
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 00c7****-rivj.cn-hangzhou.extreme.nas.aliyuncs.com:/share
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateEaiEciRequestEci) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciRequestEci) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequestEci) SetContainer(v *CreateEaiEciRequestEciContainer) *CreateEaiEciRequestEci {
	s.Container = v
	return s
}

func (s *CreateEaiEciRequestEci) SetEipId(v string) *CreateEaiEciRequestEci {
	s.EipId = &v
	return s
}

func (s *CreateEaiEciRequestEci) SetName(v string) *CreateEaiEciRequestEci {
	s.Name = &v
	return s
}

func (s *CreateEaiEciRequestEci) SetType(v string) *CreateEaiEciRequestEci {
	s.Type = &v
	return s
}

func (s *CreateEaiEciRequestEci) SetVolume(v string) *CreateEaiEciRequestEci {
	s.Volume = &v
	return s
}

type CreateEaiEciRequestEciContainer struct {
	// example:
	//
	// 100
	Arg *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	// example:
	//
	// sleep
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /mnt/eais=eais,/models=eais/models
	Volumes *string `json:"Volumes,omitempty" xml:"Volumes,omitempty"`
}

func (s CreateEaiEciRequestEciContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciRequestEciContainer) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequestEciContainer) SetArg(v string) *CreateEaiEciRequestEciContainer {
	s.Arg = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetCommand(v string) *CreateEaiEciRequestEciContainer {
	s.Command = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetImage(v string) *CreateEaiEciRequestEciContainer {
	s.Image = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetName(v string) *CreateEaiEciRequestEciContainer {
	s.Name = &v
	return s
}

func (s *CreateEaiEciRequestEciContainer) SetVolumes(v string) *CreateEaiEciRequestEciContainer {
	s.Volumes = &v
	return s
}

type CreateEaiEciRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEciRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEciRequestTag) SetKey(v string) *CreateEaiEciRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEciRequestTag) SetValue(v string) *CreateEaiEciRequestTag {
	s.Value = &v
	return s
}

type CreateEaiEciShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// eais-test01
	EaisName *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType  *string `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EciShrink *string `json:"Eci,omitempty" xml:"Eci,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiEciShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEciShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEciShrinkRequest) SetClientToken(v string) *CreateEaiEciShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetEaisName(v string) *CreateEaiEciShrinkRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetEaisType(v string) *CreateEaiEciShrinkRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetEciShrink(v string) *CreateEaiEciShrinkRequest {
	s.EciShrink = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetRegionId(v string) *CreateEaiEciShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetResourceGroupId(v string) *CreateEaiEciShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetSecurityGroupId(v string) *CreateEaiEciShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetTag(v []*CreateEaiEciShrinkRequestTag) *CreateEaiEciShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetVSwitchId(v string) *CreateEaiEciShrinkRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiEciShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEciShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEciShrinkRequestTag) SetKey(v string) *CreateEaiEciShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEciShrinkRequestTag) SetValue(v string) *CreateEaiEciShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateEaiEciResponseBody struct {
	// example:
	//
	// eci-2zeh03ygxlrzmfi6****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiEciResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiEciResponseBody) SetClientInstanceId(v string) *CreateEaiEciResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *CreateEaiEciResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiEciResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiEciResponseBody) SetRequestId(v string) *CreateEaiEciResponseBody {
	s.RequestId = &v
	return s
}

type CreateEaiEciResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiEciResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaiEciResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEciResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiEciResponse) SetHeaders(v map[string]*string) *CreateEaiEciResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiEciResponse) SetStatusCode(v int32) *CreateEaiEciResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiEciResponse) SetBody(v *CreateEaiEciResponseBody) *CreateEaiEciResponse {
	s.Body = v
	return s
}

type CreateEaiEcsRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// eais-test01
	EaisName *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType *string                 `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	Ecs      *CreateEaiEcsRequestEcs `json:"Ecs,omitempty" xml:"Ecs,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiEcsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEcsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEcsRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsRequest) SetClientToken(v string) *CreateEaiEcsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEcsRequest) SetEaisName(v string) *CreateEaiEcsRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEcsRequest) SetEaisType(v string) *CreateEaiEcsRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEcsRequest) SetEcs(v *CreateEaiEcsRequestEcs) *CreateEaiEcsRequest {
	s.Ecs = v
	return s
}

func (s *CreateEaiEcsRequest) SetRegionId(v string) *CreateEaiEcsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEcsRequest) SetResourceGroupId(v string) *CreateEaiEcsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEcsRequest) SetSecurityGroupId(v string) *CreateEaiEcsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEcsRequest) SetTag(v []*CreateEaiEcsRequestTag) *CreateEaiEcsRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEcsRequest) SetVSwitchId(v string) *CreateEaiEcsRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiEcsRequestEcs struct {
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200324.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 10
	InternetMaxBandwidthIn *string `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// example:
	//
	// 10
	InternetMaxBandwidthOut *string `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// example:
	//
	// 40
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// example:
	//
	// ecs.g7.4xlarge
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateEaiEcsRequestEcs) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEcsRequestEcs) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsRequestEcs) SetImageId(v string) *CreateEaiEcsRequestEcs {
	s.ImageId = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetInternetMaxBandwidthIn(v string) *CreateEaiEcsRequestEcs {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetInternetMaxBandwidthOut(v string) *CreateEaiEcsRequestEcs {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetName(v string) *CreateEaiEcsRequestEcs {
	s.Name = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetPassword(v string) *CreateEaiEcsRequestEcs {
	s.Password = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetSystemDiskCategory(v string) *CreateEaiEcsRequestEcs {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetSystemDiskSize(v int64) *CreateEaiEcsRequestEcs {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetType(v string) *CreateEaiEcsRequestEcs {
	s.Type = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetZoneId(v string) *CreateEaiEcsRequestEcs {
	s.ZoneId = &v
	return s
}

type CreateEaiEcsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEcsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEcsRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsRequestTag) SetKey(v string) *CreateEaiEcsRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEcsRequestTag) SetValue(v string) *CreateEaiEcsRequestTag {
	s.Value = &v
	return s
}

type CreateEaiEcsShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// eais-test01
	EaisName *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType  *string `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EcsShrink *string `json:"Ecs,omitempty" xml:"Ecs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiEcsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEcsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEcsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsShrinkRequest) SetClientToken(v string) *CreateEaiEcsShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetEaisName(v string) *CreateEaiEcsShrinkRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetEaisType(v string) *CreateEaiEcsShrinkRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetEcsShrink(v string) *CreateEaiEcsShrinkRequest {
	s.EcsShrink = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetRegionId(v string) *CreateEaiEcsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetResourceGroupId(v string) *CreateEaiEcsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetSecurityGroupId(v string) *CreateEaiEcsShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetTag(v []*CreateEaiEcsShrinkRequestTag) *CreateEaiEcsShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetVSwitchId(v string) *CreateEaiEcsShrinkRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiEcsShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEcsShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEcsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsShrinkRequestTag) SetKey(v string) *CreateEaiEcsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEcsShrinkRequestTag) SetValue(v string) *CreateEaiEcsShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateEaiEcsResponseBody struct {
	// example:
	//
	// i-bp1hjrvleawl4ogb****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiEcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEcsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsResponseBody) SetClientInstanceId(v string) *CreateEaiEcsResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *CreateEaiEcsResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiEcsResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiEcsResponseBody) SetRequestId(v string) *CreateEaiEcsResponseBody {
	s.RequestId = &v
	return s
}

type CreateEaiEcsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiEcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaiEcsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiEcsResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsResponse) SetHeaders(v map[string]*string) *CreateEaiEcsResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiEcsResponse) SetStatusCode(v int32) *CreateEaiEcsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiEcsResponse) SetBody(v *CreateEaiEcsResponseBody) *CreateEaiEcsResponse {
	s.Body = v
	return s
}

type CreateEaiJupyterRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EaisName    *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType       *string                                  `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EnvironmentVar []*CreateEaiJupyterRequestEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                       `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiJupyterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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

func (s *CreateEaiJupyterRequest) SetEaisName(v string) *CreateEaiJupyterRequest {
	s.EaisName = &v
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

func (s *CreateEaiJupyterRequest) SetResourceGroupId(v string) *CreateEaiJupyterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetSecurityGroupId(v string) *CreateEaiJupyterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetTag(v []*CreateEaiJupyterRequestTag) *CreateEaiJupyterRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiJupyterRequest) SetVSwitchId(v string) *CreateEaiJupyterRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiJupyterRequestEnvironmentVar struct {
	// example:
	//
	// MY_USER_NAME
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test123
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

type CreateEaiJupyterRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiJupyterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiJupyterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterRequestTag) SetKey(v string) *CreateEaiJupyterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiJupyterRequestTag) SetValue(v string) *CreateEaiJupyterRequestTag {
	s.Value = &v
	return s
}

type CreateEaiJupyterShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EaisName    *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType             *string `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EnvironmentVarShrink *string `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                             `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiJupyterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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

func (s *CreateEaiJupyterShrinkRequest) SetEaisName(v string) *CreateEaiJupyterShrinkRequest {
	s.EaisName = &v
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

func (s *CreateEaiJupyterShrinkRequest) SetResourceGroupId(v string) *CreateEaiJupyterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetSecurityGroupId(v string) *CreateEaiJupyterShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetTag(v []*CreateEaiJupyterShrinkRequestTag) *CreateEaiJupyterShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetVSwitchId(v string) *CreateEaiJupyterShrinkRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaiJupyterShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiJupyterShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaiJupyterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterShrinkRequestTag) SetKey(v string) *CreateEaiJupyterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequestTag) SetValue(v string) *CreateEaiJupyterShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateEaiJupyterResponseBody struct {
	// example:
	//
	// eais-hz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// A655AB0E-31BB-45AD-9255-FCE93F6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiJupyterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateEaisEiRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// test_ei
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-bp16jgp51ttnkbdr****
	SecurityGroupId *string                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaisEiRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp17wmd1wb6fwlimk****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaisEiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEaisEiRequest) GoString() string {
	return s.String()
}

func (s *CreateEaisEiRequest) SetClientToken(v string) *CreateEaisEiRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaisEiRequest) SetInstanceName(v string) *CreateEaisEiRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateEaisEiRequest) SetInstanceType(v string) *CreateEaisEiRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateEaisEiRequest) SetRegionId(v string) *CreateEaisEiRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaisEiRequest) SetResourceGroupId(v string) *CreateEaisEiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaisEiRequest) SetSecurityGroupId(v string) *CreateEaisEiRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaisEiRequest) SetTag(v []*CreateEaisEiRequestTag) *CreateEaisEiRequest {
	s.Tag = v
	return s
}

func (s *CreateEaisEiRequest) SetVSwitchId(v string) *CreateEaisEiRequest {
	s.VSwitchId = &v
	return s
}

type CreateEaisEiRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaisEiRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateEaisEiRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaisEiRequestTag) SetKey(v string) *CreateEaisEiRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaisEiRequestTag) SetValue(v string) *CreateEaisEiRequestTag {
	s.Value = &v
	return s
}

type CreateEaisEiResponseBody struct {
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaisEiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaisEiResponseBody) SetEiInstanceId(v string) *CreateEaisEiResponseBody {
	s.EiInstanceId = &v
	return s
}

func (s *CreateEaisEiResponseBody) SetRequestId(v string) *CreateEaisEiResponseBody {
	s.RequestId = &v
	return s
}

type CreateEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaisEiResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEaisEiResponse) GoString() string {
	return s.String()
}

func (s *CreateEaisEiResponse) SetHeaders(v map[string]*string) *CreateEaisEiResponse {
	s.Headers = v
	return s
}

func (s *CreateEaisEiResponse) SetStatusCode(v int32) *CreateEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaisEiResponse) SetBody(v *CreateEaisEiResponseBody) *CreateEaisEiResponse {
	s.Body = v
	return s
}

type DeleteEaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// F23AEEC7-4D98-4657-A104-0269270*****
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// i-bp1fvhi60e1zizsp****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais-hza1ahi0uuw0re33****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// AD4EA714-A35B-4710-BF92-8275BCDDB69F
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEaiAllResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEaisEiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaisEiRequest) GoString() string {
	return s.String()
}

func (s *DeleteEaisEiRequest) SetEiInstanceId(v string) *DeleteEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *DeleteEaisEiRequest) SetForce(v bool) *DeleteEaisEiRequest {
	s.Force = &v
	return s
}

func (s *DeleteEaisEiRequest) SetRegionId(v string) *DeleteEaisEiRequest {
	s.RegionId = &v
	return s
}

type DeleteEaisEiResponseBody struct {
	// example:
	//
	// F23AEEC7-4D98-4657-A104-02692701****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEaisEiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEaisEiResponseBody) SetRequestId(v string) *DeleteEaisEiResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEaisEiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEaisEiResponse) GoString() string {
	return s.String()
}

func (s *DeleteEaisEiResponse) SetHeaders(v map[string]*string) *DeleteEaisEiResponse {
	s.Headers = v
	return s
}

func (s *DeleteEaisEiResponse) SetStatusCode(v int32) *DeleteEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEaisEiResponse) SetBody(v *DeleteEaisEiResponseBody) *DeleteEaisEiResponse {
	s.Body = v
	return s
}

type DescribeEaisRequest struct {
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// ["eais-id1", "eais-id2"]
	ElasticAcceleratedInstanceIds *string `json:"ElasticAcceleratedInstanceIds,omitempty" xml:"ElasticAcceleratedInstanceIds,omitempty"`
	// example:
	//
	// eais*
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// eais.ei-a6.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// InUse
	Status *string                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*DescribeEaisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEaisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisRequest) GoString() string {
	return s.String()
}

func (s *DescribeEaisRequest) SetClientInstanceId(v string) *DescribeEaisRequest {
	s.ClientInstanceId = &v
	return s
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

func (s *DescribeEaisRequest) SetResourceGroupId(v string) *DescribeEaisRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEaisRequest) SetStatus(v string) *DescribeEaisRequest {
	s.Status = &v
	return s
}

func (s *DescribeEaisRequest) SetTag(v []*DescribeEaisRequestTag) *DescribeEaisRequest {
	s.Tag = v
	return s
}

type DescribeEaisRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEaisRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeEaisRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeEaisRequestTag) SetKey(v string) *DescribeEaisRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeEaisRequestTag) SetValue(v string) *DescribeEaisRequestTag {
	s.Value = &v
	return s
}

type DescribeEaisResponseBody struct {
	Instances *DescribeEaisResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1E23D585-BBD8-436F-9615-54CACD6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// jupyter
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// i-wz93g6pyat2g****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// test1
	ClientInstanceName *string `json:"ClientInstanceName,omitempty" xml:"ClientInstanceName,omitempty"`
	// example:
	//
	// ecs.g5ne.large
	ClientInstanceType *string `json:"ClientInstanceType,omitempty" xml:"ClientInstanceType,omitempty"`
	// example:
	//
	// 2020-11-11T03:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// eais-sz8t15a7gt7****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// testName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// eais.ei-a6.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// http://121.41.**.24:8888
	JupyterUrl  *string `json:"JupyterUrl,omitempty" xml:"JupyterUrl,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// sg-bp1gppir818lx4******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// InUse
	Status *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeEaisResponseBodyInstancesInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// vsw-bp1sd131hfmd76r******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeEaisResponseBodyInstancesInstance) SetPaymentType(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.PaymentType = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetRegionId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
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
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// TestValue
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEaisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF1840
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// eais.cn-shenzhen.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type DetachEaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 04DEB304-2436-4CB9-BB63-468BCEA03D9D
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DetachEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachEaisEiRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachEaisEiRequest) GoString() string {
	return s.String()
}

func (s *DetachEaisEiRequest) SetEiInstanceId(v string) *DetachEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *DetachEaisEiRequest) SetRegionId(v string) *DetachEaisEiRequest {
	s.RegionId = &v
	return s
}

type DetachEaisEiResponseBody struct {
	// example:
	//
	// 04DEB304-2436-4CB9-BB63-468BCEA0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEaisEiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEaisEiResponseBody) SetRequestId(v string) *DetachEaisEiResponseBody {
	s.RequestId = &v
	return s
}

type DetachEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachEaisEiResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachEaisEiResponse) GoString() string {
	return s.String()
}

func (s *DetachEaisEiResponse) SetHeaders(v map[string]*string) *DetachEaisEiResponse {
	s.Headers = v
	return s
}

func (s *DetachEaisEiResponse) SetStatusCode(v int32) *DetachEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEaisEiResponse) SetBody(v *DetachEaisEiResponseBody) *DetachEaisEiResponse {
	s.Body = v
	return s
}

type GetInstanceMetricsRequest struct {
	// example:
	//
	// 2022-11-22T16:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais-hznzre6ffmz9num4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MemoryUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2022-11-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 5m
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s GetInstanceMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsRequest) SetEndTime(v string) *GetInstanceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetInstanceId(v string) *GetInstanceMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetMetricType(v string) *GetInstanceMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetRegionId(v string) *GetInstanceMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetStartTime(v string) *GetInstanceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetTimeStep(v string) *GetInstanceMetricsRequest {
	s.TimeStep = &v
	return s
}

type GetInstanceMetricsResponseBody struct {
	// example:
	//
	// eais-bj8b53it29hfhj******
	InstanceId *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PodMetrics []*GetInstanceMetricsResponseBodyPodMetrics `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBody) SetInstanceId(v string) *GetInstanceMetricsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetPodMetrics(v []*GetInstanceMetricsResponseBodyPodMetrics) *GetInstanceMetricsResponseBody {
	s.PodMetrics = v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetRequestId(v string) *GetInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceMetricsResponseBodyPodMetrics struct {
	Metrics []*GetInstanceMetricsResponseBodyPodMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// Pod ID。
	//
	// example:
	//
	// eais-hznzre6ffmz9num4****-579b587ddf-9txr6
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetMetrics(v []*GetInstanceMetricsResponseBodyPodMetricsMetrics) *GetInstanceMetricsResponseBodyPodMetrics {
	s.Metrics = v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetPodId(v string) *GetInstanceMetricsResponseBodyPodMetrics {
	s.PodId = &v
	return s
}

type GetInstanceMetricsResponseBodyPodMetricsMetrics struct {
	// example:
	//
	// 1669107528450
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 4.536552540058814
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetTime(v string) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Time = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetValue(v string) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Value = &v
	return s
}

type GetInstanceMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponse) SetHeaders(v map[string]*string) *GetInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceMetricsResponse) SetStatusCode(v int32) *GetInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceMetricsResponse) SetBody(v *GetInstanceMetricsResponseBody) *GetInstanceMetricsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4884
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4885
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// F23AEEC7-4D98-4657-A104-02692701****
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// example:
	//
	// eais-hzs4h26yyt5xkcke****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type StartEaiJupyterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hze3x2gv9wimdj0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartEaiJupyterRequest) String() string {
	return tea.Prettify(s)
}

func (s StartEaiJupyterRequest) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterRequest) SetInstanceId(v string) *StartEaiJupyterRequest {
	s.InstanceId = &v
	return s
}

func (s *StartEaiJupyterRequest) SetRegionId(v string) *StartEaiJupyterRequest {
	s.RegionId = &v
	return s
}

type StartEaiJupyterResponseBody struct {
	AccessDeniedDetail *StartEaiJupyterResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 04DEB304-2436-4CB9-BB63-468BCEA0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartEaiJupyterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartEaiJupyterResponseBody) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterResponseBody) SetAccessDeniedDetail(v *StartEaiJupyterResponseBodyAccessDeniedDetail) *StartEaiJupyterResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *StartEaiJupyterResponseBody) SetRequestId(v string) *StartEaiJupyterResponseBody {
	s.RequestId = &v
	return s
}

type StartEaiJupyterResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// eais:StartEaiJupyter
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 20560152949032****
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 170718266783****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQFmfh3BZn4dwUQyNzY4MDVELTgzQkUtNTBEOC04QjQyLTNGM0U1QUI5MjhBRA==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ExplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s StartEaiJupyterResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s StartEaiJupyterResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthAction(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetPolicyType(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type StartEaiJupyterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEaiJupyterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEaiJupyterResponse) String() string {
	return tea.Prettify(s)
}

func (s StartEaiJupyterResponse) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterResponse) SetHeaders(v map[string]*string) *StartEaiJupyterResponse {
	s.Headers = v
	return s
}

func (s *StartEaiJupyterResponse) SetStatusCode(v int32) *StartEaiJupyterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEaiJupyterResponse) SetBody(v *StartEaiJupyterResponseBody) *StartEaiJupyterResponse {
	s.Body = v
	return s
}

type StartEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartEaisEiRequest) String() string {
	return tea.Prettify(s)
}

func (s StartEaisEiRequest) GoString() string {
	return s.String()
}

func (s *StartEaisEiRequest) SetEiInstanceId(v string) *StartEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *StartEaisEiRequest) SetRegionId(v string) *StartEaisEiRequest {
	s.RegionId = &v
	return s
}

type StartEaisEiResponseBody struct {
	// example:
	//
	// F23AEEC7-4D98-4657-A104-02692701****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartEaisEiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *StartEaisEiResponseBody) SetRequestId(v string) *StartEaisEiResponseBody {
	s.RequestId = &v
	return s
}

type StartEaisEiResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEaisEiResponse) String() string {
	return tea.Prettify(s)
}

func (s StartEaisEiResponse) GoString() string {
	return s.String()
}

func (s *StartEaisEiResponse) SetHeaders(v map[string]*string) *StartEaisEiResponse {
	s.Headers = v
	return s
}

func (s *StartEaisEiResponse) SetStatusCode(v int32) *StartEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEaisEiResponse) SetBody(v *StartEaisEiResponseBody) *StartEaisEiResponse {
	s.Body = v
	return s
}

type StopEaiJupyterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hze3x2gv9wimdj0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopEaiJupyterRequest) String() string {
	return tea.Prettify(s)
}

func (s StopEaiJupyterRequest) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterRequest) SetInstanceId(v string) *StopEaiJupyterRequest {
	s.InstanceId = &v
	return s
}

func (s *StopEaiJupyterRequest) SetRegionId(v string) *StopEaiJupyterRequest {
	s.RegionId = &v
	return s
}

type StopEaiJupyterResponseBody struct {
	AccessDeniedDetail *StopEaiJupyterResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopEaiJupyterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopEaiJupyterResponseBody) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterResponseBody) SetAccessDeniedDetail(v *StopEaiJupyterResponseBodyAccessDeniedDetail) *StopEaiJupyterResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *StopEaiJupyterResponseBody) SetRequestId(v string) *StopEaiJupyterResponseBody {
	s.RequestId = &v
	return s
}

type StopEaiJupyterResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// eais:StopEaiJupyter
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 20560152949032****
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 170718266783****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQFmj0FOZo9BTjMyQTFDRkIzLUE5MTItNUIwNC1BQzkxLTcyMUFFQTUyQjhGQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ExplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s StopEaiJupyterResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s StopEaiJupyterResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthAction(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetPolicyType(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type StopEaiJupyterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopEaiJupyterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopEaiJupyterResponse) String() string {
	return tea.Prettify(s)
}

func (s StopEaiJupyterResponse) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterResponse) SetHeaders(v map[string]*string) *StopEaiJupyterResponse {
	s.Headers = v
	return s
}

func (s *StopEaiJupyterResponse) SetStatusCode(v int32) *StopEaiJupyterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopEaiJupyterResponse) SetBody(v *StopEaiJupyterResponseBody) *StopEaiJupyterResponse {
	s.Body = v
	return s
}

type StopEaisEiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopEaisEiRequest) String() string {
	return tea.Prettify(s)
}

func (s StopEaisEiRequest) GoString() string {
	return s.String()
}

func (s *StopEaisEiRequest) SetEiInstanceId(v string) *StopEaisEiRequest {
	s.EiInstanceId = &v
	return s
}

func (s *StopEaisEiRequest) SetRegionId(v string) *StopEaisEiRequest {
	s.RegionId = &v
	return s
}

type StopEaisEiResponseBody struct {
	// example:
	//
	// F23AEEC7-4D98-4657-A104-02692701****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopEaisEiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *StopEaisEiResponseBody) SetRequestId(v string) *StopEaisEiResponseBody {
	s.RequestId = &v
	return s
}

type StopEaisEiResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopEaisEiResponse) String() string {
	return tea.Prettify(s)
}

func (s StopEaisEiResponse) GoString() string {
	return s.String()
}

func (s *StopEaisEiResponse) SetHeaders(v map[string]*string) *StopEaisEiResponse {
	s.Headers = v
	return s
}

func (s *StopEaisEiResponse) SetStatusCode(v int32) *StopEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *StopEaisEiResponse) SetBody(v *StopEaisEiResponseBody) *StopEaisEiResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// example:
	//
	// 1E23D585-BBD8-436F-9615-54CACD67****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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

// Summary:
//
// 将弹性加速计算实例挂载到ECS实例上
//
// @param request - AttachEaiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachEaiResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AttachEaiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AttachEaiResponse{}
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
// 将弹性加速计算实例挂载到ECS实例上
//
// @param request - AttachEaiRequest
//
// @return AttachEaiResponse
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

// Summary:
//
// 将EI绑定到ECS或ECI实例上
//
// @param request - AttachEaisEiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachEaisEiResponse
func (client *Client) AttachEaisEiWithOptions(request *AttachEaisEiRequest, runtime *util.RuntimeOptions) (_result *AttachEaisEiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInstanceId)) {
		query["ClientInstanceId"] = request.ClientInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EiInstanceId)) {
		query["EiInstanceId"] = request.EiInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EiInstanceType)) {
		query["EiInstanceType"] = request.EiInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachEaisEi"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AttachEaisEiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AttachEaisEiResponse{}
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
// 将EI绑定到ECS或ECI实例上
//
// @param request - AttachEaisEiRequest
//
// @return AttachEaisEiResponse
func (client *Client) AttachEaisEi(request *AttachEaisEiRequest) (_result *AttachEaisEiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachEaisEiResponse{}
	_body, _err := client.AttachEaisEiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资源转组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChangeResourceGroupResponse{}
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
// 资源转组
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个弹性加速计算实例
//
// @param request - CreateEaiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEaiResponse
func (client *Client) CreateEaiWithOptions(request *CreateEaiRequest, runtime *util.RuntimeOptions) (_result *CreateEaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEaiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEaiResponse{}
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
// 创建一个弹性加速计算实例
//
// @param request - CreateEaiRequest
//
// @return CreateEaiResponse
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

// Summary:
//
// 创建一个EAIS实例和ECI实例并绑定
//
// @param tmpReq - CreateEaiEciRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEaiEciResponse
func (client *Client) CreateEaiEciWithOptions(tmpReq *CreateEaiEciRequest, runtime *util.RuntimeOptions) (_result *CreateEaiEciResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEaiEciShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Eci)) {
		request.EciShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Eci, tea.String("Eci"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EaisName)) {
		query["EaisName"] = request.EaisName
	}

	if !tea.BoolValue(util.IsUnset(request.EaisType)) {
		query["EaisType"] = request.EaisType
	}

	if !tea.BoolValue(util.IsUnset(request.EciShrink)) {
		query["Eci"] = request.EciShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEaiEci"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEaiEciResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEaiEciResponse{}
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
// 创建一个EAIS实例和ECI实例并绑定
//
// @param request - CreateEaiEciRequest
//
// @return CreateEaiEciResponse
func (client *Client) CreateEaiEci(request *CreateEaiEciRequest) (_result *CreateEaiEciResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEaiEciResponse{}
	_body, _err := client.CreateEaiEciWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个EAIS实例和ECS实例并绑定
//
// @param tmpReq - CreateEaiEcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEaiEcsResponse
func (client *Client) CreateEaiEcsWithOptions(tmpReq *CreateEaiEcsRequest, runtime *util.RuntimeOptions) (_result *CreateEaiEcsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEaiEcsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Ecs)) {
		request.EcsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ecs, tea.String("Ecs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EaisName)) {
		query["EaisName"] = request.EaisName
	}

	if !tea.BoolValue(util.IsUnset(request.EaisType)) {
		query["EaisType"] = request.EaisType
	}

	if !tea.BoolValue(util.IsUnset(request.EcsShrink)) {
		query["Ecs"] = request.EcsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEaiEcs"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEaiEcsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEaiEcsResponse{}
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
// 创建一个EAIS实例和ECS实例并绑定
//
// @param request - CreateEaiEcsRequest
//
// @return CreateEaiEcsResponse
func (client *Client) CreateEaiEcs(request *CreateEaiEcsRequest) (_result *CreateEaiEcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEaiEcsResponse{}
	_body, _err := client.CreateEaiEcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个EAIS Jupyter环境
//
// @param tmpReq - CreateEaiJupyterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEaiJupyterResponse
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

	if !tea.BoolValue(util.IsUnset(request.EaisName)) {
		query["EaisName"] = request.EaisName
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEaiJupyterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEaiJupyterResponse{}
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
// 创建一个EAIS Jupyter环境
//
// @param request - CreateEaiJupyterRequest
//
// @return CreateEaiJupyterResponse
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

// Summary:
//
// 创建一个弹性加速计算实例
//
// @param request - CreateEaisEiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEaisEiResponse
func (client *Client) CreateEaisEiWithOptions(request *CreateEaisEiRequest, runtime *util.RuntimeOptions) (_result *CreateEaisEiResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEaisEi"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEaisEiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEaisEiResponse{}
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
// 创建一个弹性加速计算实例
//
// @param request - CreateEaisEiRequest
//
// @return CreateEaisEiResponse
func (client *Client) CreateEaisEi(request *CreateEaisEiRequest) (_result *CreateEaisEiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEaisEiResponse{}
	_body, _err := client.CreateEaisEiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 释放一个弹性加速计算实例
//
// @param request - DeleteEaiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEaiResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteEaiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteEaiResponse{}
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
// 释放一个弹性加速计算实例
//
// @param request - DeleteEaiRequest
//
// @return DeleteEaiResponse
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

// Summary:
//
// 释放一个弹性加速计算实例以及与其绑定的ECS或ECI实例
//
// @param request - DeleteEaiAllRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEaiAllResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteEaiAllResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteEaiAllResponse{}
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
// 释放一个弹性加速计算实例以及与其绑定的ECS或ECI实例
//
// @param request - DeleteEaiAllRequest
//
// @return DeleteEaiAllResponse
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

// Summary:
//
// 释放弹性加速计算实例
//
// @param request - DeleteEaisEiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEaisEiResponse
func (client *Client) DeleteEaisEiWithOptions(request *DeleteEaisEiRequest, runtime *util.RuntimeOptions) (_result *DeleteEaisEiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EiInstanceId)) {
		query["EiInstanceId"] = request.EiInstanceId
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
		Action:      tea.String("DeleteEaisEi"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteEaisEiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteEaisEiResponse{}
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
// 释放弹性加速计算实例
//
// @param request - DeleteEaisEiRequest
//
// @return DeleteEaisEiResponse
func (client *Client) DeleteEaisEi(request *DeleteEaisEiRequest) (_result *DeleteEaisEiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEaisEiResponse{}
	_body, _err := client.DeleteEaisEiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一个或多个弹性加速计算实例的详细信息
//
// @param request - DescribeEaisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEaisResponse
func (client *Client) DescribeEaisWithOptions(request *DescribeEaisRequest, runtime *util.RuntimeOptions) (_result *DescribeEaisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInstanceId)) {
		query["ClientInstanceId"] = request.ClientInstanceId
	}

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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeEaisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeEaisResponse{}
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
// 查询一个或多个弹性加速计算实例的详细信息
//
// @param request - DescribeEaisRequest
//
// @return DescribeEaisResponse
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

// Summary:
//
// 查询您可以使用的阿里云地域
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
		Version:     tea.String("2019-06-24"),
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
// 查询您可以使用的阿里云地域
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
// 从ECS实例上卸载弹性加速计算实例
//
// @param request - DetachEaiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachEaiResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DetachEaiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DetachEaiResponse{}
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
// 从ECS实例上卸载弹性加速计算实例
//
// @param request - DetachEaiRequest
//
// @return DetachEaiResponse
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

// Summary:
//
// 将EI实例与ECS或ECI实例解绑
//
// @param request - DetachEaisEiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachEaisEiResponse
func (client *Client) DetachEaisEiWithOptions(request *DetachEaisEiRequest, runtime *util.RuntimeOptions) (_result *DetachEaisEiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EiInstanceId)) {
		query["EiInstanceId"] = request.EiInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachEaisEi"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DetachEaisEiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DetachEaisEiResponse{}
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
// 将EI实例与ECS或ECI实例解绑
//
// @param request - DetachEaisEiRequest
//
// @return DetachEaisEiResponse
func (client *Client) DetachEaisEi(request *DetachEaisEiRequest) (_result *DetachEaisEiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachEaisEiResponse{}
	_body, _err := client.DetachEaisEiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取EAIS实例级别的监控数据
//
// @param request - GetInstanceMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceMetricsResponse
func (client *Client) GetInstanceMetricsWithOptions(request *GetInstanceMetricsRequest, runtime *util.RuntimeOptions) (_result *GetInstanceMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceMetrics"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceMetricsResponse{}
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
// 获取EAIS实例级别的监控数据
//
// @param request - GetInstanceMetricsRequest
//
// @return GetInstanceMetricsResponse
func (client *Client) GetInstanceMetrics(request *GetInstanceMetricsRequest) (_result *GetInstanceMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceMetricsResponse{}
	_body, _err := client.GetInstanceMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询标签列表
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
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
// 查询标签列表
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动一个部署了notebook的弹性加速计算实例
//
// @param request - StartEaiJupyterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartEaiJupyterResponse
func (client *Client) StartEaiJupyterWithOptions(request *StartEaiJupyterRequest, runtime *util.RuntimeOptions) (_result *StartEaiJupyterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartEaiJupyter"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartEaiJupyterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartEaiJupyterResponse{}
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
// 启动一个部署了notebook的弹性加速计算实例
//
// @param request - StartEaiJupyterRequest
//
// @return StartEaiJupyterResponse
func (client *Client) StartEaiJupyter(request *StartEaiJupyterRequest) (_result *StartEaiJupyterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartEaiJupyterResponse{}
	_body, _err := client.StartEaiJupyterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动一个弹性加速计算实例
//
// @param request - StartEaisEiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartEaisEiResponse
func (client *Client) StartEaisEiWithOptions(request *StartEaisEiRequest, runtime *util.RuntimeOptions) (_result *StartEaisEiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EiInstanceId)) {
		query["EiInstanceId"] = request.EiInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartEaisEi"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartEaisEiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartEaisEiResponse{}
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
// 启动一个弹性加速计算实例
//
// @param request - StartEaisEiRequest
//
// @return StartEaisEiResponse
func (client *Client) StartEaisEi(request *StartEaisEiRequest) (_result *StartEaisEiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartEaisEiResponse{}
	_body, _err := client.StartEaisEiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止一个部署了notebook的弹性加速计算实例
//
// @param request - StopEaiJupyterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopEaiJupyterResponse
func (client *Client) StopEaiJupyterWithOptions(request *StopEaiJupyterRequest, runtime *util.RuntimeOptions) (_result *StopEaiJupyterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopEaiJupyter"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopEaiJupyterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopEaiJupyterResponse{}
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
// 停止一个部署了notebook的弹性加速计算实例
//
// @param request - StopEaiJupyterRequest
//
// @return StopEaiJupyterResponse
func (client *Client) StopEaiJupyter(request *StopEaiJupyterRequest) (_result *StopEaiJupyterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopEaiJupyterResponse{}
	_body, _err := client.StopEaiJupyterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止一个弹性加速计算实例
//
// @param request - StopEaisEiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopEaisEiResponse
func (client *Client) StopEaisEiWithOptions(request *StopEaisEiRequest, runtime *util.RuntimeOptions) (_result *StopEaisEiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EiInstanceId)) {
		query["EiInstanceId"] = request.EiInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopEaisEi"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopEaisEiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopEaisEiResponse{}
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
// 停止一个弹性加速计算实例
//
// @param request - StopEaisEiRequest
//
// @return StopEaisEiResponse
func (client *Client) StopEaisEi(request *StopEaisEiRequest) (_result *StopEaisEiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopEaisEiResponse{}
	_body, _err := client.StopEaisEiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为弹性加速计算实例创建并绑定标签
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
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
// 为弹性加速计算实例创建并绑定标签
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑并删除标签
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-06-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
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
// 解绑并删除标签
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
