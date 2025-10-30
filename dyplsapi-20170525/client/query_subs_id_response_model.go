// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySubsIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySubsIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySubsIdResponse
	GetStatusCode() *int32
	SetBody(v *QuerySubsIdResponseBody) *QuerySubsIdResponse
	GetBody() *QuerySubsIdResponseBody
}

type QuerySubsIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySubsIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySubsIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySubsIdResponse) GoString() string {
	return s.String()
}

func (s *QuerySubsIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySubsIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySubsIdResponse) GetBody() *QuerySubsIdResponseBody {
	return s.Body
}

func (s *QuerySubsIdResponse) SetHeaders(v map[string]*string) *QuerySubsIdResponse {
	s.Headers = v
	return s
}

func (s *QuerySubsIdResponse) SetStatusCode(v int32) *QuerySubsIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySubsIdResponse) SetBody(v *QuerySubsIdResponseBody) *QuerySubsIdResponse {
	s.Body = v
	return s
}

func (s *QuerySubsIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
