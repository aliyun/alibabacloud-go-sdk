// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferClusterBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferClusterBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferClusterBackupResponse
	GetStatusCode() *int32
	SetBody(v *TransferClusterBackupResponseBody) *TransferClusterBackupResponse
	GetBody() *TransferClusterBackupResponseBody
}

type TransferClusterBackupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferClusterBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferClusterBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferClusterBackupResponse) GoString() string {
	return s.String()
}

func (s *TransferClusterBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferClusterBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferClusterBackupResponse) GetBody() *TransferClusterBackupResponseBody {
	return s.Body
}

func (s *TransferClusterBackupResponse) SetHeaders(v map[string]*string) *TransferClusterBackupResponse {
	s.Headers = v
	return s
}

func (s *TransferClusterBackupResponse) SetStatusCode(v int32) *TransferClusterBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferClusterBackupResponse) SetBody(v *TransferClusterBackupResponseBody) *TransferClusterBackupResponse {
	s.Body = v
	return s
}

func (s *TransferClusterBackupResponse) Validate() error {
	return dara.Validate(s)
}
