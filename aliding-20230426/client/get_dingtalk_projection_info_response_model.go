// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingtalkProjectionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingtalkProjectionInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDingtalkProjectionInfoResponseBody) *GetDingtalkProjectionInfoResponse
	GetBody() *GetDingtalkProjectionInfoResponseBody
}

type GetDingtalkProjectionInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingtalkProjectionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingtalkProjectionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingtalkProjectionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingtalkProjectionInfoResponse) GetBody() *GetDingtalkProjectionInfoResponseBody {
	return s.Body
}

func (s *GetDingtalkProjectionInfoResponse) SetHeaders(v map[string]*string) *GetDingtalkProjectionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDingtalkProjectionInfoResponse) SetStatusCode(v int32) *GetDingtalkProjectionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingtalkProjectionInfoResponse) SetBody(v *GetDingtalkProjectionInfoResponseBody) *GetDingtalkProjectionInfoResponse {
	s.Body = v
	return s
}

func (s *GetDingtalkProjectionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
