// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySiteMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySiteMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySiteMonitorResponse
	GetStatusCode() *int32
	SetBody(v *ModifySiteMonitorResponseBody) *ModifySiteMonitorResponse
	GetBody() *ModifySiteMonitorResponseBody
}

type ModifySiteMonitorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySiteMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySiteMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySiteMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySiteMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySiteMonitorResponse) GetBody() *ModifySiteMonitorResponseBody {
	return s.Body
}

func (s *ModifySiteMonitorResponse) SetHeaders(v map[string]*string) *ModifySiteMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifySiteMonitorResponse) SetStatusCode(v int32) *ModifySiteMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySiteMonitorResponse) SetBody(v *ModifySiteMonitorResponseBody) *ModifySiteMonitorResponse {
	s.Body = v
	return s
}

func (s *ModifySiteMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
