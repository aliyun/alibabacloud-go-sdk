// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *AddDnatEntryResponseBody) *AddDnatEntryResponse
	GetBody() *AddDnatEntryResponseBody
}

type AddDnatEntryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDnatEntryResponse) GoString() string {
	return s.String()
}

func (s *AddDnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDnatEntryResponse) GetBody() *AddDnatEntryResponseBody {
	return s.Body
}

func (s *AddDnatEntryResponse) SetHeaders(v map[string]*string) *AddDnatEntryResponse {
	s.Headers = v
	return s
}

func (s *AddDnatEntryResponse) SetStatusCode(v int32) *AddDnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDnatEntryResponse) SetBody(v *AddDnatEntryResponseBody) *AddDnatEntryResponse {
	s.Body = v
	return s
}

func (s *AddDnatEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
