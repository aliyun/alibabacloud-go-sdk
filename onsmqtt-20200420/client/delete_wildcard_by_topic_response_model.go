// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWildcardByTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWildcardByTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWildcardByTopicResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWildcardByTopicResponseBody) *DeleteWildcardByTopicResponse
	GetBody() *DeleteWildcardByTopicResponseBody
}

type DeleteWildcardByTopicResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWildcardByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWildcardByTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWildcardByTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteWildcardByTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWildcardByTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWildcardByTopicResponse) GetBody() *DeleteWildcardByTopicResponseBody {
	return s.Body
}

func (s *DeleteWildcardByTopicResponse) SetHeaders(v map[string]*string) *DeleteWildcardByTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteWildcardByTopicResponse) SetStatusCode(v int32) *DeleteWildcardByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWildcardByTopicResponse) SetBody(v *DeleteWildcardByTopicResponseBody) *DeleteWildcardByTopicResponse {
	s.Body = v
	return s
}

func (s *DeleteWildcardByTopicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
