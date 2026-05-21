// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDatasetImportDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenDatasetImportDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenDatasetImportDataResponse
	GetStatusCode() *int32
	SetBody(v *OpenDatasetImportDataResponseBody) *OpenDatasetImportDataResponse
	GetBody() *OpenDatasetImportDataResponseBody
}

type OpenDatasetImportDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDatasetImportDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDatasetImportDataResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetImportDataResponse) GoString() string {
	return s.String()
}

func (s *OpenDatasetImportDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenDatasetImportDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenDatasetImportDataResponse) GetBody() *OpenDatasetImportDataResponseBody {
	return s.Body
}

func (s *OpenDatasetImportDataResponse) SetHeaders(v map[string]*string) *OpenDatasetImportDataResponse {
	s.Headers = v
	return s
}

func (s *OpenDatasetImportDataResponse) SetStatusCode(v int32) *OpenDatasetImportDataResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDatasetImportDataResponse) SetBody(v *OpenDatasetImportDataResponseBody) *OpenDatasetImportDataResponse {
	s.Body = v
	return s
}

func (s *OpenDatasetImportDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
