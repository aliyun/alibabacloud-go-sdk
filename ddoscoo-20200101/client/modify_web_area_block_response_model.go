// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAreaBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebAreaBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebAreaBlockResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebAreaBlockResponseBody) *ModifyWebAreaBlockResponse
	GetBody() *ModifyWebAreaBlockResponseBody
}

type ModifyWebAreaBlockResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebAreaBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebAreaBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAreaBlockResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebAreaBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebAreaBlockResponse) GetBody() *ModifyWebAreaBlockResponseBody {
	return s.Body
}

func (s *ModifyWebAreaBlockResponse) SetHeaders(v map[string]*string) *ModifyWebAreaBlockResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAreaBlockResponse) SetStatusCode(v int32) *ModifyWebAreaBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebAreaBlockResponse) SetBody(v *ModifyWebAreaBlockResponseBody) *ModifyWebAreaBlockResponse {
	s.Body = v
	return s
}

func (s *ModifyWebAreaBlockResponse) Validate() error {
	return dara.Validate(s)
}
