// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicAttributesResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicAttributesResponseBody) *GetTopicAttributesResponse
	GetBody() *GetTopicAttributesResponseBody
}

type GetTopicAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicAttributesResponse) GetBody() *GetTopicAttributesResponseBody {
	return s.Body
}

func (s *GetTopicAttributesResponse) SetHeaders(v map[string]*string) *GetTopicAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetTopicAttributesResponse) SetStatusCode(v int32) *GetTopicAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicAttributesResponse) SetBody(v *GetTopicAttributesResponseBody) *GetTopicAttributesResponse {
	s.Body = v
	return s
}

func (s *GetTopicAttributesResponse) Validate() error {
	return dara.Validate(s)
}
