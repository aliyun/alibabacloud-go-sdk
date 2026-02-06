// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmptyNumberNoMoreCallsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEmptyNumberNoMoreCallsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEmptyNumberNoMoreCallsInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEmptyNumberNoMoreCallsInfoResponseBody) *ModifyEmptyNumberNoMoreCallsInfoResponse
	GetBody() *ModifyEmptyNumberNoMoreCallsInfoResponseBody
}

type ModifyEmptyNumberNoMoreCallsInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEmptyNumberNoMoreCallsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEmptyNumberNoMoreCallsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmptyNumberNoMoreCallsInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponse) GetBody() *ModifyEmptyNumberNoMoreCallsInfoResponseBody {
	return s.Body
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponse) SetHeaders(v map[string]*string) *ModifyEmptyNumberNoMoreCallsInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponse) SetStatusCode(v int32) *ModifyEmptyNumberNoMoreCallsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponse) SetBody(v *ModifyEmptyNumberNoMoreCallsInfoResponseBody) *ModifyEmptyNumberNoMoreCallsInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
