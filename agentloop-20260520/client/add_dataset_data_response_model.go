// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatasetDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDatasetDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDatasetDataResponse
	GetStatusCode() *int32
	SetBody(v *AddDatasetDataResponseBody) *AddDatasetDataResponse
	GetBody() *AddDatasetDataResponseBody
}

type AddDatasetDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDatasetDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDatasetDataResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDataResponse) GoString() string {
	return s.String()
}

func (s *AddDatasetDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDatasetDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDatasetDataResponse) GetBody() *AddDatasetDataResponseBody {
	return s.Body
}

func (s *AddDatasetDataResponse) SetHeaders(v map[string]*string) *AddDatasetDataResponse {
	s.Headers = v
	return s
}

func (s *AddDatasetDataResponse) SetStatusCode(v int32) *AddDatasetDataResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDatasetDataResponse) SetBody(v *AddDatasetDataResponseBody) *AddDatasetDataResponse {
	s.Body = v
	return s
}

func (s *AddDatasetDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
