// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadCommonContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadCommonContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadCommonContactResponse
	GetStatusCode() *int32
	SetBody(v *ReadCommonContactResponseBody) *ReadCommonContactResponse
	GetBody() *ReadCommonContactResponseBody
}

type ReadCommonContactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadCommonContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadCommonContactResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadCommonContactResponse) GoString() string {
	return s.String()
}

func (s *ReadCommonContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadCommonContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadCommonContactResponse) GetBody() *ReadCommonContactResponseBody {
	return s.Body
}

func (s *ReadCommonContactResponse) SetHeaders(v map[string]*string) *ReadCommonContactResponse {
	s.Headers = v
	return s
}

func (s *ReadCommonContactResponse) SetStatusCode(v int32) *ReadCommonContactResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadCommonContactResponse) SetBody(v *ReadCommonContactResponseBody) *ReadCommonContactResponse {
	s.Body = v
	return s
}

func (s *ReadCommonContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
