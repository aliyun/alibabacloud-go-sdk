// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDialogLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDialogLogResponse
	GetStatusCode() *int32
	SetBody(v *GetDialogLogResponseBody) *GetDialogLogResponse
	GetBody() *GetDialogLogResponseBody
}

type GetDialogLogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDialogLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDialogLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDialogLogResponse) GoString() string {
	return s.String()
}

func (s *GetDialogLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDialogLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDialogLogResponse) GetBody() *GetDialogLogResponseBody {
	return s.Body
}

func (s *GetDialogLogResponse) SetHeaders(v map[string]*string) *GetDialogLogResponse {
	s.Headers = v
	return s
}

func (s *GetDialogLogResponse) SetStatusCode(v int32) *GetDialogLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDialogLogResponse) SetBody(v *GetDialogLogResponseBody) *GetDialogLogResponse {
	s.Body = v
	return s
}

func (s *GetDialogLogResponse) Validate() error {
	return dara.Validate(s)
}
