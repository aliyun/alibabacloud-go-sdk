// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnDeliverTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCdnDeliverTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCdnDeliverTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCdnDeliverTaskResponseBody) *DeleteCdnDeliverTaskResponse
	GetBody() *DeleteCdnDeliverTaskResponseBody
}

type DeleteCdnDeliverTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCdnDeliverTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteCdnDeliverTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCdnDeliverTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCdnDeliverTaskResponse) GetBody() *DeleteCdnDeliverTaskResponseBody {
	return s.Body
}

func (s *DeleteCdnDeliverTaskResponse) SetHeaders(v map[string]*string) *DeleteCdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteCdnDeliverTaskResponse) SetStatusCode(v int32) *DeleteCdnDeliverTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCdnDeliverTaskResponse) SetBody(v *DeleteCdnDeliverTaskResponseBody) *DeleteCdnDeliverTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteCdnDeliverTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
