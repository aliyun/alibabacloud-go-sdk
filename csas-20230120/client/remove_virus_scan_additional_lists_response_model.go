// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVirusScanAdditionalListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveVirusScanAdditionalListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveVirusScanAdditionalListsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveVirusScanAdditionalListsResponseBody) *RemoveVirusScanAdditionalListsResponse
	GetBody() *RemoveVirusScanAdditionalListsResponseBody
}

type RemoveVirusScanAdditionalListsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveVirusScanAdditionalListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveVirusScanAdditionalListsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveVirusScanAdditionalListsResponse) GoString() string {
	return s.String()
}

func (s *RemoveVirusScanAdditionalListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveVirusScanAdditionalListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveVirusScanAdditionalListsResponse) GetBody() *RemoveVirusScanAdditionalListsResponseBody {
	return s.Body
}

func (s *RemoveVirusScanAdditionalListsResponse) SetHeaders(v map[string]*string) *RemoveVirusScanAdditionalListsResponse {
	s.Headers = v
	return s
}

func (s *RemoveVirusScanAdditionalListsResponse) SetStatusCode(v int32) *RemoveVirusScanAdditionalListsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVirusScanAdditionalListsResponse) SetBody(v *RemoveVirusScanAdditionalListsResponseBody) *RemoveVirusScanAdditionalListsResponse {
	s.Body = v
	return s
}

func (s *RemoveVirusScanAdditionalListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
