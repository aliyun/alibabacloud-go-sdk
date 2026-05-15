// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotlineInQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHotlineInQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHotlineInQueueResponse
	GetStatusCode() *int32
	SetBody(v *QueryHotlineInQueueResponseBody) *QueryHotlineInQueueResponse
	GetBody() *QueryHotlineInQueueResponseBody
}

type QueryHotlineInQueueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHotlineInQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotlineInQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineInQueueResponse) GoString() string {
	return s.String()
}

func (s *QueryHotlineInQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHotlineInQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHotlineInQueueResponse) GetBody() *QueryHotlineInQueueResponseBody {
	return s.Body
}

func (s *QueryHotlineInQueueResponse) SetHeaders(v map[string]*string) *QueryHotlineInQueueResponse {
	s.Headers = v
	return s
}

func (s *QueryHotlineInQueueResponse) SetStatusCode(v int32) *QueryHotlineInQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotlineInQueueResponse) SetBody(v *QueryHotlineInQueueResponseBody) *QueryHotlineInQueueResponse {
	s.Body = v
	return s
}

func (s *QueryHotlineInQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
