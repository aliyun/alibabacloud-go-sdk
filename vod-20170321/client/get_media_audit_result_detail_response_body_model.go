// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaAuditResultDetail(v *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) *GetMediaAuditResultDetailResponseBody
	GetMediaAuditResultDetail() *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail
	SetRequestId(v string) *GetMediaAuditResultDetailResponseBody
	GetRequestId() *string
}

type GetMediaAuditResultDetailResponseBody struct {
	// The details of the review results.
	MediaAuditResultDetail *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail `json:"MediaAuditResultDetail,omitempty" xml:"MediaAuditResultDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6438BD76-D523-46FC-956F-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditResultDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponseBody) GetMediaAuditResultDetail() *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail {
	return s.MediaAuditResultDetail
}

func (s *GetMediaAuditResultDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaAuditResultDetailResponseBody) SetMediaAuditResultDetail(v *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) *GetMediaAuditResultDetailResponseBody {
	s.MediaAuditResultDetail = v
	return s
}

func (s *GetMediaAuditResultDetailResponseBody) SetRequestId(v string) *GetMediaAuditResultDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBody) Validate() error {
	if s.MediaAuditResultDetail != nil {
		if err := s.MediaAuditResultDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail struct {
	// The list of video review result details.
	List []*GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of video review result screenshots.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) GetList() []*GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	return s.List
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) GetTotal() *int32 {
	return s.Total
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) SetList(v []*GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail {
	s.List = v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) SetTotal(v int32) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail {
	s.Total = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList struct {
	// The classification of the ad review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text contains politically sensitive content.
	//
	// - **porn**: Text contains pornographic content.
	//
	// - **abuse**: Text contains abusive content.
	//
	// - **terrorism**: Text contains terrorism-related content.
	//
	// - **contraband**: Text contains prohibited content.
	//
	// - **spam**: Text contains other spam content.
	//
	// - **npx**: Psoriasis ads.
	//
	// - **qrcode**: Contains a QR code.
	//
	// - **programCode**: Contains a mini program code.
	//
	// example:
	//
	// normal
	AdLabel *string `json:"AdLabel,omitempty" xml:"AdLabel,omitempty"`
	// The hit score of the video screenshot for the ad review result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100
	AdScore *string `json:"AdScore,omitempty" xml:"AdScore,omitempty"`
	// The classification of the undesirable scene review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **meaningless**: The image has no content (for example, a black screen or white screen).
	//
	// - **PIP**: Picture-in-Picture (PiP).
	//
	// - **smoking**: Smoking.
	//
	// - **drivelive**: In-car live streaming.
	//
	// example:
	//
	// normal
	LiveLabel *string `json:"LiveLabel,omitempty" xml:"LiveLabel,omitempty"`
	// The hit score of the video screenshot for the undesirable scene review result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100
	LiveScore *string `json:"LiveScore,omitempty" xml:"LiveScore,omitempty"`
	// The classification of the logo review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **TV**: Contains a regulated logo.
	//
	// - **trademark**: Contains a trademark.
	//
	// example:
	//
	// normal
	LogoLabel *string `json:"LogoLabel,omitempty" xml:"LogoLabel,omitempty"`
	// The hit score of the video screenshot for the logo review result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100
	LogoScore *string `json:"LogoScore,omitempty" xml:"LogoScore,omitempty"`
	// The classification of the pornography review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **porn**: Pornographic.
	//
	// - **sexy**: Sexy.
	//
	// example:
	//
	// normal
	PornLabel *string `json:"PornLabel,omitempty" xml:"PornLabel,omitempty"`
	// The hit score of the video screenshot for the pornography review result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100.00
	PornScore *string `json:"PornScore,omitempty" xml:"PornScore,omitempty"`
	// The classification of the terrorism review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **bloody**: Bloody.
	//
	// - **explosion**: Explosion and smoke.
	//
	// - **outfit**: Special attire.
	//
	// - **logo**: Special logo.
	//
	// - **weapon**: Weapon.
	//
	// - **politics**: Politically sensitive.
	//
	// - **violence**: Fighting.
	//
	// - **crowd**: Crowd gathering.
	//
	// - **parade**: Parade.
	//
	// - **carcrash**: Car crash scene.
	//
	// - **flag**: Flag.
	//
	// - **location**: Landmark.
	//
	// - **others**: Others.
	//
	// example:
	//
	// normal
	TerrorismLabel *string `json:"TerrorismLabel,omitempty" xml:"TerrorismLabel,omitempty"`
	// The hit score of the video screenshot for the terrorism review result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100.00
	TerrorismScore *string `json:"TerrorismScore,omitempty" xml:"TerrorismScore,omitempty"`
	// The position of the video screenshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video screenshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetAdLabel() *string {
	return s.AdLabel
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetAdScore() *string {
	return s.AdScore
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetLiveLabel() *string {
	return s.LiveLabel
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetLiveScore() *string {
	return s.LiveScore
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetLogoLabel() *string {
	return s.LogoLabel
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetLogoScore() *string {
	return s.LogoScore
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetPornLabel() *string {
	return s.PornLabel
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetPornScore() *string {
	return s.PornScore
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetTerrorismLabel() *string {
	return s.TerrorismLabel
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetTerrorismScore() *string {
	return s.TerrorismScore
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetAdLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.AdLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetAdScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.AdScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLiveLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LiveLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLiveScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LiveScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLogoLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LogoLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLogoScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LogoScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetPornLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.PornLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetPornScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.PornScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetTerrorismLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.TerrorismLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetTerrorismScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.TerrorismScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetTimestamp(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetUrl(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) Validate() error {
	return dara.Validate(s)
}
