// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourceMeasureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushResourceMeasureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushResourceMeasureResponse
	GetStatusCode() *int32
	SetBody(v *PushResourceMeasureResponseBody) *PushResourceMeasureResponse
	GetBody() *PushResourceMeasureResponseBody
}

type PushResourceMeasureResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushResourceMeasureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushResourceMeasureResponse) String() string {
	return dara.Prettify(s)
}

func (s PushResourceMeasureResponse) GoString() string {
	return s.String()
}

func (s *PushResourceMeasureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushResourceMeasureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushResourceMeasureResponse) GetBody() *PushResourceMeasureResponseBody {
	return s.Body
}

func (s *PushResourceMeasureResponse) SetHeaders(v map[string]*string) *PushResourceMeasureResponse {
	s.Headers = v
	return s
}

func (s *PushResourceMeasureResponse) SetStatusCode(v int32) *PushResourceMeasureResponse {
	s.StatusCode = &v
	return s
}

func (s *PushResourceMeasureResponse) SetBody(v *PushResourceMeasureResponseBody) *PushResourceMeasureResponse {
	s.Body = v
	return s
}

func (s *PushResourceMeasureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
