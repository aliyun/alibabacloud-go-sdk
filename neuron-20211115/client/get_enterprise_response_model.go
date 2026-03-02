// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *Enterprise) *GetEnterpriseResponse
	GetBody() *Enterprise
}

type GetEnterpriseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Enterprise        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEnterpriseResponse) GetBody() *Enterprise {
	return s.Body
}

func (s *GetEnterpriseResponse) SetHeaders(v map[string]*string) *GetEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseResponse) SetStatusCode(v int32) *GetEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseResponse) SetBody(v *Enterprise) *GetEnterpriseResponse {
	s.Body = v
	return s
}

func (s *GetEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
