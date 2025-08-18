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
	SetProductType(v string) *GetConnectionTicketRequest
	GetProductType() *string
	SetTaskId(v string) *GetConnectionTicketRequest
	GetTaskId() *string
}

type GetConnectionTicketRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The application ID.
	//
	// >  This parameter is required for the first call to this operation and optional for subsequent calls to the operation.
	//
	// example:
	//
	// ca-e4s0puhmwi7v****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery groups.
	//
	// > 	- If you configure this parameter, the system assigns application instances only among the specified authorized delivery groups.
	//
	// > 	- This parameter is required if you configure `AppInstanceId` or `AppInstancePersistentId`.
	AppInstanceGroupIdList []*string `json:"AppInstanceGroupIdList,omitempty" xml:"AppInstanceGroupIdList,omitempty" type:"Repeated"`
	// The ID of the application instance.
	//
	// > 	- If you configure this parameter, the system attempts to assign only the specified application instance.
	//
	// > 	- If you configure this parameter, you must also configure `AppInstanceGroupIdList` and the number of delivery groups specified by `AppInstanceGroupIdList` must be 1.
	//
	// example:
	//
	// ai-1rznfnrvsa99d****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the persistent session.
	//
	// example:
	//
	// p-0bxls9m3cl7s****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	AppPolicyId             *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// The parameters that are configured to start the application. For information about how to obtain these parameters, see [Obtain parameters configured to install and start an application](https://help.aliyun.com/document_detail/426045.html).
	//
	// example:
	//
	// /q /n
	AppStartParam *string `json:"AppStartParam,omitempty" xml:"AppStartParam,omitempty"`
	// The application version. If you configure this parameter, only an application of the specified version is started. If you do not configure this parameter, an application of a random authorized version is started.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The region ID.
	//
	// >  If you configure this parameter, the system assigns application instances only among the delivery groups that reside in the specified region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The ID of the convenience account.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The product type.
	//
	// Valid values:
	//
	// 	- CloudApp: App Streaming
	//
	// 	- AndroidCloud: Cloud Phone
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The task ID.
	//
	// >  This parameter is required for calls other than the first call to this operation. You can use this parameter to query the task status and connection credential.
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
