// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMiniGameQualificationByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMiniGameQualificationByAppResponseBody
	GetAccessDeniedDetail() *string
	SetContent(v []*QueryMiniGameQualificationByAppResponseBodyContent) *QueryMiniGameQualificationByAppResponseBody
	GetContent() []*QueryMiniGameQualificationByAppResponseBodyContent
	SetErrorCode(v string) *QueryMiniGameQualificationByAppResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *QueryMiniGameQualificationByAppResponseBody
	GetRequestId() *string
	SetResultMsg(v string) *QueryMiniGameQualificationByAppResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *QueryMiniGameQualificationByAppResponseBody
	GetSuccess() *bool
}

type QueryMiniGameQualificationByAppResponseBody struct {
	AccessDeniedDetail *string                                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Content            []*QueryMiniGameQualificationByAppResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	ErrorCode          *string                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId          *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg          *string                                               `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success            *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMiniGameQualificationByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameQualificationByAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMiniGameQualificationByAppResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMiniGameQualificationByAppResponseBody) GetContent() []*QueryMiniGameQualificationByAppResponseBodyContent {
	return s.Content
}

func (s *QueryMiniGameQualificationByAppResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMiniGameQualificationByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMiniGameQualificationByAppResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMiniGameQualificationByAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMiniGameQualificationByAppResponseBody) SetAccessDeniedDetail(v string) *QueryMiniGameQualificationByAppResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBody) SetContent(v []*QueryMiniGameQualificationByAppResponseBodyContent) *QueryMiniGameQualificationByAppResponseBody {
	s.Content = v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBody) SetErrorCode(v string) *QueryMiniGameQualificationByAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBody) SetRequestId(v string) *QueryMiniGameQualificationByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBody) SetResultMsg(v string) *QueryMiniGameQualificationByAppResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBody) SetSuccess(v bool) *QueryMiniGameQualificationByAppResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBody) Validate() error {
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

type QueryMiniGameQualificationByAppResponseBodyContent struct {
	BusinessLicenseName         *string                                                                          `json:"BusinessLicenseName,omitempty" xml:"BusinessLicenseName,omitempty"`
	BusinessLicenseUrl          *string                                                                          `json:"BusinessLicenseUrl,omitempty" xml:"BusinessLicenseUrl,omitempty"`
	CopyrightAuthorizationFiles []*QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles `json:"CopyrightAuthorizationFiles,omitempty" xml:"CopyrightAuthorizationFiles,omitempty" type:"Repeated"`
	CopyrightFileName           *string                                                                          `json:"CopyrightFileName,omitempty" xml:"CopyrightFileName,omitempty"`
	CopyrightRegistrationNumber *string                                                                          `json:"CopyrightRegistrationNumber,omitempty" xml:"CopyrightRegistrationNumber,omitempty"`
	CopyrightUrl                *string                                                                          `json:"CopyrightUrl,omitempty" xml:"CopyrightUrl,omitempty"`
	Developer                   *string                                                                          `json:"Developer,omitempty" xml:"Developer,omitempty"`
	GameType                    *string                                                                          `json:"GameType,omitempty" xml:"GameType,omitempty"`
	GameVersionId               *string                                                                          `json:"GameVersionId,omitempty" xml:"GameVersionId,omitempty"`
	GmtModified                 *int64                                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IcpFileName                 *string                                                                          `json:"IcpFileName,omitempty" xml:"IcpFileName,omitempty"`
	IcpUrl                      *string                                                                          `json:"IcpUrl,omitempty" xml:"IcpUrl,omitempty"`
	IdCode                      *string                                                                          `json:"IdCode,omitempty" xml:"IdCode,omitempty"`
	IpAuthorizationFiles        []*QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles        `json:"IpAuthorizationFiles,omitempty" xml:"IpAuthorizationFiles,omitempty" type:"Repeated"`
	MiniProgramCode             *string                                                                          `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	MiniProgramId               *int64                                                                           `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	OnlineGamePubReFileName     *string                                                                          `json:"OnlineGamePubReFileName,omitempty" xml:"OnlineGamePubReFileName,omitempty"`
	OnlineGamePubReUrl          *string                                                                          `json:"OnlineGamePubReUrl,omitempty" xml:"OnlineGamePubReUrl,omitempty"`
	PrivacyPolicyUrl            *string                                                                          `json:"PrivacyPolicyUrl,omitempty" xml:"PrivacyPolicyUrl,omitempty"`
	Publisher                   *string                                                                          `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	ReviewStatus                *string                                                                          `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
}

func (s QueryMiniGameQualificationByAppResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameQualificationByAppResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetBusinessLicenseName() *string {
	return s.BusinessLicenseName
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetBusinessLicenseUrl() *string {
	return s.BusinessLicenseUrl
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetCopyrightAuthorizationFiles() []*QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles {
	return s.CopyrightAuthorizationFiles
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetCopyrightFileName() *string {
	return s.CopyrightFileName
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetCopyrightRegistrationNumber() *string {
	return s.CopyrightRegistrationNumber
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetCopyrightUrl() *string {
	return s.CopyrightUrl
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetDeveloper() *string {
	return s.Developer
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetGameType() *string {
	return s.GameType
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetGameVersionId() *string {
	return s.GameVersionId
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetIcpFileName() *string {
	return s.IcpFileName
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetIcpUrl() *string {
	return s.IcpUrl
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetIdCode() *string {
	return s.IdCode
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetIpAuthorizationFiles() []*QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles {
	return s.IpAuthorizationFiles
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetMiniProgramCode() *string {
	return s.MiniProgramCode
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetMiniProgramId() *int64 {
	return s.MiniProgramId
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetOnlineGamePubReFileName() *string {
	return s.OnlineGamePubReFileName
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetOnlineGamePubReUrl() *string {
	return s.OnlineGamePubReUrl
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetPrivacyPolicyUrl() *string {
	return s.PrivacyPolicyUrl
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetPublisher() *string {
	return s.Publisher
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) GetReviewStatus() *string {
	return s.ReviewStatus
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetBusinessLicenseName(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.BusinessLicenseName = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetBusinessLicenseUrl(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.BusinessLicenseUrl = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetCopyrightAuthorizationFiles(v []*QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.CopyrightAuthorizationFiles = v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetCopyrightFileName(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.CopyrightFileName = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetCopyrightRegistrationNumber(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.CopyrightRegistrationNumber = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetCopyrightUrl(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.CopyrightUrl = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetDeveloper(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.Developer = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetGameType(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.GameType = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetGameVersionId(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.GameVersionId = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetGmtModified(v int64) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetIcpFileName(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.IcpFileName = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetIcpUrl(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.IcpUrl = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetIdCode(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.IdCode = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetIpAuthorizationFiles(v []*QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.IpAuthorizationFiles = v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetMiniProgramCode(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.MiniProgramCode = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetMiniProgramId(v int64) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.MiniProgramId = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetOnlineGamePubReFileName(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.OnlineGamePubReFileName = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetOnlineGamePubReUrl(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.OnlineGamePubReUrl = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetPrivacyPolicyUrl(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.PrivacyPolicyUrl = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetPublisher(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.Publisher = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) SetReviewStatus(v string) *QueryMiniGameQualificationByAppResponseBodyContent {
	s.ReviewStatus = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContent) Validate() error {
	if s.CopyrightAuthorizationFiles != nil {
		for _, item := range s.CopyrightAuthorizationFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpAuthorizationFiles != nil {
		for _, item := range s.IpAuthorizationFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) GoString() string {
	return s.String()
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) GetName() *string {
	return s.Name
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) GetUrl() *string {
	return s.Url
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) SetName(v string) *QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles {
	s.Name = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) SetUrl(v string) *QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles {
	s.Url = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentCopyrightAuthorizationFiles) Validate() error {
	return dara.Validate(s)
}

type QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) GoString() string {
	return s.String()
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) GetName() *string {
	return s.Name
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) GetUrl() *string {
	return s.Url
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) SetName(v string) *QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles {
	s.Name = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) SetUrl(v string) *QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles {
	s.Url = &v
	return s
}

func (s *QueryMiniGameQualificationByAppResponseBodyContentIpAuthorizationFiles) Validate() error {
	return dara.Validate(s)
}
