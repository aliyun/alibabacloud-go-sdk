// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaCensorJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaCensorJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaCensorJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaCensorJobListResponseBody) *QueryMediaCensorJobListResponse
	GetBody() *QueryMediaCensorJobListResponseBody
}

type QueryMediaCensorJobListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaCensorJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaCensorJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaCensorJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaCensorJobListResponse) GetBody() *QueryMediaCensorJobListResponseBody {
	return s.Body
}

func (s *QueryMediaCensorJobListResponse) SetHeaders(v map[string]*string) *QueryMediaCensorJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaCensorJobListResponse) SetStatusCode(v int32) *QueryMediaCensorJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaCensorJobListResponse) SetBody(v *QueryMediaCensorJobListResponseBody) *QueryMediaCensorJobListResponse {
	s.Body = v
	return s
}

func (s *QueryMediaCensorJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
