// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPictureSearchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPictureSearchJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryPictureSearchJobResponseBody) *QueryPictureSearchJobResponse
	GetBody() *QueryPictureSearchJobResponseBody
}

type QueryPictureSearchJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPictureSearchJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPictureSearchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchJobResponse) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPictureSearchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPictureSearchJobResponse) GetBody() *QueryPictureSearchJobResponseBody {
	return s.Body
}

func (s *QueryPictureSearchJobResponse) SetHeaders(v map[string]*string) *QueryPictureSearchJobResponse {
	s.Headers = v
	return s
}

func (s *QueryPictureSearchJobResponse) SetStatusCode(v int32) *QueryPictureSearchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPictureSearchJobResponse) SetBody(v *QueryPictureSearchJobResponseBody) *QueryPictureSearchJobResponse {
	s.Body = v
	return s
}

func (s *QueryPictureSearchJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
