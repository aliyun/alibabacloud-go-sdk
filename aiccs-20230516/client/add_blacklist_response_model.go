// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *AddBlacklistResponseBody) *AddBlacklistResponse
	GetBody() *AddBlacklistResponseBody
}

type AddBlacklistResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBlacklistResponse) GetBody() *AddBlacklistResponseBody {
	return s.Body
}

func (s *AddBlacklistResponse) SetHeaders(v map[string]*string) *AddBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddBlacklistResponse) SetStatusCode(v int32) *AddBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBlacklistResponse) SetBody(v *AddBlacklistResponseBody) *AddBlacklistResponse {
	s.Body = v
	return s
}

func (s *AddBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
