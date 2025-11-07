// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNisInspectionResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNisInspectionResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListNisInspectionResourceTypeResponseBody) *ListNisInspectionResourceTypeResponse
	GetBody() *ListNisInspectionResourceTypeResponseBody
}

type ListNisInspectionResourceTypeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNisInspectionResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNisInspectionResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *ListNisInspectionResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNisInspectionResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNisInspectionResourceTypeResponse) GetBody() *ListNisInspectionResourceTypeResponseBody {
	return s.Body
}

func (s *ListNisInspectionResourceTypeResponse) SetHeaders(v map[string]*string) *ListNisInspectionResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *ListNisInspectionResourceTypeResponse) SetStatusCode(v int32) *ListNisInspectionResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNisInspectionResourceTypeResponse) SetBody(v *ListNisInspectionResourceTypeResponseBody) *ListNisInspectionResourceTypeResponse {
	s.Body = v
	return s
}

func (s *ListNisInspectionResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
