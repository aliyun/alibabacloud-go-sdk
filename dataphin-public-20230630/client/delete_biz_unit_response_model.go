// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBizUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBizUnitResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBizUnitResponseBody) *DeleteBizUnitResponse
	GetBody() *DeleteBizUnitResponseBody
}

type DeleteBizUnitResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBizUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBizUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBizUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBizUnitResponse) GetBody() *DeleteBizUnitResponseBody {
	return s.Body
}

func (s *DeleteBizUnitResponse) SetHeaders(v map[string]*string) *DeleteBizUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizUnitResponse) SetStatusCode(v int32) *DeleteBizUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBizUnitResponse) SetBody(v *DeleteBizUnitResponseBody) *DeleteBizUnitResponse {
	s.Body = v
	return s
}

func (s *DeleteBizUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
