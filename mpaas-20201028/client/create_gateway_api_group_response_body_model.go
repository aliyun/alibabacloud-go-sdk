// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateGatewayApiGroupResponseBody
	GetAccessDeniedDetail() *string
	SetApiGroup(v *CreateGatewayApiGroupResponseBodyApiGroup) *CreateGatewayApiGroupResponseBody
	GetApiGroup() *CreateGatewayApiGroupResponseBodyApiGroup
	SetErrorField(v string) *CreateGatewayApiGroupResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *CreateGatewayApiGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateGatewayApiGroupResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateGatewayApiGroupResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateGatewayApiGroupResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *CreateGatewayApiGroupResponseBody
	GetSuccess() *bool
}

type CreateGatewayApiGroupResponseBody struct {
	AccessDeniedDetail *string                                    `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ApiGroup           *CreateGatewayApiGroupResponseBodyApiGroup `json:"ApiGroup,omitempty" xml:"ApiGroup,omitempty" type:"Struct"`
	ErrorField         *string                                    `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateGatewayApiGroupResponseBody) GetApiGroup() *CreateGatewayApiGroupResponseBodyApiGroup {
	return s.ApiGroup
}

func (s *CreateGatewayApiGroupResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *CreateGatewayApiGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateGatewayApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayApiGroupResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateGatewayApiGroupResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateGatewayApiGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGatewayApiGroupResponseBody) SetAccessDeniedDetail(v string) *CreateGatewayApiGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) SetApiGroup(v *CreateGatewayApiGroupResponseBodyApiGroup) *CreateGatewayApiGroupResponseBody {
	s.ApiGroup = v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) SetErrorField(v string) *CreateGatewayApiGroupResponseBody {
	s.ErrorField = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) SetHttpStatusCode(v int32) *CreateGatewayApiGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) SetRequestId(v string) *CreateGatewayApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) SetResultCode(v string) *CreateGatewayApiGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) SetResultMessage(v string) *CreateGatewayApiGroupResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) SetSuccess(v bool) *CreateGatewayApiGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBody) Validate() error {
	if s.ApiGroup != nil {
		if err := s.ApiGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGatewayApiGroupResponseBodyApiGroup struct {
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

func (s CreateGatewayApiGroupResponseBodyApiGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiGroupResponseBodyApiGroup) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetCenters() map[string]interface{} {
	return s.Centers
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetHost() *string {
	return s.Host
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetCenters(v map[string]interface{}) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.Centers = v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetDescription(v string) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetDomainName(v string) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.DomainName = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetEnableMTLS(v bool) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.EnableMTLS = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetGroupId(v int64) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.GroupId = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetGroupName(v string) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.GroupName = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetHost(v string) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.Host = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetHttpsModel(v bool) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.HttpsModel = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetInstanceId(v string) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.InstanceId = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetInstancePort(v int32) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.InstancePort = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetRegionMode(v bool) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.RegionMode = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetTimeout(v int32) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.Timeout = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetUpstreamType(v string) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.UpstreamType = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetVpcId(v string) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) SetVpcModel(v bool) *CreateGatewayApiGroupResponseBodyApiGroup {
	s.VpcModel = &v
	return s
}

func (s *CreateGatewayApiGroupResponseBodyApiGroup) Validate() error {
	return dara.Validate(s)
}
