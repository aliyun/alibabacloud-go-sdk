// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCustomImageShareAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveCustomImageShareAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveCustomImageShareAccountResponse
	GetStatusCode() *int32
	SetBody(v *RemoveCustomImageShareAccountResponseBody) *RemoveCustomImageShareAccountResponse
	GetBody() *RemoveCustomImageShareAccountResponseBody
}

type RemoveCustomImageShareAccountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCustomImageShareAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCustomImageShareAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveCustomImageShareAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveCustomImageShareAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveCustomImageShareAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveCustomImageShareAccountResponse) GetBody() *RemoveCustomImageShareAccountResponseBody {
	return s.Body
}

func (s *RemoveCustomImageShareAccountResponse) SetHeaders(v map[string]*string) *RemoveCustomImageShareAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveCustomImageShareAccountResponse) SetStatusCode(v int32) *RemoveCustomImageShareAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCustomImageShareAccountResponse) SetBody(v *RemoveCustomImageShareAccountResponseBody) *RemoveCustomImageShareAccountResponse {
	s.Body = v
	return s
}

func (s *RemoveCustomImageShareAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
