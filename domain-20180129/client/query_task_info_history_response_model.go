// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskInfoHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskInfoHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskInfoHistoryResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskInfoHistoryResponseBody) *QueryTaskInfoHistoryResponse
	GetBody() *QueryTaskInfoHistoryResponseBody
}

type QueryTaskInfoHistoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskInfoHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskInfoHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskInfoHistoryResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskInfoHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskInfoHistoryResponse) GetBody() *QueryTaskInfoHistoryResponseBody {
	return s.Body
}

func (s *QueryTaskInfoHistoryResponse) SetHeaders(v map[string]*string) *QueryTaskInfoHistoryResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskInfoHistoryResponse) SetStatusCode(v int32) *QueryTaskInfoHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponse) SetBody(v *QueryTaskInfoHistoryResponseBody) *QueryTaskInfoHistoryResponse {
	s.Body = v
	return s
}

func (s *QueryTaskInfoHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
