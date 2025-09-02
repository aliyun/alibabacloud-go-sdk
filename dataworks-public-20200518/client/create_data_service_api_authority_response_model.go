// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataServiceApiAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataServiceApiAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataServiceApiAuthorityResponseBody) *CreateDataServiceApiAuthorityResponse
	GetBody() *CreateDataServiceApiAuthorityResponseBody
}

type CreateDataServiceApiAuthorityResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataServiceApiAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataServiceApiAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiAuthorityResponse) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataServiceApiAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataServiceApiAuthorityResponse) GetBody() *CreateDataServiceApiAuthorityResponseBody {
	return s.Body
}

func (s *CreateDataServiceApiAuthorityResponse) SetHeaders(v map[string]*string) *CreateDataServiceApiAuthorityResponse {
	s.Headers = v
	return s
}

func (s *CreateDataServiceApiAuthorityResponse) SetStatusCode(v int32) *CreateDataServiceApiAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataServiceApiAuthorityResponse) SetBody(v *CreateDataServiceApiAuthorityResponseBody) *CreateDataServiceApiAuthorityResponse {
	s.Body = v
	return s
}

func (s *CreateDataServiceApiAuthorityResponse) Validate() error {
	return dara.Validate(s)
}
