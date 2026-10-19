// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeDetailResponseBody) *GetNodeDetailResponse
	GetBody() *GetNodeDetailResponseBody
}

type GetNodeDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponse) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeDetailResponse) GetBody() *GetNodeDetailResponseBody {
	return s.Body
}

func (s *GetNodeDetailResponse) SetHeaders(v map[string]*string) *GetNodeDetailResponse {
	s.Headers = v
	return s
}

func (s *GetNodeDetailResponse) SetStatusCode(v int32) *GetNodeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeDetailResponse) SetBody(v *GetNodeDetailResponseBody) *GetNodeDetailResponse {
	s.Body = v
	return s
}

func (s *GetNodeDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
