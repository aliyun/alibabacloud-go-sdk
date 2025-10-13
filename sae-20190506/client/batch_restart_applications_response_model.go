// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRestartApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchRestartApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchRestartApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *BatchRestartApplicationsResponseBody) *BatchRestartApplicationsResponse
	GetBody() *BatchRestartApplicationsResponseBody
}

type BatchRestartApplicationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRestartApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRestartApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchRestartApplicationsResponse) GoString() string {
	return s.String()
}

func (s *BatchRestartApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRestartApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchRestartApplicationsResponse) GetBody() *BatchRestartApplicationsResponseBody {
	return s.Body
}

func (s *BatchRestartApplicationsResponse) SetHeaders(v map[string]*string) *BatchRestartApplicationsResponse {
	s.Headers = v
	return s
}

func (s *BatchRestartApplicationsResponse) SetStatusCode(v int32) *BatchRestartApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRestartApplicationsResponse) SetBody(v *BatchRestartApplicationsResponseBody) *BatchRestartApplicationsResponse {
	s.Body = v
	return s
}

func (s *BatchRestartApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
