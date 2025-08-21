// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSynonymsDictsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSynonymsDictsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSynonymsDictsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSynonymsDictsResponseBody) *UpdateSynonymsDictsResponse
	GetBody() *UpdateSynonymsDictsResponseBody
}

type UpdateSynonymsDictsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSynonymsDictsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSynonymsDictsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSynonymsDictsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsDictsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSynonymsDictsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSynonymsDictsResponse) GetBody() *UpdateSynonymsDictsResponseBody {
	return s.Body
}

func (s *UpdateSynonymsDictsResponse) SetHeaders(v map[string]*string) *UpdateSynonymsDictsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSynonymsDictsResponse) SetStatusCode(v int32) *UpdateSynonymsDictsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSynonymsDictsResponse) SetBody(v *UpdateSynonymsDictsResponseBody) *UpdateSynonymsDictsResponse {
	s.Body = v
	return s
}

func (s *UpdateSynonymsDictsResponse) Validate() error {
	return dara.Validate(s)
}
