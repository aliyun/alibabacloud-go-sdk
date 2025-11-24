// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMeshVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeMeshVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeMeshVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeMeshVersionResponseBody) *UpgradeMeshVersionResponse
	GetBody() *UpgradeMeshVersionResponseBody
}

type UpgradeMeshVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMeshVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMeshVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMeshVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMeshVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeMeshVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeMeshVersionResponse) GetBody() *UpgradeMeshVersionResponseBody {
	return s.Body
}

func (s *UpgradeMeshVersionResponse) SetHeaders(v map[string]*string) *UpgradeMeshVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMeshVersionResponse) SetStatusCode(v int32) *UpgradeMeshVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMeshVersionResponse) SetBody(v *UpgradeMeshVersionResponseBody) *UpgradeMeshVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeMeshVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
