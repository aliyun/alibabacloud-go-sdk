// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeSuperResolutionImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MakeSuperResolutionImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MakeSuperResolutionImageResponse
	GetStatusCode() *int32
	SetBody(v *MakeSuperResolutionImageResponseBody) *MakeSuperResolutionImageResponse
	GetBody() *MakeSuperResolutionImageResponseBody
}

type MakeSuperResolutionImageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MakeSuperResolutionImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MakeSuperResolutionImageResponse) String() string {
	return dara.Prettify(s)
}

func (s MakeSuperResolutionImageResponse) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MakeSuperResolutionImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MakeSuperResolutionImageResponse) GetBody() *MakeSuperResolutionImageResponseBody {
	return s.Body
}

func (s *MakeSuperResolutionImageResponse) SetHeaders(v map[string]*string) *MakeSuperResolutionImageResponse {
	s.Headers = v
	return s
}

func (s *MakeSuperResolutionImageResponse) SetStatusCode(v int32) *MakeSuperResolutionImageResponse {
	s.StatusCode = &v
	return s
}

func (s *MakeSuperResolutionImageResponse) SetBody(v *MakeSuperResolutionImageResponseBody) *MakeSuperResolutionImageResponse {
	s.Body = v
	return s
}

func (s *MakeSuperResolutionImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
