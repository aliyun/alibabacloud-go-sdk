// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportDriveResponse
	GetStatusCode() *int32
	SetBody(v *ImportDriveResponseBody) *ImportDriveResponse
	GetBody() *ImportDriveResponseBody
}

type ImportDriveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportDriveResponse) GoString() string {
	return s.String()
}

func (s *ImportDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportDriveResponse) GetBody() *ImportDriveResponseBody {
	return s.Body
}

func (s *ImportDriveResponse) SetHeaders(v map[string]*string) *ImportDriveResponse {
	s.Headers = v
	return s
}

func (s *ImportDriveResponse) SetStatusCode(v int32) *ImportDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportDriveResponse) SetBody(v *ImportDriveResponseBody) *ImportDriveResponse {
	s.Body = v
	return s
}

func (s *ImportDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
