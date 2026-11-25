// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutWarmPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutWarmPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutWarmPoolResponse
	GetStatusCode() *int32
	SetBody(v *PutWarmPoolResponseBody) *PutWarmPoolResponse
	GetBody() *PutWarmPoolResponseBody
}

type PutWarmPoolResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutWarmPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutWarmPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s PutWarmPoolResponse) GoString() string {
	return s.String()
}

func (s *PutWarmPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutWarmPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutWarmPoolResponse) GetBody() *PutWarmPoolResponseBody {
	return s.Body
}

func (s *PutWarmPoolResponse) SetHeaders(v map[string]*string) *PutWarmPoolResponse {
	s.Headers = v
	return s
}

func (s *PutWarmPoolResponse) SetStatusCode(v int32) *PutWarmPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *PutWarmPoolResponse) SetBody(v *PutWarmPoolResponseBody) *PutWarmPoolResponse {
	s.Body = v
	return s
}

func (s *PutWarmPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
