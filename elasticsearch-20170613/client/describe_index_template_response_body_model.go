// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIndexTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeIndexTemplateResponseBody
	GetRequestId() *string
	SetResult(v *DescribeIndexTemplateResponseBodyResult) *DescribeIndexTemplateResponseBody
	GetResult() *DescribeIndexTemplateResponseBodyResult
}

type DescribeIndexTemplateResponseBody struct {
	// example:
	//
	// 25DB38F8-82E4-4D16-82BB-FF077C7F****
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeIndexTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeIndexTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIndexTemplateResponseBody) GetResult() *DescribeIndexTemplateResponseBodyResult {
	return s.Result
}

func (s *DescribeIndexTemplateResponseBody) SetRequestId(v string) *DescribeIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIndexTemplateResponseBody) SetResult(v *DescribeIndexTemplateResponseBodyResult) *DescribeIndexTemplateResponseBody {
	s.Result = v
	return s
}

func (s *DescribeIndexTemplateResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIndexTemplateResponseBodyResult struct {
	// example:
	//
	// true
	DataStream *bool `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	// example:
	//
	// cube_default_ilm_policy
	IlmPolicy     *string   `json:"ilmPolicy,omitempty" xml:"ilmPolicy,omitempty"`
	IndexPatterns []*string `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	// example:
	//
	// data-stream-default
	IndexTemplate *string `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	// example:
	//
	// 0
	Priority *int32                                           `json:"priority,omitempty" xml:"priority,omitempty"`
	Template *DescribeIndexTemplateResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s DescribeIndexTemplateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeIndexTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponseBodyResult) GetDataStream() *bool {
	return s.DataStream
}

func (s *DescribeIndexTemplateResponseBodyResult) GetIlmPolicy() *string {
	return s.IlmPolicy
}

func (s *DescribeIndexTemplateResponseBodyResult) GetIndexPatterns() []*string {
	return s.IndexPatterns
}

func (s *DescribeIndexTemplateResponseBodyResult) GetIndexTemplate() *string {
	return s.IndexTemplate
}

func (s *DescribeIndexTemplateResponseBodyResult) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeIndexTemplateResponseBodyResult) GetTemplate() *DescribeIndexTemplateResponseBodyResultTemplate {
	return s.Template
}

func (s *DescribeIndexTemplateResponseBodyResult) SetDataStream(v bool) *DescribeIndexTemplateResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetIlmPolicy(v string) *DescribeIndexTemplateResponseBodyResult {
	s.IlmPolicy = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetIndexPatterns(v []*string) *DescribeIndexTemplateResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetIndexTemplate(v string) *DescribeIndexTemplateResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetPriority(v int32) *DescribeIndexTemplateResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetTemplate(v *DescribeIndexTemplateResponseBodyResultTemplate) *DescribeIndexTemplateResponseBodyResult {
	s.Template = v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIndexTemplateResponseBodyResultTemplate struct {
	// example:
	//
	// {\\"mydata\\":{}}
	Aliases *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	// example:
	//
	// {\\"properties\\":{\\"created_at\\":{\\"format\\":\\"EEE MMM dd HH:mm:ss Z yyyy\\",\\"type\\":\\"date\\"},\\"host_name\\":{\\"type\\":\\"keyword\\"}}}
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	// example:
	//
	// {\\"index.refresh_interval\\":\\"1s\\"}
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s DescribeIndexTemplateResponseBodyResultTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeIndexTemplateResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) GetAliases() *string {
	return s.Aliases
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) GetMappings() *string {
	return s.Mappings
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) GetSettings() *string {
	return s.Settings
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) SetAliases(v string) *DescribeIndexTemplateResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) SetMappings(v string) *DescribeIndexTemplateResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) SetSettings(v string) *DescribeIndexTemplateResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) Validate() error {
	return dara.Validate(s)
}
