// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityCheckTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityCheckTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityCheckTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityCheckTaskResultResponseBody) *GetQualityCheckTaskResultResponse
	GetBody() *GetQualityCheckTaskResultResponseBody
}

type GetQualityCheckTaskResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityCheckTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityCheckTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityCheckTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityCheckTaskResultResponse) GetBody() *GetQualityCheckTaskResultResponseBody {
	return s.Body
}

func (s *GetQualityCheckTaskResultResponse) SetHeaders(v map[string]*string) *GetQualityCheckTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetQualityCheckTaskResultResponse) SetStatusCode(v int32) *GetQualityCheckTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityCheckTaskResultResponse) SetBody(v *GetQualityCheckTaskResultResponseBody) *GetQualityCheckTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetQualityCheckTaskResultResponse) Validate() error {
	return dara.Validate(s)
}
