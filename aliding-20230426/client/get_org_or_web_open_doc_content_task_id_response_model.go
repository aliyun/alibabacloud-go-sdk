// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgOrWebOpenDocContentTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrgOrWebOpenDocContentTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetOrgOrWebOpenDocContentTaskIdResponseBody) *GetOrgOrWebOpenDocContentTaskIdResponse
	GetBody() *GetOrgOrWebOpenDocContentTaskIdResponseBody
}

type GetOrgOrWebOpenDocContentTaskIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgOrWebOpenDocContentTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgOrWebOpenDocContentTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponse) GetBody() *GetOrgOrWebOpenDocContentTaskIdResponseBody {
	return s.Body
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponse) SetHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponse) SetStatusCode(v int32) *GetOrgOrWebOpenDocContentTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponse) SetBody(v *GetOrgOrWebOpenDocContentTaskIdResponseBody) *GetOrgOrWebOpenDocContentTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
