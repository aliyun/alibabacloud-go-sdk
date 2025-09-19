// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertItemsResponseBody) *ListAlertItemsResponse
	GetBody() *ListAlertItemsResponseBody
}

type ListAlertItemsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertItemsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertItemsResponse) GetBody() *ListAlertItemsResponseBody {
	return s.Body
}

func (s *ListAlertItemsResponse) SetHeaders(v map[string]*string) *ListAlertItemsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertItemsResponse) SetStatusCode(v int32) *ListAlertItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertItemsResponse) SetBody(v *ListAlertItemsResponseBody) *ListAlertItemsResponse {
	s.Body = v
	return s
}

func (s *ListAlertItemsResponse) Validate() error {
	return dara.Validate(s)
}
