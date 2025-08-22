// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnKvWithHighCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteDcdnKvWithHighCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteDcdnKvWithHighCapacityResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteDcdnKvWithHighCapacityResponseBody) *BatchDeleteDcdnKvWithHighCapacityResponse
	GetBody() *BatchDeleteDcdnKvWithHighCapacityResponseBody
}

type BatchDeleteDcdnKvWithHighCapacityResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteDcdnKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteDcdnKvWithHighCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponse) GetBody() *BatchDeleteDcdnKvWithHighCapacityResponseBody {
	return s.Body
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *BatchDeleteDcdnKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponse) SetStatusCode(v int32) *BatchDeleteDcdnKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponse) SetBody(v *BatchDeleteDcdnKvWithHighCapacityResponseBody) *BatchDeleteDcdnKvWithHighCapacityResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponse) Validate() error {
	return dara.Validate(s)
}
