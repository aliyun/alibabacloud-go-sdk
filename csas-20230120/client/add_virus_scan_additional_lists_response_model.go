// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVirusScanAdditionalListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVirusScanAdditionalListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVirusScanAdditionalListsResponse
	GetStatusCode() *int32
	SetBody(v *AddVirusScanAdditionalListsResponseBody) *AddVirusScanAdditionalListsResponse
	GetBody() *AddVirusScanAdditionalListsResponseBody
}

type AddVirusScanAdditionalListsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVirusScanAdditionalListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVirusScanAdditionalListsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVirusScanAdditionalListsResponse) GoString() string {
	return s.String()
}

func (s *AddVirusScanAdditionalListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVirusScanAdditionalListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVirusScanAdditionalListsResponse) GetBody() *AddVirusScanAdditionalListsResponseBody {
	return s.Body
}

func (s *AddVirusScanAdditionalListsResponse) SetHeaders(v map[string]*string) *AddVirusScanAdditionalListsResponse {
	s.Headers = v
	return s
}

func (s *AddVirusScanAdditionalListsResponse) SetStatusCode(v int32) *AddVirusScanAdditionalListsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVirusScanAdditionalListsResponse) SetBody(v *AddVirusScanAdditionalListsResponseBody) *AddVirusScanAdditionalListsResponse {
	s.Body = v
	return s
}

func (s *AddVirusScanAdditionalListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
