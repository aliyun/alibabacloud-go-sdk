// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBInstanceConnectivityDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDBInstanceConnectivityDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDBInstanceConnectivityDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *GetDBInstanceConnectivityDiagnosisResponseBody) *GetDBInstanceConnectivityDiagnosisResponse
	GetBody() *GetDBInstanceConnectivityDiagnosisResponseBody
}

type GetDBInstanceConnectivityDiagnosisResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDBInstanceConnectivityDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDBInstanceConnectivityDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceConnectivityDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *GetDBInstanceConnectivityDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDBInstanceConnectivityDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDBInstanceConnectivityDiagnosisResponse) GetBody() *GetDBInstanceConnectivityDiagnosisResponseBody {
	return s.Body
}

func (s *GetDBInstanceConnectivityDiagnosisResponse) SetHeaders(v map[string]*string) *GetDBInstanceConnectivityDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponse) SetStatusCode(v int32) *GetDBInstanceConnectivityDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponse) SetBody(v *GetDBInstanceConnectivityDiagnosisResponseBody) *GetDBInstanceConnectivityDiagnosisResponse {
	s.Body = v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponse) Validate() error {
	return dara.Validate(s)
}
