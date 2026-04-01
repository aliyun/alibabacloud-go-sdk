// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePluginResponse
	GetStatusCode() *int32
	SetBody(v *RemovePluginResponseBody) *RemovePluginResponse
	GetBody() *RemovePluginResponseBody
}

type RemovePluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePluginResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePluginResponse) GoString() string {
	return s.String()
}

func (s *RemovePluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePluginResponse) GetBody() *RemovePluginResponseBody {
	return s.Body
}

func (s *RemovePluginResponse) SetHeaders(v map[string]*string) *RemovePluginResponse {
	s.Headers = v
	return s
}

func (s *RemovePluginResponse) SetStatusCode(v int32) *RemovePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePluginResponse) SetBody(v *RemovePluginResponseBody) *RemovePluginResponse {
	s.Body = v
	return s
}

func (s *RemovePluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
