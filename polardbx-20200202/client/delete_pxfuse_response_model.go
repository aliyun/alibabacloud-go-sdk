// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfuseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePxfuseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePxfuseResponse
	GetStatusCode() *int32
	SetBody(v *DeletePxfuseResponseBody) *DeletePxfuseResponse
	GetBody() *DeletePxfuseResponseBody
}

type DeletePxfuseResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePxfuseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePxfuseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfuseResponse) GoString() string {
	return s.String()
}

func (s *DeletePxfuseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePxfuseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePxfuseResponse) GetBody() *DeletePxfuseResponseBody {
	return s.Body
}

func (s *DeletePxfuseResponse) SetHeaders(v map[string]*string) *DeletePxfuseResponse {
	s.Headers = v
	return s
}

func (s *DeletePxfuseResponse) SetStatusCode(v int32) *DeletePxfuseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePxfuseResponse) SetBody(v *DeletePxfuseResponseBody) *DeletePxfuseResponse {
	s.Body = v
	return s
}

func (s *DeletePxfuseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
