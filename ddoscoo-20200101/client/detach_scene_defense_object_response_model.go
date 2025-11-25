// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSceneDefenseObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachSceneDefenseObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachSceneDefenseObjectResponse
	GetStatusCode() *int32
	SetBody(v *DetachSceneDefenseObjectResponseBody) *DetachSceneDefenseObjectResponse
	GetBody() *DetachSceneDefenseObjectResponseBody
}

type DetachSceneDefenseObjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachSceneDefenseObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachSceneDefenseObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachSceneDefenseObjectResponse) GoString() string {
	return s.String()
}

func (s *DetachSceneDefenseObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachSceneDefenseObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachSceneDefenseObjectResponse) GetBody() *DetachSceneDefenseObjectResponseBody {
	return s.Body
}

func (s *DetachSceneDefenseObjectResponse) SetHeaders(v map[string]*string) *DetachSceneDefenseObjectResponse {
	s.Headers = v
	return s
}

func (s *DetachSceneDefenseObjectResponse) SetStatusCode(v int32) *DetachSceneDefenseObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachSceneDefenseObjectResponse) SetBody(v *DetachSceneDefenseObjectResponseBody) *DetachSceneDefenseObjectResponse {
	s.Body = v
	return s
}

func (s *DetachSceneDefenseObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
