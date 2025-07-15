// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneGlobalReplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertInterveneGlobalReplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertInterveneGlobalReplyResponse
	GetStatusCode() *int32
	SetBody(v *InsertInterveneGlobalReplyResponseBody) *InsertInterveneGlobalReplyResponse
	GetBody() *InsertInterveneGlobalReplyResponseBody
}

type InsertInterveneGlobalReplyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertInterveneGlobalReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertInterveneGlobalReplyResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponse) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertInterveneGlobalReplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertInterveneGlobalReplyResponse) GetBody() *InsertInterveneGlobalReplyResponseBody {
	return s.Body
}

func (s *InsertInterveneGlobalReplyResponse) SetHeaders(v map[string]*string) *InsertInterveneGlobalReplyResponse {
	s.Headers = v
	return s
}

func (s *InsertInterveneGlobalReplyResponse) SetStatusCode(v int32) *InsertInterveneGlobalReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponse) SetBody(v *InsertInterveneGlobalReplyResponseBody) *InsertInterveneGlobalReplyResponse {
	s.Body = v
	return s
}

func (s *InsertInterveneGlobalReplyResponse) Validate() error {
	return dara.Validate(s)
}
