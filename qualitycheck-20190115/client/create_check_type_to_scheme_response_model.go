// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckTypeToSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCheckTypeToSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCheckTypeToSchemeResponse
	GetStatusCode() *int32
	SetBody(v *CreateCheckTypeToSchemeResponseBody) *CreateCheckTypeToSchemeResponse
	GetBody() *CreateCheckTypeToSchemeResponseBody
}

type CreateCheckTypeToSchemeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCheckTypeToSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCheckTypeToSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckTypeToSchemeResponse) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCheckTypeToSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCheckTypeToSchemeResponse) GetBody() *CreateCheckTypeToSchemeResponseBody {
	return s.Body
}

func (s *CreateCheckTypeToSchemeResponse) SetHeaders(v map[string]*string) *CreateCheckTypeToSchemeResponse {
	s.Headers = v
	return s
}

func (s *CreateCheckTypeToSchemeResponse) SetStatusCode(v int32) *CreateCheckTypeToSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponse) SetBody(v *CreateCheckTypeToSchemeResponseBody) *CreateCheckTypeToSchemeResponse {
	s.Body = v
	return s
}

func (s *CreateCheckTypeToSchemeResponse) Validate() error {
	return dara.Validate(s)
}
