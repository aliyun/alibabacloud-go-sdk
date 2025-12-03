// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSignatureApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSignatureApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSignatureApisResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSignatureApisResponseBody) *RemoveSignatureApisResponse
	GetBody() *RemoveSignatureApisResponseBody
}

type RemoveSignatureApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSignatureApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSignatureApisResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSignatureApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveSignatureApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSignatureApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSignatureApisResponse) GetBody() *RemoveSignatureApisResponseBody {
	return s.Body
}

func (s *RemoveSignatureApisResponse) SetHeaders(v map[string]*string) *RemoveSignatureApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveSignatureApisResponse) SetStatusCode(v int32) *RemoveSignatureApisResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSignatureApisResponse) SetBody(v *RemoveSignatureApisResponseBody) *RemoveSignatureApisResponse {
	s.Body = v
	return s
}

func (s *RemoveSignatureApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
