// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndividuationProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIndividuationProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIndividuationProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIndividuationProjectResponseBody) *DeleteIndividuationProjectResponse
	GetBody() *DeleteIndividuationProjectResponseBody
}

type DeleteIndividuationProjectResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndividuationProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndividuationProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndividuationProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIndividuationProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIndividuationProjectResponse) GetBody() *DeleteIndividuationProjectResponseBody {
	return s.Body
}

func (s *DeleteIndividuationProjectResponse) SetHeaders(v map[string]*string) *DeleteIndividuationProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndividuationProjectResponse) SetStatusCode(v int32) *DeleteIndividuationProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndividuationProjectResponse) SetBody(v *DeleteIndividuationProjectResponseBody) *DeleteIndividuationProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteIndividuationProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
