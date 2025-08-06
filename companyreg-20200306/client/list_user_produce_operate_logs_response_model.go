// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProduceOperateLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserProduceOperateLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserProduceOperateLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserProduceOperateLogsResponseBody) *ListUserProduceOperateLogsResponse
	GetBody() *ListUserProduceOperateLogsResponseBody
}

type ListUserProduceOperateLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserProduceOperateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserProduceOperateLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserProduceOperateLogsResponse) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserProduceOperateLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserProduceOperateLogsResponse) GetBody() *ListUserProduceOperateLogsResponseBody {
	return s.Body
}

func (s *ListUserProduceOperateLogsResponse) SetHeaders(v map[string]*string) *ListUserProduceOperateLogsResponse {
	s.Headers = v
	return s
}

func (s *ListUserProduceOperateLogsResponse) SetStatusCode(v int32) *ListUserProduceOperateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserProduceOperateLogsResponse) SetBody(v *ListUserProduceOperateLogsResponseBody) *ListUserProduceOperateLogsResponse {
	s.Body = v
	return s
}

func (s *ListUserProduceOperateLogsResponse) Validate() error {
	return dara.Validate(s)
}
