// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPropertyValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetUserPropertyValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetUserPropertyValueResponse
	GetStatusCode() *int32
	SetBody(v *SetUserPropertyValueResponseBody) *SetUserPropertyValueResponse
	GetBody() *SetUserPropertyValueResponseBody
}

type SetUserPropertyValueResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserPropertyValueResponse) String() string {
	return dara.Prettify(s)
}

func (s SetUserPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetUserPropertyValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetUserPropertyValueResponse) GetBody() *SetUserPropertyValueResponseBody {
	return s.Body
}

func (s *SetUserPropertyValueResponse) SetHeaders(v map[string]*string) *SetUserPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *SetUserPropertyValueResponse) SetStatusCode(v int32) *SetUserPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserPropertyValueResponse) SetBody(v *SetUserPropertyValueResponseBody) *SetUserPropertyValueResponse {
	s.Body = v
	return s
}

func (s *SetUserPropertyValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
