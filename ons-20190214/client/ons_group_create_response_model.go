// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsGroupCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsGroupCreateResponse
	GetStatusCode() *int32
	SetBody(v *OnsGroupCreateResponseBody) *OnsGroupCreateResponse
	GetBody() *OnsGroupCreateResponseBody
}

type OnsGroupCreateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsGroupCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsGroupCreateResponse) GetBody() *OnsGroupCreateResponseBody {
	return s.Body
}

func (s *OnsGroupCreateResponse) SetHeaders(v map[string]*string) *OnsGroupCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupCreateResponse) SetStatusCode(v int32) *OnsGroupCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupCreateResponse) SetBody(v *OnsGroupCreateResponseBody) *OnsGroupCreateResponse {
	s.Body = v
	return s
}

func (s *OnsGroupCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
