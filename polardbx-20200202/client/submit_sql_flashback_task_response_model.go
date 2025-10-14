// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlFlashbackTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSqlFlashbackTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSqlFlashbackTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSqlFlashbackTaskResponseBody) *SubmitSqlFlashbackTaskResponse
	GetBody() *SubmitSqlFlashbackTaskResponseBody
}

type SubmitSqlFlashbackTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSqlFlashbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSqlFlashbackTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSqlFlashbackTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSqlFlashbackTaskResponse) GetBody() *SubmitSqlFlashbackTaskResponseBody {
	return s.Body
}

func (s *SubmitSqlFlashbackTaskResponse) SetHeaders(v map[string]*string) *SubmitSqlFlashbackTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSqlFlashbackTaskResponse) SetStatusCode(v int32) *SubmitSqlFlashbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponse) SetBody(v *SubmitSqlFlashbackTaskResponseBody) *SubmitSqlFlashbackTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitSqlFlashbackTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
