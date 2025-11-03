// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTargetsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTargetsResponseBody) *DeleteTargetsResponse
	GetBody() *DeleteTargetsResponseBody
}

type DeleteTargetsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTargetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTargetsResponse) GetBody() *DeleteTargetsResponseBody {
	return s.Body
}

func (s *DeleteTargetsResponse) SetHeaders(v map[string]*string) *DeleteTargetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTargetsResponse) SetStatusCode(v int32) *DeleteTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTargetsResponse) SetBody(v *DeleteTargetsResponseBody) *DeleteTargetsResponse {
	s.Body = v
	return s
}

func (s *DeleteTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
