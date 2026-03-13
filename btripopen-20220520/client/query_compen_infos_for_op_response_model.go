// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCompenInfosForOpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCompenInfosForOpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCompenInfosForOpResponse
	GetStatusCode() *int32
	SetBody(v *QueryCompenInfosForOpResponseBody) *QueryCompenInfosForOpResponse
	GetBody() *QueryCompenInfosForOpResponseBody
}

type QueryCompenInfosForOpResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCompenInfosForOpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCompenInfosForOpResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCompenInfosForOpResponse) GoString() string {
	return s.String()
}

func (s *QueryCompenInfosForOpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCompenInfosForOpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCompenInfosForOpResponse) GetBody() *QueryCompenInfosForOpResponseBody {
	return s.Body
}

func (s *QueryCompenInfosForOpResponse) SetHeaders(v map[string]*string) *QueryCompenInfosForOpResponse {
	s.Headers = v
	return s
}

func (s *QueryCompenInfosForOpResponse) SetStatusCode(v int32) *QueryCompenInfosForOpResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCompenInfosForOpResponse) SetBody(v *QueryCompenInfosForOpResponseBody) *QueryCompenInfosForOpResponse {
	s.Body = v
	return s
}

func (s *QueryCompenInfosForOpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
