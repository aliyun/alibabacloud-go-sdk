// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSwimLaneGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSwimLaneGroupListResponse
	GetStatusCode() *int32
	SetBody(v *GetSwimLaneGroupListResponseBody) *GetSwimLaneGroupListResponse
	GetBody() *GetSwimLaneGroupListResponseBody
}

type GetSwimLaneGroupListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSwimLaneGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSwimLaneGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneGroupListResponse) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSwimLaneGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSwimLaneGroupListResponse) GetBody() *GetSwimLaneGroupListResponseBody {
	return s.Body
}

func (s *GetSwimLaneGroupListResponse) SetHeaders(v map[string]*string) *GetSwimLaneGroupListResponse {
	s.Headers = v
	return s
}

func (s *GetSwimLaneGroupListResponse) SetStatusCode(v int32) *GetSwimLaneGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwimLaneGroupListResponse) SetBody(v *GetSwimLaneGroupListResponseBody) *GetSwimLaneGroupListResponse {
	s.Body = v
	return s
}

func (s *GetSwimLaneGroupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
