// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFigureClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFigureClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFigureClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFigureClusterResponseBody) *UpdateFigureClusterResponse
	GetBody() *UpdateFigureClusterResponseBody
}

type UpdateFigureClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFigureClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFigureClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFigureClusterResponse) GetBody() *UpdateFigureClusterResponseBody {
	return s.Body
}

func (s *UpdateFigureClusterResponse) SetHeaders(v map[string]*string) *UpdateFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateFigureClusterResponse) SetStatusCode(v int32) *UpdateFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFigureClusterResponse) SetBody(v *UpdateFigureClusterResponseBody) *UpdateFigureClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateFigureClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
