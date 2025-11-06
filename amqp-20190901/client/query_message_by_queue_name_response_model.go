// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageByQueueNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMessageByQueueNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMessageByQueueNameResponse
	GetStatusCode() *int32
	SetBody(v *QueryMessageByQueueNameResponseBody) *QueryMessageByQueueNameResponse
	GetBody() *QueryMessageByQueueNameResponseBody
}

type QueryMessageByQueueNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMessageByQueueNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMessageByQueueNameResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByQueueNameResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageByQueueNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMessageByQueueNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMessageByQueueNameResponse) GetBody() *QueryMessageByQueueNameResponseBody {
	return s.Body
}

func (s *QueryMessageByQueueNameResponse) SetHeaders(v map[string]*string) *QueryMessageByQueueNameResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageByQueueNameResponse) SetStatusCode(v int32) *QueryMessageByQueueNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageByQueueNameResponse) SetBody(v *QueryMessageByQueueNameResponseBody) *QueryMessageByQueueNameResponse {
	s.Body = v
	return s
}

func (s *QueryMessageByQueueNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
