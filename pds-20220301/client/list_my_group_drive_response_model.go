// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyGroupDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMyGroupDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMyGroupDriveResponse
	GetStatusCode() *int32
	SetBody(v *ListMyGroupDriveResponseBody) *ListMyGroupDriveResponse
	GetBody() *ListMyGroupDriveResponseBody
}

type ListMyGroupDriveResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyGroupDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyGroupDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMyGroupDriveResponse) GoString() string {
	return s.String()
}

func (s *ListMyGroupDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMyGroupDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMyGroupDriveResponse) GetBody() *ListMyGroupDriveResponseBody {
	return s.Body
}

func (s *ListMyGroupDriveResponse) SetHeaders(v map[string]*string) *ListMyGroupDriveResponse {
	s.Headers = v
	return s
}

func (s *ListMyGroupDriveResponse) SetStatusCode(v int32) *ListMyGroupDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyGroupDriveResponse) SetBody(v *ListMyGroupDriveResponseBody) *ListMyGroupDriveResponse {
	s.Body = v
	return s
}

func (s *ListMyGroupDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
