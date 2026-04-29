// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecallManagementTableRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecallManagementTableRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecallManagementTableRecordsResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecallManagementTableRecordsResponseBody) *QueryRecallManagementTableRecordsResponse
	GetBody() *QueryRecallManagementTableRecordsResponseBody
}

type QueryRecallManagementTableRecordsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecallManagementTableRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecallManagementTableRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecallManagementTableRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryRecallManagementTableRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecallManagementTableRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecallManagementTableRecordsResponse) GetBody() *QueryRecallManagementTableRecordsResponseBody {
	return s.Body
}

func (s *QueryRecallManagementTableRecordsResponse) SetHeaders(v map[string]*string) *QueryRecallManagementTableRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryRecallManagementTableRecordsResponse) SetStatusCode(v int32) *QueryRecallManagementTableRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecallManagementTableRecordsResponse) SetBody(v *QueryRecallManagementTableRecordsResponseBody) *QueryRecallManagementTableRecordsResponse {
	s.Body = v
	return s
}

func (s *QueryRecallManagementTableRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
