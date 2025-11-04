// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamTagListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStreamTagListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStreamTagListResponse
	GetStatusCode() *int32
	SetBody(v *GetStreamTagListResponseBody) *GetStreamTagListResponse
	GetBody() *GetStreamTagListResponseBody
}

type GetStreamTagListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStreamTagListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStreamTagListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTagListResponse) GoString() string {
	return s.String()
}

func (s *GetStreamTagListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStreamTagListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStreamTagListResponse) GetBody() *GetStreamTagListResponseBody {
	return s.Body
}

func (s *GetStreamTagListResponse) SetHeaders(v map[string]*string) *GetStreamTagListResponse {
	s.Headers = v
	return s
}

func (s *GetStreamTagListResponse) SetStatusCode(v int32) *GetStreamTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStreamTagListResponse) SetBody(v *GetStreamTagListResponseBody) *GetStreamTagListResponse {
	s.Body = v
	return s
}

func (s *GetStreamTagListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
