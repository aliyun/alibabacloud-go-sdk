// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserKubeConfigStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserKubeConfigStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserKubeConfigStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserKubeConfigStatesResponseBody) *ListUserKubeConfigStatesResponse
	GetBody() *ListUserKubeConfigStatesResponseBody
}

type ListUserKubeConfigStatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserKubeConfigStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserKubeConfigStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserKubeConfigStatesResponse) GoString() string {
	return s.String()
}

func (s *ListUserKubeConfigStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserKubeConfigStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserKubeConfigStatesResponse) GetBody() *ListUserKubeConfigStatesResponseBody {
	return s.Body
}

func (s *ListUserKubeConfigStatesResponse) SetHeaders(v map[string]*string) *ListUserKubeConfigStatesResponse {
	s.Headers = v
	return s
}

func (s *ListUserKubeConfigStatesResponse) SetStatusCode(v int32) *ListUserKubeConfigStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserKubeConfigStatesResponse) SetBody(v *ListUserKubeConfigStatesResponseBody) *ListUserKubeConfigStatesResponse {
	s.Body = v
	return s
}

func (s *ListUserKubeConfigStatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
