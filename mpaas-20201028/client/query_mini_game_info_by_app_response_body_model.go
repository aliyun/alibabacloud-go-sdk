// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMiniGameInfoByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMiniGameInfoByAppResponseBody
	GetAccessDeniedDetail() *string
	SetContent(v []*QueryMiniGameInfoByAppResponseBodyContent) *QueryMiniGameInfoByAppResponseBody
	GetContent() []*QueryMiniGameInfoByAppResponseBodyContent
	SetErrorCode(v string) *QueryMiniGameInfoByAppResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *QueryMiniGameInfoByAppResponseBody
	GetRequestId() *string
	SetResultMsg(v string) *QueryMiniGameInfoByAppResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *QueryMiniGameInfoByAppResponseBody
	GetSuccess() *bool
}

type QueryMiniGameInfoByAppResponseBody struct {
	AccessDeniedDetail *string                                      `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Content            []*QueryMiniGameInfoByAppResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	ErrorCode          *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId          *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg          *string                                      `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success            *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMiniGameInfoByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameInfoByAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMiniGameInfoByAppResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMiniGameInfoByAppResponseBody) GetContent() []*QueryMiniGameInfoByAppResponseBodyContent {
	return s.Content
}

func (s *QueryMiniGameInfoByAppResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMiniGameInfoByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMiniGameInfoByAppResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMiniGameInfoByAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMiniGameInfoByAppResponseBody) SetAccessDeniedDetail(v string) *QueryMiniGameInfoByAppResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBody) SetContent(v []*QueryMiniGameInfoByAppResponseBodyContent) *QueryMiniGameInfoByAppResponseBody {
	s.Content = v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBody) SetErrorCode(v string) *QueryMiniGameInfoByAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBody) SetRequestId(v string) *QueryMiniGameInfoByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBody) SetResultMsg(v string) *QueryMiniGameInfoByAppResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBody) SetSuccess(v bool) *QueryMiniGameInfoByAppResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMiniGameInfoByAppResponseBodyContent struct {
	GameEngine      *string                                                   `json:"GameEngine,omitempty" xml:"GameEngine,omitempty"`
	GameMaker       *string                                                   `json:"GameMaker,omitempty" xml:"GameMaker,omitempty"`
	GameTypeLevel1  *string                                                   `json:"GameTypeLevel1,omitempty" xml:"GameTypeLevel1,omitempty"`
	GameTypeLevel2  *string                                                   `json:"GameTypeLevel2,omitempty" xml:"GameTypeLevel2,omitempty"`
	GameTypeLevel3  *string                                                   `json:"GameTypeLevel3,omitempty" xml:"GameTypeLevel3,omitempty"`
	GameVersionId   *string                                                   `json:"GameVersionId,omitempty" xml:"GameVersionId,omitempty"`
	GmtModified     *int64                                                    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Icon            *string                                                   `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Introduction    *string                                                   `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	MiniProgramCode *string                                                   `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	MiniProgramId   *int64                                                    `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	MiniProgramName *string                                                   `json:"MiniProgramName,omitempty" xml:"MiniProgramName,omitempty"`
	Slogan          *string                                                   `json:"Slogan,omitempty" xml:"Slogan,omitempty"`
	UserProfileData *QueryMiniGameInfoByAppResponseBodyContentUserProfileData `json:"UserProfileData,omitempty" xml:"UserProfileData,omitempty" type:"Struct"`
	Version         *string                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryMiniGameInfoByAppResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameInfoByAppResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetGameEngine() *string {
	return s.GameEngine
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetGameMaker() *string {
	return s.GameMaker
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetGameTypeLevel1() *string {
	return s.GameTypeLevel1
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetGameTypeLevel2() *string {
	return s.GameTypeLevel2
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetGameTypeLevel3() *string {
	return s.GameTypeLevel3
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetGameVersionId() *string {
	return s.GameVersionId
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetIcon() *string {
	return s.Icon
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetIntroduction() *string {
	return s.Introduction
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetMiniProgramCode() *string {
	return s.MiniProgramCode
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetMiniProgramId() *int64 {
	return s.MiniProgramId
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetMiniProgramName() *string {
	return s.MiniProgramName
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetSlogan() *string {
	return s.Slogan
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetUserProfileData() *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	return s.UserProfileData
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) GetVersion() *string {
	return s.Version
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetGameEngine(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.GameEngine = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetGameMaker(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.GameMaker = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetGameTypeLevel1(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.GameTypeLevel1 = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetGameTypeLevel2(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.GameTypeLevel2 = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetGameTypeLevel3(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.GameTypeLevel3 = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetGameVersionId(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.GameVersionId = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetGmtModified(v int64) *QueryMiniGameInfoByAppResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetIcon(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.Icon = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetIntroduction(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.Introduction = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetMiniProgramCode(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.MiniProgramCode = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetMiniProgramId(v int64) *QueryMiniGameInfoByAppResponseBodyContent {
	s.MiniProgramId = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetMiniProgramName(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.MiniProgramName = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetSlogan(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.Slogan = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetUserProfileData(v *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) *QueryMiniGameInfoByAppResponseBodyContent {
	s.UserProfileData = v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetVersion(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.Version = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) Validate() error {
	if s.UserProfileData != nil {
		if err := s.UserProfileData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMiniGameInfoByAppResponseBodyContentUserProfileData struct {
	ActiveRegionCodes   *string `json:"ActiveRegionCodes,omitempty" xml:"ActiveRegionCodes,omitempty"`
	ActiveRegions       *string `json:"ActiveRegions,omitempty" xml:"ActiveRegions,omitempty"`
	AgeSegment          *string `json:"AgeSegment,omitempty" xml:"AgeSegment,omitempty"`
	AndroidActiveModels *string `json:"AndroidActiveModels,omitempty" xml:"AndroidActiveModels,omitempty"`
	AndroidRatio        *int32  `json:"AndroidRatio,omitempty" xml:"AndroidRatio,omitempty"`
	FemaleRatio         *int32  `json:"FemaleRatio,omitempty" xml:"FemaleRatio,omitempty"`
	HarmonyRatio        *int32  `json:"HarmonyRatio,omitempty" xml:"HarmonyRatio,omitempty"`
	IosActiveModels     *string `json:"IosActiveModels,omitempty" xml:"IosActiveModels,omitempty"`
	IosRatio            *int32  `json:"IosRatio,omitempty" xml:"IosRatio,omitempty"`
	MaleRatio           *int32  `json:"MaleRatio,omitempty" xml:"MaleRatio,omitempty"`
	ProfileRemark       *string `json:"ProfileRemark,omitempty" xml:"ProfileRemark,omitempty"`
}

func (s QueryMiniGameInfoByAppResponseBodyContentUserProfileData) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GoString() string {
	return s.String()
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetActiveRegionCodes() *string {
	return s.ActiveRegionCodes
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetActiveRegions() *string {
	return s.ActiveRegions
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetAgeSegment() *string {
	return s.AgeSegment
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetAndroidActiveModels() *string {
	return s.AndroidActiveModels
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetAndroidRatio() *int32 {
	return s.AndroidRatio
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetFemaleRatio() *int32 {
	return s.FemaleRatio
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetHarmonyRatio() *int32 {
	return s.HarmonyRatio
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetIosActiveModels() *string {
	return s.IosActiveModels
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetIosRatio() *int32 {
	return s.IosRatio
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetMaleRatio() *int32 {
	return s.MaleRatio
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) GetProfileRemark() *string {
	return s.ProfileRemark
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetActiveRegionCodes(v string) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.ActiveRegionCodes = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetActiveRegions(v string) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.ActiveRegions = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetAgeSegment(v string) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.AgeSegment = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetAndroidActiveModels(v string) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.AndroidActiveModels = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetAndroidRatio(v int32) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.AndroidRatio = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetFemaleRatio(v int32) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.FemaleRatio = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetHarmonyRatio(v int32) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.HarmonyRatio = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetIosActiveModels(v string) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.IosActiveModels = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetIosRatio(v int32) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.IosRatio = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetMaleRatio(v int32) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.MaleRatio = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) SetProfileRemark(v string) *QueryMiniGameInfoByAppResponseBodyContentUserProfileData {
	s.ProfileRemark = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContentUserProfileData) Validate() error {
	return dara.Validate(s)
}
