// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefaultLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefaultLevelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefaultLevelResponseBody) *ModifyDefaultLevelResponse
	GetBody() *ModifyDefaultLevelResponseBody
}

type ModifyDefaultLevelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefaultLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefaultLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultLevelResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefaultLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefaultLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefaultLevelResponse) GetBody() *ModifyDefaultLevelResponseBody {
	return s.Body
}

func (s *ModifyDefaultLevelResponse) SetHeaders(v map[string]*string) *ModifyDefaultLevelResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefaultLevelResponse) SetStatusCode(v int32) *ModifyDefaultLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefaultLevelResponse) SetBody(v *ModifyDefaultLevelResponseBody) *ModifyDefaultLevelResponse {
	s.Body = v
	return s
}

func (s *ModifyDefaultLevelResponse) Validate() error {
	return dara.Validate(s)
}
