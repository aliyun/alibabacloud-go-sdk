// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApDetailedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApDetailedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApDetailedConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApDetailedConfigResponseBody) *GetApDetailedConfigResponse
	GetBody() *GetApDetailedConfigResponseBody
}

type GetApDetailedConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApDetailedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApDetailedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApDetailedConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApDetailedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApDetailedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApDetailedConfigResponse) GetBody() *GetApDetailedConfigResponseBody {
	return s.Body
}

func (s *GetApDetailedConfigResponse) SetHeaders(v map[string]*string) *GetApDetailedConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApDetailedConfigResponse) SetStatusCode(v int32) *GetApDetailedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApDetailedConfigResponse) SetBody(v *GetApDetailedConfigResponseBody) *GetApDetailedConfigResponse {
	s.Body = v
	return s
}

func (s *GetApDetailedConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
