// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFirstRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveFirstRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveFirstRankResponse
	GetStatusCode() *int32
	SetBody(v *RemoveFirstRankResponseBody) *RemoveFirstRankResponse
	GetBody() *RemoveFirstRankResponseBody
}

type RemoveFirstRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFirstRankResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveFirstRankResponse) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveFirstRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveFirstRankResponse) GetBody() *RemoveFirstRankResponseBody {
	return s.Body
}

func (s *RemoveFirstRankResponse) SetHeaders(v map[string]*string) *RemoveFirstRankResponse {
	s.Headers = v
	return s
}

func (s *RemoveFirstRankResponse) SetStatusCode(v int32) *RemoveFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFirstRankResponse) SetBody(v *RemoveFirstRankResponseBody) *RemoveFirstRankResponse {
	s.Body = v
	return s
}

func (s *RemoveFirstRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
