// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticResourceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticResourceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticResourceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticResourceSpecResponseBody) *ModifyElasticResourceSpecResponse
	GetBody() *ModifyElasticResourceSpecResponseBody
}

type ModifyElasticResourceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticResourceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticResourceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticResourceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticResourceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticResourceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticResourceSpecResponse) GetBody() *ModifyElasticResourceSpecResponseBody {
	return s.Body
}

func (s *ModifyElasticResourceSpecResponse) SetHeaders(v map[string]*string) *ModifyElasticResourceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticResourceSpecResponse) SetStatusCode(v int32) *ModifyElasticResourceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticResourceSpecResponse) SetBody(v *ModifyElasticResourceSpecResponseBody) *ModifyElasticResourceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticResourceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
