// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserInRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserInRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *GetUserInRecycleBinResponseBody) *GetUserInRecycleBinResponse
	GetBody() *GetUserInRecycleBinResponseBody
}

type GetUserInRecycleBinResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserInRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserInRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserInRecycleBinResponse) GetBody() *GetUserInRecycleBinResponseBody {
	return s.Body
}

func (s *GetUserInRecycleBinResponse) SetHeaders(v map[string]*string) *GetUserInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *GetUserInRecycleBinResponse) SetStatusCode(v int32) *GetUserInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInRecycleBinResponse) SetBody(v *GetUserInRecycleBinResponseBody) *GetUserInRecycleBinResponse {
	s.Body = v
	return s
}

func (s *GetUserInRecycleBinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
