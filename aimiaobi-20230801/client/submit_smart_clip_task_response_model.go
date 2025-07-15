// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartClipTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSmartClipTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSmartClipTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSmartClipTaskResponseBody) *SubmitSmartClipTaskResponse
	GetBody() *SubmitSmartClipTaskResponseBody
}

type SubmitSmartClipTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmartClipTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmartClipTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSmartClipTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSmartClipTaskResponse) GetBody() *SubmitSmartClipTaskResponseBody {
	return s.Body
}

func (s *SubmitSmartClipTaskResponse) SetHeaders(v map[string]*string) *SubmitSmartClipTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmartClipTaskResponse) SetStatusCode(v int32) *SubmitSmartClipTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmartClipTaskResponse) SetBody(v *SubmitSmartClipTaskResponseBody) *SubmitSmartClipTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitSmartClipTaskResponse) Validate() error {
	return dara.Validate(s)
}
