// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActivityListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetActivityListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetActivityListResponse
	GetStatusCode() *int32
	SetBody(v *GetActivityListResponseBody) *GetActivityListResponse
	GetBody() *GetActivityListResponseBody
}

type GetActivityListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetActivityListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetActivityListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetActivityListResponse) GoString() string {
	return s.String()
}

func (s *GetActivityListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetActivityListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetActivityListResponse) GetBody() *GetActivityListResponseBody {
	return s.Body
}

func (s *GetActivityListResponse) SetHeaders(v map[string]*string) *GetActivityListResponse {
	s.Headers = v
	return s
}

func (s *GetActivityListResponse) SetStatusCode(v int32) *GetActivityListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActivityListResponse) SetBody(v *GetActivityListResponseBody) *GetActivityListResponse {
	s.Body = v
	return s
}

func (s *GetActivityListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
