// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBizBandWidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticBizBandWidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticBizBandWidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticBizBandWidthResponseBody) *ModifyElasticBizBandWidthResponse
	GetBody() *ModifyElasticBizBandWidthResponseBody
}

type ModifyElasticBizBandWidthResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticBizBandWidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticBizBandWidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBizBandWidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticBizBandWidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticBizBandWidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticBizBandWidthResponse) GetBody() *ModifyElasticBizBandWidthResponseBody {
	return s.Body
}

func (s *ModifyElasticBizBandWidthResponse) SetHeaders(v map[string]*string) *ModifyElasticBizBandWidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticBizBandWidthResponse) SetStatusCode(v int32) *ModifyElasticBizBandWidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticBizBandWidthResponse) SetBody(v *ModifyElasticBizBandWidthResponseBody) *ModifyElasticBizBandWidthResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticBizBandWidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
