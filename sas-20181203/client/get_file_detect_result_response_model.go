// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileDetectResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileDetectResultResponse
	GetStatusCode() *int32
	SetBody(v *GetFileDetectResultResponseBody) *GetFileDetectResultResponse
	GetBody() *GetFileDetectResultResponseBody
}

type GetFileDetectResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileDetectResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectResultResponse) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileDetectResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileDetectResultResponse) GetBody() *GetFileDetectResultResponseBody {
	return s.Body
}

func (s *GetFileDetectResultResponse) SetHeaders(v map[string]*string) *GetFileDetectResultResponse {
	s.Headers = v
	return s
}

func (s *GetFileDetectResultResponse) SetStatusCode(v int32) *GetFileDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileDetectResultResponse) SetBody(v *GetFileDetectResultResponseBody) *GetFileDetectResultResponse {
	s.Body = v
	return s
}

func (s *GetFileDetectResultResponse) Validate() error {
	return dara.Validate(s)
}
