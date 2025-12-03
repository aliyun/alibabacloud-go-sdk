// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateGeneralVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateGeneralVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateGeneralVpcResponse
	GetStatusCode() *int32
	SetBody(v *TranslateGeneralVpcResponseBody) *TranslateGeneralVpcResponse
	GetBody() *TranslateGeneralVpcResponseBody
}

type TranslateGeneralVpcResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateGeneralVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateGeneralVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralVpcResponse) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateGeneralVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateGeneralVpcResponse) GetBody() *TranslateGeneralVpcResponseBody {
	return s.Body
}

func (s *TranslateGeneralVpcResponse) SetHeaders(v map[string]*string) *TranslateGeneralVpcResponse {
	s.Headers = v
	return s
}

func (s *TranslateGeneralVpcResponse) SetStatusCode(v int32) *TranslateGeneralVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateGeneralVpcResponse) SetBody(v *TranslateGeneralVpcResponseBody) *TranslateGeneralVpcResponse {
	s.Body = v
	return s
}

func (s *TranslateGeneralVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
