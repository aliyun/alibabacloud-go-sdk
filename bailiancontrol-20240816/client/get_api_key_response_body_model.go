// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *GetApiKeyResponseBodyApiKey) *GetApiKeyResponseBody
	GetApiKey() *GetApiKeyResponseBodyApiKey
	SetCode(v string) *GetApiKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApiKeyResponseBody
	GetSuccess() *bool
}

type GetApiKeyResponseBody struct {
	ApiKey         *GetApiKeyResponseBodyApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty" type:"Struct"`
	Code           *string                      `json:"code,omitempty" xml:"code,omitempty"`
	HttpStatusCode *int32                       `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                      `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBody) GetApiKey() *GetApiKeyResponseBodyApiKey {
	return s.ApiKey
}

func (s *GetApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApiKeyResponseBody) SetApiKey(v *GetApiKeyResponseBodyApiKey) *GetApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *GetApiKeyResponseBody) SetCode(v string) *GetApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiKeyResponseBody) SetHttpStatusCode(v int32) *GetApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetApiKeyResponseBody) SetMessage(v string) *GetApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetApiKeyResponseBody) SetRequestId(v string) *GetApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiKeyResponseBody) SetSuccess(v bool) *GetApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiKeyResponseBodyApiKey struct {
	ApiKeyValue  *string                                  `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	ApikeyId     *string                                  `json:"apikeyId,omitempty" xml:"apikeyId,omitempty"`
	AuthSetModel *GetApiKeyResponseBodyApiKeyAuthSetModel `json:"authSetModel,omitempty" xml:"authSetModel,omitempty" type:"Struct"`
	Blocked      *string                                  `json:"blocked,omitempty" xml:"blocked,omitempty"`
	CreateTime   *int64                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator      *string                                  `json:"creator,omitempty" xml:"creator,omitempty"`
	Description  *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	ExpireTime   *int64                                   `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	ExtData      *string                                  `json:"extData,omitempty" xml:"extData,omitempty"`
	ParentUid    *string                                  `json:"parentUid,omitempty" xml:"parentUid,omitempty"`
	Uid          *string                                  `json:"uid,omitempty" xml:"uid,omitempty"`
	WorkspaceId  *string                                  `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetApiKeyResponseBodyApiKey) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBodyApiKey) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBodyApiKey) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *GetApiKeyResponseBodyApiKey) GetApikeyId() *string {
	return s.ApikeyId
}

func (s *GetApiKeyResponseBodyApiKey) GetAuthSetModel() *GetApiKeyResponseBodyApiKeyAuthSetModel {
	return s.AuthSetModel
}

func (s *GetApiKeyResponseBodyApiKey) GetBlocked() *string {
	return s.Blocked
}

func (s *GetApiKeyResponseBodyApiKey) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApiKeyResponseBodyApiKey) GetCreator() *string {
	return s.Creator
}

func (s *GetApiKeyResponseBodyApiKey) GetDescription() *string {
	return s.Description
}

func (s *GetApiKeyResponseBodyApiKey) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetApiKeyResponseBodyApiKey) GetExtData() *string {
	return s.ExtData
}

func (s *GetApiKeyResponseBodyApiKey) GetParentUid() *string {
	return s.ParentUid
}

func (s *GetApiKeyResponseBodyApiKey) GetUid() *string {
	return s.Uid
}

func (s *GetApiKeyResponseBodyApiKey) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetApiKeyResponseBodyApiKey) SetApiKeyValue(v string) *GetApiKeyResponseBodyApiKey {
	s.ApiKeyValue = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetApikeyId(v string) *GetApiKeyResponseBodyApiKey {
	s.ApikeyId = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetAuthSetModel(v *GetApiKeyResponseBodyApiKeyAuthSetModel) *GetApiKeyResponseBodyApiKey {
	s.AuthSetModel = v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetBlocked(v string) *GetApiKeyResponseBodyApiKey {
	s.Blocked = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetCreateTime(v int64) *GetApiKeyResponseBodyApiKey {
	s.CreateTime = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetCreator(v string) *GetApiKeyResponseBodyApiKey {
	s.Creator = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetDescription(v string) *GetApiKeyResponseBodyApiKey {
	s.Description = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetExpireTime(v int64) *GetApiKeyResponseBodyApiKey {
	s.ExpireTime = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetExtData(v string) *GetApiKeyResponseBodyApiKey {
	s.ExtData = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetParentUid(v string) *GetApiKeyResponseBodyApiKey {
	s.ParentUid = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetUid(v string) *GetApiKeyResponseBodyApiKey {
	s.Uid = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetWorkspaceId(v string) *GetApiKeyResponseBodyApiKey {
	s.WorkspaceId = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) Validate() error {
	if s.AuthSetModel != nil {
		if err := s.AuthSetModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiKeyResponseBodyApiKeyAuthSetModel struct {
	AccessIps          []*string                                                  `json:"accessIps,omitempty" xml:"accessIps,omitempty" type:"Repeated"`
	AuthAppStructure   *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure   `json:"authAppStructure,omitempty" xml:"authAppStructure,omitempty" type:"Struct"`
	AuthModelStructure *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure `json:"authModelStructure,omitempty" xml:"authModelStructure,omitempty" type:"Struct"`
	AuthSetMode        *string                                                    `json:"authSetMode,omitempty" xml:"authSetMode,omitempty"`
}

func (s GetApiKeyResponseBodyApiKeyAuthSetModel) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBodyApiKeyAuthSetModel) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) GetAccessIps() []*string {
	return s.AccessIps
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) GetAuthAppStructure() *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure {
	return s.AuthAppStructure
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) GetAuthModelStructure() *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure {
	return s.AuthModelStructure
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) GetAuthSetMode() *string {
	return s.AuthSetMode
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) SetAccessIps(v []*string) *GetApiKeyResponseBodyApiKeyAuthSetModel {
	s.AccessIps = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) SetAuthAppStructure(v *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) *GetApiKeyResponseBodyApiKeyAuthSetModel {
	s.AuthAppStructure = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) SetAuthModelStructure(v *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) *GetApiKeyResponseBodyApiKeyAuthSetModel {
	s.AuthModelStructure = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) SetAuthSetMode(v string) *GetApiKeyResponseBodyApiKeyAuthSetModel {
	s.AuthSetMode = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModel) Validate() error {
	if s.AuthAppStructure != nil {
		if err := s.AuthAppStructure.Validate(); err != nil {
			return err
		}
	}
	if s.AuthModelStructure != nil {
		if err := s.AuthModelStructure.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure struct {
	Agents               []*string `json:"agents,omitempty" xml:"agents,omitempty" type:"Repeated"`
	HighCodeApps         []*string `json:"highCodeApps,omitempty" xml:"highCodeApps,omitempty" type:"Repeated"`
	IsAllowAccessAllApps *bool     `json:"isAllowAccessAllApps,omitempty" xml:"isAllowAccessAllApps,omitempty"`
	Workflows            []*string `json:"workflows,omitempty" xml:"workflows,omitempty" type:"Repeated"`
}

func (s GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) GetAgents() []*string {
	return s.Agents
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) GetHighCodeApps() []*string {
	return s.HighCodeApps
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) GetIsAllowAccessAllApps() *bool {
	return s.IsAllowAccessAllApps
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) GetWorkflows() []*string {
	return s.Workflows
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) SetAgents(v []*string) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure {
	s.Agents = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) SetHighCodeApps(v []*string) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure {
	s.HighCodeApps = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) SetIsAllowAccessAllApps(v bool) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure {
	s.IsAllowAccessAllApps = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) SetWorkflows(v []*string) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure {
	s.Workflows = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthAppStructure) Validate() error {
	return dara.Validate(s)
}

type GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure struct {
	DefineModels           []*string `json:"defineModels,omitempty" xml:"defineModels,omitempty" type:"Repeated"`
	Deployments            []*string `json:"deployments,omitempty" xml:"deployments,omitempty" type:"Repeated"`
	IsAllowAccessAllModels *bool     `json:"isAllowAccessAllModels,omitempty" xml:"isAllowAccessAllModels,omitempty"`
	Models                 []*string `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
}

func (s GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) GetDefineModels() []*string {
	return s.DefineModels
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) GetDeployments() []*string {
	return s.Deployments
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) GetIsAllowAccessAllModels() *bool {
	return s.IsAllowAccessAllModels
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) GetModels() []*string {
	return s.Models
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) SetDefineModels(v []*string) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure {
	s.DefineModels = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) SetDeployments(v []*string) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure {
	s.Deployments = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) SetIsAllowAccessAllModels(v bool) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure {
	s.IsAllowAccessAllModels = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) SetModels(v []*string) *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure {
	s.Models = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuthSetModelAuthModelStructure) Validate() error {
	return dara.Validate(s)
}
