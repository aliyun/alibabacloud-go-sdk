// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWildcardByTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWildcardByTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWildcardByTopicResponse
	GetStatusCode() *int32
	SetBody(v *QueryWildcardByTopicResponseBody) *QueryWildcardByTopicResponse
	GetBody() *QueryWildcardByTopicResponseBody
}

type QueryWildcardByTopicResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWildcardByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWildcardByTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWildcardByTopicResponse) GoString() string {
	return s.String()
}

func (s *QueryWildcardByTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWildcardByTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWildcardByTopicResponse) GetBody() *QueryWildcardByTopicResponseBody {
	return s.Body
}

func (s *QueryWildcardByTopicResponse) SetHeaders(v map[string]*string) *QueryWildcardByTopicResponse {
	s.Headers = v
	return s
}

func (s *QueryWildcardByTopicResponse) SetStatusCode(v int32) *QueryWildcardByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWildcardByTopicResponse) SetBody(v *QueryWildcardByTopicResponseBody) *QueryWildcardByTopicResponse {
	s.Body = v
	return s
}

func (s *QueryWildcardByTopicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
