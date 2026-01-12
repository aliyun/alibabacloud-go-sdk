// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteSvcMapBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteSvcMapBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteSvcMapBindResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteSvcMapBindResponseBody) *BatchDeleteSvcMapBindResponse
	GetBody() *BatchDeleteSvcMapBindResponseBody
}

type BatchDeleteSvcMapBindResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteSvcMapBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteSvcMapBindResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteSvcMapBindResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteSvcMapBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteSvcMapBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteSvcMapBindResponse) GetBody() *BatchDeleteSvcMapBindResponseBody {
	return s.Body
}

func (s *BatchDeleteSvcMapBindResponse) SetHeaders(v map[string]*string) *BatchDeleteSvcMapBindResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteSvcMapBindResponse) SetStatusCode(v int32) *BatchDeleteSvcMapBindResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteSvcMapBindResponse) SetBody(v *BatchDeleteSvcMapBindResponseBody) *BatchDeleteSvcMapBindResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteSvcMapBindResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
