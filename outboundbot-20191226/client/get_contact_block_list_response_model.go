// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactBlockListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContactBlockListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContactBlockListResponse
	GetStatusCode() *int32
	SetBody(v *GetContactBlockListResponseBody) *GetContactBlockListResponse
	GetBody() *GetContactBlockListResponseBody
}

type GetContactBlockListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContactBlockListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContactBlockListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContactBlockListResponse) GoString() string {
	return s.String()
}

func (s *GetContactBlockListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContactBlockListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContactBlockListResponse) GetBody() *GetContactBlockListResponseBody {
	return s.Body
}

func (s *GetContactBlockListResponse) SetHeaders(v map[string]*string) *GetContactBlockListResponse {
	s.Headers = v
	return s
}

func (s *GetContactBlockListResponse) SetStatusCode(v int32) *GetContactBlockListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactBlockListResponse) SetBody(v *GetContactBlockListResponseBody) *GetContactBlockListResponse {
	s.Body = v
	return s
}

func (s *GetContactBlockListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
