// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteSampleDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteSampleDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteSampleDataResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteSampleDataResponseBody) *BatchDeleteSampleDataResponse
	GetBody() *BatchDeleteSampleDataResponseBody
}

type BatchDeleteSampleDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteSampleDataResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteSampleDataResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteSampleDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteSampleDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteSampleDataResponse) GetBody() *BatchDeleteSampleDataResponseBody {
	return s.Body
}

func (s *BatchDeleteSampleDataResponse) SetHeaders(v map[string]*string) *BatchDeleteSampleDataResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteSampleDataResponse) SetStatusCode(v int32) *BatchDeleteSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteSampleDataResponse) SetBody(v *BatchDeleteSampleDataResponseBody) *BatchDeleteSampleDataResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteSampleDataResponse) Validate() error {
	return dara.Validate(s)
}
