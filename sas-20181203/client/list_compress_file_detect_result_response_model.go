// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompressFileDetectResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCompressFileDetectResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCompressFileDetectResultResponse
	GetStatusCode() *int32
	SetBody(v *ListCompressFileDetectResultResponseBody) *ListCompressFileDetectResultResponse
	GetBody() *ListCompressFileDetectResultResponseBody
}

type ListCompressFileDetectResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCompressFileDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCompressFileDetectResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCompressFileDetectResultResponse) GoString() string {
	return s.String()
}

func (s *ListCompressFileDetectResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCompressFileDetectResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCompressFileDetectResultResponse) GetBody() *ListCompressFileDetectResultResponseBody {
	return s.Body
}

func (s *ListCompressFileDetectResultResponse) SetHeaders(v map[string]*string) *ListCompressFileDetectResultResponse {
	s.Headers = v
	return s
}

func (s *ListCompressFileDetectResultResponse) SetStatusCode(v int32) *ListCompressFileDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCompressFileDetectResultResponse) SetBody(v *ListCompressFileDetectResultResponseBody) *ListCompressFileDetectResultResponse {
	s.Body = v
	return s
}

func (s *ListCompressFileDetectResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
