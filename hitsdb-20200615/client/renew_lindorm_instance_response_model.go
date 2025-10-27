// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewLindormInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewLindormInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewLindormInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewLindormInstanceResponseBody) *RenewLindormInstanceResponse
	GetBody() *RenewLindormInstanceResponseBody
}

type RenewLindormInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewLindormInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewLindormInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewLindormInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewLindormInstanceResponse) GetBody() *RenewLindormInstanceResponseBody {
	return s.Body
}

func (s *RenewLindormInstanceResponse) SetHeaders(v map[string]*string) *RenewLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewLindormInstanceResponse) SetStatusCode(v int32) *RenewLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewLindormInstanceResponse) SetBody(v *RenewLindormInstanceResponseBody) *RenewLindormInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewLindormInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
