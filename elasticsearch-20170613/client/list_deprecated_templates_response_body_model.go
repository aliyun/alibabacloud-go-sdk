// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeprecatedTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListDeprecatedTemplatesResponseBodyHeaders) *ListDeprecatedTemplatesResponseBody
	GetHeaders() *ListDeprecatedTemplatesResponseBodyHeaders
	SetRequestId(v string) *ListDeprecatedTemplatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListDeprecatedTemplatesResponseBodyResult) *ListDeprecatedTemplatesResponseBody
	GetResult() []*ListDeprecatedTemplatesResponseBodyResult
}

type ListDeprecatedTemplatesResponseBody struct {
	Headers *ListDeprecatedTemplatesResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDeprecatedTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDeprecatedTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBody) GetHeaders() *ListDeprecatedTemplatesResponseBodyHeaders {
	return s.Headers
}

func (s *ListDeprecatedTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeprecatedTemplatesResponseBody) GetResult() []*ListDeprecatedTemplatesResponseBodyResult {
	return s.Result
}

func (s *ListDeprecatedTemplatesResponseBody) SetHeaders(v *ListDeprecatedTemplatesResponseBodyHeaders) *ListDeprecatedTemplatesResponseBody {
	s.Headers = v
	return s
}

func (s *ListDeprecatedTemplatesResponseBody) SetRequestId(v string) *ListDeprecatedTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBody) SetResult(v []*ListDeprecatedTemplatesResponseBodyResult) *ListDeprecatedTemplatesResponseBody {
	s.Result = v
	return s
}

func (s *ListDeprecatedTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeprecatedTemplatesResponseBodyHeaders struct {
	// example:
	//
	// 5
	XTotalCount *int64 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDeprecatedTemplatesResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBodyHeaders) GetXTotalCount() *int64 {
	return s.XTotalCount
}

func (s *ListDeprecatedTemplatesResponseBodyHeaders) SetXTotalCount(v int64) *ListDeprecatedTemplatesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListDeprecatedTemplatesResponseBodyResult struct {
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
	Order    *int64                                             `json:"order,omitempty" xml:"order,omitempty"`
	Template *ListDeprecatedTemplatesResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// example:
	//
	// 70000
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDeprecatedTemplatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBodyResult) GetDataStream() *bool {
	return s.DataStream
}

func (s *ListDeprecatedTemplatesResponseBodyResult) GetIndexPatterns() []*string {
	return s.IndexPatterns
}

func (s *ListDeprecatedTemplatesResponseBodyResult) GetIndexTemplate() *string {
	return s.IndexTemplate
}

func (s *ListDeprecatedTemplatesResponseBodyResult) GetOrder() *int64 {
	return s.Order
}

func (s *ListDeprecatedTemplatesResponseBodyResult) GetTemplate() *ListDeprecatedTemplatesResponseBodyResultTemplate {
	return s.Template
}

func (s *ListDeprecatedTemplatesResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetDataStream(v bool) *ListDeprecatedTemplatesResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetIndexPatterns(v []*string) *ListDeprecatedTemplatesResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetIndexTemplate(v string) *ListDeprecatedTemplatesResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetOrder(v int64) *ListDeprecatedTemplatesResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetTemplate(v *ListDeprecatedTemplatesResponseBodyResultTemplate) *ListDeprecatedTemplatesResponseBodyResult {
	s.Template = v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetVersion(v string) *ListDeprecatedTemplatesResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListDeprecatedTemplatesResponseBodyResultTemplate struct {
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

func (s ListDeprecatedTemplatesResponseBodyResultTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) GetAliases() *string {
	return s.Aliases
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) GetMappings() *string {
	return s.Mappings
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) GetSettings() *string {
	return s.Settings
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) SetAliases(v string) *ListDeprecatedTemplatesResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) SetMappings(v string) *ListDeprecatedTemplatesResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) SetSettings(v string) *ListDeprecatedTemplatesResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) Validate() error {
	return dara.Validate(s)
}
