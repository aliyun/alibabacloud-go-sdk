// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccessControlListEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAccessControlListEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAccessControlListEntryResponse
	GetStatusCode() *int32
	SetBody(v *AddAccessControlListEntryResponseBody) *AddAccessControlListEntryResponse
	GetBody() *AddAccessControlListEntryResponseBody
}

type AddAccessControlListEntryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAccessControlListEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAccessControlListEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAccessControlListEntryResponse) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAccessControlListEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAccessControlListEntryResponse) GetBody() *AddAccessControlListEntryResponseBody {
	return s.Body
}

func (s *AddAccessControlListEntryResponse) SetHeaders(v map[string]*string) *AddAccessControlListEntryResponse {
	s.Headers = v
	return s
}

func (s *AddAccessControlListEntryResponse) SetStatusCode(v int32) *AddAccessControlListEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAccessControlListEntryResponse) SetBody(v *AddAccessControlListEntryResponseBody) *AddAccessControlListEntryResponse {
	s.Body = v
	return s
}

func (s *AddAccessControlListEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
