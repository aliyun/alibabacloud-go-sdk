// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeRecallManagementServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeRecallManagementServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeRecallManagementServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *ChangeRecallManagementServiceVersionResponseBody) *ChangeRecallManagementServiceVersionResponse
	GetBody() *ChangeRecallManagementServiceVersionResponseBody
}

type ChangeRecallManagementServiceVersionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeRecallManagementServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeRecallManagementServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeRecallManagementServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *ChangeRecallManagementServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeRecallManagementServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeRecallManagementServiceVersionResponse) GetBody() *ChangeRecallManagementServiceVersionResponseBody {
	return s.Body
}

func (s *ChangeRecallManagementServiceVersionResponse) SetHeaders(v map[string]*string) *ChangeRecallManagementServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *ChangeRecallManagementServiceVersionResponse) SetStatusCode(v int32) *ChangeRecallManagementServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeRecallManagementServiceVersionResponse) SetBody(v *ChangeRecallManagementServiceVersionResponseBody) *ChangeRecallManagementServiceVersionResponse {
	s.Body = v
	return s
}

func (s *ChangeRecallManagementServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
