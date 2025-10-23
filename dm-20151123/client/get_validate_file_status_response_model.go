// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateFileStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetValidateFileStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetValidateFileStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetValidateFileStatusResponseBody) *GetValidateFileStatusResponse
	GetBody() *GetValidateFileStatusResponseBody
}

type GetValidateFileStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetValidateFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetValidateFileStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetValidateFileStatusResponse) GoString() string {
	return s.String()
}

func (s *GetValidateFileStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetValidateFileStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetValidateFileStatusResponse) GetBody() *GetValidateFileStatusResponseBody {
	return s.Body
}

func (s *GetValidateFileStatusResponse) SetHeaders(v map[string]*string) *GetValidateFileStatusResponse {
	s.Headers = v
	return s
}

func (s *GetValidateFileStatusResponse) SetStatusCode(v int32) *GetValidateFileStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetValidateFileStatusResponse) SetBody(v *GetValidateFileStatusResponseBody) *GetValidateFileStatusResponse {
	s.Body = v
	return s
}

func (s *GetValidateFileStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
