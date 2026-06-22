// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsight interface {
	dara.Model
	String() string
	GoString() string
	SetCaption(v string) *VideoInsight
	GetCaption() *string
	SetDescription(v string) *VideoInsight
	GetDescription() *string
}

type VideoInsight struct {
	// Video summary.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 视频中展示了两个不同场景：一个是静止的白色盘子、黑色瓶子和透明玻璃杯，另一个是手拿着标有“YEZOLU”的洗发水瓶在浴室中缓慢上移。
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The description of the video file.
	//
	// >  Not supported.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 无。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s VideoInsight) String() string {
	return dara.Prettify(s)
}

func (s VideoInsight) GoString() string {
	return s.String()
}

func (s *VideoInsight) GetCaption() *string {
	return s.Caption
}

func (s *VideoInsight) GetDescription() *string {
	return s.Description
}

func (s *VideoInsight) SetCaption(v string) *VideoInsight {
	s.Caption = &v
	return s
}

func (s *VideoInsight) SetDescription(v string) *VideoInsight {
	s.Description = &v
	return s
}

func (s *VideoInsight) Validate() error {
	return dara.Validate(s)
}
