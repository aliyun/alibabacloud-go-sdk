// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllCommonContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadAllCommonContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadAllCommonContactsResponse
	GetStatusCode() *int32
	SetBody(v *ReadAllCommonContactsResponseBody) *ReadAllCommonContactsResponse
	GetBody() *ReadAllCommonContactsResponseBody
}

type ReadAllCommonContactsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadAllCommonContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadAllCommonContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadAllCommonContactsResponse) GoString() string {
	return s.String()
}

func (s *ReadAllCommonContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadAllCommonContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadAllCommonContactsResponse) GetBody() *ReadAllCommonContactsResponseBody {
	return s.Body
}

func (s *ReadAllCommonContactsResponse) SetHeaders(v map[string]*string) *ReadAllCommonContactsResponse {
	s.Headers = v
	return s
}

func (s *ReadAllCommonContactsResponse) SetStatusCode(v int32) *ReadAllCommonContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadAllCommonContactsResponse) SetBody(v *ReadAllCommonContactsResponseBody) *ReadAllCommonContactsResponse {
	s.Body = v
	return s
}

func (s *ReadAllCommonContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
