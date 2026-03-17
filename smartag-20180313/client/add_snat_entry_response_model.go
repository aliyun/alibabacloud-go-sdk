// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *AddSnatEntryResponseBody) *AddSnatEntryResponse
	GetBody() *AddSnatEntryResponseBody
}

type AddSnatEntryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSnatEntryResponse) GoString() string {
	return s.String()
}

func (s *AddSnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSnatEntryResponse) GetBody() *AddSnatEntryResponseBody {
	return s.Body
}

func (s *AddSnatEntryResponse) SetHeaders(v map[string]*string) *AddSnatEntryResponse {
	s.Headers = v
	return s
}

func (s *AddSnatEntryResponse) SetStatusCode(v int32) *AddSnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSnatEntryResponse) SetBody(v *AddSnatEntryResponseBody) *AddSnatEntryResponse {
	s.Body = v
	return s
}

func (s *AddSnatEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
