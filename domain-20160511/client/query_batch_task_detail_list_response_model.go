// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBatchTaskDetailListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBatchTaskDetailListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBatchTaskDetailListResponse
	GetStatusCode() *int32
	SetBody(v *QueryBatchTaskDetailListResponseBody) *QueryBatchTaskDetailListResponse
	GetBody() *QueryBatchTaskDetailListResponseBody
}

type QueryBatchTaskDetailListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBatchTaskDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBatchTaskDetailListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskDetailListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBatchTaskDetailListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBatchTaskDetailListResponse) GetBody() *QueryBatchTaskDetailListResponseBody {
	return s.Body
}

func (s *QueryBatchTaskDetailListResponse) SetHeaders(v map[string]*string) *QueryBatchTaskDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchTaskDetailListResponse) SetStatusCode(v int32) *QueryBatchTaskDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBatchTaskDetailListResponse) SetBody(v *QueryBatchTaskDetailListResponseBody) *QueryBatchTaskDetailListResponse {
	s.Body = v
	return s
}

func (s *QueryBatchTaskDetailListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
