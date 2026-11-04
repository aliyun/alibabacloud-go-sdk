// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchJobResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPictureSearchJobResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPictureSearchJobResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryPictureSearchJobResultResponseBody) *QueryPictureSearchJobResultResponse
	GetBody() *QueryPictureSearchJobResultResponseBody
}

type QueryPictureSearchJobResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchJobResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchJobResultResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPictureSearchJobResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPictureSearchJobResultResponse) GetBody() *QueryPictureSearchJobResultResponseBody {
	return s.Body
}

func (s *QueryPictureSearchJobResultResponse) SetHeaders(v map[string]*string) *QueryPictureSearchJobResultResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchJobResultResponse) SetStatusCode(v int32) *QueryPictureSearchJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchJobResultResponse) SetBody(v *QueryPictureSearchJobResultResponseBody) *QueryPictureSearchJobResultResponse {
	s.Body = v
	return s
}

func (s *QueryPictureSearchJobResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
