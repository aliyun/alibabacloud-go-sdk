// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmptyNumberNoMoreCallsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmptyNumberNoMoreCallsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmptyNumberNoMoreCallsInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetEmptyNumberNoMoreCallsInfoResponseBody) *GetEmptyNumberNoMoreCallsInfoResponse
	GetBody() *GetEmptyNumberNoMoreCallsInfoResponseBody
}

type GetEmptyNumberNoMoreCallsInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmptyNumberNoMoreCallsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmptyNumberNoMoreCallsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmptyNumberNoMoreCallsInfoResponse) GoString() string {
	return s.String()
}

func (s *GetEmptyNumberNoMoreCallsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmptyNumberNoMoreCallsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmptyNumberNoMoreCallsInfoResponse) GetBody() *GetEmptyNumberNoMoreCallsInfoResponseBody {
	return s.Body
}

func (s *GetEmptyNumberNoMoreCallsInfoResponse) SetHeaders(v map[string]*string) *GetEmptyNumberNoMoreCallsInfoResponse {
	s.Headers = v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponse) SetStatusCode(v int32) *GetEmptyNumberNoMoreCallsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponse) SetBody(v *GetEmptyNumberNoMoreCallsInfoResponseBody) *GetEmptyNumberNoMoreCallsInfoResponse {
	s.Body = v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
