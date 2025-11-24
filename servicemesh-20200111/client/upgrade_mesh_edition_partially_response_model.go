// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMeshEditionPartiallyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeMeshEditionPartiallyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeMeshEditionPartiallyResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeMeshEditionPartiallyResponseBody) *UpgradeMeshEditionPartiallyResponse
	GetBody() *UpgradeMeshEditionPartiallyResponseBody
}

type UpgradeMeshEditionPartiallyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMeshEditionPartiallyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMeshEditionPartiallyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMeshEditionPartiallyResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMeshEditionPartiallyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeMeshEditionPartiallyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeMeshEditionPartiallyResponse) GetBody() *UpgradeMeshEditionPartiallyResponseBody {
	return s.Body
}

func (s *UpgradeMeshEditionPartiallyResponse) SetHeaders(v map[string]*string) *UpgradeMeshEditionPartiallyResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMeshEditionPartiallyResponse) SetStatusCode(v int32) *UpgradeMeshEditionPartiallyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyResponse) SetBody(v *UpgradeMeshEditionPartiallyResponseBody) *UpgradeMeshEditionPartiallyResponse {
	s.Body = v
	return s
}

func (s *UpgradeMeshEditionPartiallyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
