// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *InsertK8sApplicationResponseBody) *InsertK8sApplicationResponse
	GetBody() *InsertK8sApplicationResponseBody
}

type InsertK8sApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *InsertK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertK8sApplicationResponse) GetBody() *InsertK8sApplicationResponseBody {
	return s.Body
}

func (s *InsertK8sApplicationResponse) SetHeaders(v map[string]*string) *InsertK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *InsertK8sApplicationResponse) SetStatusCode(v int32) *InsertK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertK8sApplicationResponse) SetBody(v *InsertK8sApplicationResponseBody) *InsertK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *InsertK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
