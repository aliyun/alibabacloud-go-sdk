// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRecallManagementTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishRecallManagementTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishRecallManagementTableResponse
	GetStatusCode() *int32
	SetBody(v *PublishRecallManagementTableResponseBody) *PublishRecallManagementTableResponse
	GetBody() *PublishRecallManagementTableResponseBody
}

type PublishRecallManagementTableResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishRecallManagementTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishRecallManagementTableResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishRecallManagementTableResponse) GoString() string {
	return s.String()
}

func (s *PublishRecallManagementTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishRecallManagementTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishRecallManagementTableResponse) GetBody() *PublishRecallManagementTableResponseBody {
	return s.Body
}

func (s *PublishRecallManagementTableResponse) SetHeaders(v map[string]*string) *PublishRecallManagementTableResponse {
	s.Headers = v
	return s
}

func (s *PublishRecallManagementTableResponse) SetStatusCode(v int32) *PublishRecallManagementTableResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishRecallManagementTableResponse) SetBody(v *PublishRecallManagementTableResponseBody) *PublishRecallManagementTableResponse {
	s.Body = v
	return s
}

func (s *PublishRecallManagementTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
