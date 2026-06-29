// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRevokeSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchRevokeSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchRevokeSeatsResponse
	GetStatusCode() *int32
	SetBody(v *BatchRevokeSeatsResponseBody) *BatchRevokeSeatsResponse
	GetBody() *BatchRevokeSeatsResponseBody
}

type BatchRevokeSeatsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRevokeSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRevokeSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokeSeatsResponse) GoString() string {
	return s.String()
}

func (s *BatchRevokeSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRevokeSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchRevokeSeatsResponse) GetBody() *BatchRevokeSeatsResponseBody {
	return s.Body
}

func (s *BatchRevokeSeatsResponse) SetHeaders(v map[string]*string) *BatchRevokeSeatsResponse {
	s.Headers = v
	return s
}

func (s *BatchRevokeSeatsResponse) SetStatusCode(v int32) *BatchRevokeSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRevokeSeatsResponse) SetBody(v *BatchRevokeSeatsResponseBody) *BatchRevokeSeatsResponse {
	s.Body = v
	return s
}

func (s *BatchRevokeSeatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
