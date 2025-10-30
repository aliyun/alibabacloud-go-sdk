// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExternalDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExternalDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateExternalDataServiceResponseBody) *CreateExternalDataServiceResponse
	GetBody() *CreateExternalDataServiceResponseBody
}

type CreateExternalDataServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExternalDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExternalDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalDataServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateExternalDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExternalDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExternalDataServiceResponse) GetBody() *CreateExternalDataServiceResponseBody {
	return s.Body
}

func (s *CreateExternalDataServiceResponse) SetHeaders(v map[string]*string) *CreateExternalDataServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateExternalDataServiceResponse) SetStatusCode(v int32) *CreateExternalDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExternalDataServiceResponse) SetBody(v *CreateExternalDataServiceResponseBody) *CreateExternalDataServiceResponse {
	s.Body = v
	return s
}

func (s *CreateExternalDataServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
