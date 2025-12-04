// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateErRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateErRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *UpdateErRouteMapResponseBody) *UpdateErRouteMapResponse
	GetBody() *UpdateErRouteMapResponseBody
}

type UpdateErRouteMapResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateErRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *UpdateErRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateErRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateErRouteMapResponse) GetBody() *UpdateErRouteMapResponseBody {
	return s.Body
}

func (s *UpdateErRouteMapResponse) SetHeaders(v map[string]*string) *UpdateErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *UpdateErRouteMapResponse) SetStatusCode(v int32) *UpdateErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateErRouteMapResponse) SetBody(v *UpdateErRouteMapResponseBody) *UpdateErRouteMapResponse {
	s.Body = v
	return s
}

func (s *UpdateErRouteMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
