// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigDirResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigDirResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigDirResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigDirResponseBody) *CreateConfigDirResponse
	GetBody() *CreateConfigDirResponseBody
}

type CreateConfigDirResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigDirResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigDirResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigDirResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigDirResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigDirResponse) GetBody() *CreateConfigDirResponseBody {
	return s.Body
}

func (s *CreateConfigDirResponse) SetHeaders(v map[string]*string) *CreateConfigDirResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigDirResponse) SetStatusCode(v int32) *CreateConfigDirResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigDirResponse) SetBody(v *CreateConfigDirResponseBody) *CreateConfigDirResponse {
	s.Body = v
	return s
}

func (s *CreateConfigDirResponse) Validate() error {
	return dara.Validate(s)
}
