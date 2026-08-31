// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterConfig(v *GetComputeClusterResponseBodyClusterConfig) *GetComputeClusterResponseBody
	GetClusterConfig() *GetComputeClusterResponseBodyClusterConfig
	SetCode(v string) *GetComputeClusterResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetComputeClusterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetComputeClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetComputeClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetComputeClusterResponseBody
	GetSuccess() *bool
}

type GetComputeClusterResponseBody struct {
	// The cluster details.
	ClusterConfig *GetComputeClusterResponseBodyClusterConfig `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty" type:"Struct"`
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetComputeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeClusterResponseBody) GetClusterConfig() *GetComputeClusterResponseBodyClusterConfig {
	return s.ClusterConfig
}

func (s *GetComputeClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetComputeClusterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetComputeClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetComputeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetComputeClusterResponseBody) SetClusterConfig(v *GetComputeClusterResponseBodyClusterConfig) *GetComputeClusterResponseBody {
	s.ClusterConfig = v
	return s
}

func (s *GetComputeClusterResponseBody) SetCode(v string) *GetComputeClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetComputeClusterResponseBody) SetHttpStatusCode(v int32) *GetComputeClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetComputeClusterResponseBody) SetMessage(v string) *GetComputeClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetComputeClusterResponseBody) SetRequestId(v string) *GetComputeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeClusterResponseBody) SetSuccess(v bool) *GetComputeClusterResponseBody {
	s.Success = &v
	return s
}

func (s *GetComputeClusterResponseBody) Validate() error {
	if s.ClusterConfig != nil {
		if err := s.ClusterConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeClusterResponseBodyClusterConfig struct {
	// The cluster security control configuration.
	ClusterSafetyControl *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl `json:"ClusterSafetyControl,omitempty" xml:"ClusterSafetyControl,omitempty" type:"Struct"`
	// The cluster description.
	//
	// example:
	//
	// test
	Des *string `json:"Des,omitempty" xml:"Des,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cluster_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster owner.
	//
	// example:
	//
	// 30012211
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The cluster version.
	//
	// example:
	//
	// CDH6
	TypeVersion *string `json:"TypeVersion,omitempty" xml:"TypeVersion,omitempty"`
}

func (s GetComputeClusterResponseBodyClusterConfig) String() string {
	return dara.Prettify(s)
}

func (s GetComputeClusterResponseBodyClusterConfig) GoString() string {
	return s.String()
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetClusterSafetyControl() *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl {
	return s.ClusterSafetyControl
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetDes() *string {
	return s.Des
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetId() *int64 {
	return s.Id
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetName() *string {
	return s.Name
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetOwner() *string {
	return s.Owner
}

func (s *GetComputeClusterResponseBodyClusterConfig) GetTypeVersion() *string {
	return s.TypeVersion
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetClusterSafetyControl(v *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) *GetComputeClusterResponseBodyClusterConfig {
	s.ClusterSafetyControl = v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetDes(v string) *GetComputeClusterResponseBodyClusterConfig {
	s.Des = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetGmtCreate(v string) *GetComputeClusterResponseBodyClusterConfig {
	s.GmtCreate = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetGmtModified(v string) *GetComputeClusterResponseBodyClusterConfig {
	s.GmtModified = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetId(v int64) *GetComputeClusterResponseBodyClusterConfig {
	s.Id = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetName(v string) *GetComputeClusterResponseBodyClusterConfig {
	s.Name = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetOwner(v string) *GetComputeClusterResponseBodyClusterConfig {
	s.Owner = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) SetTypeVersion(v string) *GetComputeClusterResponseBodyClusterConfig {
	s.TypeVersion = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfig) Validate() error {
	if s.ClusterSafetyControl != nil {
		if err := s.ClusterSafetyControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeClusterResponseBodyClusterConfigClusterSafetyControl struct {
	// The control mode.
	//
	// example:
	//
	// CREATE_COMPUTE_SOURCE
	ClusterSafetyAuthType *string `json:"ClusterSafetyAuthType,omitempty" xml:"ClusterSafetyAuthType,omitempty"`
	// The list of whitelisted user group IDs.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of whitelisted user group names.
	UserGroupNames []*string `json:"UserGroupNames,omitempty" xml:"UserGroupNames,omitempty" type:"Repeated"`
	// The list of whitelisted user IDs.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// The list of whitelisted usernames.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) String() string {
	return dara.Prettify(s)
}

func (s GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) GoString() string {
	return s.String()
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) GetClusterSafetyAuthType() *string {
	return s.ClusterSafetyAuthType
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) GetUserGroupNames() []*string {
	return s.UserGroupNames
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) GetUserNames() []*string {
	return s.UserNames
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) SetClusterSafetyAuthType(v string) *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl {
	s.ClusterSafetyAuthType = &v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) SetUserGroupIds(v []*string) *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl {
	s.UserGroupIds = v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) SetUserGroupNames(v []*string) *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl {
	s.UserGroupNames = v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) SetUserIds(v []*string) *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl {
	s.UserIds = v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) SetUserNames(v []*string) *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl {
	s.UserNames = v
	return s
}

func (s *GetComputeClusterResponseBodyClusterConfigClusterSafetyControl) Validate() error {
	return dara.Validate(s)
}
