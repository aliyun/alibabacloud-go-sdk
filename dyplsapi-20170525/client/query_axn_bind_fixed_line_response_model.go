// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxnBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAxnBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAxnBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *QueryAxnBindFixedLineResponseBody) *QueryAxnBindFixedLineResponse
	GetBody() *QueryAxnBindFixedLineResponseBody
}

type QueryAxnBindFixedLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAxnBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAxnBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAxnBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAxnBindFixedLineResponse) GetBody() *QueryAxnBindFixedLineResponseBody {
	return s.Body
}

func (s *QueryAxnBindFixedLineResponse) SetHeaders(v map[string]*string) *QueryAxnBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *QueryAxnBindFixedLineResponse) SetStatusCode(v int32) *QueryAxnBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAxnBindFixedLineResponse) SetBody(v *QueryAxnBindFixedLineResponseBody) *QueryAxnBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *QueryAxnBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
