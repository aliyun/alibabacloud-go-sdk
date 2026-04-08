// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataServiceApiResponseBody) *CreateDataServiceApiResponse
	GetBody() *CreateDataServiceApiResponseBody
}

type CreateDataServiceApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataServiceApiResponse) GetBody() *CreateDataServiceApiResponseBody {
	return s.Body
}

func (s *CreateDataServiceApiResponse) SetHeaders(v map[string]*string) *CreateDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *CreateDataServiceApiResponse) SetStatusCode(v int32) *CreateDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataServiceApiResponse) SetBody(v *CreateDataServiceApiResponseBody) *CreateDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *CreateDataServiceApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
