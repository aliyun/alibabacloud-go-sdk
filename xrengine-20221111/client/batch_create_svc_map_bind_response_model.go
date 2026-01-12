// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateSvcMapBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateSvcMapBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateSvcMapBindResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateSvcMapBindResponseBody) *BatchCreateSvcMapBindResponse
	GetBody() *BatchCreateSvcMapBindResponseBody
}

type BatchCreateSvcMapBindResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateSvcMapBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateSvcMapBindResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateSvcMapBindResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateSvcMapBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateSvcMapBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateSvcMapBindResponse) GetBody() *BatchCreateSvcMapBindResponseBody {
	return s.Body
}

func (s *BatchCreateSvcMapBindResponse) SetHeaders(v map[string]*string) *BatchCreateSvcMapBindResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateSvcMapBindResponse) SetStatusCode(v int32) *BatchCreateSvcMapBindResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateSvcMapBindResponse) SetBody(v *BatchCreateSvcMapBindResponseBody) *BatchCreateSvcMapBindResponse {
	s.Body = v
	return s
}

func (s *BatchCreateSvcMapBindResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
