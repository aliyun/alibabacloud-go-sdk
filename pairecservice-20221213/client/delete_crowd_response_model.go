// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrowdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCrowdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCrowdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCrowdResponseBody) *DeleteCrowdResponse
	GetBody() *DeleteCrowdResponseBody
}

type DeleteCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCrowdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrowdResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrowdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCrowdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCrowdResponse) GetBody() *DeleteCrowdResponseBody {
	return s.Body
}

func (s *DeleteCrowdResponse) SetHeaders(v map[string]*string) *DeleteCrowdResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrowdResponse) SetStatusCode(v int32) *DeleteCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCrowdResponse) SetBody(v *DeleteCrowdResponseBody) *DeleteCrowdResponse {
	s.Body = v
	return s
}

func (s *DeleteCrowdResponse) Validate() error {
	return dara.Validate(s)
}
