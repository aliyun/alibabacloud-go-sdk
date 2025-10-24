// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpDBDeleteJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFpDBDeleteJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFpDBDeleteJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryFpDBDeleteJobListResponseBody) *QueryFpDBDeleteJobListResponse
	GetBody() *QueryFpDBDeleteJobListResponseBody
}

type QueryFpDBDeleteJobListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFpDBDeleteJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFpDBDeleteJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFpDBDeleteJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryFpDBDeleteJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFpDBDeleteJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFpDBDeleteJobListResponse) GetBody() *QueryFpDBDeleteJobListResponseBody {
	return s.Body
}

func (s *QueryFpDBDeleteJobListResponse) SetHeaders(v map[string]*string) *QueryFpDBDeleteJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryFpDBDeleteJobListResponse) SetStatusCode(v int32) *QueryFpDBDeleteJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponse) SetBody(v *QueryFpDBDeleteJobListResponseBody) *QueryFpDBDeleteJobListResponse {
	s.Body = v
	return s
}

func (s *QueryFpDBDeleteJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
