// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseBdrcServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseBdrcServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseBdrcServiceResponse
	GetStatusCode() *int32
	SetBody(v *CloseBdrcServiceResponseBody) *CloseBdrcServiceResponse
	GetBody() *CloseBdrcServiceResponseBody
}

type CloseBdrcServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseBdrcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseBdrcServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseBdrcServiceResponse) GoString() string {
	return s.String()
}

func (s *CloseBdrcServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseBdrcServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseBdrcServiceResponse) GetBody() *CloseBdrcServiceResponseBody {
	return s.Body
}

func (s *CloseBdrcServiceResponse) SetHeaders(v map[string]*string) *CloseBdrcServiceResponse {
	s.Headers = v
	return s
}

func (s *CloseBdrcServiceResponse) SetStatusCode(v int32) *CloseBdrcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseBdrcServiceResponse) SetBody(v *CloseBdrcServiceResponseBody) *CloseBdrcServiceResponse {
	s.Body = v
	return s
}

func (s *CloseBdrcServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
