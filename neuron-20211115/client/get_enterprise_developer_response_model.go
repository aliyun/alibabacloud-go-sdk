// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEnterpriseDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEnterpriseDeveloperResponse
	GetStatusCode() *int32
	SetBody(v *Developer) *GetEnterpriseDeveloperResponse
	GetBody() *Developer
}

type GetEnterpriseDeveloperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Developer         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnterpriseDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseDeveloperResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEnterpriseDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEnterpriseDeveloperResponse) GetBody() *Developer {
	return s.Body
}

func (s *GetEnterpriseDeveloperResponse) SetHeaders(v map[string]*string) *GetEnterpriseDeveloperResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDeveloperResponse) SetStatusCode(v int32) *GetEnterpriseDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseDeveloperResponse) SetBody(v *Developer) *GetEnterpriseDeveloperResponse {
	s.Body = v
	return s
}

func (s *GetEnterpriseDeveloperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
