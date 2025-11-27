// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCElasticScalingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCElasticScalingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCElasticScalingResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCElasticScalingResponseBody) *ModifyRCElasticScalingResponse
	GetBody() *ModifyRCElasticScalingResponseBody
}

type ModifyRCElasticScalingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCElasticScalingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCElasticScalingResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCElasticScalingResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCElasticScalingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCElasticScalingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCElasticScalingResponse) GetBody() *ModifyRCElasticScalingResponseBody {
	return s.Body
}

func (s *ModifyRCElasticScalingResponse) SetHeaders(v map[string]*string) *ModifyRCElasticScalingResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCElasticScalingResponse) SetStatusCode(v int32) *ModifyRCElasticScalingResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCElasticScalingResponse) SetBody(v *ModifyRCElasticScalingResponseBody) *ModifyRCElasticScalingResponse {
	s.Body = v
	return s
}

func (s *ModifyRCElasticScalingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
