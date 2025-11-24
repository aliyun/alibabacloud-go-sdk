// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSwimLaneDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSwimLaneDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSwimLaneDetailResponseBody) *GetSwimLaneDetailResponse
	GetBody() *GetSwimLaneDetailResponseBody
}

type GetSwimLaneDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSwimLaneDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSwimLaneDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSwimLaneDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSwimLaneDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSwimLaneDetailResponse) GetBody() *GetSwimLaneDetailResponseBody {
	return s.Body
}

func (s *GetSwimLaneDetailResponse) SetHeaders(v map[string]*string) *GetSwimLaneDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSwimLaneDetailResponse) SetStatusCode(v int32) *GetSwimLaneDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwimLaneDetailResponse) SetBody(v *GetSwimLaneDetailResponseBody) *GetSwimLaneDetailResponse {
	s.Body = v
	return s
}

func (s *GetSwimLaneDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
