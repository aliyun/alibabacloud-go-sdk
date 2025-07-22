// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMPULayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMPULayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMPULayoutResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMPULayoutResponseBody) *DeleteMPULayoutResponse
	GetBody() *DeleteMPULayoutResponseBody
}

type DeleteMPULayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMPULayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMPULayoutResponse) GoString() string {
	return s.String()
}

func (s *DeleteMPULayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMPULayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMPULayoutResponse) GetBody() *DeleteMPULayoutResponseBody {
	return s.Body
}

func (s *DeleteMPULayoutResponse) SetHeaders(v map[string]*string) *DeleteMPULayoutResponse {
	s.Headers = v
	return s
}

func (s *DeleteMPULayoutResponse) SetStatusCode(v int32) *DeleteMPULayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMPULayoutResponse) SetBody(v *DeleteMPULayoutResponseBody) *DeleteMPULayoutResponse {
	s.Body = v
	return s
}

func (s *DeleteMPULayoutResponse) Validate() error {
	return dara.Validate(s)
}
