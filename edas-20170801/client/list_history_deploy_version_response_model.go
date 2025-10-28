// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoryDeployVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHistoryDeployVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHistoryDeployVersionResponse
	GetStatusCode() *int32
	SetBody(v *ListHistoryDeployVersionResponseBody) *ListHistoryDeployVersionResponse
	GetBody() *ListHistoryDeployVersionResponseBody
}

type ListHistoryDeployVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHistoryDeployVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHistoryDeployVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHistoryDeployVersionResponse) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHistoryDeployVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHistoryDeployVersionResponse) GetBody() *ListHistoryDeployVersionResponseBody {
	return s.Body
}

func (s *ListHistoryDeployVersionResponse) SetHeaders(v map[string]*string) *ListHistoryDeployVersionResponse {
	s.Headers = v
	return s
}

func (s *ListHistoryDeployVersionResponse) SetStatusCode(v int32) *ListHistoryDeployVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHistoryDeployVersionResponse) SetBody(v *ListHistoryDeployVersionResponseBody) *ListHistoryDeployVersionResponse {
	s.Body = v
	return s
}

func (s *ListHistoryDeployVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
