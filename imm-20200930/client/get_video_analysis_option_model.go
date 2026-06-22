// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoAnalysisOption interface {
	dara.Model
	String() string
	GoString() string
	SetChapterSummary(v bool) *GetVideoAnalysisOption
	GetChapterSummary() *bool
	SetKeyword(v bool) *GetVideoAnalysisOption
	GetKeyword() *bool
	SetPPT(v bool) *GetVideoAnalysisOption
	GetPPT() *bool
	SetQuestion(v bool) *GetVideoAnalysisOption
	GetQuestion() *bool
	SetSummary(v bool) *GetVideoAnalysisOption
	GetSummary() *bool
	SetTranscript(v bool) *GetVideoAnalysisOption
	GetTranscript() *bool
	SetTranscriptChapterSummary(v bool) *GetVideoAnalysisOption
	GetTranscriptChapterSummary() *bool
	SetTranscriptSummary(v bool) *GetVideoAnalysisOption
	GetTranscriptSummary() *bool
}

type GetVideoAnalysisOption struct {
	// Specifies whether to generate a chapter summary for each segment of the video.
	ChapterSummary *bool `json:"ChapterSummary,omitempty" xml:"ChapterSummary,omitempty"`
	// Specifies whether to extract relevant keywords from the video.
	Keyword *bool `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Specifies whether to generate a presentation (PPT) file based on the video analysis.
	PPT *bool `json:"PPT,omitempty" xml:"PPT,omitempty"`
	// Specifies whether to generate potential questions about the video.
	Question *bool `json:"Question,omitempty" xml:"Question,omitempty"`
	// Specifies whether to generate a concise video summary.
	Summary *bool `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Specifies whether to transcribe the spoken audio in the video to text.
	Transcript *bool `json:"Transcript,omitempty" xml:"Transcript,omitempty"`
	// Specifies whether to generate a chapter-based summary from the video transcript.
	TranscriptChapterSummary *bool `json:"TranscriptChapterSummary,omitempty" xml:"TranscriptChapterSummary,omitempty"`
	// Specifies whether to generate a summary of the video transcript.
	TranscriptSummary *bool `json:"TranscriptSummary,omitempty" xml:"TranscriptSummary,omitempty"`
}

func (s GetVideoAnalysisOption) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisOption) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisOption) GetChapterSummary() *bool {
	return s.ChapterSummary
}

func (s *GetVideoAnalysisOption) GetKeyword() *bool {
	return s.Keyword
}

func (s *GetVideoAnalysisOption) GetPPT() *bool {
	return s.PPT
}

func (s *GetVideoAnalysisOption) GetQuestion() *bool {
	return s.Question
}

func (s *GetVideoAnalysisOption) GetSummary() *bool {
	return s.Summary
}

func (s *GetVideoAnalysisOption) GetTranscript() *bool {
	return s.Transcript
}

func (s *GetVideoAnalysisOption) GetTranscriptChapterSummary() *bool {
	return s.TranscriptChapterSummary
}

func (s *GetVideoAnalysisOption) GetTranscriptSummary() *bool {
	return s.TranscriptSummary
}

func (s *GetVideoAnalysisOption) SetChapterSummary(v bool) *GetVideoAnalysisOption {
	s.ChapterSummary = &v
	return s
}

func (s *GetVideoAnalysisOption) SetKeyword(v bool) *GetVideoAnalysisOption {
	s.Keyword = &v
	return s
}

func (s *GetVideoAnalysisOption) SetPPT(v bool) *GetVideoAnalysisOption {
	s.PPT = &v
	return s
}

func (s *GetVideoAnalysisOption) SetQuestion(v bool) *GetVideoAnalysisOption {
	s.Question = &v
	return s
}

func (s *GetVideoAnalysisOption) SetSummary(v bool) *GetVideoAnalysisOption {
	s.Summary = &v
	return s
}

func (s *GetVideoAnalysisOption) SetTranscript(v bool) *GetVideoAnalysisOption {
	s.Transcript = &v
	return s
}

func (s *GetVideoAnalysisOption) SetTranscriptChapterSummary(v bool) *GetVideoAnalysisOption {
	s.TranscriptChapterSummary = &v
	return s
}

func (s *GetVideoAnalysisOption) SetTranscriptSummary(v bool) *GetVideoAnalysisOption {
	s.TranscriptSummary = &v
	return s
}

func (s *GetVideoAnalysisOption) Validate() error {
	return dara.Validate(s)
}
