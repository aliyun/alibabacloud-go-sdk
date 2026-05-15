// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOutbounNumListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOutbounNumListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOutbounNumListResponse
	GetStatusCode() *int32
	SetBody(v *GetOutbounNumListResponseBody) *GetOutbounNumListResponse
	GetBody() *GetOutbounNumListResponseBody
}

type GetOutbounNumListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOutbounNumListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOutbounNumListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOutbounNumListResponse) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOutbounNumListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOutbounNumListResponse) GetBody() *GetOutbounNumListResponseBody {
	return s.Body
}

func (s *GetOutbounNumListResponse) SetHeaders(v map[string]*string) *GetOutbounNumListResponse {
	s.Headers = v
	return s
}

func (s *GetOutbounNumListResponse) SetStatusCode(v int32) *GetOutbounNumListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOutbounNumListResponse) SetBody(v *GetOutbounNumListResponseBody) *GetOutbounNumListResponse {
	s.Body = v
	return s
}

func (s *GetOutbounNumListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
