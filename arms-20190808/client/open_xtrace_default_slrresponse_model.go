// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenXtraceDefaultSLRResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenXtraceDefaultSLRResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenXtraceDefaultSLRResponse
	GetStatusCode() *int32
	SetBody(v *OpenXtraceDefaultSLRResponseBody) *OpenXtraceDefaultSLRResponse
	GetBody() *OpenXtraceDefaultSLRResponseBody
}

type OpenXtraceDefaultSLRResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenXtraceDefaultSLRResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenXtraceDefaultSLRResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenXtraceDefaultSLRResponse) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenXtraceDefaultSLRResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenXtraceDefaultSLRResponse) GetBody() *OpenXtraceDefaultSLRResponseBody {
	return s.Body
}

func (s *OpenXtraceDefaultSLRResponse) SetHeaders(v map[string]*string) *OpenXtraceDefaultSLRResponse {
	s.Headers = v
	return s
}

func (s *OpenXtraceDefaultSLRResponse) SetStatusCode(v int32) *OpenXtraceDefaultSLRResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenXtraceDefaultSLRResponse) SetBody(v *OpenXtraceDefaultSLRResponseBody) *OpenXtraceDefaultSLRResponse {
	s.Body = v
	return s
}

func (s *OpenXtraceDefaultSLRResponse) Validate() error {
	return dara.Validate(s)
}
