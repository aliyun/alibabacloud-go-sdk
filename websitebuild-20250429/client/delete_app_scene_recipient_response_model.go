// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSceneRecipientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppSceneRecipientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppSceneRecipientResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppSceneRecipientResponseBody) *DeleteAppSceneRecipientResponse
	GetBody() *DeleteAppSceneRecipientResponseBody
}

type DeleteAppSceneRecipientResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppSceneRecipientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppSceneRecipientResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSceneRecipientResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppSceneRecipientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppSceneRecipientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppSceneRecipientResponse) GetBody() *DeleteAppSceneRecipientResponseBody {
	return s.Body
}

func (s *DeleteAppSceneRecipientResponse) SetHeaders(v map[string]*string) *DeleteAppSceneRecipientResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppSceneRecipientResponse) SetStatusCode(v int32) *DeleteAppSceneRecipientResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppSceneRecipientResponse) SetBody(v *DeleteAppSceneRecipientResponseBody) *DeleteAppSceneRecipientResponse {
	s.Body = v
	return s
}

func (s *DeleteAppSceneRecipientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
