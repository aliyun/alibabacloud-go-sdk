// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdleInstanceCullerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdleInstanceCullerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdleInstanceCullerResponse
	GetStatusCode() *int32
	SetBody(v *GetIdleInstanceCullerResponseBody) *GetIdleInstanceCullerResponse
	GetBody() *GetIdleInstanceCullerResponseBody
}

type GetIdleInstanceCullerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdleInstanceCullerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdleInstanceCullerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdleInstanceCullerResponse) GoString() string {
	return s.String()
}

func (s *GetIdleInstanceCullerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdleInstanceCullerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdleInstanceCullerResponse) GetBody() *GetIdleInstanceCullerResponseBody {
	return s.Body
}

func (s *GetIdleInstanceCullerResponse) SetHeaders(v map[string]*string) *GetIdleInstanceCullerResponse {
	s.Headers = v
	return s
}

func (s *GetIdleInstanceCullerResponse) SetStatusCode(v int32) *GetIdleInstanceCullerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdleInstanceCullerResponse) SetBody(v *GetIdleInstanceCullerResponseBody) *GetIdleInstanceCullerResponse {
	s.Body = v
	return s
}

func (s *GetIdleInstanceCullerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
