// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssociatedTransferSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssociatedTransferSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssociatedTransferSettingResponse
	GetStatusCode() *int32
	SetBody(v *ListAssociatedTransferSettingResponseBody) *ListAssociatedTransferSettingResponse
	GetBody() *ListAssociatedTransferSettingResponseBody
}

type ListAssociatedTransferSettingResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssociatedTransferSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssociatedTransferSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedTransferSettingResponse) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssociatedTransferSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssociatedTransferSettingResponse) GetBody() *ListAssociatedTransferSettingResponseBody {
	return s.Body
}

func (s *ListAssociatedTransferSettingResponse) SetHeaders(v map[string]*string) *ListAssociatedTransferSettingResponse {
	s.Headers = v
	return s
}

func (s *ListAssociatedTransferSettingResponse) SetStatusCode(v int32) *ListAssociatedTransferSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssociatedTransferSettingResponse) SetBody(v *ListAssociatedTransferSettingResponseBody) *ListAssociatedTransferSettingResponse {
	s.Body = v
	return s
}

func (s *ListAssociatedTransferSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
