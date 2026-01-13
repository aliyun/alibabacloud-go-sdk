// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVCUInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVCUInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVCUInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVCUInstanceResponseBody) *DeleteVCUInstanceResponse
	GetBody() *DeleteVCUInstanceResponseBody
}

type DeleteVCUInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVCUInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVCUInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVCUInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVCUInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVCUInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVCUInstanceResponse) GetBody() *DeleteVCUInstanceResponseBody {
	return s.Body
}

func (s *DeleteVCUInstanceResponse) SetHeaders(v map[string]*string) *DeleteVCUInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVCUInstanceResponse) SetStatusCode(v int32) *DeleteVCUInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVCUInstanceResponse) SetBody(v *DeleteVCUInstanceResponseBody) *DeleteVCUInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteVCUInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
