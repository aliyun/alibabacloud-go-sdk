// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyParametersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse
	GetBody() *ModifyParametersResponseBody
}

type ModifyParametersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyParametersResponse) GetBody() *ModifyParametersResponseBody {
	return s.Body
}

func (s *ModifyParametersResponse) SetHeaders(v map[string]*string) *ModifyParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyParametersResponse) SetStatusCode(v int32) *ModifyParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParametersResponse) SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse {
	s.Body = v
	return s
}

func (s *ModifyParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
