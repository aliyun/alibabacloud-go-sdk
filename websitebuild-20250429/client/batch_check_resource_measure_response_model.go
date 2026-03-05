// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCheckResourceMeasureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCheckResourceMeasureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCheckResourceMeasureResponse
	GetStatusCode() *int32
	SetBody(v *BatchCheckResourceMeasureResponseBody) *BatchCheckResourceMeasureResponse
	GetBody() *BatchCheckResourceMeasureResponseBody
}

type BatchCheckResourceMeasureResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCheckResourceMeasureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCheckResourceMeasureResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCheckResourceMeasureResponse) GoString() string {
	return s.String()
}

func (s *BatchCheckResourceMeasureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCheckResourceMeasureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCheckResourceMeasureResponse) GetBody() *BatchCheckResourceMeasureResponseBody {
	return s.Body
}

func (s *BatchCheckResourceMeasureResponse) SetHeaders(v map[string]*string) *BatchCheckResourceMeasureResponse {
	s.Headers = v
	return s
}

func (s *BatchCheckResourceMeasureResponse) SetStatusCode(v int32) *BatchCheckResourceMeasureResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCheckResourceMeasureResponse) SetBody(v *BatchCheckResourceMeasureResponseBody) *BatchCheckResourceMeasureResponse {
	s.Body = v
	return s
}

func (s *BatchCheckResourceMeasureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
