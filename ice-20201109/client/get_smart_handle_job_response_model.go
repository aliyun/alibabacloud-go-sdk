// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartHandleJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmartHandleJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmartHandleJobResponse
	GetStatusCode() *int32
	SetBody(v *GetSmartHandleJobResponseBody) *GetSmartHandleJobResponse
	GetBody() *GetSmartHandleJobResponseBody
}

type GetSmartHandleJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmartHandleJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmartHandleJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHandleJobResponse) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmartHandleJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmartHandleJobResponse) GetBody() *GetSmartHandleJobResponseBody {
	return s.Body
}

func (s *GetSmartHandleJobResponse) SetHeaders(v map[string]*string) *GetSmartHandleJobResponse {
	s.Headers = v
	return s
}

func (s *GetSmartHandleJobResponse) SetStatusCode(v int32) *GetSmartHandleJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmartHandleJobResponse) SetBody(v *GetSmartHandleJobResponseBody) *GetSmartHandleJobResponse {
	s.Body = v
	return s
}

func (s *GetSmartHandleJobResponse) Validate() error {
	return dara.Validate(s)
}
