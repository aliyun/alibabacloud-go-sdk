// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAvailableAlertGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmAvailableAlertGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmAvailableAlertGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmAvailableAlertGroupsResponseBody) *ListCloudGtmAvailableAlertGroupsResponse
	GetBody() *ListCloudGtmAvailableAlertGroupsResponseBody
}

type ListCloudGtmAvailableAlertGroupsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmAvailableAlertGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmAvailableAlertGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAvailableAlertGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAvailableAlertGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmAvailableAlertGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmAvailableAlertGroupsResponse) GetBody() *ListCloudGtmAvailableAlertGroupsResponseBody {
	return s.Body
}

func (s *ListCloudGtmAvailableAlertGroupsResponse) SetHeaders(v map[string]*string) *ListCloudGtmAvailableAlertGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsResponse) SetStatusCode(v int32) *ListCloudGtmAvailableAlertGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsResponse) SetBody(v *ListCloudGtmAvailableAlertGroupsResponseBody) *ListCloudGtmAvailableAlertGroupsResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsResponse) Validate() error {
	return dara.Validate(s)
}
