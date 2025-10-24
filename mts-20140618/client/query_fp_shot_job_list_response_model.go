// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpShotJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFpShotJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFpShotJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryFpShotJobListResponseBody) *QueryFpShotJobListResponse
	GetBody() *QueryFpShotJobListResponseBody
}

type QueryFpShotJobListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFpShotJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFpShotJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFpShotJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFpShotJobListResponse) GetBody() *QueryFpShotJobListResponseBody {
	return s.Body
}

func (s *QueryFpShotJobListResponse) SetHeaders(v map[string]*string) *QueryFpShotJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryFpShotJobListResponse) SetStatusCode(v int32) *QueryFpShotJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFpShotJobListResponse) SetBody(v *QueryFpShotJobListResponseBody) *QueryFpShotJobListResponse {
	s.Body = v
	return s
}

func (s *QueryFpShotJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
