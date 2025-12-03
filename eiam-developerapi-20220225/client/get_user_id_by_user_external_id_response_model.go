// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUserExternalIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserIdByUserExternalIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserIdByUserExternalIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUserIdByUserExternalIdResponseBody) *GetUserIdByUserExternalIdResponse
	GetBody() *GetUserIdByUserExternalIdResponseBody
}

type GetUserIdByUserExternalIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByUserExternalIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByUserExternalIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUserExternalIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserIdByUserExternalIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserIdByUserExternalIdResponse) GetBody() *GetUserIdByUserExternalIdResponseBody {
	return s.Body
}

func (s *GetUserIdByUserExternalIdResponse) SetHeaders(v map[string]*string) *GetUserIdByUserExternalIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByUserExternalIdResponse) SetStatusCode(v int32) *GetUserIdByUserExternalIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByUserExternalIdResponse) SetBody(v *GetUserIdByUserExternalIdResponseBody) *GetUserIdByUserExternalIdResponse {
	s.Body = v
	return s
}

func (s *GetUserIdByUserExternalIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
