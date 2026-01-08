// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViberByRequestNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetViberByRequestNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetViberByRequestNoResponse
	GetStatusCode() *int32
	SetBody(v *GetViberByRequestNoResponseBody) *GetViberByRequestNoResponse
	GetBody() *GetViberByRequestNoResponseBody
}

type GetViberByRequestNoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetViberByRequestNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetViberByRequestNoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetViberByRequestNoResponse) GoString() string {
	return s.String()
}

func (s *GetViberByRequestNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetViberByRequestNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetViberByRequestNoResponse) GetBody() *GetViberByRequestNoResponseBody {
	return s.Body
}

func (s *GetViberByRequestNoResponse) SetHeaders(v map[string]*string) *GetViberByRequestNoResponse {
	s.Headers = v
	return s
}

func (s *GetViberByRequestNoResponse) SetStatusCode(v int32) *GetViberByRequestNoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetViberByRequestNoResponse) SetBody(v *GetViberByRequestNoResponseBody) *GetViberByRequestNoResponse {
	s.Body = v
	return s
}

func (s *GetViberByRequestNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
