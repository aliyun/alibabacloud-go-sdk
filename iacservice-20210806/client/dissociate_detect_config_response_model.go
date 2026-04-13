// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateDetectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateDetectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateDetectConfigResponse
	GetStatusCode() *int32
	SetBody(v *DissociateDetectConfigResponseBody) *DissociateDetectConfigResponse
	GetBody() *DissociateDetectConfigResponseBody
}

type DissociateDetectConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateDetectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateDetectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateDetectConfigResponse) GoString() string {
	return s.String()
}

func (s *DissociateDetectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateDetectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateDetectConfigResponse) GetBody() *DissociateDetectConfigResponseBody {
	return s.Body
}

func (s *DissociateDetectConfigResponse) SetHeaders(v map[string]*string) *DissociateDetectConfigResponse {
	s.Headers = v
	return s
}

func (s *DissociateDetectConfigResponse) SetStatusCode(v int32) *DissociateDetectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateDetectConfigResponse) SetBody(v *DissociateDetectConfigResponseBody) *DissociateDetectConfigResponse {
	s.Body = v
	return s
}

func (s *DissociateDetectConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
