// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInspectionTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInspectionTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInspectionTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListInspectionTemplatesResponseBody) *ListInspectionTemplatesResponse
	GetBody() *ListInspectionTemplatesResponseBody
}

type ListInspectionTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInspectionTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInspectionTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInspectionTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListInspectionTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInspectionTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInspectionTemplatesResponse) GetBody() *ListInspectionTemplatesResponseBody {
	return s.Body
}

func (s *ListInspectionTemplatesResponse) SetHeaders(v map[string]*string) *ListInspectionTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListInspectionTemplatesResponse) SetStatusCode(v int32) *ListInspectionTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInspectionTemplatesResponse) SetBody(v *ListInspectionTemplatesResponseBody) *ListInspectionTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListInspectionTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
