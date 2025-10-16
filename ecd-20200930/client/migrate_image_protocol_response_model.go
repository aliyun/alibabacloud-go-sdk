// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateImageProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateImageProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateImageProtocolResponse
	GetStatusCode() *int32
	SetBody(v *MigrateImageProtocolResponseBody) *MigrateImageProtocolResponse
	GetBody() *MigrateImageProtocolResponseBody
}

type MigrateImageProtocolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateImageProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateImageProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateImageProtocolResponse) GoString() string {
	return s.String()
}

func (s *MigrateImageProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateImageProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateImageProtocolResponse) GetBody() *MigrateImageProtocolResponseBody {
	return s.Body
}

func (s *MigrateImageProtocolResponse) SetHeaders(v map[string]*string) *MigrateImageProtocolResponse {
	s.Headers = v
	return s
}

func (s *MigrateImageProtocolResponse) SetStatusCode(v int32) *MigrateImageProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateImageProtocolResponse) SetBody(v *MigrateImageProtocolResponseBody) *MigrateImageProtocolResponse {
	s.Body = v
	return s
}

func (s *MigrateImageProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
