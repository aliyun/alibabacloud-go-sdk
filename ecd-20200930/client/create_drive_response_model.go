// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDriveResponse
	GetStatusCode() *int32
	SetBody(v *CreateDriveResponseBody) *CreateDriveResponse
	GetBody() *CreateDriveResponseBody
}

type CreateDriveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDriveResponse) GoString() string {
	return s.String()
}

func (s *CreateDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDriveResponse) GetBody() *CreateDriveResponseBody {
	return s.Body
}

func (s *CreateDriveResponse) SetHeaders(v map[string]*string) *CreateDriveResponse {
	s.Headers = v
	return s
}

func (s *CreateDriveResponse) SetStatusCode(v int32) *CreateDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDriveResponse) SetBody(v *CreateDriveResponseBody) *CreateDriveResponse {
	s.Body = v
	return s
}

func (s *CreateDriveResponse) Validate() error {
	return dara.Validate(s)
}
