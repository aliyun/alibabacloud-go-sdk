// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaterialInspectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MaterialInspectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MaterialInspectionResponse
	GetStatusCode() *int32
	SetBody(v *MaterialInspectionResponseBody) *MaterialInspectionResponse
	GetBody() *MaterialInspectionResponseBody
}

type MaterialInspectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MaterialInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MaterialInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s MaterialInspectionResponse) GoString() string {
	return s.String()
}

func (s *MaterialInspectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MaterialInspectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MaterialInspectionResponse) GetBody() *MaterialInspectionResponseBody {
	return s.Body
}

func (s *MaterialInspectionResponse) SetHeaders(v map[string]*string) *MaterialInspectionResponse {
	s.Headers = v
	return s
}

func (s *MaterialInspectionResponse) SetStatusCode(v int32) *MaterialInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *MaterialInspectionResponse) SetBody(v *MaterialInspectionResponseBody) *MaterialInspectionResponse {
	s.Body = v
	return s
}

func (s *MaterialInspectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
