// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateModuleVersionResponseBody) *CreateModuleVersionResponse
	GetBody() *CreateModuleVersionResponseBody
}

type CreateModuleVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateModuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModuleVersionResponse) GetBody() *CreateModuleVersionResponseBody {
	return s.Body
}

func (s *CreateModuleVersionResponse) SetHeaders(v map[string]*string) *CreateModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateModuleVersionResponse) SetStatusCode(v int32) *CreateModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModuleVersionResponse) SetBody(v *CreateModuleVersionResponseBody) *CreateModuleVersionResponse {
	s.Body = v
	return s
}

func (s *CreateModuleVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
