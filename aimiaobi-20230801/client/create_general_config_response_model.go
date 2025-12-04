// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGeneralConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGeneralConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGeneralConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateGeneralConfigResponseBody) *CreateGeneralConfigResponse
	GetBody() *CreateGeneralConfigResponseBody
}

type CreateGeneralConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGeneralConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGeneralConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneralConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateGeneralConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGeneralConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGeneralConfigResponse) GetBody() *CreateGeneralConfigResponseBody {
	return s.Body
}

func (s *CreateGeneralConfigResponse) SetHeaders(v map[string]*string) *CreateGeneralConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateGeneralConfigResponse) SetStatusCode(v int32) *CreateGeneralConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGeneralConfigResponse) SetBody(v *CreateGeneralConfigResponseBody) *CreateGeneralConfigResponse {
	s.Body = v
	return s
}

func (s *CreateGeneralConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
