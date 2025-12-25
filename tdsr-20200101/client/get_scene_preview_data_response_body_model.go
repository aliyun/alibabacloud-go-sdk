// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetScenePreviewDataResponseBodyAccessDeniedDetail) *GetScenePreviewDataResponseBody
	GetAccessDeniedDetail() *GetScenePreviewDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetScenePreviewDataResponseBody
	GetCode() *int64
	SetData(v *GetScenePreviewDataResponseBodyData) *GetScenePreviewDataResponseBody
	GetData() *GetScenePreviewDataResponseBodyData
	SetMessage(v string) *GetScenePreviewDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScenePreviewDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScenePreviewDataResponseBody
	GetSuccess() *bool
}

type GetScenePreviewDataResponseBody struct {
	AccessDeniedDetail *GetScenePreviewDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 0：成功，其他：失败
	Code *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScenePreviewDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4B76FF387
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true/false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePreviewDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBody) GetAccessDeniedDetail() *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetScenePreviewDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetScenePreviewDataResponseBody) GetData() *GetScenePreviewDataResponseBodyData {
	return s.Data
}

func (s *GetScenePreviewDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScenePreviewDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScenePreviewDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScenePreviewDataResponseBody) SetAccessDeniedDetail(v *GetScenePreviewDataResponseBodyAccessDeniedDetail) *GetScenePreviewDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetCode(v int64) *GetScenePreviewDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetData(v *GetScenePreviewDataResponseBodyData) *GetScenePreviewDataResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetMessage(v string) *GetScenePreviewDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetRequestId(v string) *GetScenePreviewDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetSuccess(v bool) *GetScenePreviewDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetScenePreviewDataResponseBody) Validate() error {
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

type GetScenePreviewDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetScenePreviewDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetScenePreviewDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetScenePreviewDataResponseBodyData struct {
	Model *GetScenePreviewDataResponseBodyDataModel  `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Tags  []*GetScenePreviewDataResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetScenePreviewDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyData) GetModel() *GetScenePreviewDataResponseBodyDataModel {
	return s.Model
}

func (s *GetScenePreviewDataResponseBodyData) GetTags() []*GetScenePreviewDataResponseBodyDataTags {
	return s.Tags
}

func (s *GetScenePreviewDataResponseBodyData) SetModel(v *GetScenePreviewDataResponseBodyDataModel) *GetScenePreviewDataResponseBodyData {
	s.Model = v
	return s
}

func (s *GetScenePreviewDataResponseBodyData) SetTags(v []*GetScenePreviewDataResponseBodyDataTags) *GetScenePreviewDataResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetScenePreviewDataResponseBodyData) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScenePreviewDataResponseBodyDataModel struct {
	ModelPath        *string                                             `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	PanoList         []*GetScenePreviewDataResponseBodyDataModelPanoList `json:"PanoList,omitempty" xml:"PanoList,omitempty" type:"Repeated"`
	TextureModelPath *string                                             `json:"TextureModelPath,omitempty" xml:"TextureModelPath,omitempty"`
	TexturePanoPath  *string                                             `json:"TexturePanoPath,omitempty" xml:"TexturePanoPath,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataModel) GetModelPath() *string {
	return s.ModelPath
}

func (s *GetScenePreviewDataResponseBodyDataModel) GetPanoList() []*GetScenePreviewDataResponseBodyDataModelPanoList {
	return s.PanoList
}

func (s *GetScenePreviewDataResponseBodyDataModel) GetTextureModelPath() *string {
	return s.TextureModelPath
}

func (s *GetScenePreviewDataResponseBodyDataModel) GetTexturePanoPath() *string {
	return s.TexturePanoPath
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetModelPath(v string) *GetScenePreviewDataResponseBodyDataModel {
	s.ModelPath = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetPanoList(v []*GetScenePreviewDataResponseBodyDataModelPanoList) *GetScenePreviewDataResponseBodyDataModel {
	s.PanoList = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetTextureModelPath(v string) *GetScenePreviewDataResponseBodyDataModel {
	s.TextureModelPath = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetTexturePanoPath(v string) *GetScenePreviewDataResponseBodyDataModel {
	s.TexturePanoPath = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModel) Validate() error {
	if s.PanoList != nil {
		for _, item := range s.PanoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScenePreviewDataResponseBodyDataModelPanoList struct {
	CurRoomPicList []*string `json:"CurRoomPicList,omitempty" xml:"CurRoomPicList,omitempty" type:"Repeated"`
	// example:
	//
	// true/false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 1****
	FloorIdx *string `json:"FloorIdx,omitempty" xml:"FloorIdx,omitempty"`
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true/false
	MainImage  *bool                                                     `json:"MainImage,omitempty" xml:"MainImage,omitempty"`
	Neighbours []*string                                                 `json:"Neighbours,omitempty" xml:"Neighbours,omitempty" type:"Repeated"`
	Position   *GetScenePreviewDataResponseBodyDataModelPanoListPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Struct"`
	// example:
	//
	// location_93132801658010****
	RawName  *string `json:"RawName,omitempty" xml:"RawName,omitempty"`
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// 1****
	RoomIdx *string `json:"RoomIdx,omitempty" xml:"RoomIdx,omitempty"`
	// example:
	//
	// a7RqCd3kLMgglmn****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// token
	//
	// example:
	//
	// sIPGWRGLJHEIQE****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// 93132801658010****
	VirtualId   *string `json:"VirtualId,omitempty" xml:"VirtualId,omitempty"`
	VirtualName *string `json:"VirtualName,omitempty" xml:"VirtualName,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataModelPanoList) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataModelPanoList) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetCurRoomPicList() []*string {
	return s.CurRoomPicList
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetFloorIdx() *string {
	return s.FloorIdx
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetId() *string {
	return s.Id
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetMainImage() *bool {
	return s.MainImage
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetNeighbours() []*string {
	return s.Neighbours
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetPosition() *GetScenePreviewDataResponseBodyDataModelPanoListPosition {
	return s.Position
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetRawName() *string {
	return s.RawName
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetResource() *string {
	return s.Resource
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetRoomIdx() *string {
	return s.RoomIdx
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetToken() *string {
	return s.Token
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetVirtualId() *string {
	return s.VirtualId
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) GetVirtualName() *string {
	return s.VirtualName
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetCurRoomPicList(v []*string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.CurRoomPicList = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetEnabled(v bool) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Enabled = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetFloorIdx(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.FloorIdx = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetId(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Id = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetMainImage(v bool) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.MainImage = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetNeighbours(v []*string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Neighbours = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetPosition(v *GetScenePreviewDataResponseBodyDataModelPanoListPosition) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Position = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetRawName(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.RawName = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetResource(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Resource = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetRoomIdx(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.RoomIdx = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetSubSceneId(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.SubSceneId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetToken(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Token = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetVirtualId(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.VirtualId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetVirtualName(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.VirtualName = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) Validate() error {
	if s.Position != nil {
		if err := s.Position.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenePreviewDataResponseBodyDataModelPanoListPosition struct {
	Rotation  []*float64 `json:"Rotation,omitempty" xml:"Rotation,omitempty" type:"Repeated"`
	Spot      []*float64 `json:"Spot,omitempty" xml:"Spot,omitempty" type:"Repeated"`
	Viewpoint []*float64 `json:"Viewpoint,omitempty" xml:"Viewpoint,omitempty" type:"Repeated"`
}

func (s GetScenePreviewDataResponseBodyDataModelPanoListPosition) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataModelPanoListPosition) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) GetRotation() []*float64 {
	return s.Rotation
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) GetSpot() []*float64 {
	return s.Spot
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) GetViewpoint() []*float64 {
	return s.Viewpoint
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) SetRotation(v []*float64) *GetScenePreviewDataResponseBodyDataModelPanoListPosition {
	s.Rotation = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) SetSpot(v []*float64) *GetScenePreviewDataResponseBodyDataModelPanoListPosition {
	s.Spot = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) SetViewpoint(v []*float64) *GetScenePreviewDataResponseBodyDataModelPanoListPosition {
	s.Viewpoint = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) Validate() error {
	return dara.Validate(s)
}

type GetScenePreviewDataResponseBodyDataTags struct {
	Config *GetScenePreviewDataResponseBodyDataTagsConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 1****
	Id               *string    `json:"Id,omitempty" xml:"Id,omitempty"`
	Position         []*float64 `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	PositionPanoCube []*float64 `json:"PositionPanoCube,omitempty" xml:"PositionPanoCube,omitempty" type:"Repeated"`
	// example:
	//
	// IMAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataTags) GetConfig() *GetScenePreviewDataResponseBodyDataTagsConfig {
	return s.Config
}

func (s *GetScenePreviewDataResponseBodyDataTags) GetId() *string {
	return s.Id
}

func (s *GetScenePreviewDataResponseBodyDataTags) GetPosition() []*float64 {
	return s.Position
}

func (s *GetScenePreviewDataResponseBodyDataTags) GetPositionPanoCube() []*float64 {
	return s.PositionPanoCube
}

func (s *GetScenePreviewDataResponseBodyDataTags) GetType() *string {
	return s.Type
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetConfig(v *GetScenePreviewDataResponseBodyDataTagsConfig) *GetScenePreviewDataResponseBodyDataTags {
	s.Config = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetId(v string) *GetScenePreviewDataResponseBodyDataTags {
	s.Id = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetPosition(v []*float64) *GetScenePreviewDataResponseBodyDataTags {
	s.Position = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetPositionPanoCube(v []*float64) *GetScenePreviewDataResponseBodyDataTags {
	s.PositionPanoCube = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetType(v string) *GetScenePreviewDataResponseBodyDataTags {
	s.Type = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenePreviewDataResponseBodyDataTagsConfig struct {
	// example:
	//
	// #00000
	BackgroundColor *string                                                    `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	ButtonConfig    *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig `json:"ButtonConfig,omitempty" xml:"ButtonConfig,omitempty" type:"Struct"`
	Content         *string                                                    `json:"Content,omitempty" xml:"Content,omitempty"`
	FormImgSize     []*int64                                                   `json:"FormImgSize,omitempty" xml:"FormImgSize,omitempty" type:"Repeated"`
	// example:
	//
	// true/false
	FormJumpType *bool `json:"FormJumpType,omitempty" xml:"FormJumpType,omitempty"`
	// example:
	//
	// default
	FormSelectImgType *string   `json:"FormSelectImgType,omitempty" xml:"FormSelectImgType,omitempty"`
	Images            []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// true/false
	IsTagVisibleBy3d *bool `json:"IsTagVisibleBy3d,omitempty" xml:"IsTagVisibleBy3d,omitempty"`
	// example:
	//
	// http://www.example.com/***
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// example:
	//
	// 1****
	PanoId           *string    `json:"PanoId,omitempty" xml:"PanoId,omitempty"`
	Position         []*float64 `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	PositionPanoCube []*float64 `json:"PositionPanoCube,omitempty" xml:"PositionPanoCube,omitempty" type:"Repeated"`
	RelatedPanoIds   []*string  `json:"RelatedPanoIds,omitempty" xml:"RelatedPanoIds,omitempty" type:"Repeated"`
	// example:
	//
	// 323****
	SceneId *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com/****.mp4
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataTagsConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataTagsConfig) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetBackgroundColor() *string {
	return s.BackgroundColor
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetButtonConfig() *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig {
	return s.ButtonConfig
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetContent() *string {
	return s.Content
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetFormImgSize() []*int64 {
	return s.FormImgSize
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetFormJumpType() *bool {
	return s.FormJumpType
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetFormSelectImgType() *string {
	return s.FormSelectImgType
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetImages() []*string {
	return s.Images
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetIsTagVisibleBy3d() *bool {
	return s.IsTagVisibleBy3d
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetLink() *string {
	return s.Link
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetPanoId() *string {
	return s.PanoId
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetPosition() []*float64 {
	return s.Position
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetPositionPanoCube() []*float64 {
	return s.PositionPanoCube
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetRelatedPanoIds() []*string {
	return s.RelatedPanoIds
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetSceneId() *int64 {
	return s.SceneId
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetTitle() *string {
	return s.Title
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) GetVideo() *string {
	return s.Video
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetBackgroundColor(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.BackgroundColor = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetButtonConfig(v *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.ButtonConfig = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetContent(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Content = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetFormImgSize(v []*int64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.FormImgSize = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetFormJumpType(v bool) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.FormJumpType = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetFormSelectImgType(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.FormSelectImgType = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetImages(v []*string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Images = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetIsTagVisibleBy3d(v bool) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.IsTagVisibleBy3d = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetLink(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Link = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetPanoId(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.PanoId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetPosition(v []*float64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Position = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetPositionPanoCube(v []*float64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.PositionPanoCube = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetRelatedPanoIds(v []*string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.RelatedPanoIds = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetSceneId(v int64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.SceneId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetTitle(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Title = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetVideo(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Video = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) Validate() error {
	if s.ButtonConfig != nil {
		if err := s.ButtonConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig struct {
	CustomText *string `json:"CustomText,omitempty" xml:"CustomText,omitempty"`
	// example:
	//
	// CLICK_CHECK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) GetCustomText() *string {
	return s.CustomText
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) GetType() *string {
	return s.Type
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) SetCustomText(v string) *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig {
	s.CustomText = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) SetType(v string) *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig {
	s.Type = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) Validate() error {
	return dara.Validate(s)
}
