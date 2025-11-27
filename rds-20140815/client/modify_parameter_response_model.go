// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyParameterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyParameterResponseBody) *ModifyParameterResponse
	GetBody() *ModifyParameterResponseBody
}

type ModifyParameterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyParameterResponse) GetBody() *ModifyParameterResponseBody {
	return s.Body
}

func (s *ModifyParameterResponse) SetHeaders(v map[string]*string) *ModifyParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyParameterResponse) SetStatusCode(v int32) *ModifyParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParameterResponse) SetBody(v *ModifyParameterResponseBody) *ModifyParameterResponse {
	s.Body = v
	return s
}

func (s *ModifyParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
