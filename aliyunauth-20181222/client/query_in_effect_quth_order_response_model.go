// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInEffectQuthOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInEffectQuthOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInEffectQuthOrderResponse
	GetStatusCode() *int32
	SetBody(v *QueryInEffectQuthOrderResponseBody) *QueryInEffectQuthOrderResponse
	GetBody() *QueryInEffectQuthOrderResponseBody
}

type QueryInEffectQuthOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInEffectQuthOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInEffectQuthOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInEffectQuthOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryInEffectQuthOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInEffectQuthOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInEffectQuthOrderResponse) GetBody() *QueryInEffectQuthOrderResponseBody {
	return s.Body
}

func (s *QueryInEffectQuthOrderResponse) SetHeaders(v map[string]*string) *QueryInEffectQuthOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryInEffectQuthOrderResponse) SetStatusCode(v int32) *QueryInEffectQuthOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInEffectQuthOrderResponse) SetBody(v *QueryInEffectQuthOrderResponseBody) *QueryInEffectQuthOrderResponse {
	s.Body = v
	return s
}

func (s *QueryInEffectQuthOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
