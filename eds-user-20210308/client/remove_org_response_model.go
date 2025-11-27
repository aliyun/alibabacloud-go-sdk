// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveOrgResponse
	GetStatusCode() *int32
	SetBody(v *RemoveOrgResponseBody) *RemoveOrgResponse
	GetBody() *RemoveOrgResponseBody
}

type RemoveOrgResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveOrgResponse) GoString() string {
	return s.String()
}

func (s *RemoveOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveOrgResponse) GetBody() *RemoveOrgResponseBody {
	return s.Body
}

func (s *RemoveOrgResponse) SetHeaders(v map[string]*string) *RemoveOrgResponse {
	s.Headers = v
	return s
}

func (s *RemoveOrgResponse) SetStatusCode(v int32) *RemoveOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveOrgResponse) SetBody(v *RemoveOrgResponseBody) *RemoveOrgResponse {
	s.Body = v
	return s
}

func (s *RemoveOrgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
