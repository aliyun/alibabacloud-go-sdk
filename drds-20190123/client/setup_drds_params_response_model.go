// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupDrdsParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetupDrdsParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetupDrdsParamsResponse
	GetStatusCode() *int32
	SetBody(v *SetupDrdsParamsResponseBody) *SetupDrdsParamsResponse
	GetBody() *SetupDrdsParamsResponseBody
}

type SetupDrdsParamsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupDrdsParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupDrdsParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetupDrdsParamsResponse) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetupDrdsParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetupDrdsParamsResponse) GetBody() *SetupDrdsParamsResponseBody {
	return s.Body
}

func (s *SetupDrdsParamsResponse) SetHeaders(v map[string]*string) *SetupDrdsParamsResponse {
	s.Headers = v
	return s
}

func (s *SetupDrdsParamsResponse) SetStatusCode(v int32) *SetupDrdsParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupDrdsParamsResponse) SetBody(v *SetupDrdsParamsResponseBody) *SetupDrdsParamsResponse {
	s.Body = v
	return s
}

func (s *SetupDrdsParamsResponse) Validate() error {
	return dara.Validate(s)
}
