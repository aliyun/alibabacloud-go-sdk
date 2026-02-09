// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClipsBuildInResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClipsBuildInResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClipsBuildInResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetClipsBuildInResourceResponseBody) *GetClipsBuildInResourceResponse
	GetBody() *GetClipsBuildInResourceResponseBody
}

type GetClipsBuildInResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClipsBuildInResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClipsBuildInResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClipsBuildInResourceResponse) GoString() string {
	return s.String()
}

func (s *GetClipsBuildInResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClipsBuildInResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClipsBuildInResourceResponse) GetBody() *GetClipsBuildInResourceResponseBody {
	return s.Body
}

func (s *GetClipsBuildInResourceResponse) SetHeaders(v map[string]*string) *GetClipsBuildInResourceResponse {
	s.Headers = v
	return s
}

func (s *GetClipsBuildInResourceResponse) SetStatusCode(v int32) *GetClipsBuildInResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClipsBuildInResourceResponse) SetBody(v *GetClipsBuildInResourceResponseBody) *GetClipsBuildInResourceResponse {
	s.Body = v
	return s
}

func (s *GetClipsBuildInResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
