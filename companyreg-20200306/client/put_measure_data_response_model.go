// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMeasureDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutMeasureDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutMeasureDataResponse
	GetStatusCode() *int32
	SetBody(v *PutMeasureDataResponseBody) *PutMeasureDataResponse
	GetBody() *PutMeasureDataResponseBody
}

type PutMeasureDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMeasureDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMeasureDataResponse) String() string {
	return dara.Prettify(s)
}

func (s PutMeasureDataResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutMeasureDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutMeasureDataResponse) GetBody() *PutMeasureDataResponseBody {
	return s.Body
}

func (s *PutMeasureDataResponse) SetHeaders(v map[string]*string) *PutMeasureDataResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureDataResponse) SetStatusCode(v int32) *PutMeasureDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureDataResponse) SetBody(v *PutMeasureDataResponseBody) *PutMeasureDataResponse {
	s.Body = v
	return s
}

func (s *PutMeasureDataResponse) Validate() error {
	return dara.Validate(s)
}
