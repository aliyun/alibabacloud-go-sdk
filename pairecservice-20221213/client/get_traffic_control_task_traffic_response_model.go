// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTaskTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrafficControlTaskTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrafficControlTaskTrafficResponse
	GetStatusCode() *int32
	SetBody(v *GetTrafficControlTaskTrafficResponseBody) *GetTrafficControlTaskTrafficResponse
	GetBody() *GetTrafficControlTaskTrafficResponseBody
}

type GetTrafficControlTaskTrafficResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrafficControlTaskTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrafficControlTaskTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrafficControlTaskTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrafficControlTaskTrafficResponse) GetBody() *GetTrafficControlTaskTrafficResponseBody {
	return s.Body
}

func (s *GetTrafficControlTaskTrafficResponse) SetHeaders(v map[string]*string) *GetTrafficControlTaskTrafficResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponse) SetStatusCode(v int32) *GetTrafficControlTaskTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrafficControlTaskTrafficResponse) SetBody(v *GetTrafficControlTaskTrafficResponseBody) *GetTrafficControlTaskTrafficResponse {
	s.Body = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
