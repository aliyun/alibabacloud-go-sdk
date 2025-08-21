// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDSEntityValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDSEntityValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDSEntityValueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDSEntityValueResponseBody) *DeleteDSEntityValueResponse
	GetBody() *DeleteDSEntityValueResponseBody
}

type DeleteDSEntityValueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDSEntityValueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDSEntityValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDSEntityValueResponse) GetBody() *DeleteDSEntityValueResponseBody {
	return s.Body
}

func (s *DeleteDSEntityValueResponse) SetHeaders(v map[string]*string) *DeleteDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *DeleteDSEntityValueResponse) SetStatusCode(v int32) *DeleteDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDSEntityValueResponse) SetBody(v *DeleteDSEntityValueResponseBody) *DeleteDSEntityValueResponse {
	s.Body = v
	return s
}

func (s *DeleteDSEntityValueResponse) Validate() error {
	return dara.Validate(s)
}
