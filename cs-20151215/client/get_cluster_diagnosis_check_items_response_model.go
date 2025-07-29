// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDiagnosisCheckItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterDiagnosisCheckItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterDiagnosisCheckItemsResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterDiagnosisCheckItemsResponseBody) *GetClusterDiagnosisCheckItemsResponse
	GetBody() *GetClusterDiagnosisCheckItemsResponseBody
}

type GetClusterDiagnosisCheckItemsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterDiagnosisCheckItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterDiagnosisCheckItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDiagnosisCheckItemsResponse) GoString() string {
	return s.String()
}

func (s *GetClusterDiagnosisCheckItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterDiagnosisCheckItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterDiagnosisCheckItemsResponse) GetBody() *GetClusterDiagnosisCheckItemsResponseBody {
	return s.Body
}

func (s *GetClusterDiagnosisCheckItemsResponse) SetHeaders(v map[string]*string) *GetClusterDiagnosisCheckItemsResponse {
	s.Headers = v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponse) SetStatusCode(v int32) *GetClusterDiagnosisCheckItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponse) SetBody(v *GetClusterDiagnosisCheckItemsResponseBody) *GetClusterDiagnosisCheckItemsResponse {
	s.Body = v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponse) Validate() error {
	return dara.Validate(s)
}
