// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfficeConversionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOfficeConversionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOfficeConversionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateOfficeConversionTaskResponseBody) *CreateOfficeConversionTaskResponse
	GetBody() *CreateOfficeConversionTaskResponseBody
}

type CreateOfficeConversionTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOfficeConversionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOfficeConversionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOfficeConversionTaskResponse) GetBody() *CreateOfficeConversionTaskResponseBody {
	return s.Body
}

func (s *CreateOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *CreateOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOfficeConversionTaskResponse) SetStatusCode(v int32) *CreateOfficeConversionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOfficeConversionTaskResponse) SetBody(v *CreateOfficeConversionTaskResponseBody) *CreateOfficeConversionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateOfficeConversionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
