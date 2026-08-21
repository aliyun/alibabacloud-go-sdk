// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaAuditResultTimeline(v *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) *GetMediaAuditResultTimelineResponseBody
	GetMediaAuditResultTimeline() *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline
	SetRequestId(v string) *GetMediaAuditResultTimelineResponseBody
	GetRequestId() *string
}

type GetMediaAuditResultTimelineResponseBody struct {
	// The collection of automated review result timelines.
	MediaAuditResultTimeline *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline `json:"MediaAuditResultTimeline,omitempty" xml:"MediaAuditResultTimeline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6438BD76-D523-46FC-956F-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBody) GetMediaAuditResultTimeline() *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	return s.MediaAuditResultTimeline
}

func (s *GetMediaAuditResultTimelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaAuditResultTimelineResponseBody) SetMediaAuditResultTimeline(v *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) *GetMediaAuditResultTimelineResponseBody {
	s.MediaAuditResultTimeline = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBody) SetRequestId(v string) *GetMediaAuditResultTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBody) Validate() error {
	if s.MediaAuditResultTimeline != nil {
		if err := s.MediaAuditResultTimeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline struct {
	// The collection of advertisement detection timelines.
	Ad []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd `json:"Ad,omitempty" xml:"Ad,omitempty" type:"Repeated"`
	// The collection of undesirable scene timelines.
	Live []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive `json:"Live,omitempty" xml:"Live,omitempty" type:"Repeated"`
	// The collection of logo detection timelines.
	Logo []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
	// The collection of pornography detection timelines.
	Porn []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn `json:"Porn,omitempty" xml:"Porn,omitempty" type:"Repeated"`
	// The collection of terrorism detection timelines.
	Terrorism []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism `json:"Terrorism,omitempty" xml:"Terrorism,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) GetAd() []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd {
	return s.Ad
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) GetLive() []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive {
	return s.Live
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) GetLogo() []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo {
	return s.Logo
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) GetPorn() []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn {
	return s.Porn
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) GetTerrorism() []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism {
	return s.Terrorism
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetAd(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Ad = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetLive(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Live = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetLogo(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Logo = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetPorn(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Porn = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetTerrorism(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Terrorism = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) Validate() error {
	if s.Ad != nil {
		for _, item := range s.Ad {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Live != nil {
		for _, item := range s.Live {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Logo != nil {
		for _, item := range s.Logo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Porn != nil {
		for _, item := range s.Porn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Terrorism != nil {
		for _, item := range s.Terrorism {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd struct {
	// The classification of the advertisement detection result. Valid values:
	//
	// - **normal**: normal.
	//
	// - **ad**: other advertisements.
	//
	// - **politics**: text contains political content.
	//
	// - **porn**: text contains pornographic content.
	//
	// - **abuse**: text contains abusive content.
	//
	// - **terrorism**: text contains terrorism-related content.
	//
	// - **contraband**: text contains prohibited content.
	//
	// - **spam**: text contains other spam content.
	//
	// - **npx**: psoriasis advertisements.
	//
	// - **qrcode**: contains a QR code.
	//
	// - **programCode**: contains a mini program code.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The hit score of the video screenshot for the advertisement detection result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video screenshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 10
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive struct {
	// The classification of the undesirable content detection result. Valid values:
	//
	// - **normal**: normal.
	//
	// - **meaningless**: no content in the image (such as a black screen or white screen).
	//
	// - **PIP**: Picture-in-Picture (PiP).
	//
	// - **smoking**: smoking.
	//
	// - **drivelive**: in-car live streaming.
	//
	// example:
	//
	// smoking
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The hit score of the video screenshot for the undesirable content detection result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video screenshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 12
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo struct {
	// The classification of the logo detection result. Valid values:
	//
	// - **normal**: normal.
	//
	// - **TV**: contains a regulated logo.
	//
	// - **trademark**: contains a trademark.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The hit score of the video screenshot for the logo detection result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video screenshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 13
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn struct {
	// The classification of the pornography detection result. Valid values:
	//
	// - **porn**: pornographic content.
	//
	// - **sexy**: sexually suggestive content.
	//
	// - **normal**: normal content.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The hit score of the video screenshot for the pornography detection result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100.00
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video screenshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism struct {
	// The terrorism and political content detection result. Valid values:
	//
	// - **normal**: normal.
	//
	// - **bloody**: bloody content.
	//
	// - **explosion**: explosions and smoke.
	//
	// - **outfit**: special attire.
	//
	// - **logo**: special logos.
	//
	// - **weapon**: weapons.
	//
	// - **politics**: political content.
	//
	// - **violence**: fighting.
	//
	// - **crowd**: crowd gathering.
	//
	// - **parade**: parades.
	//
	// - **carcrash**: car accident scenes.
	//
	// - **flag**: flags.
	//
	// - **location**: landmarks.
	//
	// - **others**: others.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The hit score of the video screenshot for the terrorism and political content detection result. Value range: `[0-100]`, with a precision of 10 decimal places. The hit result indicates the probability of the corresponding classification label. A higher value indicates higher accuracy.
	//
	// example:
	//
	// 100.00
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) Validate() error {
	return dara.Validate(s)
}
