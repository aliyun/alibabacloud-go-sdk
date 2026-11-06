// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCorpNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCorpNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCorpNumberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCorpNumberResponseBody) *DeleteCorpNumberResponse
	GetBody() *DeleteCorpNumberResponseBody
}

type DeleteCorpNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCorpNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCorpNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCorpNumberResponse) GoString() string {
	return s.String()
}

func (s *DeleteCorpNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCorpNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCorpNumberResponse) GetBody() *DeleteCorpNumberResponseBody {
	return s.Body
}

func (s *DeleteCorpNumberResponse) SetHeaders(v map[string]*string) *DeleteCorpNumberResponse {
	s.Headers = v
	return s
}

func (s *DeleteCorpNumberResponse) SetStatusCode(v int32) *DeleteCorpNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCorpNumberResponse) SetBody(v *DeleteCorpNumberResponseBody) *DeleteCorpNumberResponse {
	s.Body = v
	return s
}

func (s *DeleteCorpNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
