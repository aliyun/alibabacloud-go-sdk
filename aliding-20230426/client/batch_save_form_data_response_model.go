// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveFormDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSaveFormDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSaveFormDataResponse
	GetStatusCode() *int32
	SetBody(v *BatchSaveFormDataResponseBody) *BatchSaveFormDataResponse
	GetBody() *BatchSaveFormDataResponseBody
}

type BatchSaveFormDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSaveFormDataResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveFormDataResponse) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSaveFormDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSaveFormDataResponse) GetBody() *BatchSaveFormDataResponseBody {
	return s.Body
}

func (s *BatchSaveFormDataResponse) SetHeaders(v map[string]*string) *BatchSaveFormDataResponse {
	s.Headers = v
	return s
}

func (s *BatchSaveFormDataResponse) SetStatusCode(v int32) *BatchSaveFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSaveFormDataResponse) SetBody(v *BatchSaveFormDataResponseBody) *BatchSaveFormDataResponse {
	s.Body = v
	return s
}

func (s *BatchSaveFormDataResponse) Validate() error {
	return dara.Validate(s)
}
