// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SynchronizeAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SynchronizeAppResponse
	GetStatusCode() *int32
	SetBody(v *SynchronizeAppResponseBody) *SynchronizeAppResponse
	GetBody() *SynchronizeAppResponseBody
}

type SynchronizeAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynchronizeAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SynchronizeAppResponse) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeAppResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SynchronizeAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SynchronizeAppResponse) GetBody() *SynchronizeAppResponseBody {
	return s.Body
}

func (s *SynchronizeAppResponse) SetHeaders(v map[string]*string) *SynchronizeAppResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeAppResponse) SetStatusCode(v int32) *SynchronizeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SynchronizeAppResponse) SetBody(v *SynchronizeAppResponseBody) *SynchronizeAppResponse {
	s.Body = v
	return s
}

func (s *SynchronizeAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
