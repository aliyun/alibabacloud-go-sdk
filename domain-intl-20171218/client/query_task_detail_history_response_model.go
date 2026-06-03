// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskDetailHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskDetailHistoryResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskDetailHistoryResponseBody) *QueryTaskDetailHistoryResponse
	GetBody() *QueryTaskDetailHistoryResponseBody
}

type QueryTaskDetailHistoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskDetailHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskDetailHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailHistoryResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskDetailHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskDetailHistoryResponse) GetBody() *QueryTaskDetailHistoryResponseBody {
	return s.Body
}

func (s *QueryTaskDetailHistoryResponse) SetHeaders(v map[string]*string) *QueryTaskDetailHistoryResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskDetailHistoryResponse) SetStatusCode(v int32) *QueryTaskDetailHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponse) SetBody(v *QueryTaskDetailHistoryResponseBody) *QueryTaskDetailHistoryResponse {
	s.Body = v
	return s
}

func (s *QueryTaskDetailHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
