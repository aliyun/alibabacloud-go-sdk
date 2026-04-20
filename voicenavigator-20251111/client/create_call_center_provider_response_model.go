// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallCenterProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCallCenterProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCallCenterProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateCallCenterProviderResponseBody) *CreateCallCenterProviderResponse
	GetBody() *CreateCallCenterProviderResponseBody
}

type CreateCallCenterProviderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallCenterProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCallCenterProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCallCenterProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateCallCenterProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCallCenterProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCallCenterProviderResponse) GetBody() *CreateCallCenterProviderResponseBody {
	return s.Body
}

func (s *CreateCallCenterProviderResponse) SetHeaders(v map[string]*string) *CreateCallCenterProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateCallCenterProviderResponse) SetStatusCode(v int32) *CreateCallCenterProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallCenterProviderResponse) SetBody(v *CreateCallCenterProviderResponseBody) *CreateCallCenterProviderResponse {
	s.Body = v
	return s
}

func (s *CreateCallCenterProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
