// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnRealTimeLogProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnRealTimeLogProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnRealTimeLogProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnRealTimeLogProjectResponseBody) *DeleteDcdnRealTimeLogProjectResponse
	GetBody() *DeleteDcdnRealTimeLogProjectResponseBody
}

type DeleteDcdnRealTimeLogProjectResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnRealTimeLogProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnRealTimeLogProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnRealTimeLogProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnRealTimeLogProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnRealTimeLogProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnRealTimeLogProjectResponse) GetBody() *DeleteDcdnRealTimeLogProjectResponseBody {
	return s.Body
}

func (s *DeleteDcdnRealTimeLogProjectResponse) SetHeaders(v map[string]*string) *DeleteDcdnRealTimeLogProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectResponse) SetStatusCode(v int32) *DeleteDcdnRealTimeLogProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectResponse) SetBody(v *DeleteDcdnRealTimeLogProjectResponseBody) *DeleteDcdnRealTimeLogProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectResponse) Validate() error {
	return dara.Validate(s)
}
