// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChangeSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChangeSetResponse
	GetStatusCode() *int32
	SetBody(v *GetChangeSetResponseBody) *GetChangeSetResponse
	GetBody() *GetChangeSetResponseBody
}

type GetChangeSetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChangeSetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChangeSetResponse) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChangeSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChangeSetResponse) GetBody() *GetChangeSetResponseBody {
	return s.Body
}

func (s *GetChangeSetResponse) SetHeaders(v map[string]*string) *GetChangeSetResponse {
	s.Headers = v
	return s
}

func (s *GetChangeSetResponse) SetStatusCode(v int32) *GetChangeSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChangeSetResponse) SetBody(v *GetChangeSetResponseBody) *GetChangeSetResponse {
	s.Body = v
	return s
}

func (s *GetChangeSetResponse) Validate() error {
	return dara.Validate(s)
}
