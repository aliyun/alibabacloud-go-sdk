// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDriveSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDriveSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDriveSpaceResponse
	GetStatusCode() *int32
	SetBody(v *AddDriveSpaceResponseBody) *AddDriveSpaceResponse
	GetBody() *AddDriveSpaceResponseBody
}

type AddDriveSpaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDriveSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDriveSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceResponse) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDriveSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDriveSpaceResponse) GetBody() *AddDriveSpaceResponseBody {
	return s.Body
}

func (s *AddDriveSpaceResponse) SetHeaders(v map[string]*string) *AddDriveSpaceResponse {
	s.Headers = v
	return s
}

func (s *AddDriveSpaceResponse) SetStatusCode(v int32) *AddDriveSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDriveSpaceResponse) SetBody(v *AddDriveSpaceResponseBody) *AddDriveSpaceResponse {
	s.Body = v
	return s
}

func (s *AddDriveSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
