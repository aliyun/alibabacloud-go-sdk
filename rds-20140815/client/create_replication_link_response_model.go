// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReplicationLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReplicationLinkResponse
	GetStatusCode() *int32
	SetBody(v *CreateReplicationLinkResponseBody) *CreateReplicationLinkResponse
	GetBody() *CreateReplicationLinkResponseBody
}

type CreateReplicationLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReplicationLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReplicationLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateReplicationLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReplicationLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReplicationLinkResponse) GetBody() *CreateReplicationLinkResponseBody {
	return s.Body
}

func (s *CreateReplicationLinkResponse) SetHeaders(v map[string]*string) *CreateReplicationLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateReplicationLinkResponse) SetStatusCode(v int32) *CreateReplicationLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReplicationLinkResponse) SetBody(v *CreateReplicationLinkResponseBody) *CreateReplicationLinkResponse {
	s.Body = v
	return s
}

func (s *CreateReplicationLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
