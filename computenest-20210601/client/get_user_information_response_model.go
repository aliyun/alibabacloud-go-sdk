// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserInformationResponse
	GetStatusCode() *int32
	SetBody(v *GetUserInformationResponseBody) *GetUserInformationResponse
	GetBody() *GetUserInformationResponseBody
}

type GetUserInformationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserInformationResponse) GoString() string {
	return s.String()
}

func (s *GetUserInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserInformationResponse) GetBody() *GetUserInformationResponseBody {
	return s.Body
}

func (s *GetUserInformationResponse) SetHeaders(v map[string]*string) *GetUserInformationResponse {
	s.Headers = v
	return s
}

func (s *GetUserInformationResponse) SetStatusCode(v int32) *GetUserInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInformationResponse) SetBody(v *GetUserInformationResponseBody) *GetUserInformationResponse {
	s.Body = v
	return s
}

func (s *GetUserInformationResponse) Validate() error {
	return dara.Validate(s)
}
