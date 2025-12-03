// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTestcaseListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTestcaseListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTestcaseListResponse
	GetStatusCode() *int32
	SetBody(v *GetTestcaseListResponseBody) *GetTestcaseListResponse
	GetBody() *GetTestcaseListResponseBody
}

type GetTestcaseListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTestcaseListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTestcaseListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListResponse) GoString() string {
	return s.String()
}

func (s *GetTestcaseListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTestcaseListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTestcaseListResponse) GetBody() *GetTestcaseListResponseBody {
	return s.Body
}

func (s *GetTestcaseListResponse) SetHeaders(v map[string]*string) *GetTestcaseListResponse {
	s.Headers = v
	return s
}

func (s *GetTestcaseListResponse) SetStatusCode(v int32) *GetTestcaseListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTestcaseListResponse) SetBody(v *GetTestcaseListResponseBody) *GetTestcaseListResponse {
	s.Body = v
	return s
}

func (s *GetTestcaseListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
