// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJinaAiReaderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetJinaAiReaderResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetJinaAiReaderResponseBody
	GetRequestId() *string
	SetResult(v *GetJinaAiReaderResponseBodyResult) *GetJinaAiReaderResponseBody
	GetResult() *GetJinaAiReaderResponseBodyResult
	SetUsage(v *GetJinaAiReaderResponseBodyUsage) *GetJinaAiReaderResponseBody
	GetUsage() *GetJinaAiReaderResponseBodyUsage
}

type GetJinaAiReaderResponseBody struct {
	Latency   *int32                             `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                            `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetJinaAiReaderResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetJinaAiReaderResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetJinaAiReaderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJinaAiReaderResponseBody) GoString() string {
	return s.String()
}

func (s *GetJinaAiReaderResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetJinaAiReaderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJinaAiReaderResponseBody) GetResult() *GetJinaAiReaderResponseBodyResult {
	return s.Result
}

func (s *GetJinaAiReaderResponseBody) GetUsage() *GetJinaAiReaderResponseBodyUsage {
	return s.Usage
}

func (s *GetJinaAiReaderResponseBody) SetLatency(v int32) *GetJinaAiReaderResponseBody {
	s.Latency = &v
	return s
}

func (s *GetJinaAiReaderResponseBody) SetRequestId(v string) *GetJinaAiReaderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJinaAiReaderResponseBody) SetResult(v *GetJinaAiReaderResponseBodyResult) *GetJinaAiReaderResponseBody {
	s.Result = v
	return s
}

func (s *GetJinaAiReaderResponseBody) SetUsage(v *GetJinaAiReaderResponseBodyUsage) *GetJinaAiReaderResponseBody {
	s.Usage = v
	return s
}

func (s *GetJinaAiReaderResponseBody) Validate() error {
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

type GetJinaAiReaderResponseBodyResult struct {
	Content *string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta    map[string]*string `json:"meta,omitempty" xml:"meta,omitempty"`
}

func (s GetJinaAiReaderResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetJinaAiReaderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetJinaAiReaderResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetJinaAiReaderResponseBodyResult) GetMeta() map[string]*string {
	return s.Meta
}

func (s *GetJinaAiReaderResponseBodyResult) SetContent(v string) *GetJinaAiReaderResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetJinaAiReaderResponseBodyResult) SetMeta(v map[string]*string) *GetJinaAiReaderResponseBodyResult {
	s.Meta = v
	return s
}

func (s *GetJinaAiReaderResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetJinaAiReaderResponseBodyUsage struct {
	TokenCount *int32 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetJinaAiReaderResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetJinaAiReaderResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetJinaAiReaderResponseBodyUsage) GetTokenCount() *int32 {
	return s.TokenCount
}

func (s *GetJinaAiReaderResponseBodyUsage) SetTokenCount(v int32) *GetJinaAiReaderResponseBodyUsage {
	s.TokenCount = &v
	return s
}

func (s *GetJinaAiReaderResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
