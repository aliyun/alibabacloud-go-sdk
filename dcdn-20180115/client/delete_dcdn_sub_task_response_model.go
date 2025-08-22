// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnSubTaskResponseBody) *DeleteDcdnSubTaskResponse
	GetBody() *DeleteDcdnSubTaskResponseBody
}

type DeleteDcdnSubTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnSubTaskResponse) GetBody() *DeleteDcdnSubTaskResponseBody {
	return s.Body
}

func (s *DeleteDcdnSubTaskResponse) SetHeaders(v map[string]*string) *DeleteDcdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnSubTaskResponse) SetStatusCode(v int32) *DeleteDcdnSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnSubTaskResponse) SetBody(v *DeleteDcdnSubTaskResponseBody) *DeleteDcdnSubTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnSubTaskResponse) Validate() error {
	return dara.Validate(s)
}
