// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *Enterprise) *CreateEnterpriseResponse
	GetBody() *Enterprise
}

type CreateEnterpriseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Enterprise        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnterpriseResponse) GetBody() *Enterprise {
	return s.Body
}

func (s *CreateEnterpriseResponse) SetHeaders(v map[string]*string) *CreateEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseResponse) SetStatusCode(v int32) *CreateEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnterpriseResponse) SetBody(v *Enterprise) *CreateEnterpriseResponse {
	s.Body = v
	return s
}

func (s *CreateEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
