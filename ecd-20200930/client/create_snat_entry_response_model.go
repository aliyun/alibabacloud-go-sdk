// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateSnatEntryResponseBody) *CreateSnatEntryResponse
	GetBody() *CreateSnatEntryResponseBody
}

type CreateSnatEntryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSnatEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateSnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSnatEntryResponse) GetBody() *CreateSnatEntryResponseBody {
	return s.Body
}

func (s *CreateSnatEntryResponse) SetHeaders(v map[string]*string) *CreateSnatEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateSnatEntryResponse) SetStatusCode(v int32) *CreateSnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnatEntryResponse) SetBody(v *CreateSnatEntryResponseBody) *CreateSnatEntryResponse {
	s.Body = v
	return s
}

func (s *CreateSnatEntryResponse) Validate() error {
	return dara.Validate(s)
}
