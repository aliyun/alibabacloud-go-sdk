// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ConvertInstanceResponseBody) *ConvertInstanceResponse
	GetBody() *ConvertInstanceResponseBody
}

type ConvertInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertInstanceResponse) GetBody() *ConvertInstanceResponseBody {
	return s.Body
}

func (s *ConvertInstanceResponse) SetHeaders(v map[string]*string) *ConvertInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConvertInstanceResponse) SetStatusCode(v int32) *ConvertInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertInstanceResponse) SetBody(v *ConvertInstanceResponseBody) *ConvertInstanceResponse {
	s.Body = v
	return s
}

func (s *ConvertInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
