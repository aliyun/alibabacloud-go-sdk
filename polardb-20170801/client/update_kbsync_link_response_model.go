// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKBSyncLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKBSyncLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKBSyncLinkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKBSyncLinkResponseBody) *UpdateKBSyncLinkResponse
	GetBody() *UpdateKBSyncLinkResponseBody
}

type UpdateKBSyncLinkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKBSyncLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKBSyncLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKBSyncLinkResponse) GoString() string {
	return s.String()
}

func (s *UpdateKBSyncLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKBSyncLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKBSyncLinkResponse) GetBody() *UpdateKBSyncLinkResponseBody {
	return s.Body
}

func (s *UpdateKBSyncLinkResponse) SetHeaders(v map[string]*string) *UpdateKBSyncLinkResponse {
	s.Headers = v
	return s
}

func (s *UpdateKBSyncLinkResponse) SetStatusCode(v int32) *UpdateKBSyncLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKBSyncLinkResponse) SetBody(v *UpdateKBSyncLinkResponseBody) *UpdateKBSyncLinkResponse {
	s.Body = v
	return s
}

func (s *UpdateKBSyncLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
