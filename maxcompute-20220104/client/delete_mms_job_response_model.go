// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMmsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMmsJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMmsJobResponseBody) *DeleteMmsJobResponse
	GetBody() *DeleteMmsJobResponseBody
}

type DeleteMmsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMmsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmsJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteMmsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMmsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMmsJobResponse) GetBody() *DeleteMmsJobResponseBody {
	return s.Body
}

func (s *DeleteMmsJobResponse) SetHeaders(v map[string]*string) *DeleteMmsJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteMmsJobResponse) SetStatusCode(v int32) *DeleteMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMmsJobResponse) SetBody(v *DeleteMmsJobResponseBody) *DeleteMmsJobResponse {
	s.Body = v
	return s
}

func (s *DeleteMmsJobResponse) Validate() error {
	return dara.Validate(s)
}
