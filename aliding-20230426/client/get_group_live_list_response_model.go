// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupLiveListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGroupLiveListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGroupLiveListResponse
	GetStatusCode() *int32
	SetBody(v *GetGroupLiveListResponseBody) *GetGroupLiveListResponse
	GetBody() *GetGroupLiveListResponseBody
}

type GetGroupLiveListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupLiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupLiveListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListResponse) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGroupLiveListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGroupLiveListResponse) GetBody() *GetGroupLiveListResponseBody {
	return s.Body
}

func (s *GetGroupLiveListResponse) SetHeaders(v map[string]*string) *GetGroupLiveListResponse {
	s.Headers = v
	return s
}

func (s *GetGroupLiveListResponse) SetStatusCode(v int32) *GetGroupLiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupLiveListResponse) SetBody(v *GetGroupLiveListResponseBody) *GetGroupLiveListResponse {
	s.Body = v
	return s
}

func (s *GetGroupLiveListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
