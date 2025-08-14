// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChannelAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChannelAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChannelAlertsResponse
	GetStatusCode() *int32
	SetBody(v *ListChannelAlertsResponseBody) *ListChannelAlertsResponse
	GetBody() *ListChannelAlertsResponseBody
}

type ListChannelAlertsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChannelAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChannelAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChannelAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListChannelAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChannelAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChannelAlertsResponse) GetBody() *ListChannelAlertsResponseBody {
	return s.Body
}

func (s *ListChannelAlertsResponse) SetHeaders(v map[string]*string) *ListChannelAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListChannelAlertsResponse) SetStatusCode(v int32) *ListChannelAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChannelAlertsResponse) SetBody(v *ListChannelAlertsResponseBody) *ListChannelAlertsResponse {
	s.Body = v
	return s
}

func (s *ListChannelAlertsResponse) Validate() error {
	return dara.Validate(s)
}
