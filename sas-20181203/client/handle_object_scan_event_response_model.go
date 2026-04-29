// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleObjectScanEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleObjectScanEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleObjectScanEventResponse
	GetStatusCode() *int32
	SetBody(v *HandleObjectScanEventResponseBody) *HandleObjectScanEventResponse
	GetBody() *HandleObjectScanEventResponseBody
}

type HandleObjectScanEventResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleObjectScanEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleObjectScanEventResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleObjectScanEventResponse) GoString() string {
	return s.String()
}

func (s *HandleObjectScanEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleObjectScanEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleObjectScanEventResponse) GetBody() *HandleObjectScanEventResponseBody {
	return s.Body
}

func (s *HandleObjectScanEventResponse) SetHeaders(v map[string]*string) *HandleObjectScanEventResponse {
	s.Headers = v
	return s
}

func (s *HandleObjectScanEventResponse) SetStatusCode(v int32) *HandleObjectScanEventResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleObjectScanEventResponse) SetBody(v *HandleObjectScanEventResponseBody) *HandleObjectScanEventResponse {
	s.Body = v
	return s
}

func (s *HandleObjectScanEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
