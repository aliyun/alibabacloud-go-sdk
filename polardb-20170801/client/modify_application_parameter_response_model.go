// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationParameterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationParameterResponseBody) *ModifyApplicationParameterResponse
	GetBody() *ModifyApplicationParameterResponseBody
}

type ModifyApplicationParameterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationParameterResponse) GetBody() *ModifyApplicationParameterResponseBody {
	return s.Body
}

func (s *ModifyApplicationParameterResponse) SetHeaders(v map[string]*string) *ModifyApplicationParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationParameterResponse) SetStatusCode(v int32) *ModifyApplicationParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationParameterResponse) SetBody(v *ModifyApplicationParameterResponseBody) *ModifyApplicationParameterResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
