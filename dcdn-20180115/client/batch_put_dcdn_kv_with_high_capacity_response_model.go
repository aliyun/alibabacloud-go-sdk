// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutDcdnKvWithHighCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchPutDcdnKvWithHighCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchPutDcdnKvWithHighCapacityResponse
	GetStatusCode() *int32
	SetBody(v *BatchPutDcdnKvWithHighCapacityResponseBody) *BatchPutDcdnKvWithHighCapacityResponse
	GetBody() *BatchPutDcdnKvWithHighCapacityResponseBody
}

type BatchPutDcdnKvWithHighCapacityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchPutDcdnKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchPutDcdnKvWithHighCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvWithHighCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchPutDcdnKvWithHighCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchPutDcdnKvWithHighCapacityResponse) GetBody() *BatchPutDcdnKvWithHighCapacityResponseBody {
	return s.Body
}

func (s *BatchPutDcdnKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *BatchPutDcdnKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityResponse) SetStatusCode(v int32) *BatchPutDcdnKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityResponse) SetBody(v *BatchPutDcdnKvWithHighCapacityResponseBody) *BatchPutDcdnKvWithHighCapacityResponse {
	s.Body = v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityResponse) Validate() error {
	return dara.Validate(s)
}
