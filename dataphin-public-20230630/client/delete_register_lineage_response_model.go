// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegisterLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRegisterLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRegisterLineageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRegisterLineageResponseBody) *DeleteRegisterLineageResponse
	GetBody() *DeleteRegisterLineageResponseBody
}

type DeleteRegisterLineageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRegisterLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRegisterLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRegisterLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRegisterLineageResponse) GetBody() *DeleteRegisterLineageResponseBody {
	return s.Body
}

func (s *DeleteRegisterLineageResponse) SetHeaders(v map[string]*string) *DeleteRegisterLineageResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegisterLineageResponse) SetStatusCode(v int32) *DeleteRegisterLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegisterLineageResponse) SetBody(v *DeleteRegisterLineageResponseBody) *DeleteRegisterLineageResponse {
	s.Body = v
	return s
}

func (s *DeleteRegisterLineageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
