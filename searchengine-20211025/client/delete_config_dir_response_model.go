// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigDirResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigDirResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigDirResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigDirResponseBody) *DeleteConfigDirResponse
	GetBody() *DeleteConfigDirResponseBody
}

type DeleteConfigDirResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigDirResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigDirResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigDirResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigDirResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigDirResponse) GetBody() *DeleteConfigDirResponseBody {
	return s.Body
}

func (s *DeleteConfigDirResponse) SetHeaders(v map[string]*string) *DeleteConfigDirResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigDirResponse) SetStatusCode(v int32) *DeleteConfigDirResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigDirResponse) SetBody(v *DeleteConfigDirResponseBody) *DeleteConfigDirResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigDirResponse) Validate() error {
	return dara.Validate(s)
}
