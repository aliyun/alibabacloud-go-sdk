// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessKeyLeakDealResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccessKeyLeakDealResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccessKeyLeakDealResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccessKeyLeakDealResponseBody) *ModifyAccessKeyLeakDealResponse
	GetBody() *ModifyAccessKeyLeakDealResponseBody
}

type ModifyAccessKeyLeakDealResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccessKeyLeakDealResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccessKeyLeakDealResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessKeyLeakDealResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessKeyLeakDealResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccessKeyLeakDealResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccessKeyLeakDealResponse) GetBody() *ModifyAccessKeyLeakDealResponseBody {
	return s.Body
}

func (s *ModifyAccessKeyLeakDealResponse) SetHeaders(v map[string]*string) *ModifyAccessKeyLeakDealResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessKeyLeakDealResponse) SetStatusCode(v int32) *ModifyAccessKeyLeakDealResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessKeyLeakDealResponse) SetBody(v *ModifyAccessKeyLeakDealResponseBody) *ModifyAccessKeyLeakDealResponse {
	s.Body = v
	return s
}

func (s *ModifyAccessKeyLeakDealResponse) Validate() error {
	return dara.Validate(s)
}
