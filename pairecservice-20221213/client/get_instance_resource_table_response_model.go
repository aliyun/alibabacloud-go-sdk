// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResourceTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceResourceTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceResourceTableResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceResourceTableResponseBody) *GetInstanceResourceTableResponse
	GetBody() *GetInstanceResourceTableResponseBody
}

type GetInstanceResourceTableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResourceTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResourceTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResourceTableResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceResourceTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceResourceTableResponse) GetBody() *GetInstanceResourceTableResponseBody {
	return s.Body
}

func (s *GetInstanceResourceTableResponse) SetHeaders(v map[string]*string) *GetInstanceResourceTableResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResourceTableResponse) SetStatusCode(v int32) *GetInstanceResourceTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResourceTableResponse) SetBody(v *GetInstanceResourceTableResponseBody) *GetInstanceResourceTableResponse {
	s.Body = v
	return s
}

func (s *GetInstanceResourceTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
