// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *VideoGenerationRequestInput) *VideoGenerationRequest
	GetInput() *VideoGenerationRequestInput
	SetIntent(v *VideoGenerationRequestIntent) *VideoGenerationRequest
	GetIntent() *VideoGenerationRequestIntent
	SetOutput(v *VideoGenerationRequestOutput) *VideoGenerationRequest
	GetOutput() *VideoGenerationRequestOutput
}

type VideoGenerationRequest struct {
	// The product input.
	//
	// This parameter is required.
	Input *VideoGenerationRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The intent parameters. Currently unavailable.
	Intent *VideoGenerationRequestIntent `json:"Intent,omitempty" xml:"Intent,omitempty" type:"Struct"`
	// The output parameters.
	//
	// This parameter is required.
	Output *VideoGenerationRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
}

func (s VideoGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequest) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequest) GetInput() *VideoGenerationRequestInput {
	return s.Input
}

func (s *VideoGenerationRequest) GetIntent() *VideoGenerationRequestIntent {
	return s.Intent
}

func (s *VideoGenerationRequest) GetOutput() *VideoGenerationRequestOutput {
	return s.Output
}

func (s *VideoGenerationRequest) SetInput(v *VideoGenerationRequestInput) *VideoGenerationRequest {
	s.Input = v
	return s
}

func (s *VideoGenerationRequest) SetIntent(v *VideoGenerationRequestIntent) *VideoGenerationRequest {
	s.Intent = v
	return s
}

func (s *VideoGenerationRequest) SetOutput(v *VideoGenerationRequestOutput) *VideoGenerationRequest {
	s.Output = v
	return s
}

func (s *VideoGenerationRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Intent != nil {
		if err := s.Intent.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoGenerationRequestInput struct {
	// The asset binding list.
	AssetBindings []*VideoGenerationRequestInputAssetBindings `json:"AssetBindings,omitempty" xml:"AssetBindings,omitempty" type:"Repeated"`
	// The extended information.
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The list of product image URLs (1 to 6 images). The URLs must be publicly accessible.
	//
	// This parameter is required.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The product title. A maximum of the first 60 characters are used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026 New Slimming Women\\"s Summer Dress with Mid-Length Design, High-Quality Waist Definition for a Slender Look
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s VideoGenerationRequestInput) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequestInput) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequestInput) GetAssetBindings() []*VideoGenerationRequestInputAssetBindings {
	return s.AssetBindings
}

func (s *VideoGenerationRequestInput) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *VideoGenerationRequestInput) GetImages() []*string {
	return s.Images
}

func (s *VideoGenerationRequestInput) GetTitle() *string {
	return s.Title
}

func (s *VideoGenerationRequestInput) SetAssetBindings(v []*VideoGenerationRequestInputAssetBindings) *VideoGenerationRequestInput {
	s.AssetBindings = v
	return s
}

func (s *VideoGenerationRequestInput) SetExtra(v map[string]interface{}) *VideoGenerationRequestInput {
	s.Extra = v
	return s
}

func (s *VideoGenerationRequestInput) SetImages(v []*string) *VideoGenerationRequestInput {
	s.Images = v
	return s
}

func (s *VideoGenerationRequestInput) SetTitle(v string) *VideoGenerationRequestInput {
	s.Title = &v
	return s
}

func (s *VideoGenerationRequestInput) Validate() error {
	if s.AssetBindings != nil {
		for _, item := range s.AssetBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VideoGenerationRequestInputAssetBindings struct {
	// The asset index.
	//
	// example:
	//
	// 0
	AssetIndex *int32 `json:"AssetIndex,omitempty" xml:"AssetIndex,omitempty"`
	// The asset description.
	//
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The asset usage.
	//
	// example:
	//
	// -
	Slot *string `json:"Slot,omitempty" xml:"Slot,omitempty"`
}

func (s VideoGenerationRequestInputAssetBindings) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequestInputAssetBindings) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequestInputAssetBindings) GetAssetIndex() *int32 {
	return s.AssetIndex
}

func (s *VideoGenerationRequestInputAssetBindings) GetDescription() *string {
	return s.Description
}

func (s *VideoGenerationRequestInputAssetBindings) GetSlot() *string {
	return s.Slot
}

func (s *VideoGenerationRequestInputAssetBindings) SetAssetIndex(v int32) *VideoGenerationRequestInputAssetBindings {
	s.AssetIndex = &v
	return s
}

func (s *VideoGenerationRequestInputAssetBindings) SetDescription(v string) *VideoGenerationRequestInputAssetBindings {
	s.Description = &v
	return s
}

func (s *VideoGenerationRequestInputAssetBindings) SetSlot(v string) *VideoGenerationRequestInputAssetBindings {
	s.Slot = &v
	return s
}

func (s *VideoGenerationRequestInputAssetBindings) Validate() error {
	return dara.Validate(s)
}

type VideoGenerationRequestIntent struct {
	// The distribution channel.
	//
	// example:
	//
	// -
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The business goal. Valid values:
	//
	// camera_motion: Camera movement mode. Generates video based on fixed 360° camera movement logic.
	//
	// scripted_video: Scripted mode. Provides a script or prompt, and the system generates video based on the script.
	//
	// auto_video: Unscripted mode. No script is provided. The system automatically plans the script and then generates the video.
	//
	// example:
	//
	// auto_video
	Goal *string `json:"Goal,omitempty" xml:"Goal,omitempty"`
	// The script or prompt.
	//
	// example:
	//
	// 【2. Shot Breakdown】
	//
	// [0-1s]【Eye-level close-up, static camera】Summer commute trousers folded on a clean tabletop; one corner is quickly unfolded to reveal the solid-color fabric and crisp crease line. The frame focuses on the trouser leg\\"s surface smoothness and sharp appearance. [BGM/SFX: Upbeat, crisp commute-vibe background music begins]
	//
	// [1-2s]【Low-angle mid-shot, static camera】The full pair of trousers hangs freely in the air with a slight swing, showcasing the natural drape of the lightweight fabric and the complete straight-leg silhouette, then returns to a still, hanging state.
	//
	// [2-3s]【Eye-level mid-shot, static camera】A commute-styled model stands in a fitting area, holding the folded trousers at waist level with both hands, performing one up-and-down sizing gesture before holding them steady, displaying the overall proportions of the trousers in their folded state and the expected fit. [[Voiceover]: Who says wearing trousers in summer has to be stuffy? Most likely you just haven\\"t picked the right pair.]
	//
	// [3-5s]【Eye-level full shot, slight pullback】The commute-styled model, now wearing the trousers paired with a clean commute top, takes two steps forward in a modern office building corridor, dynamically showcasing the front straight-leg silhouette, trouser leg lines, and commute outfit coordination.
	//
	// [5-7s]【Low-angle mid-shot, static camera】The commute-styled model shifts to an angled side stance, one hand in the front pocket and the other in the back pocket, displaying the cut from waist and hip down to the knee, the solid-color surface, and the clean lateral lines.
	//
	// [7-9s]【Eye-level mid-shot, static camera】The commute-styled model lifts one foot onto a low step, one hand brushes down the lower-leg trouser fabric and lightly lifts the hem, briefly exposing the ankle, showcasing the hem edge, side-seam details, and the wearing state during movement. [[Voiceover]: These ice-skin trousers are lightweight and breathable, solid-color straight-leg — looking sharp and at ease even when walking around at work.]
	//
	// [9-11s]【Overhead close-up, static camera】The lens focuses closely on the waistband area; the commute-styled model presses both hands along the waistband contour and smooths it, then pauses to display the seams and actual wearing state, clearly presenting the waistband shape and structural details.
	//
	// [11-13s]【Eye-level close-up, static camera】The commute-styled model gently lifts the fabric on both sides of the thigh with both hands, then raises one knee with a slight lateral turn before returning to a natural stance, demonstrating the actual range of motion during knee-lift and side-turn, as well as how the straight-leg trouser recovers after the foot lands.
	//
	// [13-15s]【Eye-level close-up, static camera】The commute-styled model enters the frame from the rear side, places one hand in the back pocket, then lightly traces along the back pocket edge and rear waist seam, showcasing the back pocket contour, pocket opening edge, and rear waist seam details up close. [[Voiceover]: If you\\"re always on the move, there\\"s room for knee lifts and side turns — commuting just got easier.]
	//
	// [15-17s]【Eye-level mid-shot, static camera】The commute-styled model faces away to display the rear silhouette of the trousers, hands naturally in pockets with a slight weight shift, then turns slightly to an angled side view so both the rear and lateral trouser lines are visible, fully presenting the rear cut and the overall line from hip and thigh down to the hem.
	//
	// [17-19s]【Eye-level full shot, smooth follow】The commute-styled model first walks sideways to the right, then transitions into a forward stride to complete the full presentation, finishing in the modern office building corridor in full commute attire, clearly showing the front straight-leg silhouette, clean lines, and overall fit. [[Voiceover]: If you like this fit, go check out the details.]
	//
	// 【3. Negative/Constraint Instructions】
	//
	// The entire video centers on the summer commute trousers; no unrelated products or multi-person interactions are allowed.
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s VideoGenerationRequestIntent) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequestIntent) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequestIntent) GetChannel() *string {
	return s.Channel
}

func (s *VideoGenerationRequestIntent) GetGoal() *string {
	return s.Goal
}

func (s *VideoGenerationRequestIntent) GetScript() *string {
	return s.Script
}

func (s *VideoGenerationRequestIntent) SetChannel(v string) *VideoGenerationRequestIntent {
	s.Channel = &v
	return s
}

func (s *VideoGenerationRequestIntent) SetGoal(v string) *VideoGenerationRequestIntent {
	s.Goal = &v
	return s
}

func (s *VideoGenerationRequestIntent) SetScript(v string) *VideoGenerationRequestIntent {
	s.Script = &v
	return s
}

func (s *VideoGenerationRequestIntent) Validate() error {
	return dara.Validate(s)
}

type VideoGenerationRequestOutput struct {
	// The video duration in seconds. Currently supports integers between 5 and 15. More options will be available in the future.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The output resolution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1080p
	Quality *string `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The video aspect ratio.
	//
	// example:
	//
	// 9:16
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s VideoGenerationRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequestOutput) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequestOutput) GetDuration() *int64 {
	return s.Duration
}

func (s *VideoGenerationRequestOutput) GetQuality() *string {
	return s.Quality
}

func (s *VideoGenerationRequestOutput) GetRatio() *string {
	return s.Ratio
}

func (s *VideoGenerationRequestOutput) SetDuration(v int64) *VideoGenerationRequestOutput {
	s.Duration = &v
	return s
}

func (s *VideoGenerationRequestOutput) SetQuality(v string) *VideoGenerationRequestOutput {
	s.Quality = &v
	return s
}

func (s *VideoGenerationRequestOutput) SetRatio(v string) *VideoGenerationRequestOutput {
	s.Ratio = &v
	return s
}

func (s *VideoGenerationRequestOutput) Validate() error {
	return dara.Validate(s)
}
