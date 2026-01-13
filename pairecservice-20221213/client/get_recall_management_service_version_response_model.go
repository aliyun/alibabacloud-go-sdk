// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecallManagementServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecallManagementServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetRecallManagementServiceVersionResponseBody) *GetRecallManagementServiceVersionResponse
	GetBody() *GetRecallManagementServiceVersionResponseBody
}

type GetRecallManagementServiceVersionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecallManagementServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecallManagementServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecallManagementServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecallManagementServiceVersionResponse) GetBody() *GetRecallManagementServiceVersionResponseBody {
	return s.Body
}

func (s *GetRecallManagementServiceVersionResponse) SetHeaders(v map[string]*string) *GetRecallManagementServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *GetRecallManagementServiceVersionResponse) SetStatusCode(v int32) *GetRecallManagementServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponse) SetBody(v *GetRecallManagementServiceVersionResponseBody) *GetRecallManagementServiceVersionResponse {
	s.Body = v
	return s
}

func (s *GetRecallManagementServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
