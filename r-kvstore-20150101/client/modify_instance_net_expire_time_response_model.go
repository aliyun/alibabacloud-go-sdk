// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceNetExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceNetExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceNetExpireTimeResponseBody) *ModifyInstanceNetExpireTimeResponse
	GetBody() *ModifyInstanceNetExpireTimeResponseBody
}

type ModifyInstanceNetExpireTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceNetExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceNetExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceNetExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceNetExpireTimeResponse) GetBody() *ModifyInstanceNetExpireTimeResponseBody {
	return s.Body
}

func (s *ModifyInstanceNetExpireTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceNetExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponse) SetStatusCode(v int32) *ModifyInstanceNetExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponse) SetBody(v *ModifyInstanceNetExpireTimeResponseBody) *ModifyInstanceNetExpireTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
