// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebPathResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebPathResponseBody) *ModifyWebPathResponse
	GetBody() *ModifyWebPathResponseBody
}

type ModifyWebPathResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebPathResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPathResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebPathResponse) GetBody() *ModifyWebPathResponseBody {
	return s.Body
}

func (s *ModifyWebPathResponse) SetHeaders(v map[string]*string) *ModifyWebPathResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebPathResponse) SetStatusCode(v int32) *ModifyWebPathResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebPathResponse) SetBody(v *ModifyWebPathResponseBody) *ModifyWebPathResponse {
	s.Body = v
	return s
}

func (s *ModifyWebPathResponse) Validate() error {
	return dara.Validate(s)
}
