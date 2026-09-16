// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeUmodelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDigitalEmployeeUmodelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDigitalEmployeeUmodelResponse
	GetStatusCode() *int32
	SetBody(v *GetDigitalEmployeeUmodelResponseBody) *GetDigitalEmployeeUmodelResponse
	GetBody() *GetDigitalEmployeeUmodelResponseBody
}

type GetDigitalEmployeeUmodelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDigitalEmployeeUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDigitalEmployeeUmodelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeUmodelResponse) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeUmodelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDigitalEmployeeUmodelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDigitalEmployeeUmodelResponse) GetBody() *GetDigitalEmployeeUmodelResponseBody {
	return s.Body
}

func (s *GetDigitalEmployeeUmodelResponse) SetHeaders(v map[string]*string) *GetDigitalEmployeeUmodelResponse {
	s.Headers = v
	return s
}

func (s *GetDigitalEmployeeUmodelResponse) SetStatusCode(v int32) *GetDigitalEmployeeUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDigitalEmployeeUmodelResponse) SetBody(v *GetDigitalEmployeeUmodelResponseBody) *GetDigitalEmployeeUmodelResponse {
	s.Body = v
	return s
}

func (s *GetDigitalEmployeeUmodelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
