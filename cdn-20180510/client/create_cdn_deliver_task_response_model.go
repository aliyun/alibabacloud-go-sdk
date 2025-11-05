// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnDeliverTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCdnDeliverTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCdnDeliverTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCdnDeliverTaskResponseBody) *CreateCdnDeliverTaskResponse
	GetBody() *CreateCdnDeliverTaskResponseBody
}

type CreateCdnDeliverTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCdnDeliverTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCdnDeliverTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCdnDeliverTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCdnDeliverTaskResponse) GetBody() *CreateCdnDeliverTaskResponseBody {
	return s.Body
}

func (s *CreateCdnDeliverTaskResponse) SetHeaders(v map[string]*string) *CreateCdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCdnDeliverTaskResponse) SetStatusCode(v int32) *CreateCdnDeliverTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCdnDeliverTaskResponse) SetBody(v *CreateCdnDeliverTaskResponseBody) *CreateCdnDeliverTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCdnDeliverTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
