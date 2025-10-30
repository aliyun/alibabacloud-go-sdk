// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceResponseBody) *GetResourceResponse
	GetBody() *GetResourceResponseBody
}

type GetResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceResponse) GetBody() *GetResourceResponseBody {
	return s.Body
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetStatusCode(v int32) *GetResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

func (s *GetResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
