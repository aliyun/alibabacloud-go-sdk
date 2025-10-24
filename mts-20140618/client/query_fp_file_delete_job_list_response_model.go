// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpFileDeleteJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFpFileDeleteJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFpFileDeleteJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryFpFileDeleteJobListResponseBody) *QueryFpFileDeleteJobListResponse
	GetBody() *QueryFpFileDeleteJobListResponseBody
}

type QueryFpFileDeleteJobListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFpFileDeleteJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFpFileDeleteJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFpFileDeleteJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryFpFileDeleteJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFpFileDeleteJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFpFileDeleteJobListResponse) GetBody() *QueryFpFileDeleteJobListResponseBody {
	return s.Body
}

func (s *QueryFpFileDeleteJobListResponse) SetHeaders(v map[string]*string) *QueryFpFileDeleteJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryFpFileDeleteJobListResponse) SetStatusCode(v int32) *QueryFpFileDeleteJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponse) SetBody(v *QueryFpFileDeleteJobListResponseBody) *QueryFpFileDeleteJobListResponse {
	s.Body = v
	return s
}

func (s *QueryFpFileDeleteJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
