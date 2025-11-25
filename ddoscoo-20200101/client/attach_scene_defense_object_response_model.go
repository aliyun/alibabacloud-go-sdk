// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSceneDefenseObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachSceneDefenseObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachSceneDefenseObjectResponse
	GetStatusCode() *int32
	SetBody(v *AttachSceneDefenseObjectResponseBody) *AttachSceneDefenseObjectResponse
	GetBody() *AttachSceneDefenseObjectResponseBody
}

type AttachSceneDefenseObjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachSceneDefenseObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachSceneDefenseObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachSceneDefenseObjectResponse) GoString() string {
	return s.String()
}

func (s *AttachSceneDefenseObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachSceneDefenseObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachSceneDefenseObjectResponse) GetBody() *AttachSceneDefenseObjectResponseBody {
	return s.Body
}

func (s *AttachSceneDefenseObjectResponse) SetHeaders(v map[string]*string) *AttachSceneDefenseObjectResponse {
	s.Headers = v
	return s
}

func (s *AttachSceneDefenseObjectResponse) SetStatusCode(v int32) *AttachSceneDefenseObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachSceneDefenseObjectResponse) SetBody(v *AttachSceneDefenseObjectResponseBody) *AttachSceneDefenseObjectResponse {
	s.Body = v
	return s
}

func (s *AttachSceneDefenseObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
