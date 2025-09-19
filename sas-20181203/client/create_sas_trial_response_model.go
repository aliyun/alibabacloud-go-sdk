// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSasTrialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSasTrialResponse
	GetStatusCode() *int32
	SetBody(v *CreateSasTrialResponseBody) *CreateSasTrialResponse
	GetBody() *CreateSasTrialResponseBody
}

type CreateSasTrialResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSasTrialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSasTrialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialResponse) GoString() string {
	return s.String()
}

func (s *CreateSasTrialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSasTrialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSasTrialResponse) GetBody() *CreateSasTrialResponseBody {
	return s.Body
}

func (s *CreateSasTrialResponse) SetHeaders(v map[string]*string) *CreateSasTrialResponse {
	s.Headers = v
	return s
}

func (s *CreateSasTrialResponse) SetStatusCode(v int32) *CreateSasTrialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSasTrialResponse) SetBody(v *CreateSasTrialResponseBody) *CreateSasTrialResponse {
	s.Body = v
	return s
}

func (s *CreateSasTrialResponse) Validate() error {
	return dara.Validate(s)
}
