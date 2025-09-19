// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRefreshProcessInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRefreshProcessInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRefreshProcessInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRefreshProcessInfoResponseBody) *ModifyRefreshProcessInfoResponse
	GetBody() *ModifyRefreshProcessInfoResponseBody
}

type ModifyRefreshProcessInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRefreshProcessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRefreshProcessInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRefreshProcessInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyRefreshProcessInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRefreshProcessInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRefreshProcessInfoResponse) GetBody() *ModifyRefreshProcessInfoResponseBody {
	return s.Body
}

func (s *ModifyRefreshProcessInfoResponse) SetHeaders(v map[string]*string) *ModifyRefreshProcessInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyRefreshProcessInfoResponse) SetStatusCode(v int32) *ModifyRefreshProcessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRefreshProcessInfoResponse) SetBody(v *ModifyRefreshProcessInfoResponseBody) *ModifyRefreshProcessInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyRefreshProcessInfoResponse) Validate() error {
	return dara.Validate(s)
}
