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
	// Details about review results.
	MediaAuditResultDetail *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail `json:"MediaAuditResultDetail,omitempty" xml:"MediaAuditResultDetail,omitempty" type:"Struct"`
	// The ID of the request.
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
	return dara.Validate(s)
}

type GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail struct {
	// The review results returned.
	List []*GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of snapshots returned.
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
	return dara.Validate(s)
}

type GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList struct {
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content.
	//
	// 	- **ad**: other ads.
	//
	// 	- **politics**: political content in text.
	//
	// 	- **porn**: pornographic content in text.
	//
	// 	- **abuse**: verbal abuse in text.
	//
	// 	- **terrorism**: terrorist content in text.
	//
	// 	- **contraband**: prohibited content in text.
	//
	// 	- **spam**: spam content in text.
	//
	// 	- **npx**: illegal ad
	//
	// 	- **qrcode**: QR code.
	//
	// 	- **programCode**: mini program code.
	//
	// example:
	//
	// normal
	AdLabel *string `json:"AdLabel,omitempty" xml:"AdLabel,omitempty"`
	// The score of the video snapshot in the ad review result. Valid values: `[0,100]`. The value is rounded down to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	AdScore *string `json:"AdScore,omitempty" xml:"AdScore,omitempty"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content.
	//
	// 	- **meaningless**: meaningless content, such as a black or white screen.
	//
	// 	- **PIP**: picture-in-picture.
	//
	// 	- **smoking**: smoking.
	//
	// 	- **drivelive**: live broadcasting in a running vehicle.
	//
	// example:
	//
	// normal
	LiveLabel *string `json:"LiveLabel,omitempty" xml:"LiveLabel,omitempty"`
	// The score of the video snapshot in the undesirable content review result. Valid values: `[0,100]`. The value is rounded down to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	LiveScore *string `json:"LiveScore,omitempty" xml:"LiveScore,omitempty"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content.
	//
	// 	- **TV**: controlled TV station logo.
	//
	// 	- **trademark**: trademark.
	//
	// example:
	//
	// normal
	LogoLabel *string `json:"LogoLabel,omitempty" xml:"LogoLabel,omitempty"`
	// The score of the video snapshot in the logo review result. Valid values: `[0,100]`. The value is rounded down to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	LogoScore *string `json:"LogoScore,omitempty" xml:"LogoScore,omitempty"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**
	//
	// 	- **porn**
	//
	// 	- **sexy**
	//
	// example:
	//
	// normal
	PornLabel *string `json:"PornLabel,omitempty" xml:"PornLabel,omitempty"`
	// The score of the video snapshot in the pornographic content review result. Valid values: `[0,100]`. The value is rounded down to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100.00
	PornScore *string `json:"PornScore,omitempty" xml:"PornScore,omitempty"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**
	//
	// 	- **bloody**
	//
	// 	- **explosion**
	//
	// 	- **outfit**
	//
	// 	- **logo**
	//
	// 	- **weapon**
	//
	// 	- **politics**
	//
	// 	- **violence**
	//
	// 	- **crowd**
	//
	// 	- **parade**
	//
	// 	- **carcrash**
	//
	// 	- **flag**
	//
	// 	- **location**
	//
	// 	- **others**
	//
	// example:
	//
	// normal
	TerrorismLabel *string `json:"TerrorismLabel,omitempty" xml:"TerrorismLabel,omitempty"`
	// The score of the video snapshot in the terrorist content review result. Valid values: `[0,100]`. The value is rounded down to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100.00
	TerrorismScore *string `json:"TerrorismScore,omitempty" xml:"TerrorismScore,omitempty"`
	// The timestamp of the snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
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
