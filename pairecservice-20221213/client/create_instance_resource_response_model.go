// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceResourceResponseBody) *CreateInstanceResourceResponse
	GetBody() *CreateInstanceResourceResponseBody
}

type CreateInstanceResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceResourceResponse) GetBody() *CreateInstanceResourceResponseBody {
	return s.Body
}

func (s *CreateInstanceResourceResponse) SetHeaders(v map[string]*string) *CreateInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResourceResponse) SetStatusCode(v int32) *CreateInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResourceResponse) SetBody(v *CreateInstanceResourceResponseBody) *CreateInstanceResourceResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
