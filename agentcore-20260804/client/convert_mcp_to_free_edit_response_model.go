// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertMcpToFreeEditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertMcpToFreeEditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertMcpToFreeEditResponse
	GetStatusCode() *int32
	SetBody(v *ConvertMcpToFreeEditResponseBody) *ConvertMcpToFreeEditResponse
	GetBody() *ConvertMcpToFreeEditResponseBody
}

type ConvertMcpToFreeEditResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertMcpToFreeEditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertMcpToFreeEditResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponse) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertMcpToFreeEditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertMcpToFreeEditResponse) GetBody() *ConvertMcpToFreeEditResponseBody {
	return s.Body
}

func (s *ConvertMcpToFreeEditResponse) SetHeaders(v map[string]*string) *ConvertMcpToFreeEditResponse {
	s.Headers = v
	return s
}

func (s *ConvertMcpToFreeEditResponse) SetStatusCode(v int32) *ConvertMcpToFreeEditResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertMcpToFreeEditResponse) SetBody(v *ConvertMcpToFreeEditResponseBody) *ConvertMcpToFreeEditResponse {
	s.Body = v
	return s
}

func (s *ConvertMcpToFreeEditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
