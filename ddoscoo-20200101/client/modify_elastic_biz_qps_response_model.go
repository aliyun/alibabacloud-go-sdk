// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBizQpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticBizQpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticBizQpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticBizQpsResponseBody) *ModifyElasticBizQpsResponse
	GetBody() *ModifyElasticBizQpsResponseBody
}

type ModifyElasticBizQpsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticBizQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticBizQpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBizQpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticBizQpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticBizQpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticBizQpsResponse) GetBody() *ModifyElasticBizQpsResponseBody {
	return s.Body
}

func (s *ModifyElasticBizQpsResponse) SetHeaders(v map[string]*string) *ModifyElasticBizQpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticBizQpsResponse) SetStatusCode(v int32) *ModifyElasticBizQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticBizQpsResponse) SetBody(v *ModifyElasticBizQpsResponseBody) *ModifyElasticBizQpsResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticBizQpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
