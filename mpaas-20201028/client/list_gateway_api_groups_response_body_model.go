// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayApiGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListGatewayApiGroupsResponseBody
	GetAccessDeniedDetail() *string
	SetApiGroups(v []*ListGatewayApiGroupsResponseBodyApiGroups) *ListGatewayApiGroupsResponseBody
	GetApiGroups() []*ListGatewayApiGroupsResponseBodyApiGroups
	SetErrorField(v string) *ListGatewayApiGroupsResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *ListGatewayApiGroupsResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *ListGatewayApiGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayApiGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListGatewayApiGroupsResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListGatewayApiGroupsResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListGatewayApiGroupsResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *ListGatewayApiGroupsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListGatewayApiGroupsResponseBody
	GetTotal() *int64
}

type ListGatewayApiGroupsResponseBody struct {
	AccessDeniedDetail *string                                      `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ApiGroups          []*ListGatewayApiGroupsResponseBodyApiGroups `json:"ApiGroups,omitempty" xml:"ApiGroups,omitempty" type:"Repeated"`
	ErrorField         *string                                      `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	PageNumber         *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Total              *int64                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGatewayApiGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayApiGroupsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListGatewayApiGroupsResponseBody) GetApiGroups() []*ListGatewayApiGroupsResponseBodyApiGroups {
	return s.ApiGroups
}

func (s *ListGatewayApiGroupsResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *ListGatewayApiGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayApiGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayApiGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayApiGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayApiGroupsResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListGatewayApiGroupsResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListGatewayApiGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayApiGroupsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListGatewayApiGroupsResponseBody) SetAccessDeniedDetail(v string) *ListGatewayApiGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetApiGroups(v []*ListGatewayApiGroupsResponseBodyApiGroups) *ListGatewayApiGroupsResponseBody {
	s.ApiGroups = v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetErrorField(v string) *ListGatewayApiGroupsResponseBody {
	s.ErrorField = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetHttpStatusCode(v int32) *ListGatewayApiGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetPageNumber(v int32) *ListGatewayApiGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetPageSize(v int32) *ListGatewayApiGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetRequestId(v string) *ListGatewayApiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetResultCode(v string) *ListGatewayApiGroupsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetResultMessage(v string) *ListGatewayApiGroupsResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetSuccess(v bool) *ListGatewayApiGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) SetTotal(v int64) *ListGatewayApiGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBody) Validate() error {
	if s.ApiGroups != nil {
		for _, item := range s.ApiGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayApiGroupsResponseBodyApiGroups struct {
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

func (s ListGatewayApiGroupsResponseBodyApiGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApiGroupsResponseBodyApiGroups) GoString() string {
	return s.String()
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetCenters() map[string]interface{} {
	return s.Centers
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetDescription() *string {
	return s.Description
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetDomainName() *string {
	return s.DomainName
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetHost() *string {
	return s.Host
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetCenters(v map[string]interface{}) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.Centers = v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetDescription(v string) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.Description = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetDomainName(v string) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.DomainName = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetEnableMTLS(v bool) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.EnableMTLS = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetGroupId(v int64) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.GroupId = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetGroupName(v string) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.GroupName = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetHost(v string) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.Host = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetHttpsModel(v bool) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.HttpsModel = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetInstanceId(v string) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.InstanceId = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetInstancePort(v int32) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.InstancePort = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetRegionMode(v bool) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.RegionMode = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetTimeout(v int32) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.Timeout = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetUpstreamType(v string) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.UpstreamType = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetVpcId(v string) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.VpcId = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) SetVpcModel(v bool) *ListGatewayApiGroupsResponseBodyApiGroups {
	s.VpcModel = &v
	return s
}

func (s *ListGatewayApiGroupsResponseBodyApiGroups) Validate() error {
	return dara.Validate(s)
}
