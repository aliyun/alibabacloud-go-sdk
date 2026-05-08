// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndividuationTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIndividuationTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIndividuationTextResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIndividuationTextResponseBody) *DeleteIndividuationTextResponse
	GetBody() *DeleteIndividuationTextResponseBody
}

type DeleteIndividuationTextResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndividuationTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndividuationTextResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndividuationTextResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIndividuationTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIndividuationTextResponse) GetBody() *DeleteIndividuationTextResponseBody {
	return s.Body
}

func (s *DeleteIndividuationTextResponse) SetHeaders(v map[string]*string) *DeleteIndividuationTextResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndividuationTextResponse) SetStatusCode(v int32) *DeleteIndividuationTextResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndividuationTextResponse) SetBody(v *DeleteIndividuationTextResponseBody) *DeleteIndividuationTextResponse {
	s.Body = v
	return s
}

func (s *DeleteIndividuationTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
