// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaintainWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaintainWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaintainWindowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaintainWindowResponseBody) *DeleteMaintainWindowResponse
	GetBody() *DeleteMaintainWindowResponseBody
}

type DeleteMaintainWindowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaintainWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaintainWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaintainWindowResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaintainWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaintainWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaintainWindowResponse) GetBody() *DeleteMaintainWindowResponseBody {
	return s.Body
}

func (s *DeleteMaintainWindowResponse) SetHeaders(v map[string]*string) *DeleteMaintainWindowResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaintainWindowResponse) SetStatusCode(v int32) *DeleteMaintainWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaintainWindowResponse) SetBody(v *DeleteMaintainWindowResponseBody) *DeleteMaintainWindowResponse {
	s.Body = v
	return s
}

func (s *DeleteMaintainWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
