// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKBSyncLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKBSyncLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKBSyncLinkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKBSyncLinkResponseBody) *DeleteKBSyncLinkResponse
	GetBody() *DeleteKBSyncLinkResponseBody
}

type DeleteKBSyncLinkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKBSyncLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKBSyncLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKBSyncLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteKBSyncLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKBSyncLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKBSyncLinkResponse) GetBody() *DeleteKBSyncLinkResponseBody {
	return s.Body
}

func (s *DeleteKBSyncLinkResponse) SetHeaders(v map[string]*string) *DeleteKBSyncLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteKBSyncLinkResponse) SetStatusCode(v int32) *DeleteKBSyncLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKBSyncLinkResponse) SetBody(v *DeleteKBSyncLinkResponseBody) *DeleteKBSyncLinkResponse {
	s.Body = v
	return s
}

func (s *DeleteKBSyncLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
