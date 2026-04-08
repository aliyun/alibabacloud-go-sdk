// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDesensStatusListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgQueryDesensStatusListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgQueryDesensStatusListResponse
	GetStatusCode() *int32
	SetBody(v *DsgQueryDesensStatusListResponseBody) *DsgQueryDesensStatusListResponse
	GetBody() *DsgQueryDesensStatusListResponseBody
}

type DsgQueryDesensStatusListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgQueryDesensStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgQueryDesensStatusListResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDesensStatusListResponse) GoString() string {
	return s.String()
}

func (s *DsgQueryDesensStatusListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgQueryDesensStatusListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgQueryDesensStatusListResponse) GetBody() *DsgQueryDesensStatusListResponseBody {
	return s.Body
}

func (s *DsgQueryDesensStatusListResponse) SetHeaders(v map[string]*string) *DsgQueryDesensStatusListResponse {
	s.Headers = v
	return s
}

func (s *DsgQueryDesensStatusListResponse) SetStatusCode(v int32) *DsgQueryDesensStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgQueryDesensStatusListResponse) SetBody(v *DsgQueryDesensStatusListResponseBody) *DsgQueryDesensStatusListResponse {
	s.Body = v
	return s
}

func (s *DsgQueryDesensStatusListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
