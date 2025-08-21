// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RecommendTemplatesResponseBody
	GetRequestId() *string
	SetResult(v []*RecommendTemplatesResponseBodyResult) *RecommendTemplatesResponseBody
	GetResult() []*RecommendTemplatesResponseBodyResult
}

type RecommendTemplatesResponseBody struct {
	// example:
	//
	// 66B060CF-7381-49C7-9B89-7757927FDA16
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*RecommendTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s RecommendTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecommendTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecommendTemplatesResponseBody) GetResult() []*RecommendTemplatesResponseBodyResult {
	return s.Result
}

func (s *RecommendTemplatesResponseBody) SetRequestId(v string) *RecommendTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecommendTemplatesResponseBody) SetResult(v []*RecommendTemplatesResponseBodyResult) *RecommendTemplatesResponseBody {
	s.Result = v
	return s
}

func (s *RecommendTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecommendTemplatesResponseBodyResult struct {
	// example:
	//
	// {\\n\\t\\"persistent\\": {\\n\\t\\t\\"search\\": {\\n\\t\\t\\t\\"max_buckets\\": \\"10000\\"\\n\\t\\t}\\n\\t}\\n}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// **
	//
	// ****
	//
	// example:
	//
	// dynamicSettings
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s RecommendTemplatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s RecommendTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *RecommendTemplatesResponseBodyResult) GetTemplateName() *string {
	return s.TemplateName
}

func (s *RecommendTemplatesResponseBodyResult) SetContent(v string) *RecommendTemplatesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *RecommendTemplatesResponseBodyResult) SetTemplateName(v string) *RecommendTemplatesResponseBodyResult {
	s.TemplateName = &v
	return s
}

func (s *RecommendTemplatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
