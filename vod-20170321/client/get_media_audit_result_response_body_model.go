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
	// The review results.
	MediaAuditResult *GetMediaAuditResultResponseBodyMediaAuditResult `json:"MediaAuditResult,omitempty" xml:"MediaAuditResult,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The content that violates the regulations. Separate multiple values with commas (,). Valid values:
	//
	// 	- **video**
	//
	// 	- **image-cover**
	//
	// 	- **text-title**
	//
	// example:
	//
	// video
	AbnormalModules *string `json:"AbnormalModules,omitempty" xml:"AbnormalModules,omitempty"`
	// The results of audio review.
	AudioResult []*GetMediaAuditResultResponseBodyMediaAuditResultAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	// The results of image review.
	ImageResult []*GetMediaAuditResultResponseBodyMediaAuditResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	// The category of the review result. Separate multiple values with commas (,). Valid values:
	//
	// 	- **porn**
	//
	// 	- **terrorism**
	//
	// 	- **normal**
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The results of text review.
	TextResult []*GetMediaAuditResultResponseBodyMediaAuditResultTextResult `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	// The results of video review.
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
	// The category of the review result.
	//
	// 	- **normal**
	//
	// 	- **spam**
	//
	// 	- **ad**
	//
	// 	- **politics**
	//
	// 	- **terrorism**
	//
	// 	- **abuse**
	//
	// 	- **porn**
	//
	// 	- **flood**
	//
	// 	- **contraband**
	//
	// 	- **meaningless**
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scenario. The value is **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score.
	//
	// example:
	//
	// 99.91
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
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
	// The category of the review result. Separate multiple values with commas (,). Valid values:
	//
	// 	- **porn**
	//
	// 	- **terrorism**
	//
	// 	- **normal**
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Details of image review results.
	Result []*GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The type of the image. The value is **cover**.
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
	// The category of the review result.
	//
	// Valid values if scene is **porn**:
	//
	// 	- **porn**
	//
	// 	- **sexy**
	//
	// 	- **normal**
	//
	// Valid values if scene is **terrorism**:
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
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scenario. Valid values:
	//
	// 	- **terrorism**
	//
	// 	- **porn**
	//
	// example:
	//
	// porn
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score of the image of the category that is indicated by Label.
	//
	// example:
	//
	// 100.00000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
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
	// The text content for review.
	//
	// example:
	//
	// hot line 123****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The category of the review result. Valid values:
	//
	// - **spam**
	//
	// - **ad**
	//
	// - **abuse**
	//
	// - **flood**
	//
	// - **contraband**
	//
	// - **meaningless**
	//
	// - **normal**
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scenario. The value is **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score of the image of the category that is indicated by Label.
	//
	// example:
	//
	// 100.00000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// - **block**
	//
	// - **review**
	//
	// - **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The type of the text. The value is **title**.
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
	// The results of ad review.
	AdResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult `json:"AdResult,omitempty" xml:"AdResult,omitempty" type:"Struct"`
	// The category of the review result. Separate multiple values with commas (,). Valid values:
	//
	// - **porn**
	//
	// - **terrorism**
	//
	// - **normal**
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The results of undesired content review.
	LiveResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult `json:"LiveResult,omitempty" xml:"LiveResult,omitempty" type:"Struct"`
	// The results of logo review.
	LogoResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult `json:"LogoResult,omitempty" xml:"LogoResult,omitempty" type:"Struct"`
	// The results of pornographic content review.
	PornResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult `json:"PornResult,omitempty" xml:"PornResult,omitempty" type:"Struct"`
	// The recommendation for review results. Valid values:
	//
	// - **block**
	//
	// - **review**
	//
	// - **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The results of terrorist content review.
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
	// The average score of the review results.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The statistics about tag frames.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// - **ad**
	//
	// - **normal**
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// - **block**
	//
	// - **review**
	//
	// - **pass**
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the image with the highest score of the category that is indicated by Label.
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
	// The number of frames.
	//
	// example:
	//
	// 12
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the review result. Valid values:
	//
	// - **ad**
	//
	// - **normal**
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
	// The category of the review result.
	//
	// - **ad**
	//
	// - **normal**
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the image of the category that is indicated by Label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position in the video. Unit: milliseconds.
	//
	// example:
	//
	// 10
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the image.
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

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult struct {
	// The average score of the review results.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The statistics about tag frames.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// - **live**: The content contains undesirable scenes.
	//
	// - **normal**: normal content.
	//
	// example:
	//
	// live
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// - **block**
	//
	// - **review**
	//
	// - **pass**
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the image with the highest score of the category that is indicated by Label.
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
	// The number of frames.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the review result. Valid values:
	//
	// - **live**: The content contains undesirable scenes.
	//
	// - **normal**: normal content.
	//
	// example:
	//
	// live
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
	// The category of the review result. Valid values:
	//
	// - **live**: The content contains undesirable scenes.
	//
	// - **normal**: normal content.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the image of the category that is indicated by Label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position in the video. Unit: milliseconds.
	//
	// example:
	//
	// 10
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the image.
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
	// The average score of the review results.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The statistics about tag frames.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// - **logo**
	//
	// - **normal**
	//
	// example:
	//
	// logo
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// - **block**
	//
	// - **review**
	//
	// - **pass**
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the image with the highest score of the category that is indicated by Label.
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
	// The number of frames.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the review result. Valid values:
	//
	// - **logo**
	//
	// - **normal**
	//
	// example:
	//
	// logo
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
	// The category of the review result.
	//
	// - **logo**
	//
	// - **normal**
	//
	// example:
	//
	// logo
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the image of the category that is indicated by Label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position in the video. Unit: milliseconds.
	//
	// example:
	//
	// 16
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the image.
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
	// The average score of the review results.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The statistics about tag frames.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// - **porn**
	//
	// - **sexy**
	//
	// - **normal**
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the image with the highest score of the category that is indicated by Label.
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
	// The number of frames.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the review result. Valid values:
	//
	// - **porn**
	//
	// - **sexy**
	//
	// - **normal**
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
	// The category of the review result. Valid values:
	//
	// - **porn**
	//
	// - **sexy**
	//
	// - **normal**
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the image of the category that is indicated by Label.
	//
	// example:
	//
	// 100.0000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the image.
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
	// The average score of the review results.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The statistics about tag frames.
	CounterList []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// - **normal**
	//
	// - **bloody**
	//
	// - **explosion**
	//
	// - **outfit**
	//
	// - **logo**
	//
	// - **weapon**
	//
	// - **politics**
	//
	// - **violence**
	//
	// - **crowd**
	//
	// - **parade**
	//
	// - **carcrash**
	//
	// - **flag**
	//
	// - **location**
	//
	// - **others**
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// - **block**
	//
	// - **review**
	//
	// - **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the image with the highest score of the category that is indicated by Label.
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
	// The number of frames.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the review result. Valid values:
	//
	// - **normal**
	//
	// - **bloody**
	//
	// - **explosion**
	//
	// - **outfit**
	//
	// - **logo**
	//
	// - **weapon**
	//
	// - **politics**
	//
	// - **violence**
	//
	// - **crowd**
	//
	// - **parade**
	//
	// - **carcrash**
	//
	// - **flag**
	//
	// - **location**
	//
	// - **others**
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
	// The category of the review result. Valid values:
	//
	// - **normal**
	//
	// - **bloody**
	//
	// - **explosion**
	//
	// - **outfit**
	//
	// - **logo**
	//
	// - **weapon**
	//
	// - **politics**
	//
	// - **violence**
	//
	// - **crowd**
	//
	// - **parade**
	//
	// - **carcrash**
	//
	// - **flag**
	//
	// - **location**
	//
	// - **others**
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the image of the category that is indicated by Label.
	//
	// example:
	//
	// 100.000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the image.
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
