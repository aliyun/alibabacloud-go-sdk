// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTranslateByVPCResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTranslateByVPCResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTranslateByVPCResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTranslateByVPCResponseBody) *GetBatchTranslateByVPCResponse
	GetBody() *GetBatchTranslateByVPCResponseBody
}

type GetBatchTranslateByVPCResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTranslateByVPCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTranslateByVPCResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTranslateByVPCResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateByVPCResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTranslateByVPCResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTranslateByVPCResponse) GetBody() *GetBatchTranslateByVPCResponseBody {
	return s.Body
}

func (s *GetBatchTranslateByVPCResponse) SetHeaders(v map[string]*string) *GetBatchTranslateByVPCResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTranslateByVPCResponse) SetStatusCode(v int32) *GetBatchTranslateByVPCResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTranslateByVPCResponse) SetBody(v *GetBatchTranslateByVPCResponseBody) *GetBatchTranslateByVPCResponse {
	s.Body = v
	return s
}

func (s *GetBatchTranslateByVPCResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
