// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMsgIdListByQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMsgIdListByQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMsgIdListByQueueResponse
	GetStatusCode() *int32
	SetBody(v *GetMsgIdListByQueueResponseBody) *GetMsgIdListByQueueResponse
	GetBody() *GetMsgIdListByQueueResponseBody
}

type GetMsgIdListByQueueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMsgIdListByQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMsgIdListByQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMsgIdListByQueueResponse) GoString() string {
	return s.String()
}

func (s *GetMsgIdListByQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMsgIdListByQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMsgIdListByQueueResponse) GetBody() *GetMsgIdListByQueueResponseBody {
	return s.Body
}

func (s *GetMsgIdListByQueueResponse) SetHeaders(v map[string]*string) *GetMsgIdListByQueueResponse {
	s.Headers = v
	return s
}

func (s *GetMsgIdListByQueueResponse) SetStatusCode(v int32) *GetMsgIdListByQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMsgIdListByQueueResponse) SetBody(v *GetMsgIdListByQueueResponseBody) *GetMsgIdListByQueueResponse {
	s.Body = v
	return s
}

func (s *GetMsgIdListByQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
