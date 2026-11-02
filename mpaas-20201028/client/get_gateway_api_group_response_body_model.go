// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetGatewayApiGroupResponseBody
	GetAccessDeniedDetail() *string
	SetApiGroup(v *GetGatewayApiGroupResponseBodyApiGroup) *GetGatewayApiGroupResponseBody
	GetApiGroup() *GetGatewayApiGroupResponseBodyApiGroup
	SetErrorField(v string) *GetGatewayApiGroupResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *GetGatewayApiGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetGatewayApiGroupResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetGatewayApiGroupResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *GetGatewayApiGroupResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *GetGatewayApiGroupResponseBody
	GetSuccess() *bool
}

type GetGatewayApiGroupResponseBody struct {
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ApiGroup           *GetGatewayApiGroupResponseBodyApiGroup `json:"ApiGroup,omitempty" xml:"ApiGroup,omitempty" type:"Struct"`
	ErrorField         *string                                 `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayApiGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetGatewayApiGroupResponseBody) GetApiGroup() *GetGatewayApiGroupResponseBodyApiGroup {
	return s.ApiGroup
}

func (s *GetGatewayApiGroupResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *GetGatewayApiGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayApiGroupResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetGatewayApiGroupResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetGatewayApiGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayApiGroupResponseBody) SetAccessDeniedDetail(v string) *GetGatewayApiGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetGatewayApiGroupResponseBody) SetApiGroup(v *GetGatewayApiGroupResponseBodyApiGroup) *GetGatewayApiGroupResponseBody {
	s.ApiGroup = v
	return s
}

func (s *GetGatewayApiGroupResponseBody) SetErrorField(v string) *GetGatewayApiGroupResponseBody {
	s.ErrorField = &v
	return s
}

func (s *GetGatewayApiGroupResponseBody) SetHttpStatusCode(v int32) *GetGatewayApiGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayApiGroupResponseBody) SetRequestId(v string) *GetGatewayApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayApiGroupResponseBody) SetResultCode(v string) *GetGatewayApiGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetGatewayApiGroupResponseBody) SetResultMessage(v string) *GetGatewayApiGroupResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetGatewayApiGroupResponseBody) SetSuccess(v bool) *GetGatewayApiGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayApiGroupResponseBody) Validate() error {
	if s.ApiGroup != nil {
		if err := s.ApiGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayApiGroupResponseBodyApiGroup struct {
	Centers      map[string]interface{} `json:"Centers,omitempty" xml:"Centers,omitempty"`
	Description  *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName   *string                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnableMTLS   *bool                  `json:"EnableMTLS,omitempty" xml:"EnableMTLS,omitempty"`
	GroupId      *int64                 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Host         *string                `json:"Host,omitempty" xml:"Host,omitempty"`
	HttpsModel   *bool                  `json:"HttpsModel,omitempty" xml:"HttpsModel,omitempty"`
	InstanceId   *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstancePort *int32                 `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	RegionMode   *bool                  `json:"RegionMode,omitempty" xml:"RegionMode,omitempty"`
	Timeout      *int32                 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	UpstreamType *string                `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	VpcId        *string                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcModel     *bool                  `json:"VpcModel,omitempty" xml:"VpcModel,omitempty"`
}

func (s GetGatewayApiGroupResponseBodyApiGroup) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiGroupResponseBodyApiGroup) GoString() string {
	return s.String()
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetCenters() map[string]interface{} {
	return s.Centers
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetDescription() *string {
	return s.Description
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetDomainName() *string {
	return s.DomainName
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetHost() *string {
	return s.Host
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetVpcId() *string {
	return s.VpcId
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetCenters(v map[string]interface{}) *GetGatewayApiGroupResponseBodyApiGroup {
	s.Centers = v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetDescription(v string) *GetGatewayApiGroupResponseBodyApiGroup {
	s.Description = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetDomainName(v string) *GetGatewayApiGroupResponseBodyApiGroup {
	s.DomainName = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetEnableMTLS(v bool) *GetGatewayApiGroupResponseBodyApiGroup {
	s.EnableMTLS = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetGroupId(v int64) *GetGatewayApiGroupResponseBodyApiGroup {
	s.GroupId = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetGroupName(v string) *GetGatewayApiGroupResponseBodyApiGroup {
	s.GroupName = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetHost(v string) *GetGatewayApiGroupResponseBodyApiGroup {
	s.Host = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetHttpsModel(v bool) *GetGatewayApiGroupResponseBodyApiGroup {
	s.HttpsModel = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetInstanceId(v string) *GetGatewayApiGroupResponseBodyApiGroup {
	s.InstanceId = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetInstancePort(v int32) *GetGatewayApiGroupResponseBodyApiGroup {
	s.InstancePort = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetRegionMode(v bool) *GetGatewayApiGroupResponseBodyApiGroup {
	s.RegionMode = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetTimeout(v int32) *GetGatewayApiGroupResponseBodyApiGroup {
	s.Timeout = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetUpstreamType(v string) *GetGatewayApiGroupResponseBodyApiGroup {
	s.UpstreamType = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetVpcId(v string) *GetGatewayApiGroupResponseBodyApiGroup {
	s.VpcId = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) SetVpcModel(v bool) *GetGatewayApiGroupResponseBodyApiGroup {
	s.VpcModel = &v
	return s
}

func (s *GetGatewayApiGroupResponseBodyApiGroup) Validate() error {
	return dara.Validate(s)
}
