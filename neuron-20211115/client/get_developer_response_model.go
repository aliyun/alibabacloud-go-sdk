// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeveloperResponse
	GetStatusCode() *int32
	SetBody(v *Developer) *GetDeveloperResponse
	GetBody() *Developer
}

type GetDeveloperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Developer         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeveloperResponse) GoString() string {
	return s.String()
}

func (s *GetDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeveloperResponse) GetBody() *Developer {
	return s.Body
}

func (s *GetDeveloperResponse) SetHeaders(v map[string]*string) *GetDeveloperResponse {
	s.Headers = v
	return s
}

func (s *GetDeveloperResponse) SetStatusCode(v int32) *GetDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeveloperResponse) SetBody(v *Developer) *GetDeveloperResponse {
	s.Body = v
	return s
}

func (s *GetDeveloperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
