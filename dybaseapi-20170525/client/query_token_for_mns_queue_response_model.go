// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTokenForMnsQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTokenForMnsQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTokenForMnsQueueResponse
	GetStatusCode() *int32
	SetBody(v *QueryTokenForMnsQueueResponseBody) *QueryTokenForMnsQueueResponse
	GetBody() *QueryTokenForMnsQueueResponseBody
}

type QueryTokenForMnsQueueResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTokenForMnsQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTokenForMnsQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTokenForMnsQueueResponse) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTokenForMnsQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTokenForMnsQueueResponse) GetBody() *QueryTokenForMnsQueueResponseBody {
	return s.Body
}

func (s *QueryTokenForMnsQueueResponse) SetHeaders(v map[string]*string) *QueryTokenForMnsQueueResponse {
	s.Headers = v
	return s
}

func (s *QueryTokenForMnsQueueResponse) SetStatusCode(v int32) *QueryTokenForMnsQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTokenForMnsQueueResponse) SetBody(v *QueryTokenForMnsQueueResponseBody) *QueryTokenForMnsQueueResponse {
	s.Body = v
	return s
}

func (s *QueryTokenForMnsQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
