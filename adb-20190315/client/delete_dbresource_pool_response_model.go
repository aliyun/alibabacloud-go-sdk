// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourcePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBResourcePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBResourcePoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBResourcePoolResponseBody) *DeleteDBResourcePoolResponse
	GetBody() *DeleteDBResourcePoolResponseBody
}

type DeleteDBResourcePoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBResourcePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBResourcePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBResourcePoolResponse) GetBody() *DeleteDBResourcePoolResponseBody {
	return s.Body
}

func (s *DeleteDBResourcePoolResponse) SetHeaders(v map[string]*string) *DeleteDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResourcePoolResponse) SetStatusCode(v int32) *DeleteDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResourcePoolResponse) SetBody(v *DeleteDBResourcePoolResponseBody) *DeleteDBResourcePoolResponse {
	s.Body = v
	return s
}

func (s *DeleteDBResourcePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
