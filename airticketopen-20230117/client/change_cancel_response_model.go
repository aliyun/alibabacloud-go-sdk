// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeCancelResponse
	GetStatusCode() *int32
	SetBody(v *ChangeCancelResponseBody) *ChangeCancelResponse
	GetBody() *ChangeCancelResponseBody
}

type ChangeCancelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeCancelResponse) GoString() string {
	return s.String()
}

func (s *ChangeCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeCancelResponse) GetBody() *ChangeCancelResponseBody {
	return s.Body
}

func (s *ChangeCancelResponse) SetHeaders(v map[string]*string) *ChangeCancelResponse {
	s.Headers = v
	return s
}

func (s *ChangeCancelResponse) SetStatusCode(v int32) *ChangeCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCancelResponse) SetBody(v *ChangeCancelResponseBody) *ChangeCancelResponse {
	s.Body = v
	return s
}

func (s *ChangeCancelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
