// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUpdateDesensStatusListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgUpdateDesensStatusListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgUpdateDesensStatusListResponse
	GetStatusCode() *int32
	SetBody(v *DsgUpdateDesensStatusListResponseBody) *DsgUpdateDesensStatusListResponse
	GetBody() *DsgUpdateDesensStatusListResponseBody
}

type DsgUpdateDesensStatusListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgUpdateDesensStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgUpdateDesensStatusListResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgUpdateDesensStatusListResponse) GoString() string {
	return s.String()
}

func (s *DsgUpdateDesensStatusListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgUpdateDesensStatusListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgUpdateDesensStatusListResponse) GetBody() *DsgUpdateDesensStatusListResponseBody {
	return s.Body
}

func (s *DsgUpdateDesensStatusListResponse) SetHeaders(v map[string]*string) *DsgUpdateDesensStatusListResponse {
	s.Headers = v
	return s
}

func (s *DsgUpdateDesensStatusListResponse) SetStatusCode(v int32) *DsgUpdateDesensStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgUpdateDesensStatusListResponse) SetBody(v *DsgUpdateDesensStatusListResponseBody) *DsgUpdateDesensStatusListResponse {
	s.Body = v
	return s
}

func (s *DsgUpdateDesensStatusListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
