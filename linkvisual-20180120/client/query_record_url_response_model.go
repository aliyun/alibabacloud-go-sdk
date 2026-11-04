// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordUrlResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordUrlResponseBody) *QueryRecordUrlResponse
	GetBody() *QueryRecordUrlResponseBody
}

type QueryRecordUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordUrlResponse) GetBody() *QueryRecordUrlResponseBody {
	return s.Body
}

func (s *QueryRecordUrlResponse) SetHeaders(v map[string]*string) *QueryRecordUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordUrlResponse) SetStatusCode(v int32) *QueryRecordUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordUrlResponse) SetBody(v *QueryRecordUrlResponseBody) *QueryRecordUrlResponse {
	s.Body = v
	return s
}

func (s *QueryRecordUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
