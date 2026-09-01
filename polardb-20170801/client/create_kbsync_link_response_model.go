// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKBSyncLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKBSyncLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKBSyncLinkResponse
	GetStatusCode() *int32
	SetBody(v *CreateKBSyncLinkResponseBody) *CreateKBSyncLinkResponse
	GetBody() *CreateKBSyncLinkResponseBody
}

type CreateKBSyncLinkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKBSyncLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKBSyncLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKBSyncLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateKBSyncLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKBSyncLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKBSyncLinkResponse) GetBody() *CreateKBSyncLinkResponseBody {
	return s.Body
}

func (s *CreateKBSyncLinkResponse) SetHeaders(v map[string]*string) *CreateKBSyncLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateKBSyncLinkResponse) SetStatusCode(v int32) *CreateKBSyncLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKBSyncLinkResponse) SetBody(v *CreateKBSyncLinkResponseBody) *CreateKBSyncLinkResponse {
	s.Body = v
	return s
}

func (s *CreateKBSyncLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
