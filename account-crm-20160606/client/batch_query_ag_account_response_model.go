// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryAgAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryAgAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryAgAccountResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryAgAccountResponseBody) *BatchQueryAgAccountResponse
	GetBody() *BatchQueryAgAccountResponseBody
}

type BatchQueryAgAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryAgAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryAgAccountResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryAgAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryAgAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryAgAccountResponse) GetBody() *BatchQueryAgAccountResponseBody {
	return s.Body
}

func (s *BatchQueryAgAccountResponse) SetHeaders(v map[string]*string) *BatchQueryAgAccountResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryAgAccountResponse) SetStatusCode(v int32) *BatchQueryAgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryAgAccountResponse) SetBody(v *BatchQueryAgAccountResponseBody) *BatchQueryAgAccountResponse {
	s.Body = v
	return s
}

func (s *BatchQueryAgAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
