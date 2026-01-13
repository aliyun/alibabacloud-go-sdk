// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAlgorithmCustomizationScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAlgorithmCustomizationScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAlgorithmCustomizationScriptResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAlgorithmCustomizationScriptResponseBody) *GenerateAlgorithmCustomizationScriptResponse
	GetBody() *GenerateAlgorithmCustomizationScriptResponseBody
}

type GenerateAlgorithmCustomizationScriptResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAlgorithmCustomizationScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAlgorithmCustomizationScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAlgorithmCustomizationScriptResponse) GoString() string {
	return s.String()
}

func (s *GenerateAlgorithmCustomizationScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAlgorithmCustomizationScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAlgorithmCustomizationScriptResponse) GetBody() *GenerateAlgorithmCustomizationScriptResponseBody {
	return s.Body
}

func (s *GenerateAlgorithmCustomizationScriptResponse) SetHeaders(v map[string]*string) *GenerateAlgorithmCustomizationScriptResponse {
	s.Headers = v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptResponse) SetStatusCode(v int32) *GenerateAlgorithmCustomizationScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptResponse) SetBody(v *GenerateAlgorithmCustomizationScriptResponseBody) *GenerateAlgorithmCustomizationScriptResponse {
	s.Body = v
	return s
}

func (s *GenerateAlgorithmCustomizationScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
