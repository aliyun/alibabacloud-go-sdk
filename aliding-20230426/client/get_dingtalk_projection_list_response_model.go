// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingtalkProjectionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingtalkProjectionListResponse
	GetStatusCode() *int32
	SetBody(v *GetDingtalkProjectionListResponseBody) *GetDingtalkProjectionListResponse
	GetBody() *GetDingtalkProjectionListResponseBody
}

type GetDingtalkProjectionListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingtalkProjectionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingtalkProjectionListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListResponse) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingtalkProjectionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingtalkProjectionListResponse) GetBody() *GetDingtalkProjectionListResponseBody {
	return s.Body
}

func (s *GetDingtalkProjectionListResponse) SetHeaders(v map[string]*string) *GetDingtalkProjectionListResponse {
	s.Headers = v
	return s
}

func (s *GetDingtalkProjectionListResponse) SetStatusCode(v int32) *GetDingtalkProjectionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingtalkProjectionListResponse) SetBody(v *GetDingtalkProjectionListResponseBody) *GetDingtalkProjectionListResponse {
	s.Body = v
	return s
}

func (s *GetDingtalkProjectionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
