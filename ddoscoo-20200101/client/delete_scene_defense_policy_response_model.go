// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneDefensePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSceneDefensePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSceneDefensePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSceneDefensePolicyResponseBody) *DeleteSceneDefensePolicyResponse
	GetBody() *DeleteSceneDefensePolicyResponseBody
}

type DeleteSceneDefensePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSceneDefensePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSceneDefensePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSceneDefensePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSceneDefensePolicyResponse) GetBody() *DeleteSceneDefensePolicyResponseBody {
	return s.Body
}

func (s *DeleteSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *DeleteSceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSceneDefensePolicyResponse) SetStatusCode(v int32) *DeleteSceneDefensePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSceneDefensePolicyResponse) SetBody(v *DeleteSceneDefensePolicyResponseBody) *DeleteSceneDefensePolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteSceneDefensePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
