// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICloudPhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAICloudPhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAICloudPhoneResponse
	GetStatusCode() *int32
	SetBody(v *CreateAICloudPhoneResponseBody) *CreateAICloudPhoneResponse
	GetBody() *CreateAICloudPhoneResponseBody
}

type CreateAICloudPhoneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAICloudPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAICloudPhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAICloudPhoneResponse) GoString() string {
	return s.String()
}

func (s *CreateAICloudPhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAICloudPhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAICloudPhoneResponse) GetBody() *CreateAICloudPhoneResponseBody {
	return s.Body
}

func (s *CreateAICloudPhoneResponse) SetHeaders(v map[string]*string) *CreateAICloudPhoneResponse {
	s.Headers = v
	return s
}

func (s *CreateAICloudPhoneResponse) SetStatusCode(v int32) *CreateAICloudPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAICloudPhoneResponse) SetBody(v *CreateAICloudPhoneResponseBody) *CreateAICloudPhoneResponse {
	s.Body = v
	return s
}

func (s *CreateAICloudPhoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
