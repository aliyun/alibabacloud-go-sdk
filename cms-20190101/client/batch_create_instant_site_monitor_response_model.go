// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateInstantSiteMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateInstantSiteMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateInstantSiteMonitorResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateInstantSiteMonitorResponseBody) *BatchCreateInstantSiteMonitorResponse
	GetBody() *BatchCreateInstantSiteMonitorResponseBody
}

type BatchCreateInstantSiteMonitorResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateInstantSiteMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateInstantSiteMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateInstantSiteMonitorResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateInstantSiteMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateInstantSiteMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateInstantSiteMonitorResponse) GetBody() *BatchCreateInstantSiteMonitorResponseBody {
	return s.Body
}

func (s *BatchCreateInstantSiteMonitorResponse) SetHeaders(v map[string]*string) *BatchCreateInstantSiteMonitorResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponse) SetStatusCode(v int32) *BatchCreateInstantSiteMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponse) SetBody(v *BatchCreateInstantSiteMonitorResponseBody) *BatchCreateInstantSiteMonitorResponse {
	s.Body = v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
