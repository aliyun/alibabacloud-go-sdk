// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApStatusByGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApStatusByGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApStatusByGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *GetApStatusByGroupIdResponseBody) *GetApStatusByGroupIdResponse
	GetBody() *GetApStatusByGroupIdResponseBody
}

type GetApStatusByGroupIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApStatusByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApStatusByGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApStatusByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *GetApStatusByGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApStatusByGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApStatusByGroupIdResponse) GetBody() *GetApStatusByGroupIdResponseBody {
	return s.Body
}

func (s *GetApStatusByGroupIdResponse) SetHeaders(v map[string]*string) *GetApStatusByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *GetApStatusByGroupIdResponse) SetStatusCode(v int32) *GetApStatusByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApStatusByGroupIdResponse) SetBody(v *GetApStatusByGroupIdResponseBody) *GetApStatusByGroupIdResponse {
	s.Body = v
	return s
}

func (s *GetApStatusByGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
