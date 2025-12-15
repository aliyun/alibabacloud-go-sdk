// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceParameterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceParameterResponseBody) *ModifyInstanceParameterResponse
	GetBody() *ModifyInstanceParameterResponseBody
}

type ModifyInstanceParameterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceParameterResponse) GetBody() *ModifyInstanceParameterResponseBody {
	return s.Body
}

func (s *ModifyInstanceParameterResponse) SetHeaders(v map[string]*string) *ModifyInstanceParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceParameterResponse) SetStatusCode(v int32) *ModifyInstanceParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceParameterResponse) SetBody(v *ModifyInstanceParameterResponseBody) *ModifyInstanceParameterResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
