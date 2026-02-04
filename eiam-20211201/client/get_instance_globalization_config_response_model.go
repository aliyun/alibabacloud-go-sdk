// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGlobalizationConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceGlobalizationConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceGlobalizationConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceGlobalizationConfigResponseBody) *GetInstanceGlobalizationConfigResponse
	GetBody() *GetInstanceGlobalizationConfigResponseBody
}

type GetInstanceGlobalizationConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceGlobalizationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceGlobalizationConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGlobalizationConfigResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceGlobalizationConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceGlobalizationConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceGlobalizationConfigResponse) GetBody() *GetInstanceGlobalizationConfigResponseBody {
	return s.Body
}

func (s *GetInstanceGlobalizationConfigResponse) SetHeaders(v map[string]*string) *GetInstanceGlobalizationConfigResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceGlobalizationConfigResponse) SetStatusCode(v int32) *GetInstanceGlobalizationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceGlobalizationConfigResponse) SetBody(v *GetInstanceGlobalizationConfigResponseBody) *GetInstanceGlobalizationConfigResponse {
	s.Body = v
	return s
}

func (s *GetInstanceGlobalizationConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
