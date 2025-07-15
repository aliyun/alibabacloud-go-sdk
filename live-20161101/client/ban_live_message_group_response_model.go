// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBanLiveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BanLiveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BanLiveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *BanLiveMessageGroupResponseBody) *BanLiveMessageGroupResponse
	GetBody() *BanLiveMessageGroupResponseBody
}

type BanLiveMessageGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BanLiveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BanLiveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s BanLiveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *BanLiveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BanLiveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BanLiveMessageGroupResponse) GetBody() *BanLiveMessageGroupResponseBody {
	return s.Body
}

func (s *BanLiveMessageGroupResponse) SetHeaders(v map[string]*string) *BanLiveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *BanLiveMessageGroupResponse) SetStatusCode(v int32) *BanLiveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *BanLiveMessageGroupResponse) SetBody(v *BanLiveMessageGroupResponseBody) *BanLiveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *BanLiveMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
