// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBandWidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticBandWidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticBandWidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticBandWidthResponseBody) *ModifyElasticBandWidthResponse
	GetBody() *ModifyElasticBandWidthResponseBody
}

type ModifyElasticBandWidthResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticBandWidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticBandWidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBandWidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticBandWidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticBandWidthResponse) GetBody() *ModifyElasticBandWidthResponseBody {
	return s.Body
}

func (s *ModifyElasticBandWidthResponse) SetHeaders(v map[string]*string) *ModifyElasticBandWidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticBandWidthResponse) SetStatusCode(v int32) *ModifyElasticBandWidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticBandWidthResponse) SetBody(v *ModifyElasticBandWidthResponseBody) *ModifyElasticBandWidthResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticBandWidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
