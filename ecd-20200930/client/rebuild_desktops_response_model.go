// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebuildDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebuildDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *RebuildDesktopsResponseBody) *RebuildDesktopsResponse
	GetBody() *RebuildDesktopsResponseBody
}

type RebuildDesktopsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s RebuildDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebuildDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebuildDesktopsResponse) GetBody() *RebuildDesktopsResponseBody {
	return s.Body
}

func (s *RebuildDesktopsResponse) SetHeaders(v map[string]*string) *RebuildDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebuildDesktopsResponse) SetStatusCode(v int32) *RebuildDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildDesktopsResponse) SetBody(v *RebuildDesktopsResponseBody) *RebuildDesktopsResponse {
	s.Body = v
	return s
}

func (s *RebuildDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
