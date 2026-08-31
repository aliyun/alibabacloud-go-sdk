// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryColumnarLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryColumnarLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryColumnarLogResponse
	GetStatusCode() *int32
	SetBody(v *QueryColumnarLogResponseBody) *QueryColumnarLogResponse
	GetBody() *QueryColumnarLogResponseBody
}

type QueryColumnarLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryColumnarLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryColumnarLogResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryColumnarLogResponse) GoString() string {
	return s.String()
}

func (s *QueryColumnarLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryColumnarLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryColumnarLogResponse) GetBody() *QueryColumnarLogResponseBody {
	return s.Body
}

func (s *QueryColumnarLogResponse) SetHeaders(v map[string]*string) *QueryColumnarLogResponse {
	s.Headers = v
	return s
}

func (s *QueryColumnarLogResponse) SetStatusCode(v int32) *QueryColumnarLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryColumnarLogResponse) SetBody(v *QueryColumnarLogResponseBody) *QueryColumnarLogResponse {
	s.Body = v
	return s
}

func (s *QueryColumnarLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
