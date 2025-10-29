// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListIndexTemplatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListIndexTemplatesResponseBodyResult) *ListIndexTemplatesResponseBody
	GetResult() []*ListIndexTemplatesResponseBodyResult
}

type ListIndexTemplatesResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListIndexTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListIndexTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndexTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndexTemplatesResponseBody) GetResult() []*ListIndexTemplatesResponseBodyResult {
	return s.Result
}

func (s *ListIndexTemplatesResponseBody) SetRequestId(v string) *ListIndexTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexTemplatesResponseBody) SetResult(v []*ListIndexTemplatesResponseBodyResult) *ListIndexTemplatesResponseBody {
	s.Result = v
	return s
}

func (s *ListIndexTemplatesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIndexTemplatesResponseBodyResult struct {
	// example:
	//
	// true
	DataStream *bool `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	// example:
	//
	// my_ilm_policy
	IlmPolicy     *string   `json:"ilmPolicy,omitempty" xml:"ilmPolicy,omitempty"`
	IndexPatterns []*string `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	// example:
	//
	// my-template
	IndexTemplate *string `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	// example:
	//
	// 100
	Priority *int32                                        `json:"priority,omitempty" xml:"priority,omitempty"`
	Template *ListIndexTemplatesResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s ListIndexTemplatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListIndexTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponseBodyResult) GetDataStream() *bool {
	return s.DataStream
}

func (s *ListIndexTemplatesResponseBodyResult) GetIlmPolicy() *string {
	return s.IlmPolicy
}

func (s *ListIndexTemplatesResponseBodyResult) GetIndexPatterns() []*string {
	return s.IndexPatterns
}

func (s *ListIndexTemplatesResponseBodyResult) GetIndexTemplate() *string {
	return s.IndexTemplate
}

func (s *ListIndexTemplatesResponseBodyResult) GetPriority() *int32 {
	return s.Priority
}

func (s *ListIndexTemplatesResponseBodyResult) GetTemplate() *ListIndexTemplatesResponseBodyResultTemplate {
	return s.Template
}

func (s *ListIndexTemplatesResponseBodyResult) SetDataStream(v bool) *ListIndexTemplatesResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetIlmPolicy(v string) *ListIndexTemplatesResponseBodyResult {
	s.IlmPolicy = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetIndexPatterns(v []*string) *ListIndexTemplatesResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetIndexTemplate(v string) *ListIndexTemplatesResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetPriority(v int32) *ListIndexTemplatesResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetTemplate(v *ListIndexTemplatesResponseBodyResultTemplate) *ListIndexTemplatesResponseBodyResult {
	s.Template = v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIndexTemplatesResponseBodyResultTemplate struct {
	// example:
	//
	// {\\"index.number_of_shards\\":\\"1\\"}
	Aliases *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	// example:
	//
	// {\\"properties\\":{\\"created_at\\":{\\"format\\":\\"EEE MMM dd HH:mm:ss Z yyyy\\",\\"type\\":\\"date\\"},\\"host_name\\":{\\"type\\":\\"keyword\\"}}}
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	// example:
	//
	// {\\"mydata\\":{}}
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s ListIndexTemplatesResponseBodyResultTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListIndexTemplatesResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) GetAliases() *string {
	return s.Aliases
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) GetMappings() *string {
	return s.Mappings
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) GetSettings() *string {
	return s.Settings
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) SetAliases(v string) *ListIndexTemplatesResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) SetMappings(v string) *ListIndexTemplatesResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) SetSettings(v string) *ListIndexTemplatesResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) Validate() error {
	return dara.Validate(s)
}
