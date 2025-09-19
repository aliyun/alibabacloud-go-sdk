// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSasModuleTrialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSasModuleTrialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSasModuleTrialResponse
	GetStatusCode() *int32
	SetBody(v *AddSasModuleTrialResponseBody) *AddSasModuleTrialResponse
	GetBody() *AddSasModuleTrialResponseBody
}

type AddSasModuleTrialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSasModuleTrialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSasModuleTrialResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSasModuleTrialResponse) GoString() string {
	return s.String()
}

func (s *AddSasModuleTrialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSasModuleTrialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSasModuleTrialResponse) GetBody() *AddSasModuleTrialResponseBody {
	return s.Body
}

func (s *AddSasModuleTrialResponse) SetHeaders(v map[string]*string) *AddSasModuleTrialResponse {
	s.Headers = v
	return s
}

func (s *AddSasModuleTrialResponse) SetStatusCode(v int32) *AddSasModuleTrialResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSasModuleTrialResponse) SetBody(v *AddSasModuleTrialResponseBody) *AddSasModuleTrialResponse {
	s.Body = v
	return s
}

func (s *AddSasModuleTrialResponse) Validate() error {
	return dara.Validate(s)
}
