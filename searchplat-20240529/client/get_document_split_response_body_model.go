// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentSplitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetDocumentSplitResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetDocumentSplitResponseBody
	GetRequestId() *string
	SetResult(v *GetDocumentSplitResponseBodyResult) *GetDocumentSplitResponseBody
	GetResult() *GetDocumentSplitResponseBodyResult
	SetUsage(v *GetDocumentSplitResponseBodyUsage) *GetDocumentSplitResponseBody
	GetUsage() *GetDocumentSplitResponseBodyUsage
}

type GetDocumentSplitResponseBody struct {
	Latency   *int32                              `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                             `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetDocumentSplitResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetDocumentSplitResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetDocumentSplitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetDocumentSplitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentSplitResponseBody) GetResult() *GetDocumentSplitResponseBodyResult {
	return s.Result
}

func (s *GetDocumentSplitResponseBody) GetUsage() *GetDocumentSplitResponseBodyUsage {
	return s.Usage
}

func (s *GetDocumentSplitResponseBody) SetLatency(v int32) *GetDocumentSplitResponseBody {
	s.Latency = &v
	return s
}

func (s *GetDocumentSplitResponseBody) SetRequestId(v string) *GetDocumentSplitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentSplitResponseBody) SetResult(v *GetDocumentSplitResponseBodyResult) *GetDocumentSplitResponseBody {
	s.Result = v
	return s
}

func (s *GetDocumentSplitResponseBody) SetUsage(v *GetDocumentSplitResponseBodyUsage) *GetDocumentSplitResponseBody {
	s.Usage = v
	return s
}

func (s *GetDocumentSplitResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentSplitResponseBodyResult struct {
	Chunks    []*GetDocumentSplitResponseBodyResultChunks    `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	Nodes     []map[string]*string                           `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	RichTexts []*GetDocumentSplitResponseBodyResultRichTexts `json:"rich_texts,omitempty" xml:"rich_texts,omitempty" type:"Repeated"`
	Sentences []*GetDocumentSplitResponseBodyResultSentences `json:"sentences,omitempty" xml:"sentences,omitempty" type:"Repeated"`
}

func (s GetDocumentSplitResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResult) GetChunks() []*GetDocumentSplitResponseBodyResultChunks {
	return s.Chunks
}

func (s *GetDocumentSplitResponseBodyResult) GetNodes() []map[string]*string {
	return s.Nodes
}

func (s *GetDocumentSplitResponseBodyResult) GetRichTexts() []*GetDocumentSplitResponseBodyResultRichTexts {
	return s.RichTexts
}

func (s *GetDocumentSplitResponseBodyResult) GetSentences() []*GetDocumentSplitResponseBodyResultSentences {
	return s.Sentences
}

func (s *GetDocumentSplitResponseBodyResult) SetChunks(v []*GetDocumentSplitResponseBodyResultChunks) *GetDocumentSplitResponseBodyResult {
	s.Chunks = v
	return s
}

func (s *GetDocumentSplitResponseBodyResult) SetNodes(v []map[string]*string) *GetDocumentSplitResponseBodyResult {
	s.Nodes = v
	return s
}

func (s *GetDocumentSplitResponseBodyResult) SetRichTexts(v []*GetDocumentSplitResponseBodyResultRichTexts) *GetDocumentSplitResponseBodyResult {
	s.RichTexts = v
	return s
}

func (s *GetDocumentSplitResponseBodyResult) SetSentences(v []*GetDocumentSplitResponseBodyResultSentences) *GetDocumentSplitResponseBodyResult {
	s.Sentences = v
	return s
}

func (s *GetDocumentSplitResponseBodyResult) Validate() error {
	if s.Chunks != nil {
		for _, item := range s.Chunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RichTexts != nil {
		for _, item := range s.RichTexts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Sentences != nil {
		for _, item := range s.Sentences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDocumentSplitResponseBodyResultChunks struct {
	Content *string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta    map[string]*string `json:"meta,omitempty" xml:"meta,omitempty"`
}

func (s GetDocumentSplitResponseBodyResultChunks) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResultChunks) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResultChunks) GetContent() *string {
	return s.Content
}

func (s *GetDocumentSplitResponseBodyResultChunks) GetMeta() map[string]*string {
	return s.Meta
}

func (s *GetDocumentSplitResponseBodyResultChunks) SetContent(v string) *GetDocumentSplitResponseBodyResultChunks {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitResponseBodyResultChunks) SetMeta(v map[string]*string) *GetDocumentSplitResponseBodyResultChunks {
	s.Meta = v
	return s
}

func (s *GetDocumentSplitResponseBodyResultChunks) Validate() error {
	return dara.Validate(s)
}

type GetDocumentSplitResponseBodyResultRichTexts struct {
	Content *string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta    map[string]*string `json:"meta,omitempty" xml:"meta,omitempty"`
}

func (s GetDocumentSplitResponseBodyResultRichTexts) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResultRichTexts) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResultRichTexts) GetContent() *string {
	return s.Content
}

func (s *GetDocumentSplitResponseBodyResultRichTexts) GetMeta() map[string]*string {
	return s.Meta
}

func (s *GetDocumentSplitResponseBodyResultRichTexts) SetContent(v string) *GetDocumentSplitResponseBodyResultRichTexts {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitResponseBodyResultRichTexts) SetMeta(v map[string]*string) *GetDocumentSplitResponseBodyResultRichTexts {
	s.Meta = v
	return s
}

func (s *GetDocumentSplitResponseBodyResultRichTexts) Validate() error {
	return dara.Validate(s)
}

type GetDocumentSplitResponseBodyResultSentences struct {
	Content *string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta    map[string]*string `json:"meta,omitempty" xml:"meta,omitempty"`
}

func (s GetDocumentSplitResponseBodyResultSentences) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResultSentences) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResultSentences) GetContent() *string {
	return s.Content
}

func (s *GetDocumentSplitResponseBodyResultSentences) GetMeta() map[string]*string {
	return s.Meta
}

func (s *GetDocumentSplitResponseBodyResultSentences) SetContent(v string) *GetDocumentSplitResponseBodyResultSentences {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitResponseBodyResultSentences) SetMeta(v map[string]*string) *GetDocumentSplitResponseBodyResultSentences {
	s.Meta = v
	return s
}

func (s *GetDocumentSplitResponseBodyResultSentences) Validate() error {
	return dara.Validate(s)
}

type GetDocumentSplitResponseBodyUsage struct {
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetDocumentSplitResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyUsage) GetTokenCount() *int64 {
	return s.TokenCount
}

func (s *GetDocumentSplitResponseBodyUsage) SetTokenCount(v int64) *GetDocumentSplitResponseBodyUsage {
	s.TokenCount = &v
	return s
}

func (s *GetDocumentSplitResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
