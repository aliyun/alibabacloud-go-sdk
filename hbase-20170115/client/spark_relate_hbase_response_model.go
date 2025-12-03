// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkRelateHBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SparkRelateHBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SparkRelateHBaseResponse
	GetStatusCode() *int32
	SetBody(v *SparkRelateHBaseResponseBody) *SparkRelateHBaseResponse
	GetBody() *SparkRelateHBaseResponseBody
}

type SparkRelateHBaseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SparkRelateHBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SparkRelateHBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s SparkRelateHBaseResponse) GoString() string {
	return s.String()
}

func (s *SparkRelateHBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SparkRelateHBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SparkRelateHBaseResponse) GetBody() *SparkRelateHBaseResponseBody {
	return s.Body
}

func (s *SparkRelateHBaseResponse) SetHeaders(v map[string]*string) *SparkRelateHBaseResponse {
	s.Headers = v
	return s
}

func (s *SparkRelateHBaseResponse) SetStatusCode(v int32) *SparkRelateHBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *SparkRelateHBaseResponse) SetBody(v *SparkRelateHBaseResponseBody) *SparkRelateHBaseResponse {
	s.Body = v
	return s
}

func (s *SparkRelateHBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
