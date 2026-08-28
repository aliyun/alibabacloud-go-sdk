// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyncMCPServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSyncMCPServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSyncMCPServerResponse
	GetStatusCode() *int32
	SetBody(v *ListSyncMCPServerResponseBody) *ListSyncMCPServerResponse
	GetBody() *ListSyncMCPServerResponseBody
}

type ListSyncMCPServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSyncMCPServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSyncMCPServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSyncMCPServerResponse) GoString() string {
	return s.String()
}

func (s *ListSyncMCPServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSyncMCPServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSyncMCPServerResponse) GetBody() *ListSyncMCPServerResponseBody {
	return s.Body
}

func (s *ListSyncMCPServerResponse) SetHeaders(v map[string]*string) *ListSyncMCPServerResponse {
	s.Headers = v
	return s
}

func (s *ListSyncMCPServerResponse) SetStatusCode(v int32) *ListSyncMCPServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSyncMCPServerResponse) SetBody(v *ListSyncMCPServerResponseBody) *ListSyncMCPServerResponse {
	s.Body = v
	return s
}

func (s *ListSyncMCPServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
