// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBatchTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBatchTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBatchTaskListResponse
	GetStatusCode() *int32
	SetBody(v *QueryBatchTaskListResponseBody) *QueryBatchTaskListResponse
	GetBody() *QueryBatchTaskListResponseBody
}

type QueryBatchTaskListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBatchTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBatchTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBatchTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBatchTaskListResponse) GetBody() *QueryBatchTaskListResponseBody {
	return s.Body
}

func (s *QueryBatchTaskListResponse) SetHeaders(v map[string]*string) *QueryBatchTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchTaskListResponse) SetStatusCode(v int32) *QueryBatchTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBatchTaskListResponse) SetBody(v *QueryBatchTaskListResponseBody) *QueryBatchTaskListResponse {
	s.Body = v
	return s
}

func (s *QueryBatchTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
