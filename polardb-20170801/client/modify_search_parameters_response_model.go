// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearchParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySearchParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySearchParametersResponse
	GetStatusCode() *int32
	SetBody(v *ModifySearchParametersResponseBody) *ModifySearchParametersResponse
	GetBody() *ModifySearchParametersResponseBody
}

type ModifySearchParametersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySearchParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySearchParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySearchParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifySearchParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySearchParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySearchParametersResponse) GetBody() *ModifySearchParametersResponseBody {
	return s.Body
}

func (s *ModifySearchParametersResponse) SetHeaders(v map[string]*string) *ModifySearchParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifySearchParametersResponse) SetStatusCode(v int32) *ModifySearchParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySearchParametersResponse) SetBody(v *ModifySearchParametersResponseBody) *ModifySearchParametersResponse {
	s.Body = v
	return s
}

func (s *ModifySearchParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
