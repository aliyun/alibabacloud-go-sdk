// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAGDpiAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmartAGDpiAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmartAGDpiAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetSmartAGDpiAttributeResponseBody) *GetSmartAGDpiAttributeResponse
	GetBody() *GetSmartAGDpiAttributeResponseBody
}

type GetSmartAGDpiAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmartAGDpiAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmartAGDpiAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAGDpiAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetSmartAGDpiAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmartAGDpiAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmartAGDpiAttributeResponse) GetBody() *GetSmartAGDpiAttributeResponseBody {
	return s.Body
}

func (s *GetSmartAGDpiAttributeResponse) SetHeaders(v map[string]*string) *GetSmartAGDpiAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetSmartAGDpiAttributeResponse) SetStatusCode(v int32) *GetSmartAGDpiAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmartAGDpiAttributeResponse) SetBody(v *GetSmartAGDpiAttributeResponseBody) *GetSmartAGDpiAttributeResponse {
	s.Body = v
	return s
}

func (s *GetSmartAGDpiAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
