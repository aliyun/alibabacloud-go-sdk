// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDefaultDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDefaultDriveResponse
	GetStatusCode() *int32
	SetBody(v *Drive) *GetDefaultDriveResponse
	GetBody() *Drive
}

type GetDefaultDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultDriveResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDefaultDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDefaultDriveResponse) GetBody() *Drive {
	return s.Body
}

func (s *GetDefaultDriveResponse) SetHeaders(v map[string]*string) *GetDefaultDriveResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultDriveResponse) SetStatusCode(v int32) *GetDefaultDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultDriveResponse) SetBody(v *Drive) *GetDefaultDriveResponse {
	s.Body = v
	return s
}

func (s *GetDefaultDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
