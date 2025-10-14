// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvWithHighCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteKvWithHighCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteKvWithHighCapacityResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteKvWithHighCapacityResponseBody) *BatchDeleteKvWithHighCapacityResponse
	GetBody() *BatchDeleteKvWithHighCapacityResponseBody
}

type BatchDeleteKvWithHighCapacityResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteKvWithHighCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteKvWithHighCapacityResponse) GetBody() *BatchDeleteKvWithHighCapacityResponseBody {
	return s.Body
}

func (s *BatchDeleteKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *BatchDeleteKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponse) SetStatusCode(v int32) *BatchDeleteKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponse) SetBody(v *BatchDeleteKvWithHighCapacityResponseBody) *BatchDeleteKvWithHighCapacityResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
