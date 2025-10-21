// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyDefaultOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyDefaultOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyDefaultOptionsResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyDefaultOptionsResponseBody) *GetPolicyDefaultOptionsResponse
	GetBody() *GetPolicyDefaultOptionsResponseBody
}

type GetPolicyDefaultOptionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyDefaultOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyDefaultOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyDefaultOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyDefaultOptionsResponse) GetBody() *GetPolicyDefaultOptionsResponseBody {
	return s.Body
}

func (s *GetPolicyDefaultOptionsResponse) SetHeaders(v map[string]*string) *GetPolicyDefaultOptionsResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyDefaultOptionsResponse) SetStatusCode(v int32) *GetPolicyDefaultOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyDefaultOptionsResponse) SetBody(v *GetPolicyDefaultOptionsResponseBody) *GetPolicyDefaultOptionsResponse {
	s.Body = v
	return s
}

func (s *GetPolicyDefaultOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
