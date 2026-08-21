// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePrivateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeletePrivateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeletePrivateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeletePrivateAccessApplicationResponseBody) *BatchDeletePrivateAccessApplicationResponse
	GetBody() *BatchDeletePrivateAccessApplicationResponseBody
}

type BatchDeletePrivateAccessApplicationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeletePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeletePrivateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *BatchDeletePrivateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeletePrivateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeletePrivateAccessApplicationResponse) GetBody() *BatchDeletePrivateAccessApplicationResponseBody {
	return s.Body
}

func (s *BatchDeletePrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *BatchDeletePrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *BatchDeletePrivateAccessApplicationResponse) SetStatusCode(v int32) *BatchDeletePrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeletePrivateAccessApplicationResponse) SetBody(v *BatchDeletePrivateAccessApplicationResponseBody) *BatchDeletePrivateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *BatchDeletePrivateAccessApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
