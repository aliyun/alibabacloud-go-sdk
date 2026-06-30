// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelOperatorApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelOperatorApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelOperatorApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListModelOperatorApiKeysResponseBody) *ListModelOperatorApiKeysResponse
	GetBody() *ListModelOperatorApiKeysResponseBody
}

type ListModelOperatorApiKeysResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelOperatorApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelOperatorApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorApiKeysResponse) GoString() string {
	return s.String()
}

func (s *ListModelOperatorApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelOperatorApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelOperatorApiKeysResponse) GetBody() *ListModelOperatorApiKeysResponseBody {
	return s.Body
}

func (s *ListModelOperatorApiKeysResponse) SetHeaders(v map[string]*string) *ListModelOperatorApiKeysResponse {
	s.Headers = v
	return s
}

func (s *ListModelOperatorApiKeysResponse) SetStatusCode(v int32) *ListModelOperatorApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelOperatorApiKeysResponse) SetBody(v *ListModelOperatorApiKeysResponseBody) *ListModelOperatorApiKeysResponse {
	s.Body = v
	return s
}

func (s *ListModelOperatorApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
