// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveResourcesResponse
	GetStatusCode() *int32
	SetBody(v *MoveResourcesResponseBody) *MoveResourcesResponse
	GetBody() *MoveResourcesResponseBody
}

type MoveResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveResourcesResponse) GoString() string {
	return s.String()
}

func (s *MoveResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveResourcesResponse) GetBody() *MoveResourcesResponseBody {
	return s.Body
}

func (s *MoveResourcesResponse) SetHeaders(v map[string]*string) *MoveResourcesResponse {
	s.Headers = v
	return s
}

func (s *MoveResourcesResponse) SetStatusCode(v int32) *MoveResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourcesResponse) SetBody(v *MoveResourcesResponseBody) *MoveResourcesResponse {
	s.Body = v
	return s
}

func (s *MoveResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
