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
	GameEngine      *string `json:"GameEngine,omitempty" xml:"GameEngine,omitempty"`
	GameMaker       *string `json:"GameMaker,omitempty" xml:"GameMaker,omitempty"`
	GameTypeLevel1  *string `json:"GameTypeLevel1,omitempty" xml:"GameTypeLevel1,omitempty"`
	GameTypeLevel2  *string `json:"GameTypeLevel2,omitempty" xml:"GameTypeLevel2,omitempty"`
	GameTypeLevel3  *string `json:"GameTypeLevel3,omitempty" xml:"GameTypeLevel3,omitempty"`
	GameVersionId   *string `json:"GameVersionId,omitempty" xml:"GameVersionId,omitempty"`
	GmtModified     *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Icon            *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Introduction    *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	MiniProgramCode *string `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	MiniProgramId   *int64  `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	MiniProgramName *string `json:"MiniProgramName,omitempty" xml:"MiniProgramName,omitempty"`
	Slogan          *string `json:"Slogan,omitempty" xml:"Slogan,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

func (s *QueryMiniGameInfoByAppResponseBodyContent) SetVersion(v string) *QueryMiniGameInfoByAppResponseBodyContent {
	s.Version = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
