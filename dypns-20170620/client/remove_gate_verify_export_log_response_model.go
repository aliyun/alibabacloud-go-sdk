// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGateVerifyExportLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveGateVerifyExportLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveGateVerifyExportLogResponse
	GetStatusCode() *int32
	SetBody(v *RemoveGateVerifyExportLogResponseBody) *RemoveGateVerifyExportLogResponse
	GetBody() *RemoveGateVerifyExportLogResponseBody
}

type RemoveGateVerifyExportLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveGateVerifyExportLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveGateVerifyExportLogResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveGateVerifyExportLogResponse) GoString() string {
	return s.String()
}

func (s *RemoveGateVerifyExportLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveGateVerifyExportLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveGateVerifyExportLogResponse) GetBody() *RemoveGateVerifyExportLogResponseBody {
	return s.Body
}

func (s *RemoveGateVerifyExportLogResponse) SetHeaders(v map[string]*string) *RemoveGateVerifyExportLogResponse {
	s.Headers = v
	return s
}

func (s *RemoveGateVerifyExportLogResponse) SetStatusCode(v int32) *RemoveGateVerifyExportLogResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveGateVerifyExportLogResponse) SetBody(v *RemoveGateVerifyExportLogResponseBody) *RemoveGateVerifyExportLogResponse {
	s.Body = v
	return s
}

func (s *RemoveGateVerifyExportLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
