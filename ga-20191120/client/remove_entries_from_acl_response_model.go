// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntriesFromAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveEntriesFromAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveEntriesFromAclResponse
	GetStatusCode() *int32
	SetBody(v *RemoveEntriesFromAclResponseBody) *RemoveEntriesFromAclResponse
	GetBody() *RemoveEntriesFromAclResponseBody
}

type RemoveEntriesFromAclResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveEntriesFromAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveEntriesFromAclResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntriesFromAclResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveEntriesFromAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveEntriesFromAclResponse) GetBody() *RemoveEntriesFromAclResponseBody {
	return s.Body
}

func (s *RemoveEntriesFromAclResponse) SetHeaders(v map[string]*string) *RemoveEntriesFromAclResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntriesFromAclResponse) SetStatusCode(v int32) *RemoveEntriesFromAclResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEntriesFromAclResponse) SetBody(v *RemoveEntriesFromAclResponseBody) *RemoveEntriesFromAclResponse {
	s.Body = v
	return s
}

func (s *RemoveEntriesFromAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
