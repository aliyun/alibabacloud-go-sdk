// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolicStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumSymbolicStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumSymbolicStackResponse
	GetStatusCode() *int32
	SetBody(v *GetRumSymbolicStackResponseBody) *GetRumSymbolicStackResponse
	GetBody() *GetRumSymbolicStackResponseBody
}

type GetRumSymbolicStackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumSymbolicStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumSymbolicStackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolicStackResponse) GoString() string {
	return s.String()
}

func (s *GetRumSymbolicStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumSymbolicStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumSymbolicStackResponse) GetBody() *GetRumSymbolicStackResponseBody {
	return s.Body
}

func (s *GetRumSymbolicStackResponse) SetHeaders(v map[string]*string) *GetRumSymbolicStackResponse {
	s.Headers = v
	return s
}

func (s *GetRumSymbolicStackResponse) SetStatusCode(v int32) *GetRumSymbolicStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumSymbolicStackResponse) SetBody(v *GetRumSymbolicStackResponseBody) *GetRumSymbolicStackResponse {
	s.Body = v
	return s
}

func (s *GetRumSymbolicStackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
