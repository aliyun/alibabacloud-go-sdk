// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BlockObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BlockObjectResponse
	GetStatusCode() *int32
	SetBody(v *BlockObjectResponseBody) *BlockObjectResponse
	GetBody() *BlockObjectResponseBody
}

type BlockObjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BlockObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BlockObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s BlockObjectResponse) GoString() string {
	return s.String()
}

func (s *BlockObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BlockObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BlockObjectResponse) GetBody() *BlockObjectResponseBody {
	return s.Body
}

func (s *BlockObjectResponse) SetHeaders(v map[string]*string) *BlockObjectResponse {
	s.Headers = v
	return s
}

func (s *BlockObjectResponse) SetStatusCode(v int32) *BlockObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *BlockObjectResponse) SetBody(v *BlockObjectResponseBody) *BlockObjectResponse {
	s.Body = v
	return s
}

func (s *BlockObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
