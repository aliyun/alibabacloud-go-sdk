// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAudienceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomAudienceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomAudienceUserResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomAudienceUserResponseBody) *AddCustomAudienceUserResponse
	GetBody() *AddCustomAudienceUserResponseBody
}

type AddCustomAudienceUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomAudienceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomAudienceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAudienceUserResponse) GoString() string {
	return s.String()
}

func (s *AddCustomAudienceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomAudienceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomAudienceUserResponse) GetBody() *AddCustomAudienceUserResponseBody {
	return s.Body
}

func (s *AddCustomAudienceUserResponse) SetHeaders(v map[string]*string) *AddCustomAudienceUserResponse {
	s.Headers = v
	return s
}

func (s *AddCustomAudienceUserResponse) SetStatusCode(v int32) *AddCustomAudienceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomAudienceUserResponse) SetBody(v *AddCustomAudienceUserResponseBody) *AddCustomAudienceUserResponse {
	s.Body = v
	return s
}

func (s *AddCustomAudienceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
