// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckSqlFlashbackTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreCheckSqlFlashbackTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreCheckSqlFlashbackTaskResponse
	GetStatusCode() *int32
	SetBody(v *PreCheckSqlFlashbackTaskResponseBody) *PreCheckSqlFlashbackTaskResponse
	GetBody() *PreCheckSqlFlashbackTaskResponseBody
}

type PreCheckSqlFlashbackTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreCheckSqlFlashbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreCheckSqlFlashbackTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PreCheckSqlFlashbackTaskResponse) GoString() string {
	return s.String()
}

func (s *PreCheckSqlFlashbackTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreCheckSqlFlashbackTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreCheckSqlFlashbackTaskResponse) GetBody() *PreCheckSqlFlashbackTaskResponseBody {
	return s.Body
}

func (s *PreCheckSqlFlashbackTaskResponse) SetHeaders(v map[string]*string) *PreCheckSqlFlashbackTaskResponse {
	s.Headers = v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponse) SetStatusCode(v int32) *PreCheckSqlFlashbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponse) SetBody(v *PreCheckSqlFlashbackTaskResponseBody) *PreCheckSqlFlashbackTaskResponse {
	s.Body = v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
