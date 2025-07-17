// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateContactResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateContactResponseBody) *CreateOrUpdateContactResponse
	GetBody() *CreateOrUpdateContactResponseBody
}

type CreateOrUpdateContactResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateContactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateContactResponse) GetBody() *CreateOrUpdateContactResponseBody {
	return s.Body
}

func (s *CreateOrUpdateContactResponse) SetHeaders(v map[string]*string) *CreateOrUpdateContactResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateContactResponse) SetStatusCode(v int32) *CreateOrUpdateContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateContactResponse) SetBody(v *CreateOrUpdateContactResponseBody) *CreateOrUpdateContactResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateContactResponse) Validate() error {
	return dara.Validate(s)
}
