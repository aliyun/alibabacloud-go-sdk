// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPartitionResponse
	GetStatusCode() *int32
	SetBody(v *GetPartitionResponseBody) *GetPartitionResponse
	GetBody() *GetPartitionResponseBody
}

type GetPartitionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPartitionResponse) GoString() string {
	return s.String()
}

func (s *GetPartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPartitionResponse) GetBody() *GetPartitionResponseBody {
	return s.Body
}

func (s *GetPartitionResponse) SetHeaders(v map[string]*string) *GetPartitionResponse {
	s.Headers = v
	return s
}

func (s *GetPartitionResponse) SetStatusCode(v int32) *GetPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPartitionResponse) SetBody(v *GetPartitionResponseBody) *GetPartitionResponse {
	s.Body = v
	return s
}

func (s *GetPartitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
