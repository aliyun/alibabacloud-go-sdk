// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchSimilarArticlesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSearchSimilarArticlesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSearchSimilarArticlesResponse
	GetStatusCode() *int32
	SetId(v string) *RunSearchSimilarArticlesResponse
	GetId() *string
	SetEvent(v string) *RunSearchSimilarArticlesResponse
	GetEvent() *string
	SetBody(v *RunSearchSimilarArticlesResponseBody) *RunSearchSimilarArticlesResponse
	GetBody() *RunSearchSimilarArticlesResponseBody
}

type RunSearchSimilarArticlesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                               `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                               `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunSearchSimilarArticlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSearchSimilarArticlesResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponse) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSearchSimilarArticlesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSearchSimilarArticlesResponse) GetId() *string {
	return s.Id
}

func (s *RunSearchSimilarArticlesResponse) GetEvent() *string {
	return s.Event
}

func (s *RunSearchSimilarArticlesResponse) GetBody() *RunSearchSimilarArticlesResponseBody {
	return s.Body
}

func (s *RunSearchSimilarArticlesResponse) SetHeaders(v map[string]*string) *RunSearchSimilarArticlesResponse {
	s.Headers = v
	return s
}

func (s *RunSearchSimilarArticlesResponse) SetStatusCode(v int32) *RunSearchSimilarArticlesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSearchSimilarArticlesResponse) SetId(v string) *RunSearchSimilarArticlesResponse {
	s.Id = &v
	return s
}

func (s *RunSearchSimilarArticlesResponse) SetEvent(v string) *RunSearchSimilarArticlesResponse {
	s.Event = &v
	return s
}

func (s *RunSearchSimilarArticlesResponse) SetBody(v *RunSearchSimilarArticlesResponseBody) *RunSearchSimilarArticlesResponse {
	s.Body = v
	return s
}

func (s *RunSearchSimilarArticlesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
