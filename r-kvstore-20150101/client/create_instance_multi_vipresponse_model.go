// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceMultiVIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceMultiVIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceMultiVIPResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceMultiVIPResponseBody) *CreateInstanceMultiVIPResponse
	GetBody() *CreateInstanceMultiVIPResponseBody
}

type CreateInstanceMultiVIPResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceMultiVIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceMultiVIPResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceMultiVIPResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceMultiVIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceMultiVIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceMultiVIPResponse) GetBody() *CreateInstanceMultiVIPResponseBody {
	return s.Body
}

func (s *CreateInstanceMultiVIPResponse) SetHeaders(v map[string]*string) *CreateInstanceMultiVIPResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceMultiVIPResponse) SetStatusCode(v int32) *CreateInstanceMultiVIPResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceMultiVIPResponse) SetBody(v *CreateInstanceMultiVIPResponseBody) *CreateInstanceMultiVIPResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceMultiVIPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
