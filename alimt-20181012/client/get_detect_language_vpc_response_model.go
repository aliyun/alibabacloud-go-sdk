// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectLanguageVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDetectLanguageVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDetectLanguageVpcResponse
	GetStatusCode() *int32
	SetBody(v *GetDetectLanguageVpcResponseBody) *GetDetectLanguageVpcResponse
	GetBody() *GetDetectLanguageVpcResponseBody
}

type GetDetectLanguageVpcResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectLanguageVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectLanguageVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDetectLanguageVpcResponse) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDetectLanguageVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDetectLanguageVpcResponse) GetBody() *GetDetectLanguageVpcResponseBody {
	return s.Body
}

func (s *GetDetectLanguageVpcResponse) SetHeaders(v map[string]*string) *GetDetectLanguageVpcResponse {
	s.Headers = v
	return s
}

func (s *GetDetectLanguageVpcResponse) SetStatusCode(v int32) *GetDetectLanguageVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectLanguageVpcResponse) SetBody(v *GetDetectLanguageVpcResponseBody) *GetDetectLanguageVpcResponse {
	s.Body = v
	return s
}

func (s *GetDetectLanguageVpcResponse) Validate() error {
	return dara.Validate(s)
}
