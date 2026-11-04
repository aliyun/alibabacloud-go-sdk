// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStreamPushJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryStreamPushJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryStreamPushJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryStreamPushJobListResponseBody) *QueryStreamPushJobListResponse
	GetBody() *QueryStreamPushJobListResponseBody
}

type QueryStreamPushJobListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStreamPushJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStreamPushJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryStreamPushJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryStreamPushJobListResponse) GetBody() *QueryStreamPushJobListResponseBody {
	return s.Body
}

func (s *QueryStreamPushJobListResponse) SetHeaders(v map[string]*string) *QueryStreamPushJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryStreamPushJobListResponse) SetStatusCode(v int32) *QueryStreamPushJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStreamPushJobListResponse) SetBody(v *QueryStreamPushJobListResponseBody) *QueryStreamPushJobListResponse {
	s.Body = v
	return s
}

func (s *QueryStreamPushJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
