// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenForUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkloadAccessTokenForUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkloadAccessTokenForUserIdResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkloadAccessTokenForUserIdResponseBody) *GetWorkloadAccessTokenForUserIdResponse
	GetBody() *GetWorkloadAccessTokenForUserIdResponseBody
}

type GetWorkloadAccessTokenForUserIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkloadAccessTokenForUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkloadAccessTokenForUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkloadAccessTokenForUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkloadAccessTokenForUserIdResponse) GetBody() *GetWorkloadAccessTokenForUserIdResponseBody {
	return s.Body
}

func (s *GetWorkloadAccessTokenForUserIdResponse) SetHeaders(v map[string]*string) *GetWorkloadAccessTokenForUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponse) SetStatusCode(v int32) *GetWorkloadAccessTokenForUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponse) SetBody(v *GetWorkloadAccessTokenForUserIdResponseBody) *GetWorkloadAccessTokenForUserIdResponse {
	s.Body = v
	return s
}

func (s *GetWorkloadAccessTokenForUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
