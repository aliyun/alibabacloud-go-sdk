// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySentenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifySentenceResponseBody
	GetCode() *string
	SetData(v *VerifySentenceResponseBodyData) *VerifySentenceResponseBody
	GetData() *VerifySentenceResponseBodyData
	SetIncorrectWords(v int32) *VerifySentenceResponseBody
	GetIncorrectWords() *int32
	SetMessage(v string) *VerifySentenceResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifySentenceResponseBody
	GetRequestId() *string
	SetSourceRole(v int32) *VerifySentenceResponseBody
	GetSourceRole() *int32
	SetSuccess(v bool) *VerifySentenceResponseBody
	GetSuccess() *bool
	SetTargetRole(v int32) *VerifySentenceResponseBody
	GetTargetRole() *int32
}

type VerifySentenceResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *VerifySentenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2
	IncorrectWords *int32 `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	SourceRole *int32 `json:"SourceRole,omitempty" xml:"SourceRole,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TargetRole *int32 `json:"TargetRole,omitempty" xml:"TargetRole,omitempty"`
}

func (s VerifySentenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifySentenceResponseBody) GetData() *VerifySentenceResponseBodyData {
	return s.Data
}

func (s *VerifySentenceResponseBody) GetIncorrectWords() *int32 {
	return s.IncorrectWords
}

func (s *VerifySentenceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifySentenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifySentenceResponseBody) GetSourceRole() *int32 {
	return s.SourceRole
}

func (s *VerifySentenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifySentenceResponseBody) GetTargetRole() *int32 {
	return s.TargetRole
}

func (s *VerifySentenceResponseBody) SetCode(v string) *VerifySentenceResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySentenceResponseBody) SetData(v *VerifySentenceResponseBodyData) *VerifySentenceResponseBody {
	s.Data = v
	return s
}

func (s *VerifySentenceResponseBody) SetIncorrectWords(v int32) *VerifySentenceResponseBody {
	s.IncorrectWords = &v
	return s
}

func (s *VerifySentenceResponseBody) SetMessage(v string) *VerifySentenceResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySentenceResponseBody) SetRequestId(v string) *VerifySentenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySentenceResponseBody) SetSourceRole(v int32) *VerifySentenceResponseBody {
	s.SourceRole = &v
	return s
}

func (s *VerifySentenceResponseBody) SetSuccess(v bool) *VerifySentenceResponseBody {
	s.Success = &v
	return s
}

func (s *VerifySentenceResponseBody) SetTargetRole(v int32) *VerifySentenceResponseBody {
	s.TargetRole = &v
	return s
}

func (s *VerifySentenceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifySentenceResponseBodyData struct {
	Delta []*VerifySentenceResponseBodyDataDelta `json:"Delta,omitempty" xml:"Delta,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyData) GetDelta() []*VerifySentenceResponseBodyDataDelta {
	return s.Delta
}

func (s *VerifySentenceResponseBodyData) SetDelta(v []*VerifySentenceResponseBodyDataDelta) *VerifySentenceResponseBodyData {
	s.Delta = v
	return s
}

func (s *VerifySentenceResponseBodyData) Validate() error {
	if s.Delta != nil {
		for _, item := range s.Delta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VerifySentenceResponseBodyDataDelta struct {
	Source *VerifySentenceResponseBodyDataDeltaSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Target *VerifySentenceResponseBodyDataDeltaTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// CHANGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s VerifySentenceResponseBodyDataDelta) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDelta) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDelta) GetSource() *VerifySentenceResponseBodyDataDeltaSource {
	return s.Source
}

func (s *VerifySentenceResponseBodyDataDelta) GetTarget() *VerifySentenceResponseBodyDataDeltaTarget {
	return s.Target
}

func (s *VerifySentenceResponseBodyDataDelta) GetType() *string {
	return s.Type
}

func (s *VerifySentenceResponseBodyDataDelta) SetSource(v *VerifySentenceResponseBodyDataDeltaSource) *VerifySentenceResponseBodyDataDelta {
	s.Source = v
	return s
}

func (s *VerifySentenceResponseBodyDataDelta) SetTarget(v *VerifySentenceResponseBodyDataDeltaTarget) *VerifySentenceResponseBodyDataDelta {
	s.Target = v
	return s
}

func (s *VerifySentenceResponseBodyDataDelta) SetType(v string) *VerifySentenceResponseBodyDataDelta {
	s.Type = &v
	return s
}

func (s *VerifySentenceResponseBodyDataDelta) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifySentenceResponseBodyDataDeltaSource struct {
	Line *VerifySentenceResponseBodyDataDeltaSourceLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s VerifySentenceResponseBodyDataDeltaSource) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaSource) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaSource) GetLine() *VerifySentenceResponseBodyDataDeltaSourceLine {
	return s.Line
}

func (s *VerifySentenceResponseBodyDataDeltaSource) GetPosition() *int32 {
	return s.Position
}

func (s *VerifySentenceResponseBodyDataDeltaSource) SetLine(v *VerifySentenceResponseBodyDataDeltaSourceLine) *VerifySentenceResponseBodyDataDeltaSource {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaSource) SetPosition(v int32) *VerifySentenceResponseBodyDataDeltaSource {
	s.Position = &v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaSource) Validate() error {
	if s.Line != nil {
		if err := s.Line.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifySentenceResponseBodyDataDeltaSourceLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyDataDeltaSourceLine) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaSourceLine) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaSourceLine) GetLine() []*string {
	return s.Line
}

func (s *VerifySentenceResponseBodyDataDeltaSourceLine) SetLine(v []*string) *VerifySentenceResponseBodyDataDeltaSourceLine {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaSourceLine) Validate() error {
	return dara.Validate(s)
}

type VerifySentenceResponseBodyDataDeltaTarget struct {
	Line *VerifySentenceResponseBodyDataDeltaTargetLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s VerifySentenceResponseBodyDataDeltaTarget) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaTarget) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) GetLine() *VerifySentenceResponseBodyDataDeltaTargetLine {
	return s.Line
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) GetPosition() *int32 {
	return s.Position
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) SetLine(v *VerifySentenceResponseBodyDataDeltaTargetLine) *VerifySentenceResponseBodyDataDeltaTarget {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) SetPosition(v int32) *VerifySentenceResponseBodyDataDeltaTarget {
	s.Position = &v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) Validate() error {
	if s.Line != nil {
		if err := s.Line.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifySentenceResponseBodyDataDeltaTargetLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyDataDeltaTargetLine) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaTargetLine) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaTargetLine) GetLine() []*string {
	return s.Line
}

func (s *VerifySentenceResponseBodyDataDeltaTargetLine) SetLine(v []*string) *VerifySentenceResponseBodyDataDeltaTargetLine {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaTargetLine) Validate() error {
	return dara.Validate(s)
}
