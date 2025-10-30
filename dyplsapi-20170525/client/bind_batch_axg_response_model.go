// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindBatchAxgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindBatchAxgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindBatchAxgResponse
	GetStatusCode() *int32
	SetBody(v *BindBatchAxgResponseBody) *BindBatchAxgResponse
	GetBody() *BindBatchAxgResponseBody
}

type BindBatchAxgResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindBatchAxgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindBatchAxgResponse) String() string {
	return dara.Prettify(s)
}

func (s BindBatchAxgResponse) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindBatchAxgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindBatchAxgResponse) GetBody() *BindBatchAxgResponseBody {
	return s.Body
}

func (s *BindBatchAxgResponse) SetHeaders(v map[string]*string) *BindBatchAxgResponse {
	s.Headers = v
	return s
}

func (s *BindBatchAxgResponse) SetStatusCode(v int32) *BindBatchAxgResponse {
	s.StatusCode = &v
	return s
}

func (s *BindBatchAxgResponse) SetBody(v *BindBatchAxgResponseBody) *BindBatchAxgResponse {
	s.Body = v
	return s
}

func (s *BindBatchAxgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
