// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateGatewayApiGroupResponseBody
	GetAccessDeniedDetail() *string
	SetApiGroup(v *UpdateGatewayApiGroupResponseBodyApiGroup) *UpdateGatewayApiGroupResponseBody
	GetApiGroup() *UpdateGatewayApiGroupResponseBodyApiGroup
	SetErrorField(v string) *UpdateGatewayApiGroupResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *UpdateGatewayApiGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateGatewayApiGroupResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UpdateGatewayApiGroupResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *UpdateGatewayApiGroupResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *UpdateGatewayApiGroupResponseBody
	GetSuccess() *bool
}

type UpdateGatewayApiGroupResponseBody struct {
	AccessDeniedDetail *string                                    `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ApiGroup           *UpdateGatewayApiGroupResponseBodyApiGroup `json:"ApiGroup,omitempty" xml:"ApiGroup,omitempty" type:"Struct"`
	ErrorField         *string                                    `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateGatewayApiGroupResponseBody) GetApiGroup() *UpdateGatewayApiGroupResponseBodyApiGroup {
	return s.ApiGroup
}

func (s *UpdateGatewayApiGroupResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *UpdateGatewayApiGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayApiGroupResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UpdateGatewayApiGroupResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UpdateGatewayApiGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayApiGroupResponseBody) SetAccessDeniedDetail(v string) *UpdateGatewayApiGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) SetApiGroup(v *UpdateGatewayApiGroupResponseBodyApiGroup) *UpdateGatewayApiGroupResponseBody {
	s.ApiGroup = v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) SetErrorField(v string) *UpdateGatewayApiGroupResponseBody {
	s.ErrorField = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayApiGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) SetRequestId(v string) *UpdateGatewayApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) SetResultCode(v string) *UpdateGatewayApiGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) SetResultMessage(v string) *UpdateGatewayApiGroupResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) SetSuccess(v bool) *UpdateGatewayApiGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBody) Validate() error {
	if s.ApiGroup != nil {
		if err := s.ApiGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayApiGroupResponseBodyApiGroup struct {
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

func (s UpdateGatewayApiGroupResponseBodyApiGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiGroupResponseBodyApiGroup) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetCenters() map[string]interface{} {
	return s.Centers
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetHost() *string {
	return s.Host
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetCenters(v map[string]interface{}) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.Centers = v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetDescription(v string) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.Description = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetDomainName(v string) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.DomainName = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetEnableMTLS(v bool) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.EnableMTLS = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetGroupId(v int64) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.GroupId = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetGroupName(v string) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetHost(v string) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.Host = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetHttpsModel(v bool) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.HttpsModel = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetInstanceId(v string) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.InstanceId = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetInstancePort(v int32) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.InstancePort = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetRegionMode(v bool) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.RegionMode = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetTimeout(v int32) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetUpstreamType(v string) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.UpstreamType = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetVpcId(v string) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.VpcId = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) SetVpcModel(v bool) *UpdateGatewayApiGroupResponseBodyApiGroup {
	s.VpcModel = &v
	return s
}

func (s *UpdateGatewayApiGroupResponseBodyApiGroup) Validate() error {
	return dara.Validate(s)
}
