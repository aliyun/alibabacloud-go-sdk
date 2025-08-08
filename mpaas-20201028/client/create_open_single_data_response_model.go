// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSingleDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpenSingleDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpenSingleDataResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpenSingleDataResponseBody) *CreateOpenSingleDataResponse
	GetBody() *CreateOpenSingleDataResponseBody
}

type CreateOpenSingleDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpenSingleDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpenSingleDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSingleDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOpenSingleDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpenSingleDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpenSingleDataResponse) GetBody() *CreateOpenSingleDataResponseBody {
	return s.Body
}

func (s *CreateOpenSingleDataResponse) SetHeaders(v map[string]*string) *CreateOpenSingleDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOpenSingleDataResponse) SetStatusCode(v int32) *CreateOpenSingleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpenSingleDataResponse) SetBody(v *CreateOpenSingleDataResponseBody) *CreateOpenSingleDataResponse {
	s.Body = v
	return s
}

func (s *CreateOpenSingleDataResponse) Validate() error {
	return dara.Validate(s)
}
