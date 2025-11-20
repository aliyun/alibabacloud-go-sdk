// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateFormDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateFormDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateFormDataResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateFormDataResponseBody) *CreateOrUpdateFormDataResponse
	GetBody() *CreateOrUpdateFormDataResponseBody
}

type CreateOrUpdateFormDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateFormDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateFormDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateFormDataResponse) GetBody() *CreateOrUpdateFormDataResponseBody {
	return s.Body
}

func (s *CreateOrUpdateFormDataResponse) SetHeaders(v map[string]*string) *CreateOrUpdateFormDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateFormDataResponse) SetStatusCode(v int32) *CreateOrUpdateFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateFormDataResponse) SetBody(v *CreateOrUpdateFormDataResponseBody) *CreateOrUpdateFormDataResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateFormDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
