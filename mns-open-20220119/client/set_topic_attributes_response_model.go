// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTopicAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetTopicAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetTopicAttributesResponse
	GetStatusCode() *int32
	SetBody(v *SetTopicAttributesResponseBody) *SetTopicAttributesResponse
	GetBody() *SetTopicAttributesResponseBody
}

type SetTopicAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTopicAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTopicAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetTopicAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetTopicAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetTopicAttributesResponse) GetBody() *SetTopicAttributesResponseBody {
	return s.Body
}

func (s *SetTopicAttributesResponse) SetHeaders(v map[string]*string) *SetTopicAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetTopicAttributesResponse) SetStatusCode(v int32) *SetTopicAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTopicAttributesResponse) SetBody(v *SetTopicAttributesResponseBody) *SetTopicAttributesResponse {
	s.Body = v
	return s
}

func (s *SetTopicAttributesResponse) Validate() error {
	return dara.Validate(s)
}
