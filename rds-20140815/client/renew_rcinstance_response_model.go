// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewRCInstanceResponseBody) *RenewRCInstanceResponse
	GetBody() *RenewRCInstanceResponseBody
}

type RenewRCInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewRCInstanceResponse) GetBody() *RenewRCInstanceResponseBody {
	return s.Body
}

func (s *RenewRCInstanceResponse) SetHeaders(v map[string]*string) *RenewRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewRCInstanceResponse) SetStatusCode(v int32) *RenewRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewRCInstanceResponse) SetBody(v *RenewRCInstanceResponseBody) *RenewRCInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
