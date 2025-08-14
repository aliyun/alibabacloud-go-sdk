// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartVoiceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmartVoiceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmartVoiceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListSmartVoiceGroupsResponseBody) *ListSmartVoiceGroupsResponse
	GetBody() *ListSmartVoiceGroupsResponseBody
}

type ListSmartVoiceGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmartVoiceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmartVoiceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmartVoiceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSmartVoiceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmartVoiceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmartVoiceGroupsResponse) GetBody() *ListSmartVoiceGroupsResponseBody {
	return s.Body
}

func (s *ListSmartVoiceGroupsResponse) SetHeaders(v map[string]*string) *ListSmartVoiceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSmartVoiceGroupsResponse) SetStatusCode(v int32) *ListSmartVoiceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmartVoiceGroupsResponse) SetBody(v *ListSmartVoiceGroupsResponseBody) *ListSmartVoiceGroupsResponse {
	s.Body = v
	return s
}

func (s *ListSmartVoiceGroupsResponse) Validate() error {
	return dara.Validate(s)
}
