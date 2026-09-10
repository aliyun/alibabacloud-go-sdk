// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataCheckConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataCheckConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddDataCheckConfigResponseBody) *AddDataCheckConfigResponse
	GetBody() *AddDataCheckConfigResponseBody
}

type AddDataCheckConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataCheckConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *AddDataCheckConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataCheckConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataCheckConfigResponse) GetBody() *AddDataCheckConfigResponseBody {
	return s.Body
}

func (s *AddDataCheckConfigResponse) SetHeaders(v map[string]*string) *AddDataCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *AddDataCheckConfigResponse) SetStatusCode(v int32) *AddDataCheckConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataCheckConfigResponse) SetBody(v *AddDataCheckConfigResponseBody) *AddDataCheckConfigResponse {
	s.Body = v
	return s
}

func (s *AddDataCheckConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
