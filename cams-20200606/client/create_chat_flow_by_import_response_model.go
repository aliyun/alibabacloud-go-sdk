// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowByImportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatFlowByImportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatFlowByImportResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatFlowByImportResponseBody) *CreateChatFlowByImportResponse
	GetBody() *CreateChatFlowByImportResponseBody
}

type CreateChatFlowByImportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatFlowByImportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatFlowByImportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowByImportResponse) GoString() string {
	return s.String()
}

func (s *CreateChatFlowByImportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatFlowByImportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatFlowByImportResponse) GetBody() *CreateChatFlowByImportResponseBody {
	return s.Body
}

func (s *CreateChatFlowByImportResponse) SetHeaders(v map[string]*string) *CreateChatFlowByImportResponse {
	s.Headers = v
	return s
}

func (s *CreateChatFlowByImportResponse) SetStatusCode(v int32) *CreateChatFlowByImportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatFlowByImportResponse) SetBody(v *CreateChatFlowByImportResponseBody) *CreateChatFlowByImportResponse {
	s.Body = v
	return s
}

func (s *CreateChatFlowByImportResponse) Validate() error {
	return dara.Validate(s)
}
