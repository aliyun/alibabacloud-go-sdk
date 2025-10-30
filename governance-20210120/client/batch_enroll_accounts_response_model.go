// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEnrollAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchEnrollAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchEnrollAccountsResponse
	GetStatusCode() *int32
	SetBody(v *BatchEnrollAccountsResponseBody) *BatchEnrollAccountsResponse
	GetBody() *BatchEnrollAccountsResponseBody
}

type BatchEnrollAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchEnrollAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchEnrollAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchEnrollAccountsResponse) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchEnrollAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchEnrollAccountsResponse) GetBody() *BatchEnrollAccountsResponseBody {
	return s.Body
}

func (s *BatchEnrollAccountsResponse) SetHeaders(v map[string]*string) *BatchEnrollAccountsResponse {
	s.Headers = v
	return s
}

func (s *BatchEnrollAccountsResponse) SetStatusCode(v int32) *BatchEnrollAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchEnrollAccountsResponse) SetBody(v *BatchEnrollAccountsResponseBody) *BatchEnrollAccountsResponse {
	s.Body = v
	return s
}

func (s *BatchEnrollAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
