// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanAdditionalListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanAdditionalListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanAdditionalListsResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanAdditionalListsResponseBody) *ListVirusScanAdditionalListsResponse
	GetBody() *ListVirusScanAdditionalListsResponseBody
}

type ListVirusScanAdditionalListsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanAdditionalListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanAdditionalListsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanAdditionalListsResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanAdditionalListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanAdditionalListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanAdditionalListsResponse) GetBody() *ListVirusScanAdditionalListsResponseBody {
	return s.Body
}

func (s *ListVirusScanAdditionalListsResponse) SetHeaders(v map[string]*string) *ListVirusScanAdditionalListsResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanAdditionalListsResponse) SetStatusCode(v int32) *ListVirusScanAdditionalListsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanAdditionalListsResponse) SetBody(v *ListVirusScanAdditionalListsResponseBody) *ListVirusScanAdditionalListsResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanAdditionalListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
