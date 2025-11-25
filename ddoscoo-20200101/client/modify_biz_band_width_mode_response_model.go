// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBizBandWidthModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBizBandWidthModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBizBandWidthModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBizBandWidthModeResponseBody) *ModifyBizBandWidthModeResponse
	GetBody() *ModifyBizBandWidthModeResponseBody
}

type ModifyBizBandWidthModeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBizBandWidthModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBizBandWidthModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBizBandWidthModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyBizBandWidthModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBizBandWidthModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBizBandWidthModeResponse) GetBody() *ModifyBizBandWidthModeResponseBody {
	return s.Body
}

func (s *ModifyBizBandWidthModeResponse) SetHeaders(v map[string]*string) *ModifyBizBandWidthModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyBizBandWidthModeResponse) SetStatusCode(v int32) *ModifyBizBandWidthModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBizBandWidthModeResponse) SetBody(v *ModifyBizBandWidthModeResponseBody) *ModifyBizBandWidthModeResponse {
	s.Body = v
	return s
}

func (s *ModifyBizBandWidthModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
