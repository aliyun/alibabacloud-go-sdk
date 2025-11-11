// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTimeLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncCreateClipsTimeLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncCreateClipsTimeLineResponse
	GetStatusCode() *int32
	SetBody(v *AsyncCreateClipsTimeLineResponseBody) *AsyncCreateClipsTimeLineResponse
	GetBody() *AsyncCreateClipsTimeLineResponseBody
}

type AsyncCreateClipsTimeLineResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncCreateClipsTimeLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncCreateClipsTimeLineResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTimeLineResponse) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTimeLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncCreateClipsTimeLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncCreateClipsTimeLineResponse) GetBody() *AsyncCreateClipsTimeLineResponseBody {
	return s.Body
}

func (s *AsyncCreateClipsTimeLineResponse) SetHeaders(v map[string]*string) *AsyncCreateClipsTimeLineResponse {
	s.Headers = v
	return s
}

func (s *AsyncCreateClipsTimeLineResponse) SetStatusCode(v int32) *AsyncCreateClipsTimeLineResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncCreateClipsTimeLineResponse) SetBody(v *AsyncCreateClipsTimeLineResponseBody) *AsyncCreateClipsTimeLineResponse {
	s.Body = v
	return s
}

func (s *AsyncCreateClipsTimeLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
