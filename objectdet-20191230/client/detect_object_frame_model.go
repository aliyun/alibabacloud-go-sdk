// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectObjectFrame interface {
	dara.Model
	String() string
	GoString() string
	SetElements(v []*DetectObjectElement) *DetectObjectFrame
	GetElements() []*DetectObjectElement
	SetTime(v int64) *DetectObjectFrame
	GetTime() *int64
}

type DetectObjectFrame struct {
	Elements []*DetectObjectElement `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Time     *int64                 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DetectObjectFrame) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectFrame) GoString() string {
	return s.String()
}

func (s *DetectObjectFrame) GetElements() []*DetectObjectElement {
	return s.Elements
}

func (s *DetectObjectFrame) GetTime() *int64 {
	return s.Time
}

func (s *DetectObjectFrame) SetElements(v []*DetectObjectElement) *DetectObjectFrame {
	s.Elements = v
	return s
}

func (s *DetectObjectFrame) SetTime(v int64) *DetectObjectFrame {
	s.Time = &v
	return s
}

func (s *DetectObjectFrame) Validate() error {
	return dara.Validate(s)
}
