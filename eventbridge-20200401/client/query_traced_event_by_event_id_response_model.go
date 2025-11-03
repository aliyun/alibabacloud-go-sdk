// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTracedEventByEventIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTracedEventByEventIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTracedEventByEventIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryTracedEventByEventIdResponseBody) *QueryTracedEventByEventIdResponse
	GetBody() *QueryTracedEventByEventIdResponseBody
}

type QueryTracedEventByEventIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTracedEventByEventIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTracedEventByEventIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventByEventIdResponse) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTracedEventByEventIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTracedEventByEventIdResponse) GetBody() *QueryTracedEventByEventIdResponseBody {
	return s.Body
}

func (s *QueryTracedEventByEventIdResponse) SetHeaders(v map[string]*string) *QueryTracedEventByEventIdResponse {
	s.Headers = v
	return s
}

func (s *QueryTracedEventByEventIdResponse) SetStatusCode(v int32) *QueryTracedEventByEventIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTracedEventByEventIdResponse) SetBody(v *QueryTracedEventByEventIdResponseBody) *QueryTracedEventByEventIdResponse {
	s.Body = v
	return s
}

func (s *QueryTracedEventByEventIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
