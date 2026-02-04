// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceGlobalizationConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetInstanceGlobalizationConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetInstanceGlobalizationConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetInstanceGlobalizationConfigResponseBody) *SetInstanceGlobalizationConfigResponse
	GetBody() *SetInstanceGlobalizationConfigResponseBody
}

type SetInstanceGlobalizationConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInstanceGlobalizationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInstanceGlobalizationConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceGlobalizationConfigResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceGlobalizationConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetInstanceGlobalizationConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetInstanceGlobalizationConfigResponse) GetBody() *SetInstanceGlobalizationConfigResponseBody {
	return s.Body
}

func (s *SetInstanceGlobalizationConfigResponse) SetHeaders(v map[string]*string) *SetInstanceGlobalizationConfigResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceGlobalizationConfigResponse) SetStatusCode(v int32) *SetInstanceGlobalizationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceGlobalizationConfigResponse) SetBody(v *SetInstanceGlobalizationConfigResponseBody) *SetInstanceGlobalizationConfigResponse {
	s.Body = v
	return s
}

func (s *SetInstanceGlobalizationConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
