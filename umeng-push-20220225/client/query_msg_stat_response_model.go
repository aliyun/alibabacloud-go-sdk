// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMsgStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMsgStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMsgStatResponse
	GetStatusCode() *int32
	SetBody(v *QueryMsgStatResponseBody) *QueryMsgStatResponse
	GetBody() *QueryMsgStatResponseBody
}

type QueryMsgStatResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMsgStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMsgStatResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMsgStatResponse) GoString() string {
	return s.String()
}

func (s *QueryMsgStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMsgStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMsgStatResponse) GetBody() *QueryMsgStatResponseBody {
	return s.Body
}

func (s *QueryMsgStatResponse) SetHeaders(v map[string]*string) *QueryMsgStatResponse {
	s.Headers = v
	return s
}

func (s *QueryMsgStatResponse) SetStatusCode(v int32) *QueryMsgStatResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMsgStatResponse) SetBody(v *QueryMsgStatResponseBody) *QueryMsgStatResponse {
	s.Body = v
	return s
}

func (s *QueryMsgStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
