// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeprecatedTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDeprecatedTemplateResponseBody
	GetRequestId() *string
	SetResult(v *DescribeDeprecatedTemplateResponseBodyResult) *DescribeDeprecatedTemplateResponseBody
	GetResult() *DescribeDeprecatedTemplateResponseBodyResult
}

type DescribeDeprecatedTemplateResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDeprecatedTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDeprecatedTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeprecatedTemplateResponseBody) GetResult() *DescribeDeprecatedTemplateResponseBodyResult {
	return s.Result
}

func (s *DescribeDeprecatedTemplateResponseBody) SetRequestId(v string) *DescribeDeprecatedTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBody) SetResult(v *DescribeDeprecatedTemplateResponseBodyResult) *DescribeDeprecatedTemplateResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDeprecatedTemplateResponseBodyResult struct {
	// example:
	//
	// false
	DataStream    *bool     `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	IndexPatterns []*string `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	// example:
	//
	// openstore-index-template
	IndexTemplate *string `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	// example:
	//
	// 100
	Order    *int64                                                `json:"order,omitempty" xml:"order,omitempty"`
	Template *DescribeDeprecatedTemplateResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// example:
	//
	// 70000
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeDeprecatedTemplateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) GetDataStream() *bool {
	return s.DataStream
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) GetIndexPatterns() []*string {
	return s.IndexPatterns
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) GetIndexTemplate() *string {
	return s.IndexTemplate
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) GetOrder() *int64 {
	return s.Order
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) GetTemplate() *DescribeDeprecatedTemplateResponseBodyResultTemplate {
	return s.Template
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetDataStream(v bool) *DescribeDeprecatedTemplateResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetIndexPatterns(v []*string) *DescribeDeprecatedTemplateResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetIndexTemplate(v string) *DescribeDeprecatedTemplateResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetOrder(v int64) *DescribeDeprecatedTemplateResponseBodyResult {
	s.Order = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetTemplate(v *DescribeDeprecatedTemplateResponseBodyResultTemplate) *DescribeDeprecatedTemplateResponseBodyResult {
	s.Template = v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetVersion(v string) *DescribeDeprecatedTemplateResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type DescribeDeprecatedTemplateResponseBodyResultTemplate struct {
	// example:
	//
	// "{}"
	Aliases *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	// example:
	//
	// "{\\"properties\\":{\\"created_at\\":{\\"format\\":\\"EEE MMM dd HH:mm:ss Z yyyy\\",\\"type\\":\\"date\\"},\\"host_name\\":{\\"type\\":\\"keyword\\"}}}"
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	// example:
	//
	// "{\\"index.number_of_shards\\":\\"1\\"}"
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s DescribeDeprecatedTemplateResponseBodyResultTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) GetAliases() *string {
	return s.Aliases
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) GetMappings() *string {
	return s.Mappings
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) GetSettings() *string {
	return s.Settings
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) SetAliases(v string) *DescribeDeprecatedTemplateResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) SetMappings(v string) *DescribeDeprecatedTemplateResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) SetSettings(v string) *DescribeDeprecatedTemplateResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) Validate() error {
	return dara.Validate(s)
}
