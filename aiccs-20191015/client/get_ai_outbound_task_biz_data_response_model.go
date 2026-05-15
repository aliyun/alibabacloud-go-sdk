// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskBizDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiOutboundTaskBizDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiOutboundTaskBizDataResponse
	GetStatusCode() *int32
	SetBody(v *GetAiOutboundTaskBizDataResponseBody) *GetAiOutboundTaskBizDataResponse
	GetBody() *GetAiOutboundTaskBizDataResponseBody
}

type GetAiOutboundTaskBizDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiOutboundTaskBizDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiOutboundTaskBizDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskBizDataResponse) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskBizDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiOutboundTaskBizDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiOutboundTaskBizDataResponse) GetBody() *GetAiOutboundTaskBizDataResponseBody {
	return s.Body
}

func (s *GetAiOutboundTaskBizDataResponse) SetHeaders(v map[string]*string) *GetAiOutboundTaskBizDataResponse {
	s.Headers = v
	return s
}

func (s *GetAiOutboundTaskBizDataResponse) SetStatusCode(v int32) *GetAiOutboundTaskBizDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponse) SetBody(v *GetAiOutboundTaskBizDataResponseBody) *GetAiOutboundTaskBizDataResponse {
	s.Body = v
	return s
}

func (s *GetAiOutboundTaskBizDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
