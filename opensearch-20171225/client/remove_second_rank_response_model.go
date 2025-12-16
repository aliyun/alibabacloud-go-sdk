// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSecondRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSecondRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSecondRankResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSecondRankResponseBody) *RemoveSecondRankResponse
	GetBody() *RemoveSecondRankResponseBody
}

type RemoveSecondRankResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSecondRankResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSecondRankResponse) GoString() string {
	return s.String()
}

func (s *RemoveSecondRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSecondRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSecondRankResponse) GetBody() *RemoveSecondRankResponseBody {
	return s.Body
}

func (s *RemoveSecondRankResponse) SetHeaders(v map[string]*string) *RemoveSecondRankResponse {
	s.Headers = v
	return s
}

func (s *RemoveSecondRankResponse) SetStatusCode(v int32) *RemoveSecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSecondRankResponse) SetBody(v *RemoveSecondRankResponseBody) *RemoveSecondRankResponse {
	s.Body = v
	return s
}

func (s *RemoveSecondRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
