// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaInfoJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaInfoJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaInfoJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaInfoJobListResponseBody) *QueryMediaInfoJobListResponse
	GetBody() *QueryMediaInfoJobListResponseBody
}

type QueryMediaInfoJobListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaInfoJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaInfoJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaInfoJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaInfoJobListResponse) GetBody() *QueryMediaInfoJobListResponseBody {
	return s.Body
}

func (s *QueryMediaInfoJobListResponse) SetHeaders(v map[string]*string) *QueryMediaInfoJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaInfoJobListResponse) SetStatusCode(v int32) *QueryMediaInfoJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaInfoJobListResponse) SetBody(v *QueryMediaInfoJobListResponseBody) *QueryMediaInfoJobListResponse {
	s.Body = v
	return s
}

func (s *QueryMediaInfoJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
