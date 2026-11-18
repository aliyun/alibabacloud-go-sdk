// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaAuditResult(v *GetMediaAuditResultResponseBodyMediaAuditResult) *GetMediaAuditResultResponseBody
	GetMediaAuditResult() *GetMediaAuditResultResponseBodyMediaAuditResult
	SetRequestId(v string) *GetMediaAuditResultResponseBody
	GetRequestId() *string
}

type GetMediaAuditResultResponseBody struct {
	// The moderation result.
	MediaAuditResult *GetMediaAuditResultResponseBodyMediaAuditResult `json:"MediaAuditResult,omitempty" xml:"MediaAuditResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CB7D7232-1AB2-40FE-B8D5-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBody) GetMediaAuditResult() *GetMediaAuditResultResponseBodyMediaAuditResult {
	return s.MediaAuditResult
}

func (s *GetMediaAuditResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaAuditResultResponseBody) SetMediaAuditResult(v *GetMediaAuditResultResponseBodyMediaAuditResult) *GetMediaAuditResultResponseBody {
	s.MediaAuditResult = v
	return s
}

func (s *GetMediaAuditResultResponseBody) SetRequestId(v string) *GetMediaAuditResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaAuditResultResponseBody) Validate() error {
	if s.MediaAuditResult != nil {
		if err := s.MediaAuditResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResult struct {
	// The moderated content that contains violations. Multiple values are separated by commas (,). Valid values:
	//
	// - **video**: Video.
	//
	// - **image-cover**: Thumbnail.
	//
	// - **text-title**: Title.
	//
	// example:
	//
	// video
	AbnormalModules *string `json:"AbnormalModules,omitempty" xml:"AbnormalModules,omitempty"`
	// The audio moderation result.
	AudioResult []*GetMediaAuditResultResponseBodyMediaAuditResultAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	// The image moderation result.
	ImageResult []*GetMediaAuditResultResponseBodyMediaAuditResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	// The category of the moderation result. Multiple values are separated by commas (,). Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **terrorism**: Terrorist and political content.
	//
	// - **ad**: Image and text violations.
	//
	// - **live**: Inappropriate scenes.
	//
	// - **logo**: Image logos.
	//
	// - **audio**: Audio anti-spam.
	//
	// - **normal**: Normal content.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The text moderation result.
	TextResult []*GetMediaAuditResultResponseBodyMediaAuditResultTextResult `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	// The video moderation result.
	VideoResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Struct"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) GetAbnormalModules() *string {
	return s.AbnormalModules
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) GetAudioResult() []*GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	return s.AudioResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) GetImageResult() []*GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	return s.ImageResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) GetTextResult() []*GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	return s.TextResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) GetVideoResult() *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	return s.VideoResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetAbnormalModules(v string) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.AbnormalModules = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetAudioResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.AudioResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetImageResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultImageResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.ImageResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetTextResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultTextResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.TextResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetVideoResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.VideoResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) Validate() error {
	if s.AudioResult != nil {
		for _, item := range s.AudioResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageResult != nil {
		for _, item := range s.ImageResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextResult != nil {
		for _, item := range s.TextResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoResult != nil {
		if err := s.VideoResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultAudioResult struct {
	// The category of the moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **spam**: Spam.
	//
	// - **ad**: Ads.
	//
	// - **politics**: Political content.
	//
	// - **terrorism**: Terrorist content.
	//
	// - **abuse**: Abuse.
	//
	// - **porn**: Pornographic content.
	//
	// - **flood**: Flooding.
	//
	// - **contraband**: Contraband.
	//
	// - **meaningless**: Meaningless content.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The moderation scenario. Fixed value: **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The confidence score for the matched label in the audio. Valid values: `[0, 100]`. The score indicates the probability that the audio belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 99.91
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) GetScene() *string {
	return s.Scene
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetScene(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Scene = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultImageResult struct {
	// The category of the image moderation result. Multiple values are separated by commas (,). Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **terrorism**: Terrorist and political content.
	//
	// - **ad**: Image and text violations.
	//
	// - **live**: Inappropriate scenes.
	//
	// - **logo**: Image logos.
	//
	// - **normal**: Normal content.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The details of the image moderation result.
	Result []*GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The image category. Valid value: **cover*	- (thumbnail).
	//
	// example:
	//
	// cover
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the image.
	//
	// example:
	//
	// http://www.test.com/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) GetResult() []*GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	return s.Result
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) GetType() *string {
	return s.Type
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Result = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetType(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Type = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult struct {
	// The category of the moderation result.
	//
	// When the scene is **porn**, valid values are:
	//
	// - **porn**: Pornographic content.
	//
	// - **sexy**: Sexy content.
	//
	// - **normal**: Normal content.
	//
	// When the scene is **terrorism**, valid values are:
	//
	// - **normal**: Normal content.
	//
	// - **bloody**: Bloody content.
	//
	// - **explosion**: Explosions and smoke.
	//
	// - **outfit**: Special outfits.
	//
	// - **logo**: Special logos.
	//
	// - **weapon**: Weapons.
	//
	// - **politics**: Political content.
	//
	// - **violence**: Violence.
	//
	// - **crowd**: Crowds.
	//
	// - **parade**: Parades.
	//
	// - **carcrash**: Car crashes.
	//
	// - **flag**: Flags.
	//
	// - **location**: Landmarks.
	//
	// - **others**: Others.
	//
	// When the scene is **ad**, valid values are:
	//
	// - **normal**: Normal content.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Intrusive ads.
	//
	// - **qrcode**: QR codes.
	//
	// - **programCode**: Mini program codes.
	//
	// When the scene is **live**, valid values are:
	//
	// - **normal**: Normal content.
	//
	// - **meaningless**: No content in the image, such as a black or white screen.
	//
	// - **PIP**: Picture-in-Picture (PiP).
	//
	// - **smoking**: Smoking.
	//
	// - **drivelive**: Live streaming in a car.
	//
	// When the scene is **logo**, valid values are:
	//
	// - **normal**: Normal content.
	//
	// - **TV**: Regulated logos.
	//
	// - **trademark**: Trademarks.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The moderation scenario. Valid values:
	//
	// - **porn**: Pornography detection.
	//
	// - **terrorism**: Terrorist and political content detection.
	//
	// - **ad**: Image and text violation detection.
	//
	// - **live**: Inappropriate scene detection.
	//
	// - **logo**: Image logo detection.
	//
	// example:
	//
	// porn
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The confidence score for the matched label in the image. Valid values: `[0, 100]`. The score indicates the probability that the image belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100.00000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) GetScene() *string {
	return s.Scene
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetScene(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Scene = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultTextResult struct {
	// The moderated text content.
	//
	// example:
	//
	// Hotline 123****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The category of the moderation result. Valid values:
	//
	// - **spam**: Spam.
	//
	// - **ad**: Ads.
	//
	// - **abuse**: Abuse.
	//
	// - **flood**: Flooding.
	//
	// - **contraband**: Contraband.
	//
	// - **meaningless**: Meaningless content.
	//
	// - **normal**: Normal content.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The moderation scenario. Fixed value: **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The confidence score for the matched label in the text snapshot. Valid values: `[0, 100]`. The score indicates the probability that the text belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100.00000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The text categorization. Valid value: **title*	- (title).
	//
	// example:
	//
	// title
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultTextResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GetContent() *string {
	return s.Content
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GetScene() *string {
	return s.Scene
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GetType() *string {
	return s.Type
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetContent(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Content = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetScene(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Scene = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetType(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Type = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResult struct {
	// The ad moderation result.
	AdResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult `json:"AdResult,omitempty" xml:"AdResult,omitempty" type:"Struct"`
	// The GreenEnhanced moderation result.
	GreenEnhancedResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult `json:"GreenEnhancedResult,omitempty" xml:"GreenEnhancedResult,omitempty" type:"Struct"`
	// The category of the moderation result. Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **terrorism**: Terrorist and political content.
	//
	// - **ad**: Image and text violations.
	//
	// - **live**: Inappropriate scenes.
	//
	// - **logo**: Image logos.
	//
	// - **normal**: Normal content.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The inappropriate content moderation result.
	LiveResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult `json:"LiveResult,omitempty" xml:"LiveResult,omitempty" type:"Struct"`
	// The logo moderation result.
	LogoResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult `json:"LogoResult,omitempty" xml:"LogoResult,omitempty" type:"Struct"`
	// The pornography moderation result.
	PornResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult `json:"PornResult,omitempty" xml:"PornResult,omitempty" type:"Struct"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The moderation result for terrorism and political content.
	TerrorismResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult `json:"TerrorismResult,omitempty" xml:"TerrorismResult,omitempty" type:"Struct"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetAdResult() *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	return s.AdResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetGreenEnhancedResult() *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult {
	return s.GreenEnhancedResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetLiveResult() *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	return s.LiveResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetLogoResult() *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	return s.LogoResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetPornResult() *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	return s.PornResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GetTerrorismResult() *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	return s.TerrorismResult
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetAdResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.AdResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetGreenEnhancedResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.GreenEnhancedResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetLiveResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.LiveResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetLogoResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.LogoResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetPornResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.PornResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetTerrorismResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.TerrorismResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) Validate() error {
	if s.AdResult != nil {
		if err := s.AdResult.Validate(); err != nil {
			return err
		}
	}
	if s.GreenEnhancedResult != nil {
		if err := s.GreenEnhancedResult.Validate(); err != nil {
			return err
		}
	}
	if s.LiveResult != nil {
		if err := s.LiveResult.Validate(); err != nil {
			return err
		}
	}
	if s.LogoResult != nil {
		if err := s.LogoResult.Validate(); err != nil {
			return err
		}
	}
	if s.PornResult != nil {
		if err := s.PornResult.Validate(); err != nil {
			return err
		}
	}
	if s.TerrorismResult != nil {
		if err := s.TerrorismResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult struct {
	// The average confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The number of video snapshots for each category in the moderation result.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the ad moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Intrusive ads.
	//
	// - **qrcode**: QR codes.
	//
	// - **programCode**: Mini program codes.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the matched label.
	TopList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GetCounterList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList {
	return s.CounterList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GetTopList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	return s.TopList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.TopList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 12
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the ad moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Intrusive ads.
	//
	// - **qrcode**: QR codes.
	//
	// - **programCode**: Mini program codes.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList struct {
	// The category of the ad moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Intrusive ads.
	//
	// - **qrcode**: QR codes.
	//
	// - **programCode**: Mini program codes.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confidence score for the matched label in the video snapshot. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 10
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult struct {
	// The average confidence score of the matched frames. This parameter is not returned if no frames are matched.
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The aggregation of violation label counts, including Label for the Green label and Count for the number of matched frames for the label.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The union of matched native Green labels, separated by commas, such as pornographic_adultContent_tii. If no labels are matched, the value is normal.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum confidence score of the matched frames. This parameter is not returned if no frames are matched.
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The image moderation conclusion, which is mapped from frameResult.riskLevel: high is mapped to block, medium or low is mapped to review, and none is mapped to pass.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The details of the matched frames, sorted in descending order of confidence score.
	TopList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) GetCounterList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList {
	return s.CounterList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) GetTopList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList {
	return s.TopList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult {
	s.TopList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList struct {
	// The number of video snapshots for the corresponding label.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The moderation result label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList struct {
	// The moderation result category.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confidence score of the video snapshot that matches the label.
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The timestamp of the video snapshot in the video.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultGreenEnhancedResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult struct {
	// The average confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The number of video snapshots for each inappropriate content category in the moderation result.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **meaningless**: No content in the image, such as a black or white screen.
	//
	// - **PIP**: Picture-in-Picture (PiP).
	//
	// - **smoking**: Smoking.
	//
	// - **drivelive**: Live streaming in a car.
	//
	// example:
	//
	// smoking
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the matched label.
	TopList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GetCounterList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList {
	return s.CounterList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GetTopList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	return s.TopList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.TopList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **meaningless**: No content in the image, such as a black or white screen.
	//
	// - **PIP**: Picture-in-Picture (PiP).
	//
	// - **smoking**: Smoking.
	//
	// - **drivelive**: Live streaming in a car.
	//
	// example:
	//
	// smoking
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList struct {
	// The category of the moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **meaningless**: No content in the image, such as a black or white screen.
	//
	// - **PIP**: Picture-in-Picture (PiP).
	//
	// - **smoking**: Smoking.
	//
	// - **drivelive**: Live streaming in a car.
	//
	// example:
	//
	// smoking
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confidence score for the matched label in the video snapshot. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 10
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult struct {
	// The average confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The number of video snapshots for each category in the moderation result.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the logo moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **TV**: Regulated logos.
	//
	// - **trademark**: Trademarks.
	//
	// example:
	//
	// TV
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the matched label.
	TopList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GetCounterList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList {
	return s.CounterList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GetTopList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	return s.TopList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.TopList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the logo moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **TV**: Regulated logos.
	//
	// - **trademark**: Trademarks.
	//
	// example:
	//
	// TV
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList struct {
	// The category of the logo moderation result. Valid values:
	//
	// - **normal**: Normal content.
	//
	// - **TV**: Regulated logos.
	//
	// - **trademark**: Trademarks.
	//
	// example:
	//
	// TV
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confidence score for the matched label in the video snapshot. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 16
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult struct {
	// The average confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The number of video snapshots for each category in the moderation result.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The pornography moderation result. Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **sexy**: Sexy content.
	//
	// - **normal**: Normal content.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The pornography moderation suggestion. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the matched label.
	TopList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GetCounterList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList {
	return s.CounterList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GetTopList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	return s.TopList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.TopList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The pornography moderation result. Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **sexy**: Sexy content.
	//
	// - **normal**: Normal content.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList struct {
	// The pornography moderation result. Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **sexy**: Sexy content.
	//
	// - **normal**: Normal content.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confidence score for the matched label in the video snapshot. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100.0000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
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

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult struct {
	// The average confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The categories of terrorism and political content moderation results and the number of video snapshots for each category.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The moderation label for terrorism and political content. Valid values:
	//
	// - **normal**: The content is normal.
	//
	// - **bloody**: The content contains bloody scenes.
	//
	// - **explosion**: The content contains explosions and smoke.
	//
	// - **outfit**: The content contains special outfits.
	//
	// - **logo**: The content contains special logos.
	//
	// - **weapon**: The content contains weapons.
	//
	// - **politics**: The content contains political elements.
	//
	// - **violence**: The content contains violence.
	//
	// - **crowd**: The content contains crowd gatherings.
	//
	// - **parade**: The content contains parades.
	//
	// - **carcrash**: The content contains car crash scenes.
	//
	// - **flag**: The content contains flags.
	//
	// - **location**: The content contains landmarks.
	//
	// - **others**: The content contains other elements.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum confidence score for the matched label in the video snapshots. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The moderation suggestion for terrorism and political content. Valid values:
	//
	// - **block**: The content is blocked.
	//
	// - **review**: The content requires manual review.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the matched label.
	TopList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GetCounterList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList {
	return s.CounterList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GetTopList() []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	return s.TopList
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.TopList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList struct {
	// The number of video snapshots for the terrorism and political content moderation result.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The moderation label for terrorism and political content. Valid values:
	//
	// - **normal**: The content is normal.
	//
	// - **bloody**: The content contains bloody scenes.
	//
	// - **explosion**: The content contains explosions and smoke.
	//
	// - **outfit**: The content contains special outfits.
	//
	// - **logo**: The content contains special logos.
	//
	// - **weapon**: The content contains weapons.
	//
	// - **politics**: The content contains political elements.
	//
	// - **violence**: The content contains violence.
	//
	// - **crowd**: The content contains crowd gatherings.
	//
	// - **parade**: The content contains parades.
	//
	// - **carcrash**: The content contains car crash scenes.
	//
	// - **flag**: The content contains flags.
	//
	// - **location**: The content contains landmarks.
	//
	// - **others**: The content contains other elements.
	//
	// example:
	//
	// outfit
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList struct {
	// The moderation label for terrorism and political content. Valid values:
	//
	// - **normal**: The content is normal.
	//
	// - **bloody**: The content contains bloody scenes.
	//
	// - **explosion**: The content contains explosions and smoke.
	//
	// - **outfit**: The content contains special outfits.
	//
	// - **logo**: The content contains special logos.
	//
	// - **weapon**: The content contains weapons.
	//
	// - **politics**: The content contains political elements.
	//
	// - **violence**: The content contains violence.
	//
	// - **crowd**: The content contains crowd gatherings.
	//
	// - **parade**: The content contains parades.
	//
	// - **carcrash**: The content contains car crash scenes.
	//
	// - **flag**: The content contains flags.
	//
	// - **location**: The content contains landmarks.
	//
	// - **others**: The content contains other elements.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confidence score for the matched label in the video snapshot. Valid values: `[0, 100]`. The score is accurate to 10 decimal places. The score indicates the probability that the video snapshot belongs to the corresponding category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100.000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
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

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Url = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) Validate() error {
	return dara.Validate(s)
}
