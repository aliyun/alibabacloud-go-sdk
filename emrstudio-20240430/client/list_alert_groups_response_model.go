// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertGroupsResponseBody) *ListAlertGroupsResponse
	GetBody() *ListAlertGroupsResponseBody
}

type ListAlertGroupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertGroupsResponse) GetBody() *ListAlertGroupsResponseBody {
	return s.Body
}

func (s *ListAlertGroupsResponse) SetHeaders(v map[string]*string) *ListAlertGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertGroupsResponse) SetStatusCode(v int32) *ListAlertGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertGroupsResponse) SetBody(v *ListAlertGroupsResponseBody) *ListAlertGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAlertGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
