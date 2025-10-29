// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentIndicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListComponentIndicesResponseBodyHeaders) *ListComponentIndicesResponseBody
	GetHeaders() *ListComponentIndicesResponseBodyHeaders
	SetRequestId(v string) *ListComponentIndicesResponseBody
	GetRequestId() *string
	SetResult(v []*ListComponentIndicesResponseBodyResult) *ListComponentIndicesResponseBody
	GetResult() []*ListComponentIndicesResponseBodyResult
}

type ListComponentIndicesResponseBody struct {
	Headers *ListComponentIndicesResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListComponentIndicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListComponentIndicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBody) GetHeaders() *ListComponentIndicesResponseBodyHeaders {
	return s.Headers
}

func (s *ListComponentIndicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComponentIndicesResponseBody) GetResult() []*ListComponentIndicesResponseBodyResult {
	return s.Result
}

func (s *ListComponentIndicesResponseBody) SetHeaders(v *ListComponentIndicesResponseBodyHeaders) *ListComponentIndicesResponseBody {
	s.Headers = v
	return s
}

func (s *ListComponentIndicesResponseBody) SetRequestId(v string) *ListComponentIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentIndicesResponseBody) SetResult(v []*ListComponentIndicesResponseBodyResult) *ListComponentIndicesResponseBody {
	s.Result = v
	return s
}

func (s *ListComponentIndicesResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
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

type ListComponentIndicesResponseBodyHeaders struct {
	// example:
	//
	// 10
	XTotalCount *int64 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListComponentIndicesResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyHeaders) GetXTotalCount() *int64 {
	return s.XTotalCount
}

func (s *ListComponentIndicesResponseBodyHeaders) SetXTotalCount(v int64) *ListComponentIndicesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListComponentIndicesResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListComponentIndicesResponseBodyResult struct {
	Composed []*string                                      `json:"composed,omitempty" xml:"composed,omitempty" type:"Repeated"`
	Content  *ListComponentIndicesResponseBodyResultContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// synthetics-settings
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListComponentIndicesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResult) GetComposed() []*string {
	return s.Composed
}

func (s *ListComponentIndicesResponseBodyResult) GetContent() *ListComponentIndicesResponseBodyResultContent {
	return s.Content
}

func (s *ListComponentIndicesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListComponentIndicesResponseBodyResult) SetComposed(v []*string) *ListComponentIndicesResponseBodyResult {
	s.Composed = v
	return s
}

func (s *ListComponentIndicesResponseBodyResult) SetContent(v *ListComponentIndicesResponseBodyResultContent) *ListComponentIndicesResponseBodyResult {
	s.Content = v
	return s
}

func (s *ListComponentIndicesResponseBodyResult) SetName(v string) *ListComponentIndicesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListComponentIndicesResponseBodyResult) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComponentIndicesResponseBodyResultContent struct {
	// example:
	//
	// { "description": "set number of shards to one" }
	Meta     map[string]interface{}                                 `json:"_meta,omitempty" xml:"_meta,omitempty"`
	Template *ListComponentIndicesResponseBodyResultContentTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListComponentIndicesResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContent) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListComponentIndicesResponseBodyResultContent) GetTemplate() *ListComponentIndicesResponseBodyResultContentTemplate {
	return s.Template
}

func (s *ListComponentIndicesResponseBodyResultContent) GetVersion() *int64 {
	return s.Version
}

func (s *ListComponentIndicesResponseBodyResultContent) SetMeta(v map[string]interface{}) *ListComponentIndicesResponseBodyResultContent {
	s.Meta = v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContent) SetTemplate(v *ListComponentIndicesResponseBodyResultContentTemplate) *ListComponentIndicesResponseBodyResultContent {
	s.Template = v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContent) SetVersion(v int64) *ListComponentIndicesResponseBodyResultContent {
	s.Version = &v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContent) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComponentIndicesResponseBodyResultContentTemplate struct {
	Settings *ListComponentIndicesResponseBodyResultContentTemplateSettings `json:"settings,omitempty" xml:"settings,omitempty" type:"Struct"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplate) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplate) GetSettings() *ListComponentIndicesResponseBodyResultContentTemplateSettings {
	return s.Settings
}

func (s *ListComponentIndicesResponseBodyResultContentTemplate) SetSettings(v *ListComponentIndicesResponseBodyResultContentTemplateSettings) *ListComponentIndicesResponseBodyResultContentTemplate {
	s.Settings = v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContentTemplate) Validate() error {
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComponentIndicesResponseBodyResultContentTemplateSettings struct {
	Index *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex `json:"index,omitempty" xml:"index,omitempty" type:"Struct"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettings) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettings) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettings) GetIndex() *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex {
	return s.Index
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettings) SetIndex(v *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) *ListComponentIndicesResponseBodyResultContentTemplateSettings {
	s.Index = v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettings) Validate() error {
	if s.Index != nil {
		if err := s.Index.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex struct {
	// example:
	//
	// best_compression
	Codec     *string                                                                      `json:"codec,omitempty" xml:"codec,omitempty"`
	Lifecycle *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle `json:"lifecycle,omitempty" xml:"lifecycle,omitempty" type:"Struct"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) GetCodec() *string {
	return s.Codec
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) GetLifecycle() *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle {
	return s.Lifecycle
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) SetCodec(v string) *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex {
	s.Codec = &v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) SetLifecycle(v *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex {
	s.Lifecycle = v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) Validate() error {
	if s.Lifecycle != nil {
		if err := s.Lifecycle.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle struct {
	// example:
	//
	// synthetics
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) GetName() *string {
	return s.Name
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) SetName(v string) *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle {
	s.Name = &v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) Validate() error {
	return dara.Validate(s)
}
