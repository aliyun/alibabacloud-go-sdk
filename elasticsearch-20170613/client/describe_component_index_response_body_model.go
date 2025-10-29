// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeComponentIndexResponseBody
	GetRequestId() *string
	SetResult(v *DescribeComponentIndexResponseBodyResult) *DescribeComponentIndexResponseBody
	GetResult() *DescribeComponentIndexResponseBodyResult
}

type DescribeComponentIndexResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeComponentIndexResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeComponentIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentIndexResponseBody) GetResult() *DescribeComponentIndexResponseBodyResult {
	return s.Result
}

func (s *DescribeComponentIndexResponseBody) SetRequestId(v string) *DescribeComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentIndexResponseBody) SetResult(v *DescribeComponentIndexResponseBodyResult) *DescribeComponentIndexResponseBody {
	s.Result = v
	return s
}

func (s *DescribeComponentIndexResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeComponentIndexResponseBodyResult struct {
	// example:
	//
	// { "description": "set number of shards to one" }
	Meta     map[string]interface{}                            `json:"_meta,omitempty" xml:"_meta,omitempty"`
	Template *DescribeComponentIndexResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s DescribeComponentIndexResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentIndexResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponseBodyResult) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *DescribeComponentIndexResponseBodyResult) GetTemplate() *DescribeComponentIndexResponseBodyResultTemplate {
	return s.Template
}

func (s *DescribeComponentIndexResponseBodyResult) SetMeta(v map[string]interface{}) *DescribeComponentIndexResponseBodyResult {
	s.Meta = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResult) SetTemplate(v *DescribeComponentIndexResponseBodyResultTemplate) *DescribeComponentIndexResponseBodyResult {
	s.Template = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResult) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeComponentIndexResponseBodyResultTemplate struct {
	// example:
	//
	// {}
	Aliases map[string]interface{} `json:"aliases,omitempty" xml:"aliases,omitempty"`
	// example:
	//
	// { "properties": { "@timestamp": { "type": "date" } } }
	Mappings map[string]interface{} `json:"mappings,omitempty" xml:"mappings,omitempty"`
	// example:
	//
	// { "index.number_of_replicas": 0 }
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s DescribeComponentIndexResponseBodyResultTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentIndexResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) GetAliases() map[string]interface{} {
	return s.Aliases
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) GetMappings() map[string]interface{} {
	return s.Mappings
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) GetSettings() map[string]interface{} {
	return s.Settings
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) SetAliases(v map[string]interface{}) *DescribeComponentIndexResponseBodyResultTemplate {
	s.Aliases = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) SetMappings(v map[string]interface{}) *DescribeComponentIndexResponseBodyResultTemplate {
	s.Mappings = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) SetSettings(v map[string]interface{}) *DescribeComponentIndexResponseBodyResultTemplate {
	s.Settings = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) Validate() error {
	return dara.Validate(s)
}
