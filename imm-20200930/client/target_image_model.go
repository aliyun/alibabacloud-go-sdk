// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetImage interface {
	dara.Model
	String() string
	GoString() string
	SetAnimations(v []*TargetImageAnimations) *TargetImage
	GetAnimations() []*TargetImageAnimations
	SetSnapshots(v []*TargetImageSnapshots) *TargetImage
	GetSnapshots() []*TargetImageSnapshots
	SetSprites(v []*TargetImageSprites) *TargetImage
	GetSprites() []*TargetImageSprites
}

type TargetImage struct {
	// Animations
	Animations []*TargetImageAnimations `json:"Animations,omitempty" xml:"Animations,omitempty" type:"Repeated"`
	// Snapshots
	Snapshots []*TargetImageSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// Sprites
	Sprites []*TargetImageSprites `json:"Sprites,omitempty" xml:"Sprites,omitempty" type:"Repeated"`
}

func (s TargetImage) String() string {
	return dara.Prettify(s)
}

func (s TargetImage) GoString() string {
	return s.String()
}

func (s *TargetImage) GetAnimations() []*TargetImageAnimations {
	return s.Animations
}

func (s *TargetImage) GetSnapshots() []*TargetImageSnapshots {
	return s.Snapshots
}

func (s *TargetImage) GetSprites() []*TargetImageSprites {
	return s.Sprites
}

func (s *TargetImage) SetAnimations(v []*TargetImageAnimations) *TargetImage {
	s.Animations = v
	return s
}

func (s *TargetImage) SetSnapshots(v []*TargetImageSnapshots) *TargetImage {
	s.Snapshots = v
	return s
}

func (s *TargetImage) SetSprites(v []*TargetImageSprites) *TargetImage {
	s.Sprites = v
	return s
}

func (s *TargetImage) Validate() error {
	if s.Animations != nil {
		for _, item := range s.Animations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Sprites != nil {
		for _, item := range s.Sprites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TargetImageAnimations struct {
	// Format
	//
	// This parameter is required.
	//
	// example:
	//
	// gif
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Animation frame rate, in frames per second
	//
	// example:
	//
	// 25
	FrameRate *float64 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// Height
	//
	// example:
	//
	// 960
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Time interval between animation frames, in seconds
	//
	// example:
	//
	// 0.5
	Interval *float64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Total number of animation frames to generate
	//
	// example:
	//
	// 100
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// Scaling method
	//
	// example:
	//
	// crop
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	// StartTime
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The URI of the animation.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/animations
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// Output width, in pixels
	//
	// example:
	//
	// 1280
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s TargetImageAnimations) String() string {
	return dara.Prettify(s)
}

func (s TargetImageAnimations) GoString() string {
	return s.String()
}

func (s *TargetImageAnimations) GetFormat() *string {
	return s.Format
}

func (s *TargetImageAnimations) GetFrameRate() *float64 {
	return s.FrameRate
}

func (s *TargetImageAnimations) GetHeight() *float64 {
	return s.Height
}

func (s *TargetImageAnimations) GetInterval() *float64 {
	return s.Interval
}

func (s *TargetImageAnimations) GetNumber() *int32 {
	return s.Number
}

func (s *TargetImageAnimations) GetScaleType() *string {
	return s.ScaleType
}

func (s *TargetImageAnimations) GetStartTime() *float64 {
	return s.StartTime
}

func (s *TargetImageAnimations) GetURI() *string {
	return s.URI
}

func (s *TargetImageAnimations) GetWidth() *float64 {
	return s.Width
}

func (s *TargetImageAnimations) SetFormat(v string) *TargetImageAnimations {
	s.Format = &v
	return s
}

func (s *TargetImageAnimations) SetFrameRate(v float64) *TargetImageAnimations {
	s.FrameRate = &v
	return s
}

func (s *TargetImageAnimations) SetHeight(v float64) *TargetImageAnimations {
	s.Height = &v
	return s
}

func (s *TargetImageAnimations) SetInterval(v float64) *TargetImageAnimations {
	s.Interval = &v
	return s
}

func (s *TargetImageAnimations) SetNumber(v int32) *TargetImageAnimations {
	s.Number = &v
	return s
}

func (s *TargetImageAnimations) SetScaleType(v string) *TargetImageAnimations {
	s.ScaleType = &v
	return s
}

func (s *TargetImageAnimations) SetStartTime(v float64) *TargetImageAnimations {
	s.StartTime = &v
	return s
}

func (s *TargetImageAnimations) SetURI(v string) *TargetImageAnimations {
	s.URI = &v
	return s
}

func (s *TargetImageAnimations) SetWidth(v float64) *TargetImageAnimations {
	s.Width = &v
	return s
}

func (s *TargetImageAnimations) Validate() error {
	return dara.Validate(s)
}

type TargetImageSnapshots struct {
	// Image format
	//
	// This parameter is required.
	//
	// example:
	//
	// jpg
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Output height, in pixels
	//
	// example:
	//
	// 960
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Time interval between snapshots, in seconds
	//
	// example:
	//
	// 0.5
	Interval *float64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// interval
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The sequence number of the snapshot.
	//
	// example:
	//
	// 10
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// Scaling method
	//
	// example:
	//
	// crop
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	// The start time of the snapshot.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// OSS URI where snapshots are stored
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/snapshots
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The width of the snapshot.
	//
	// example:
	//
	// 1280
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s TargetImageSnapshots) String() string {
	return dara.Prettify(s)
}

func (s TargetImageSnapshots) GoString() string {
	return s.String()
}

func (s *TargetImageSnapshots) GetFormat() *string {
	return s.Format
}

func (s *TargetImageSnapshots) GetHeight() *float64 {
	return s.Height
}

func (s *TargetImageSnapshots) GetInterval() *float64 {
	return s.Interval
}

func (s *TargetImageSnapshots) GetMode() *string {
	return s.Mode
}

func (s *TargetImageSnapshots) GetNumber() *int32 {
	return s.Number
}

func (s *TargetImageSnapshots) GetScaleType() *string {
	return s.ScaleType
}

func (s *TargetImageSnapshots) GetStartTime() *float64 {
	return s.StartTime
}

func (s *TargetImageSnapshots) GetThreshold() *int32 {
	return s.Threshold
}

func (s *TargetImageSnapshots) GetURI() *string {
	return s.URI
}

func (s *TargetImageSnapshots) GetWidth() *float64 {
	return s.Width
}

func (s *TargetImageSnapshots) SetFormat(v string) *TargetImageSnapshots {
	s.Format = &v
	return s
}

func (s *TargetImageSnapshots) SetHeight(v float64) *TargetImageSnapshots {
	s.Height = &v
	return s
}

func (s *TargetImageSnapshots) SetInterval(v float64) *TargetImageSnapshots {
	s.Interval = &v
	return s
}

func (s *TargetImageSnapshots) SetMode(v string) *TargetImageSnapshots {
	s.Mode = &v
	return s
}

func (s *TargetImageSnapshots) SetNumber(v int32) *TargetImageSnapshots {
	s.Number = &v
	return s
}

func (s *TargetImageSnapshots) SetScaleType(v string) *TargetImageSnapshots {
	s.ScaleType = &v
	return s
}

func (s *TargetImageSnapshots) SetStartTime(v float64) *TargetImageSnapshots {
	s.StartTime = &v
	return s
}

func (s *TargetImageSnapshots) SetThreshold(v int32) *TargetImageSnapshots {
	s.Threshold = &v
	return s
}

func (s *TargetImageSnapshots) SetURI(v string) *TargetImageSnapshots {
	s.URI = &v
	return s
}

func (s *TargetImageSnapshots) SetWidth(v float64) *TargetImageSnapshots {
	s.Width = &v
	return s
}

func (s *TargetImageSnapshots) Validate() error {
	return dara.Validate(s)
}

type TargetImageSprites struct {
	// Image format
	//
	// This parameter is required.
	//
	// example:
	//
	// jpg
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Time interval between sprites, in seconds
	//
	// example:
	//
	// 1
	Interval *float64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Margin around the sprite grid, in pixels
	//
	// example:
	//
	// 2
	Margin *int32 `json:"Margin,omitempty" xml:"Margin,omitempty"`
	// example:
	//
	// interval
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Total number of sprites to generate
	//
	// example:
	//
	// 0
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// Padding between sprite tiles, in pixels
	//
	// example:
	//
	// 2
	Pad *int32 `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// Output height after scaling, in pixels
	//
	// example:
	//
	// 960
	ScaleHeight *float32 `json:"ScaleHeight,omitempty" xml:"ScaleHeight,omitempty"`
	// Scaling method
	//
	// example:
	//
	// crop
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	// Output width after scaling, in pixels
	//
	// example:
	//
	// 1280
	ScaleWidth *float32 `json:"ScaleWidth,omitempty" xml:"ScaleWidth,omitempty"`
	// StartTime
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Height of each sprite tile, in pixels
	//
	// example:
	//
	// 6
	TileHeight *int32 `json:"TileHeight,omitempty" xml:"TileHeight,omitempty"`
	// Width of each sprite tile, in pixels
	//
	// example:
	//
	// 6
	TileWidth *int32 `json:"TileWidth,omitempty" xml:"TileWidth,omitempty"`
	// URI
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/sprites
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s TargetImageSprites) String() string {
	return dara.Prettify(s)
}

func (s TargetImageSprites) GoString() string {
	return s.String()
}

func (s *TargetImageSprites) GetFormat() *string {
	return s.Format
}

func (s *TargetImageSprites) GetInterval() *float64 {
	return s.Interval
}

func (s *TargetImageSprites) GetMargin() *int32 {
	return s.Margin
}

func (s *TargetImageSprites) GetMode() *string {
	return s.Mode
}

func (s *TargetImageSprites) GetNumber() *int32 {
	return s.Number
}

func (s *TargetImageSprites) GetPad() *int32 {
	return s.Pad
}

func (s *TargetImageSprites) GetScaleHeight() *float32 {
	return s.ScaleHeight
}

func (s *TargetImageSprites) GetScaleType() *string {
	return s.ScaleType
}

func (s *TargetImageSprites) GetScaleWidth() *float32 {
	return s.ScaleWidth
}

func (s *TargetImageSprites) GetStartTime() *float64 {
	return s.StartTime
}

func (s *TargetImageSprites) GetThreshold() *int32 {
	return s.Threshold
}

func (s *TargetImageSprites) GetTileHeight() *int32 {
	return s.TileHeight
}

func (s *TargetImageSprites) GetTileWidth() *int32 {
	return s.TileWidth
}

func (s *TargetImageSprites) GetURI() *string {
	return s.URI
}

func (s *TargetImageSprites) SetFormat(v string) *TargetImageSprites {
	s.Format = &v
	return s
}

func (s *TargetImageSprites) SetInterval(v float64) *TargetImageSprites {
	s.Interval = &v
	return s
}

func (s *TargetImageSprites) SetMargin(v int32) *TargetImageSprites {
	s.Margin = &v
	return s
}

func (s *TargetImageSprites) SetMode(v string) *TargetImageSprites {
	s.Mode = &v
	return s
}

func (s *TargetImageSprites) SetNumber(v int32) *TargetImageSprites {
	s.Number = &v
	return s
}

func (s *TargetImageSprites) SetPad(v int32) *TargetImageSprites {
	s.Pad = &v
	return s
}

func (s *TargetImageSprites) SetScaleHeight(v float32) *TargetImageSprites {
	s.ScaleHeight = &v
	return s
}

func (s *TargetImageSprites) SetScaleType(v string) *TargetImageSprites {
	s.ScaleType = &v
	return s
}

func (s *TargetImageSprites) SetScaleWidth(v float32) *TargetImageSprites {
	s.ScaleWidth = &v
	return s
}

func (s *TargetImageSprites) SetStartTime(v float64) *TargetImageSprites {
	s.StartTime = &v
	return s
}

func (s *TargetImageSprites) SetThreshold(v int32) *TargetImageSprites {
	s.Threshold = &v
	return s
}

func (s *TargetImageSprites) SetTileHeight(v int32) *TargetImageSprites {
	s.TileHeight = &v
	return s
}

func (s *TargetImageSprites) SetTileWidth(v int32) *TargetImageSprites {
	s.TileWidth = &v
	return s
}

func (s *TargetImageSprites) SetURI(v string) *TargetImageSprites {
	s.URI = &v
	return s
}

func (s *TargetImageSprites) Validate() error {
	return dara.Validate(s)
}
