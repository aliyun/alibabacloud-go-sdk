// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertActionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertActionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertActionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertActionsResponseBody) *ListAlertActionsResponse
	GetBody() *ListAlertActionsResponseBody
}

type ListAlertActionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertActionsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertActionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertActionsResponse) GetBody() *ListAlertActionsResponseBody {
	return s.Body
}

func (s *ListAlertActionsResponse) SetHeaders(v map[string]*string) *ListAlertActionsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertActionsResponse) SetStatusCode(v int32) *ListAlertActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertActionsResponse) SetBody(v *ListAlertActionsResponseBody) *ListAlertActionsResponse {
	s.Body = v
	return s
}

func (s *ListAlertActionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
