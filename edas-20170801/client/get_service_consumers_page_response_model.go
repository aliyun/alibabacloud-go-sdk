// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConsumersPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceConsumersPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceConsumersPageResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceConsumersPageResponseBody) *GetServiceConsumersPageResponse
	GetBody() *GetServiceConsumersPageResponseBody
}

type GetServiceConsumersPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceConsumersPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceConsumersPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConsumersPageResponse) GoString() string {
	return s.String()
}

func (s *GetServiceConsumersPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceConsumersPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceConsumersPageResponse) GetBody() *GetServiceConsumersPageResponseBody {
	return s.Body
}

func (s *GetServiceConsumersPageResponse) SetHeaders(v map[string]*string) *GetServiceConsumersPageResponse {
	s.Headers = v
	return s
}

func (s *GetServiceConsumersPageResponse) SetStatusCode(v int32) *GetServiceConsumersPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceConsumersPageResponse) SetBody(v *GetServiceConsumersPageResponseBody) *GetServiceConsumersPageResponse {
	s.Body = v
	return s
}

func (s *GetServiceConsumersPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
