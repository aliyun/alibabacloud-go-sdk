// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMaterialTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMaterialTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMaterialTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMaterialTaskResponseBody) *SubmitMaterialTaskResponse
	GetBody() *SubmitMaterialTaskResponseBody
}

type SubmitMaterialTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMaterialTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMaterialTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMaterialTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitMaterialTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMaterialTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMaterialTaskResponse) GetBody() *SubmitMaterialTaskResponseBody {
	return s.Body
}

func (s *SubmitMaterialTaskResponse) SetHeaders(v map[string]*string) *SubmitMaterialTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitMaterialTaskResponse) SetStatusCode(v int32) *SubmitMaterialTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMaterialTaskResponse) SetBody(v *SubmitMaterialTaskResponseBody) *SubmitMaterialTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitMaterialTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
