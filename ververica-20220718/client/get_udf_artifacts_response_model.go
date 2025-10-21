// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfArtifactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUdfArtifactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUdfArtifactsResponse
	GetStatusCode() *int32
	SetBody(v *GetUdfArtifactsResponseBody) *GetUdfArtifactsResponse
	GetBody() *GetUdfArtifactsResponseBody
}

type GetUdfArtifactsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUdfArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUdfArtifactsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUdfArtifactsResponse) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUdfArtifactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUdfArtifactsResponse) GetBody() *GetUdfArtifactsResponseBody {
	return s.Body
}

func (s *GetUdfArtifactsResponse) SetHeaders(v map[string]*string) *GetUdfArtifactsResponse {
	s.Headers = v
	return s
}

func (s *GetUdfArtifactsResponse) SetStatusCode(v int32) *GetUdfArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUdfArtifactsResponse) SetBody(v *GetUdfArtifactsResponseBody) *GetUdfArtifactsResponse {
	s.Body = v
	return s
}

func (s *GetUdfArtifactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
