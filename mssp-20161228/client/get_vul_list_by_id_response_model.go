// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulListByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulListByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulListByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetVulListByIdResponseBody) *GetVulListByIdResponse
	GetBody() *GetVulListByIdResponseBody
}

type GetVulListByIdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulListByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulListByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulListByIdResponse) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulListByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulListByIdResponse) GetBody() *GetVulListByIdResponseBody {
	return s.Body
}

func (s *GetVulListByIdResponse) SetHeaders(v map[string]*string) *GetVulListByIdResponse {
	s.Headers = v
	return s
}

func (s *GetVulListByIdResponse) SetStatusCode(v int32) *GetVulListByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulListByIdResponse) SetBody(v *GetVulListByIdResponseBody) *GetVulListByIdResponse {
	s.Body = v
	return s
}

func (s *GetVulListByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
