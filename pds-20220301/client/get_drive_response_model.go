// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDriveResponse
	GetStatusCode() *int32
	SetBody(v *Drive) *GetDriveResponse
	GetBody() *Drive
}

type GetDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDriveResponse) GoString() string {
	return s.String()
}

func (s *GetDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDriveResponse) GetBody() *Drive {
	return s.Body
}

func (s *GetDriveResponse) SetHeaders(v map[string]*string) *GetDriveResponse {
	s.Headers = v
	return s
}

func (s *GetDriveResponse) SetStatusCode(v int32) *GetDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDriveResponse) SetBody(v *Drive) *GetDriveResponse {
	s.Body = v
	return s
}

func (s *GetDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
