// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountJobByConditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CountJobByConditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CountJobByConditionResponse
	GetStatusCode() *int32
	SetBody(v *CountJobByConditionResponseBody) *CountJobByConditionResponse
	GetBody() *CountJobByConditionResponseBody
}

type CountJobByConditionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountJobByConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountJobByConditionResponse) String() string {
	return dara.Prettify(s)
}

func (s CountJobByConditionResponse) GoString() string {
	return s.String()
}

func (s *CountJobByConditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CountJobByConditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CountJobByConditionResponse) GetBody() *CountJobByConditionResponseBody {
	return s.Body
}

func (s *CountJobByConditionResponse) SetHeaders(v map[string]*string) *CountJobByConditionResponse {
	s.Headers = v
	return s
}

func (s *CountJobByConditionResponse) SetStatusCode(v int32) *CountJobByConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *CountJobByConditionResponse) SetBody(v *CountJobByConditionResponseBody) *CountJobByConditionResponse {
	s.Body = v
	return s
}

func (s *CountJobByConditionResponse) Validate() error {
	return dara.Validate(s)
}
