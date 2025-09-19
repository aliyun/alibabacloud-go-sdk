// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectApiInvokeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileDetectApiInvokeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileDetectApiInvokeInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetFileDetectApiInvokeInfoResponseBody) *GetFileDetectApiInvokeInfoResponse
	GetBody() *GetFileDetectApiInvokeInfoResponseBody
}

type GetFileDetectApiInvokeInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileDetectApiInvokeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileDetectApiInvokeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectApiInvokeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileDetectApiInvokeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileDetectApiInvokeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileDetectApiInvokeInfoResponse) GetBody() *GetFileDetectApiInvokeInfoResponseBody {
	return s.Body
}

func (s *GetFileDetectApiInvokeInfoResponse) SetHeaders(v map[string]*string) *GetFileDetectApiInvokeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponse) SetStatusCode(v int32) *GetFileDetectApiInvokeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponse) SetBody(v *GetFileDetectApiInvokeInfoResponseBody) *GetFileDetectApiInvokeInfoResponse {
	s.Body = v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponse) Validate() error {
	return dara.Validate(s)
}
