// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCloneMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterCloneMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterCloneMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterCloneMetaResponseBody) *GetClusterCloneMetaResponse
	GetBody() *GetClusterCloneMetaResponseBody
}

type GetClusterCloneMetaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterCloneMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterCloneMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCloneMetaResponse) GoString() string {
	return s.String()
}

func (s *GetClusterCloneMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterCloneMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterCloneMetaResponse) GetBody() *GetClusterCloneMetaResponseBody {
	return s.Body
}

func (s *GetClusterCloneMetaResponse) SetHeaders(v map[string]*string) *GetClusterCloneMetaResponse {
	s.Headers = v
	return s
}

func (s *GetClusterCloneMetaResponse) SetStatusCode(v int32) *GetClusterCloneMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterCloneMetaResponse) SetBody(v *GetClusterCloneMetaResponseBody) *GetClusterCloneMetaResponse {
	s.Body = v
	return s
}

func (s *GetClusterCloneMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
