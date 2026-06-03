// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGateVerifyExportLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGateVerifyExportLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGateVerifyExportLogResponse
	GetStatusCode() *int32
	SetBody(v *CreateGateVerifyExportLogResponseBody) *CreateGateVerifyExportLogResponse
	GetBody() *CreateGateVerifyExportLogResponseBody
}

type CreateGateVerifyExportLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGateVerifyExportLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGateVerifyExportLogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGateVerifyExportLogResponse) GoString() string {
	return s.String()
}

func (s *CreateGateVerifyExportLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGateVerifyExportLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGateVerifyExportLogResponse) GetBody() *CreateGateVerifyExportLogResponseBody {
	return s.Body
}

func (s *CreateGateVerifyExportLogResponse) SetHeaders(v map[string]*string) *CreateGateVerifyExportLogResponse {
	s.Headers = v
	return s
}

func (s *CreateGateVerifyExportLogResponse) SetStatusCode(v int32) *CreateGateVerifyExportLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGateVerifyExportLogResponse) SetBody(v *CreateGateVerifyExportLogResponseBody) *CreateGateVerifyExportLogResponse {
	s.Body = v
	return s
}

func (s *CreateGateVerifyExportLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
