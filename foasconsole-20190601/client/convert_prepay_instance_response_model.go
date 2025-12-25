// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPrepayInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertPrepayInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertPrepayInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ConvertPrepayInstanceResponseBody) *ConvertPrepayInstanceResponse
	GetBody() *ConvertPrepayInstanceResponseBody
}

type ConvertPrepayInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertPrepayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertPrepayInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertPrepayInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertPrepayInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertPrepayInstanceResponse) GetBody() *ConvertPrepayInstanceResponseBody {
	return s.Body
}

func (s *ConvertPrepayInstanceResponse) SetHeaders(v map[string]*string) *ConvertPrepayInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConvertPrepayInstanceResponse) SetStatusCode(v int32) *ConvertPrepayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertPrepayInstanceResponse) SetBody(v *ConvertPrepayInstanceResponseBody) *ConvertPrepayInstanceResponse {
	s.Body = v
	return s
}

func (s *ConvertPrepayInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
