// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackupsSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveBackupsSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveBackupsSetResponse
	GetStatusCode() *int32
	SetBody(v *RemoveBackupsSetResponseBody) *RemoveBackupsSetResponse
	GetBody() *RemoveBackupsSetResponseBody
}

type RemoveBackupsSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveBackupsSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveBackupsSetResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackupsSetResponse) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveBackupsSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveBackupsSetResponse) GetBody() *RemoveBackupsSetResponseBody {
	return s.Body
}

func (s *RemoveBackupsSetResponse) SetHeaders(v map[string]*string) *RemoveBackupsSetResponse {
	s.Headers = v
	return s
}

func (s *RemoveBackupsSetResponse) SetStatusCode(v int32) *RemoveBackupsSetResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBackupsSetResponse) SetBody(v *RemoveBackupsSetResponseBody) *RemoveBackupsSetResponse {
	s.Body = v
	return s
}

func (s *RemoveBackupsSetResponse) Validate() error {
	return dara.Validate(s)
}
