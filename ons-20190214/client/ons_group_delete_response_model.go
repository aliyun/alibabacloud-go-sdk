// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsGroupDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsGroupDeleteResponse
	GetStatusCode() *int32
	SetBody(v *OnsGroupDeleteResponseBody) *OnsGroupDeleteResponse
	GetBody() *OnsGroupDeleteResponseBody
}

type OnsGroupDeleteResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsGroupDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsGroupDeleteResponse) GetBody() *OnsGroupDeleteResponseBody {
	return s.Body
}

func (s *OnsGroupDeleteResponse) SetHeaders(v map[string]*string) *OnsGroupDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupDeleteResponse) SetStatusCode(v int32) *OnsGroupDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupDeleteResponse) SetBody(v *OnsGroupDeleteResponseBody) *OnsGroupDeleteResponse {
	s.Body = v
	return s
}

func (s *OnsGroupDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
