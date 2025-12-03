// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVServerGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVServerGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVServerGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetVServerGroupAttributeResponseBody) *SetVServerGroupAttributeResponse
	GetBody() *SetVServerGroupAttributeResponseBody
}

type SetVServerGroupAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVServerGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVServerGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVServerGroupAttributeResponse) GetBody() *SetVServerGroupAttributeResponseBody {
	return s.Body
}

func (s *SetVServerGroupAttributeResponse) SetHeaders(v map[string]*string) *SetVServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetVServerGroupAttributeResponse) SetStatusCode(v int32) *SetVServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVServerGroupAttributeResponse) SetBody(v *SetVServerGroupAttributeResponseBody) *SetVServerGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *SetVServerGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
