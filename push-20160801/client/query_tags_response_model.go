// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTagsResponse
	GetStatusCode() *int32
	SetBody(v *QueryTagsResponseBody) *QueryTagsResponse
	GetBody() *QueryTagsResponseBody
}

type QueryTagsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTagsResponse) GoString() string {
	return s.String()
}

func (s *QueryTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTagsResponse) GetBody() *QueryTagsResponseBody {
	return s.Body
}

func (s *QueryTagsResponse) SetHeaders(v map[string]*string) *QueryTagsResponse {
	s.Headers = v
	return s
}

func (s *QueryTagsResponse) SetStatusCode(v int32) *QueryTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagsResponse) SetBody(v *QueryTagsResponseBody) *QueryTagsResponse {
	s.Body = v
	return s
}

func (s *QueryTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
