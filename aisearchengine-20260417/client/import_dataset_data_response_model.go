// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDatasetDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportDatasetDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportDatasetDataResponse
	GetStatusCode() *int32
	SetBody(v *ImportDatasetDataResponseBody) *ImportDatasetDataResponse
	GetBody() *ImportDatasetDataResponseBody
}

type ImportDatasetDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportDatasetDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportDatasetDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportDatasetDataResponse) GoString() string {
	return s.String()
}

func (s *ImportDatasetDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportDatasetDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportDatasetDataResponse) GetBody() *ImportDatasetDataResponseBody {
	return s.Body
}

func (s *ImportDatasetDataResponse) SetHeaders(v map[string]*string) *ImportDatasetDataResponse {
	s.Headers = v
	return s
}

func (s *ImportDatasetDataResponse) SetStatusCode(v int32) *ImportDatasetDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportDatasetDataResponse) SetBody(v *ImportDatasetDataResponseBody) *ImportDatasetDataResponse {
	s.Body = v
	return s
}

func (s *ImportDatasetDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
