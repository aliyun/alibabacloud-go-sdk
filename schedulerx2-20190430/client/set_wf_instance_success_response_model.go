// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWfInstanceSuccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWfInstanceSuccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWfInstanceSuccessResponse
	GetStatusCode() *int32
	SetBody(v *SetWfInstanceSuccessResponseBody) *SetWfInstanceSuccessResponse
	GetBody() *SetWfInstanceSuccessResponseBody
}

type SetWfInstanceSuccessResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWfInstanceSuccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWfInstanceSuccessResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWfInstanceSuccessResponse) GoString() string {
	return s.String()
}

func (s *SetWfInstanceSuccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWfInstanceSuccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWfInstanceSuccessResponse) GetBody() *SetWfInstanceSuccessResponseBody {
	return s.Body
}

func (s *SetWfInstanceSuccessResponse) SetHeaders(v map[string]*string) *SetWfInstanceSuccessResponse {
	s.Headers = v
	return s
}

func (s *SetWfInstanceSuccessResponse) SetStatusCode(v int32) *SetWfInstanceSuccessResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWfInstanceSuccessResponse) SetBody(v *SetWfInstanceSuccessResponseBody) *SetWfInstanceSuccessResponse {
	s.Body = v
	return s
}

func (s *SetWfInstanceSuccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
