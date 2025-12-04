// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGeneralConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGeneralConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGeneralConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetGeneralConfigResponseBody) *GetGeneralConfigResponse
	GetBody() *GetGeneralConfigResponseBody
}

type GetGeneralConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGeneralConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGeneralConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGeneralConfigResponse) GoString() string {
	return s.String()
}

func (s *GetGeneralConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGeneralConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGeneralConfigResponse) GetBody() *GetGeneralConfigResponseBody {
	return s.Body
}

func (s *GetGeneralConfigResponse) SetHeaders(v map[string]*string) *GetGeneralConfigResponse {
	s.Headers = v
	return s
}

func (s *GetGeneralConfigResponse) SetStatusCode(v int32) *GetGeneralConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGeneralConfigResponse) SetBody(v *GetGeneralConfigResponseBody) *GetGeneralConfigResponse {
	s.Body = v
	return s
}

func (s *GetGeneralConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
