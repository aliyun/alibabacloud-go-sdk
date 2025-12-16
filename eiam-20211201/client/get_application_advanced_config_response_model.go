// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationAdvancedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationAdvancedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationAdvancedConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationAdvancedConfigResponseBody) *GetApplicationAdvancedConfigResponse
	GetBody() *GetApplicationAdvancedConfigResponseBody
}

type GetApplicationAdvancedConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationAdvancedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationAdvancedConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationAdvancedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationAdvancedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationAdvancedConfigResponse) GetBody() *GetApplicationAdvancedConfigResponseBody {
	return s.Body
}

func (s *GetApplicationAdvancedConfigResponse) SetHeaders(v map[string]*string) *GetApplicationAdvancedConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationAdvancedConfigResponse) SetStatusCode(v int32) *GetApplicationAdvancedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationAdvancedConfigResponse) SetBody(v *GetApplicationAdvancedConfigResponseBody) *GetApplicationAdvancedConfigResponse {
	s.Body = v
	return s
}

func (s *GetApplicationAdvancedConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
