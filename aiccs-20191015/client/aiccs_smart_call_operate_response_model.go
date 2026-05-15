// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiccsSmartCallOperateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AiccsSmartCallOperateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AiccsSmartCallOperateResponse
	GetStatusCode() *int32
	SetBody(v *AiccsSmartCallOperateResponseBody) *AiccsSmartCallOperateResponse
	GetBody() *AiccsSmartCallOperateResponseBody
}

type AiccsSmartCallOperateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AiccsSmartCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AiccsSmartCallOperateResponse) String() string {
	return dara.Prettify(s)
}

func (s AiccsSmartCallOperateResponse) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallOperateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AiccsSmartCallOperateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AiccsSmartCallOperateResponse) GetBody() *AiccsSmartCallOperateResponseBody {
	return s.Body
}

func (s *AiccsSmartCallOperateResponse) SetHeaders(v map[string]*string) *AiccsSmartCallOperateResponse {
	s.Headers = v
	return s
}

func (s *AiccsSmartCallOperateResponse) SetStatusCode(v int32) *AiccsSmartCallOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *AiccsSmartCallOperateResponse) SetBody(v *AiccsSmartCallOperateResponseBody) *AiccsSmartCallOperateResponse {
	s.Body = v
	return s
}

func (s *AiccsSmartCallOperateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
