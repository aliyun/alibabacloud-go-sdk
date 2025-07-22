// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudNotePhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudNotePhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudNotePhrasesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudNotePhrasesResponseBody) *ModifyCloudNotePhrasesResponse
	GetBody() *ModifyCloudNotePhrasesResponseBody
}

type ModifyCloudNotePhrasesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudNotePhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudNotePhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudNotePhrasesResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudNotePhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudNotePhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudNotePhrasesResponse) GetBody() *ModifyCloudNotePhrasesResponseBody {
	return s.Body
}

func (s *ModifyCloudNotePhrasesResponse) SetHeaders(v map[string]*string) *ModifyCloudNotePhrasesResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudNotePhrasesResponse) SetStatusCode(v int32) *ModifyCloudNotePhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudNotePhrasesResponse) SetBody(v *ModifyCloudNotePhrasesResponseBody) *ModifyCloudNotePhrasesResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudNotePhrasesResponse) Validate() error {
	return dara.Validate(s)
}
