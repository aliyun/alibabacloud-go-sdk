// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateDetectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateDetectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateDetectConfigResponse
	GetStatusCode() *int32
	SetBody(v *AssociateDetectConfigResponseBody) *AssociateDetectConfigResponse
	GetBody() *AssociateDetectConfigResponseBody
}

type AssociateDetectConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateDetectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateDetectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateDetectConfigResponse) GoString() string {
	return s.String()
}

func (s *AssociateDetectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateDetectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateDetectConfigResponse) GetBody() *AssociateDetectConfigResponseBody {
	return s.Body
}

func (s *AssociateDetectConfigResponse) SetHeaders(v map[string]*string) *AssociateDetectConfigResponse {
	s.Headers = v
	return s
}

func (s *AssociateDetectConfigResponse) SetStatusCode(v int32) *AssociateDetectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateDetectConfigResponse) SetBody(v *AssociateDetectConfigResponseBody) *AssociateDetectConfigResponse {
	s.Body = v
	return s
}

func (s *AssociateDetectConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
