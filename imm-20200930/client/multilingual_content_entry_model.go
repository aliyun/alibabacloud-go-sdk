// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultilingualContentEntry interface {
	dara.Model
	String() string
	GoString() string
	SetCaption(v string) *MultilingualContentEntry
	GetCaption() *string
	SetDescription(v string) *MultilingualContentEntry
	GetDescription() *string
}

type MultilingualContentEntry struct {
	// The multilingual brief description.
	//
	// example:
	//
	// No personnel activity at the office desk
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The multilingual detailed description.
	//
	// example:
	//
	// This is a close-up shot of an office desk setup. In the left foreground stands a tall, cylindrical, off-white insulated tumbler. A rectangular black mousepad occupies the center of the desk, holding a black backlit mechanical keyboard. Directly behind the keyboard sits a computer monitor with its screen illuminated, displaying the operating system\\"s application dock at the bottom. To the front right of the monitor stands a red metal beverage can, surrounded by a tangle of white data cables and a charging adapter. A small, silver, rectangular device (possibly a USB drive or an adapter) rests in the gap behind the left side of the keyboard, and a tiny pink decorative object is faintly visible on the desk surface. The scene is devoid of human activity; all objects remain motionless.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s MultilingualContentEntry) String() string {
	return dara.Prettify(s)
}

func (s MultilingualContentEntry) GoString() string {
	return s.String()
}

func (s *MultilingualContentEntry) GetCaption() *string {
	return s.Caption
}

func (s *MultilingualContentEntry) GetDescription() *string {
	return s.Description
}

func (s *MultilingualContentEntry) SetCaption(v string) *MultilingualContentEntry {
	s.Caption = &v
	return s
}

func (s *MultilingualContentEntry) SetDescription(v string) *MultilingualContentEntry {
	s.Description = &v
	return s
}

func (s *MultilingualContentEntry) Validate() error {
	return dara.Validate(s)
}
