// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySparkRelateHBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySparkRelateHBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySparkRelateHBaseResponse
	GetStatusCode() *int32
	SetBody(v *QuerySparkRelateHBaseResponseBody) *QuerySparkRelateHBaseResponse
	GetBody() *QuerySparkRelateHBaseResponseBody
}

type QuerySparkRelateHBaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySparkRelateHBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySparkRelateHBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySparkRelateHBaseResponse) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySparkRelateHBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySparkRelateHBaseResponse) GetBody() *QuerySparkRelateHBaseResponseBody {
	return s.Body
}

func (s *QuerySparkRelateHBaseResponse) SetHeaders(v map[string]*string) *QuerySparkRelateHBaseResponse {
	s.Headers = v
	return s
}

func (s *QuerySparkRelateHBaseResponse) SetStatusCode(v int32) *QuerySparkRelateHBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySparkRelateHBaseResponse) SetBody(v *QuerySparkRelateHBaseResponseBody) *QuerySparkRelateHBaseResponse {
	s.Body = v
	return s
}

func (s *QuerySparkRelateHBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
