// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPlayInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlayInfo(v *GetVideoPlayInfoResponseBodyPlayInfo) *GetVideoPlayInfoResponseBody
	GetPlayInfo() *GetVideoPlayInfoResponseBodyPlayInfo
	SetRequestId(v string) *GetVideoPlayInfoResponseBody
	GetRequestId() *string
	SetVideoInfo(v *GetVideoPlayInfoResponseBodyVideoInfo) *GetVideoPlayInfoResponseBody
	GetVideoInfo() *GetVideoPlayInfoResponseBodyVideoInfo
}

type GetVideoPlayInfoResponseBody struct {
	PlayInfo  *GetVideoPlayInfoResponseBodyPlayInfo  `json:"PlayInfo,omitempty" xml:"PlayInfo,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoInfo *GetVideoPlayInfoResponseBodyVideoInfo `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s GetVideoPlayInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoPlayInfoResponseBody) GetPlayInfo() *GetVideoPlayInfoResponseBodyPlayInfo {
	return s.PlayInfo
}

func (s *GetVideoPlayInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoPlayInfoResponseBody) GetVideoInfo() *GetVideoPlayInfoResponseBodyVideoInfo {
	return s.VideoInfo
}

func (s *GetVideoPlayInfoResponseBody) SetPlayInfo(v *GetVideoPlayInfoResponseBodyPlayInfo) *GetVideoPlayInfoResponseBody {
	s.PlayInfo = v
	return s
}

func (s *GetVideoPlayInfoResponseBody) SetRequestId(v string) *GetVideoPlayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoPlayInfoResponseBody) SetVideoInfo(v *GetVideoPlayInfoResponseBodyVideoInfo) *GetVideoPlayInfoResponseBody {
	s.VideoInfo = v
	return s
}

func (s *GetVideoPlayInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoPlayInfoResponseBodyPlayInfo struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	AuthInfo        *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	PlayDomain      *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetVideoPlayInfoResponseBodyPlayInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayInfoResponseBodyPlayInfo) GoString() string {
	return s.String()
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) GetRegion() *string {
	return s.Region
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) SetAccessKeyId(v string) *GetVideoPlayInfoResponseBodyPlayInfo {
	s.AccessKeyId = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) SetAccessKeySecret(v string) *GetVideoPlayInfoResponseBodyPlayInfo {
	s.AccessKeySecret = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) SetAuthInfo(v string) *GetVideoPlayInfoResponseBodyPlayInfo {
	s.AuthInfo = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) SetPlayDomain(v string) *GetVideoPlayInfoResponseBodyPlayInfo {
	s.PlayDomain = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) SetRegion(v string) *GetVideoPlayInfoResponseBodyPlayInfo {
	s.Region = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) SetSecurityToken(v string) *GetVideoPlayInfoResponseBodyPlayInfo {
	s.SecurityToken = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyPlayInfo) Validate() error {
	return dara.Validate(s)
}

type GetVideoPlayInfoResponseBodyVideoInfo struct {
	CoverURL   *string  `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CustomerId *int64   `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Duration   *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Status     *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title      *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId    *string  `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoPlayInfoResponseBodyVideoInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayInfoResponseBodyVideoInfo) GoString() string {
	return s.String()
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) GetDuration() *float32 {
	return s.Duration
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) GetStatus() *string {
	return s.Status
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) GetTitle() *string {
	return s.Title
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) SetCoverURL(v string) *GetVideoPlayInfoResponseBodyVideoInfo {
	s.CoverURL = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) SetCustomerId(v int64) *GetVideoPlayInfoResponseBodyVideoInfo {
	s.CustomerId = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) SetDuration(v float32) *GetVideoPlayInfoResponseBodyVideoInfo {
	s.Duration = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) SetStatus(v string) *GetVideoPlayInfoResponseBodyVideoInfo {
	s.Status = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) SetTitle(v string) *GetVideoPlayInfoResponseBodyVideoInfo {
	s.Title = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) SetVideoId(v string) *GetVideoPlayInfoResponseBodyVideoInfo {
	s.VideoId = &v
	return s
}

func (s *GetVideoPlayInfoResponseBodyVideoInfo) Validate() error {
	return dara.Validate(s)
}
