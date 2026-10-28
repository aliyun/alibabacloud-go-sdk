// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertHistoriesResponseBody) *ListAlertHistoriesResponse
	GetBody() *ListAlertHistoriesResponseBody
}

type ListAlertHistoriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAlertHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertHistoriesResponse) GetBody() *ListAlertHistoriesResponseBody {
	return s.Body
}

func (s *ListAlertHistoriesResponse) SetHeaders(v map[string]*string) *ListAlertHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAlertHistoriesResponse) SetStatusCode(v int32) *ListAlertHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertHistoriesResponse) SetBody(v *ListAlertHistoriesResponseBody) *ListAlertHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListAlertHistoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
