// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceResourceResponseBody) *GetInstanceResourceResponse
	GetBody() *GetInstanceResourceResponseBody
}

type GetInstanceResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceResourceResponse) GetBody() *GetInstanceResourceResponseBody {
	return s.Body
}

func (s *GetInstanceResourceResponse) SetHeaders(v map[string]*string) *GetInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResourceResponse) SetStatusCode(v int32) *GetInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResourceResponse) SetBody(v *GetInstanceResourceResponseBody) *GetInstanceResourceResponse {
	s.Body = v
	return s
}

func (s *GetInstanceResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
