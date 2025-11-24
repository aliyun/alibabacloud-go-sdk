// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSwimLaneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSwimLaneListResponse
	GetStatusCode() *int32
	SetBody(v *GetSwimLaneListResponseBody) *GetSwimLaneListResponse
	GetBody() *GetSwimLaneListResponseBody
}

type GetSwimLaneListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSwimLaneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSwimLaneListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneListResponse) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSwimLaneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSwimLaneListResponse) GetBody() *GetSwimLaneListResponseBody {
	return s.Body
}

func (s *GetSwimLaneListResponse) SetHeaders(v map[string]*string) *GetSwimLaneListResponse {
	s.Headers = v
	return s
}

func (s *GetSwimLaneListResponse) SetStatusCode(v int32) *GetSwimLaneListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwimLaneListResponse) SetBody(v *GetSwimLaneListResponseBody) *GetSwimLaneListResponse {
	s.Body = v
	return s
}

func (s *GetSwimLaneListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
