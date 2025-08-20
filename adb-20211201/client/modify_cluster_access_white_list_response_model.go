// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAccessWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterAccessWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterAccessWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterAccessWhiteListResponseBody) *ModifyClusterAccessWhiteListResponse
	GetBody() *ModifyClusterAccessWhiteListResponseBody
}

type ModifyClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterAccessWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterAccessWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterAccessWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterAccessWhiteListResponse) GetBody() *ModifyClusterAccessWhiteListResponseBody {
	return s.Body
}

func (s *ModifyClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *ModifyClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterAccessWhiteListResponse) SetStatusCode(v int32) *ModifyClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterAccessWhiteListResponse) SetBody(v *ModifyClusterAccessWhiteListResponseBody) *ModifyClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterAccessWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
