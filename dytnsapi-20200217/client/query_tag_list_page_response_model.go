// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagListPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTagListPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTagListPageResponse
	GetStatusCode() *int32
	SetBody(v *QueryTagListPageResponseBody) *QueryTagListPageResponse
	GetBody() *QueryTagListPageResponseBody
}

type QueryTagListPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagListPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTagListPageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTagListPageResponse) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTagListPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTagListPageResponse) GetBody() *QueryTagListPageResponseBody {
	return s.Body
}

func (s *QueryTagListPageResponse) SetHeaders(v map[string]*string) *QueryTagListPageResponse {
	s.Headers = v
	return s
}

func (s *QueryTagListPageResponse) SetStatusCode(v int32) *QueryTagListPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagListPageResponse) SetBody(v *QueryTagListPageResponseBody) *QueryTagListPageResponse {
	s.Body = v
	return s
}

func (s *QueryTagListPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
