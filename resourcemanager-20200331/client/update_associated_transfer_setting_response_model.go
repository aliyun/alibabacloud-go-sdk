// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssociatedTransferSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAssociatedTransferSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAssociatedTransferSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAssociatedTransferSettingResponseBody) *UpdateAssociatedTransferSettingResponse
	GetBody() *UpdateAssociatedTransferSettingResponseBody
}

type UpdateAssociatedTransferSettingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAssociatedTransferSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssociatedTransferSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssociatedTransferSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAssociatedTransferSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAssociatedTransferSettingResponse) GetBody() *UpdateAssociatedTransferSettingResponseBody {
	return s.Body
}

func (s *UpdateAssociatedTransferSettingResponse) SetHeaders(v map[string]*string) *UpdateAssociatedTransferSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssociatedTransferSettingResponse) SetStatusCode(v int32) *UpdateAssociatedTransferSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssociatedTransferSettingResponse) SetBody(v *UpdateAssociatedTransferSettingResponseBody) *UpdateAssociatedTransferSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateAssociatedTransferSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
