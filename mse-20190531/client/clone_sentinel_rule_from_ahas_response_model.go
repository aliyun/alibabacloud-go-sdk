// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneSentinelRuleFromAhasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneSentinelRuleFromAhasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneSentinelRuleFromAhasResponse
	GetStatusCode() *int32
	SetBody(v *CloneSentinelRuleFromAhasResponseBody) *CloneSentinelRuleFromAhasResponse
	GetBody() *CloneSentinelRuleFromAhasResponseBody
}

type CloneSentinelRuleFromAhasResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneSentinelRuleFromAhasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneSentinelRuleFromAhasResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneSentinelRuleFromAhasResponse) GoString() string {
	return s.String()
}

func (s *CloneSentinelRuleFromAhasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneSentinelRuleFromAhasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneSentinelRuleFromAhasResponse) GetBody() *CloneSentinelRuleFromAhasResponseBody {
	return s.Body
}

func (s *CloneSentinelRuleFromAhasResponse) SetHeaders(v map[string]*string) *CloneSentinelRuleFromAhasResponse {
	s.Headers = v
	return s
}

func (s *CloneSentinelRuleFromAhasResponse) SetStatusCode(v int32) *CloneSentinelRuleFromAhasResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneSentinelRuleFromAhasResponse) SetBody(v *CloneSentinelRuleFromAhasResponseBody) *CloneSentinelRuleFromAhasResponse {
	s.Body = v
	return s
}

func (s *CloneSentinelRuleFromAhasResponse) Validate() error {
	return dara.Validate(s)
}
