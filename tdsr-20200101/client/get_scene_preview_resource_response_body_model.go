// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetScenePreviewResourceResponseBodyAccessDeniedDetail) *GetScenePreviewResourceResponseBody
	GetAccessDeniedDetail() *GetScenePreviewResourceResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetScenePreviewResourceResponseBody
	GetCode() *int64
	SetData(v *GetScenePreviewResourceResponseBodyData) *GetScenePreviewResourceResponseBody
	GetData() *GetScenePreviewResourceResponseBodyData
	SetMessage(v string) *GetScenePreviewResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScenePreviewResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScenePreviewResourceResponseBody
	GetSuccess() *bool
}

type GetScenePreviewResourceResponseBody struct {
	AccessDeniedDetail *GetScenePreviewResourceResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 0：成功，其他：失败
	Code *int64                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScenePreviewResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4B76*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true/false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePreviewResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponseBody) GetAccessDeniedDetail() *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetScenePreviewResourceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetScenePreviewResourceResponseBody) GetData() *GetScenePreviewResourceResponseBodyData {
	return s.Data
}

func (s *GetScenePreviewResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScenePreviewResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScenePreviewResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScenePreviewResourceResponseBody) SetAccessDeniedDetail(v *GetScenePreviewResourceResponseBodyAccessDeniedDetail) *GetScenePreviewResourceResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetCode(v int64) *GetScenePreviewResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetData(v *GetScenePreviewResourceResponseBodyData) *GetScenePreviewResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetMessage(v string) *GetScenePreviewResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetRequestId(v string) *GetScenePreviewResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetSuccess(v bool) *GetScenePreviewResourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetScenePreviewResourceResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenePreviewResourceResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetScenePreviewResourceResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewResourceResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetScenePreviewResourceResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetScenePreviewResourceResponseBodyData struct {
	Name              *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceDirectory *GetScenePreviewResourceResponseBodyDataResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetScenePreviewResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetScenePreviewResourceResponseBodyData) GetResourceDirectory() *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	return s.ResourceDirectory
}

func (s *GetScenePreviewResourceResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetScenePreviewResourceResponseBodyData) SetName(v string) *GetScenePreviewResourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyData) SetResourceDirectory(v *GetScenePreviewResourceResponseBodyDataResourceDirectory) *GetScenePreviewResourceResponseBodyData {
	s.ResourceDirectory = v
	return s
}

func (s *GetScenePreviewResourceResponseBodyData) SetVersion(v string) *GetScenePreviewResourceResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyData) Validate() error {
	if s.ResourceDirectory != nil {
		if err := s.ResourceDirectory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenePreviewResourceResponseBodyDataResourceDirectory struct {
	// example:
	//
	// hotspotTag.json
	HotspotTagConfig *string `json:"HotspotTagConfig,omitempty" xml:"HotspotTagConfig,omitempty"`
	// example:
	//
	// config.json
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// example:
	//
	// orthomap/orthomap.json
	OrthomapConfig *string `json:"OrthomapConfig,omitempty" xml:"OrthomapConfig,omitempty"`
	// example:
	//
	// A.e.YKPYuuYuituy****
	RootPath *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
}

func (s GetScenePreviewResourceResponseBodyDataResourceDirectory) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewResourceResponseBodyDataResourceDirectory) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) GetHotspotTagConfig() *string {
	return s.HotspotTagConfig
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) GetOrthomapConfig() *string {
	return s.OrthomapConfig
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) GetRootPath() *string {
	return s.RootPath
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetHotspotTagConfig(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.HotspotTagConfig = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetModelConfig(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.ModelConfig = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetOrthomapConfig(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.OrthomapConfig = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetRootPath(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.RootPath = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) Validate() error {
	return dara.Validate(s)
}
