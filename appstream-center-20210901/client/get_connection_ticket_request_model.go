// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *GetConnectionTicketRequest
	GetAccessType() *string
	SetAppId(v string) *GetConnectionTicketRequest
	GetAppId() *string
	SetAppInstanceGroupIdList(v []*string) *GetConnectionTicketRequest
	GetAppInstanceGroupIdList() []*string
	SetAppInstanceGroupSetId(v string) *GetConnectionTicketRequest
	GetAppInstanceGroupSetId() *string
	SetAppInstanceId(v string) *GetConnectionTicketRequest
	GetAppInstanceId() *string
	SetAppInstancePersistentId(v string) *GetConnectionTicketRequest
	GetAppInstancePersistentId() *string
	SetAppPolicyId(v string) *GetConnectionTicketRequest
	GetAppPolicyId() *string
	SetAppStartParam(v string) *GetConnectionTicketRequest
	GetAppStartParam() *string
	SetAppVersion(v string) *GetConnectionTicketRequest
	GetAppVersion() *string
	SetBizRegionId(v string) *GetConnectionTicketRequest
	GetBizRegionId() *string
	SetEndUserId(v string) *GetConnectionTicketRequest
	GetEndUserId() *string
	SetEnvironmentConfig(v string) *GetConnectionTicketRequest
	GetEnvironmentConfig() *string
	SetProductType(v string) *GetConnectionTicketRequest
	GetProductType() *string
	SetTaskId(v string) *GetConnectionTicketRequest
	GetTaskId() *string
}

type GetConnectionTicketRequest struct {
	// The access type.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The application ID.
	//
	// > This parameter is required for the initial call and optional for subsequent calls.
	//
	// example:
	//
	// ca-e4s0puhmwi7v****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The list of delivery groups.
	//
	// > - If you specify this parameter, application instances are allocated only from the specified and authorized delivery groups.
	//
	// > - If you specify the `AppInstanceId` or `AppInstancePersistentId` parameter, this parameter is required.
	AppInstanceGroupIdList []*string `json:"AppInstanceGroupIdList,omitempty" xml:"AppInstanceGroupIdList,omitempty" type:"Repeated"`
	// The delivery group set ID used to obtain the connection credential.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The application instance ID.
	//
	// >
	//
	// > - If you specify this parameter, only the specified application instance is allocated.
	//
	// > - If you specify this parameter, you must also specify the `AppInstanceGroupIdList` parameter, and the number of delivery group IDs in `AppInstanceGroupIdList` must be 1.
	//
	// example:
	//
	// ai-1rznfnrvsa99d****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The persistent session ID.
	//
	// example:
	//
	// p-0bxls9m3cl7s****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-0clfzcy0adpcf****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// The application startup parameter. This parameter is optional. You can refer to the method for specifying startup parameters in the image creation documentation and manually verify the startup parameters during image creation. This field is suitable for startup parameters with variable content, allowing API callers to set them flexibly. For more information about how to obtain startup parameters, see [How to obtain application installation parameters and startup parameters](https://help.aliyun.com/document_detail/426045.html).
	//
	// example:
	//
	// /q /n
	AppStartParam *string `json:"AppStartParam,omitempty" xml:"AppStartParam,omitempty"`
	// The application version. If you specify this parameter, only the specified version of the application is opened. If you do not specify this parameter, any authorized version of the application is opened.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The region ID.
	//
	// > If you specify this parameter, application instances are allocated only from delivery groups in the specified region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The environment configuration.
	//
	// example:
	//
	// {"userConfigReenter":"NATIVE"}
	EnvironmentConfig *string `json:"EnvironmentConfig,omitempty" xml:"EnvironmentConfig,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The node ID.
	//
	// > This parameter is required for non-initial calls. Use this parameter query to invoke the node status and connection credential retrieval.
	//
	// example:
	//
	// 28778acb-a469-4bc0-8e0f****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *GetConnectionTicketRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetConnectionTicketRequest) GetAppInstanceGroupIdList() []*string {
	return s.AppInstanceGroupIdList
}

func (s *GetConnectionTicketRequest) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *GetConnectionTicketRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *GetConnectionTicketRequest) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *GetConnectionTicketRequest) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *GetConnectionTicketRequest) GetAppStartParam() *string {
	return s.AppStartParam
}

func (s *GetConnectionTicketRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetConnectionTicketRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetConnectionTicketRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetConnectionTicketRequest) GetEnvironmentConfig() *string {
	return s.EnvironmentConfig
}

func (s *GetConnectionTicketRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetConnectionTicketRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetConnectionTicketRequest) SetAccessType(v string) *GetConnectionTicketRequest {
	s.AccessType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppId(v string) *GetConnectionTicketRequest {
	s.AppId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceGroupIdList(v []*string) *GetConnectionTicketRequest {
	s.AppInstanceGroupIdList = v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceGroupSetId(v string) *GetConnectionTicketRequest {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceId(v string) *GetConnectionTicketRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstancePersistentId(v string) *GetConnectionTicketRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppPolicyId(v string) *GetConnectionTicketRequest {
	s.AppPolicyId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppStartParam(v string) *GetConnectionTicketRequest {
	s.AppStartParam = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppVersion(v string) *GetConnectionTicketRequest {
	s.AppVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetBizRegionId(v string) *GetConnectionTicketRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEnvironmentConfig(v string) *GetConnectionTicketRequest {
	s.EnvironmentConfig = &v
	return s
}

func (s *GetConnectionTicketRequest) SetProductType(v string) *GetConnectionTicketRequest {
	s.ProductType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) Validate() error {
	return dara.Validate(s)
}
