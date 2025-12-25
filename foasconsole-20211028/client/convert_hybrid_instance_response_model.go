// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertHybridInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertHybridInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertHybridInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ConvertHybridInstanceResponseBody) *ConvertHybridInstanceResponse
	GetBody() *ConvertHybridInstanceResponseBody
}

type ConvertHybridInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertHybridInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertHybridInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertHybridInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConvertHybridInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertHybridInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertHybridInstanceResponse) GetBody() *ConvertHybridInstanceResponseBody {
	return s.Body
}

func (s *ConvertHybridInstanceResponse) SetHeaders(v map[string]*string) *ConvertHybridInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConvertHybridInstanceResponse) SetStatusCode(v int32) *ConvertHybridInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertHybridInstanceResponse) SetBody(v *ConvertHybridInstanceResponseBody) *ConvertHybridInstanceResponse {
	s.Body = v
	return s
}

func (s *ConvertHybridInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
