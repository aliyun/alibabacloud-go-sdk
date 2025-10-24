// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsPartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsPartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsPartitionResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsPartitionResponseBody) *GetMmsPartitionResponse
	GetBody() *GetMmsPartitionResponseBody
}

type GetMmsPartitionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsPartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsPartitionResponse) GoString() string {
	return s.String()
}

func (s *GetMmsPartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsPartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsPartitionResponse) GetBody() *GetMmsPartitionResponseBody {
	return s.Body
}

func (s *GetMmsPartitionResponse) SetHeaders(v map[string]*string) *GetMmsPartitionResponse {
	s.Headers = v
	return s
}

func (s *GetMmsPartitionResponse) SetStatusCode(v int32) *GetMmsPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsPartitionResponse) SetBody(v *GetMmsPartitionResponseBody) *GetMmsPartitionResponse {
	s.Body = v
	return s
}

func (s *GetMmsPartitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
