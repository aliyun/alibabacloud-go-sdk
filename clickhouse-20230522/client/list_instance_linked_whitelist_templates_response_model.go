// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceLinkedWhitelistTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceLinkedWhitelistTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceLinkedWhitelistTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceLinkedWhitelistTemplatesResponseBody) *ListInstanceLinkedWhitelistTemplatesResponse
	GetBody() *ListInstanceLinkedWhitelistTemplatesResponseBody
}

type ListInstanceLinkedWhitelistTemplatesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceLinkedWhitelistTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceLinkedWhitelistTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLinkedWhitelistTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceLinkedWhitelistTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceLinkedWhitelistTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceLinkedWhitelistTemplatesResponse) GetBody() *ListInstanceLinkedWhitelistTemplatesResponseBody {
	return s.Body
}

func (s *ListInstanceLinkedWhitelistTemplatesResponse) SetHeaders(v map[string]*string) *ListInstanceLinkedWhitelistTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponse) SetStatusCode(v int32) *ListInstanceLinkedWhitelistTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponse) SetBody(v *ListInstanceLinkedWhitelistTemplatesResponseBody) *ListInstanceLinkedWhitelistTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
