// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertActionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertActionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertActionsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertActionsResponseBody) *DeleteAlertActionsResponse
	GetBody() *DeleteAlertActionsResponseBody
}

type DeleteAlertActionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertActionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertActionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertActionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertActionsResponse) GetBody() *DeleteAlertActionsResponseBody {
	return s.Body
}

func (s *DeleteAlertActionsResponse) SetHeaders(v map[string]*string) *DeleteAlertActionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertActionsResponse) SetStatusCode(v int32) *DeleteAlertActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertActionsResponse) SetBody(v *DeleteAlertActionsResponseBody) *DeleteAlertActionsResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertActionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
