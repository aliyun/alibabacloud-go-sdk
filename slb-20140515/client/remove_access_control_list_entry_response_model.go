// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAccessControlListEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAccessControlListEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAccessControlListEntryResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAccessControlListEntryResponseBody) *RemoveAccessControlListEntryResponse
	GetBody() *RemoveAccessControlListEntryResponseBody
}

type RemoveAccessControlListEntryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAccessControlListEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAccessControlListEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAccessControlListEntryResponse) GoString() string {
	return s.String()
}

func (s *RemoveAccessControlListEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAccessControlListEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAccessControlListEntryResponse) GetBody() *RemoveAccessControlListEntryResponseBody {
	return s.Body
}

func (s *RemoveAccessControlListEntryResponse) SetHeaders(v map[string]*string) *RemoveAccessControlListEntryResponse {
	s.Headers = v
	return s
}

func (s *RemoveAccessControlListEntryResponse) SetStatusCode(v int32) *RemoveAccessControlListEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAccessControlListEntryResponse) SetBody(v *RemoveAccessControlListEntryResponseBody) *RemoveAccessControlListEntryResponse {
	s.Body = v
	return s
}

func (s *RemoveAccessControlListEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
