// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *AddWhitelistResponseBody) *AddWhitelistResponse
	GetBody() *AddWhitelistResponseBody
}

type AddWhitelistResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWhitelistResponse) GoString() string {
	return s.String()
}

func (s *AddWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWhitelistResponse) GetBody() *AddWhitelistResponseBody {
	return s.Body
}

func (s *AddWhitelistResponse) SetHeaders(v map[string]*string) *AddWhitelistResponse {
	s.Headers = v
	return s
}

func (s *AddWhitelistResponse) SetStatusCode(v int32) *AddWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWhitelistResponse) SetBody(v *AddWhitelistResponseBody) *AddWhitelistResponse {
	s.Body = v
	return s
}

func (s *AddWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
