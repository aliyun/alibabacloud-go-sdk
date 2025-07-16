// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgLiveListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrgLiveListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrgLiveListResponse
	GetStatusCode() *int32
	SetBody(v *GetOrgLiveListResponseBody) *GetOrgLiveListResponse
	GetBody() *GetOrgLiveListResponseBody
}

type GetOrgLiveListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgLiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgLiveListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListResponse) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrgLiveListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrgLiveListResponse) GetBody() *GetOrgLiveListResponseBody {
	return s.Body
}

func (s *GetOrgLiveListResponse) SetHeaders(v map[string]*string) *GetOrgLiveListResponse {
	s.Headers = v
	return s
}

func (s *GetOrgLiveListResponse) SetStatusCode(v int32) *GetOrgLiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgLiveListResponse) SetBody(v *GetOrgLiveListResponseBody) *GetOrgLiveListResponse {
	s.Body = v
	return s
}

func (s *GetOrgLiveListResponse) Validate() error {
	return dara.Validate(s)
}
