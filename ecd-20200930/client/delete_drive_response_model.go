// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDriveResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDriveResponseBody) *DeleteDriveResponse
	GetBody() *DeleteDriveResponseBody
}

type DeleteDriveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveResponse) GoString() string {
	return s.String()
}

func (s *DeleteDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDriveResponse) GetBody() *DeleteDriveResponseBody {
	return s.Body
}

func (s *DeleteDriveResponse) SetHeaders(v map[string]*string) *DeleteDriveResponse {
	s.Headers = v
	return s
}

func (s *DeleteDriveResponse) SetStatusCode(v int32) *DeleteDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDriveResponse) SetBody(v *DeleteDriveResponseBody) *DeleteDriveResponse {
	s.Body = v
	return s
}

func (s *DeleteDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
