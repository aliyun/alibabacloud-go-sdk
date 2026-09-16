// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeEntityDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDigitalEmployeeEntityDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDigitalEmployeeEntityDataResponse
	GetStatusCode() *int32
	SetBody(v *GetDigitalEmployeeEntityDataResponseBody) *GetDigitalEmployeeEntityDataResponse
	GetBody() *GetDigitalEmployeeEntityDataResponseBody
}

type GetDigitalEmployeeEntityDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDigitalEmployeeEntityDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDigitalEmployeeEntityDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeEntityDataResponse) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeEntityDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDigitalEmployeeEntityDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDigitalEmployeeEntityDataResponse) GetBody() *GetDigitalEmployeeEntityDataResponseBody {
	return s.Body
}

func (s *GetDigitalEmployeeEntityDataResponse) SetHeaders(v map[string]*string) *GetDigitalEmployeeEntityDataResponse {
	s.Headers = v
	return s
}

func (s *GetDigitalEmployeeEntityDataResponse) SetStatusCode(v int32) *GetDigitalEmployeeEntityDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDigitalEmployeeEntityDataResponse) SetBody(v *GetDigitalEmployeeEntityDataResponseBody) *GetDigitalEmployeeEntityDataResponse {
	s.Body = v
	return s
}

func (s *GetDigitalEmployeeEntityDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
