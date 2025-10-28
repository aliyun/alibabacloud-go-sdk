// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefineRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserDefineRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserDefineRegionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserDefineRegionResponseBody) *DeleteUserDefineRegionResponse
	GetBody() *DeleteUserDefineRegionResponseBody
}

type DeleteUserDefineRegionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserDefineRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserDefineRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefineRegionResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserDefineRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserDefineRegionResponse) GetBody() *DeleteUserDefineRegionResponseBody {
	return s.Body
}

func (s *DeleteUserDefineRegionResponse) SetHeaders(v map[string]*string) *DeleteUserDefineRegionResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDefineRegionResponse) SetStatusCode(v int32) *DeleteUserDefineRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDefineRegionResponse) SetBody(v *DeleteUserDefineRegionResponseBody) *DeleteUserDefineRegionResponse {
	s.Body = v
	return s
}

func (s *DeleteUserDefineRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
