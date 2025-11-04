// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartSysAvatarModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmartSysAvatarModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmartSysAvatarModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListSmartSysAvatarModelsResponseBody) *ListSmartSysAvatarModelsResponse
	GetBody() *ListSmartSysAvatarModelsResponseBody
}

type ListSmartSysAvatarModelsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmartSysAvatarModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmartSysAvatarModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmartSysAvatarModelsResponse) GoString() string {
	return s.String()
}

func (s *ListSmartSysAvatarModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmartSysAvatarModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmartSysAvatarModelsResponse) GetBody() *ListSmartSysAvatarModelsResponseBody {
	return s.Body
}

func (s *ListSmartSysAvatarModelsResponse) SetHeaders(v map[string]*string) *ListSmartSysAvatarModelsResponse {
	s.Headers = v
	return s
}

func (s *ListSmartSysAvatarModelsResponse) SetStatusCode(v int32) *ListSmartSysAvatarModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponse) SetBody(v *ListSmartSysAvatarModelsResponseBody) *ListSmartSysAvatarModelsResponse {
	s.Body = v
	return s
}

func (s *ListSmartSysAvatarModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
