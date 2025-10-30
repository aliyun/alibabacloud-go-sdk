// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExternalDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExternalDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExternalDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExternalDataServiceResponseBody) *DeleteExternalDataServiceResponse
	GetBody() *DeleteExternalDataServiceResponseBody
}

type DeleteExternalDataServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExternalDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExternalDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExternalDataServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteExternalDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExternalDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExternalDataServiceResponse) GetBody() *DeleteExternalDataServiceResponseBody {
	return s.Body
}

func (s *DeleteExternalDataServiceResponse) SetHeaders(v map[string]*string) *DeleteExternalDataServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteExternalDataServiceResponse) SetStatusCode(v int32) *DeleteExternalDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExternalDataServiceResponse) SetBody(v *DeleteExternalDataServiceResponseBody) *DeleteExternalDataServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteExternalDataServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
