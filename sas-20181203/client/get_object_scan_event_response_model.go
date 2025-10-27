// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectScanEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectScanEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectScanEventResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectScanEventResponseBody) *GetObjectScanEventResponse
	GetBody() *GetObjectScanEventResponseBody
}

type GetObjectScanEventResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectScanEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectScanEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectScanEventResponse) GoString() string {
	return s.String()
}

func (s *GetObjectScanEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectScanEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectScanEventResponse) GetBody() *GetObjectScanEventResponseBody {
	return s.Body
}

func (s *GetObjectScanEventResponse) SetHeaders(v map[string]*string) *GetObjectScanEventResponse {
	s.Headers = v
	return s
}

func (s *GetObjectScanEventResponse) SetStatusCode(v int32) *GetObjectScanEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectScanEventResponse) SetBody(v *GetObjectScanEventResponseBody) *GetObjectScanEventResponse {
	s.Body = v
	return s
}

func (s *GetObjectScanEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
