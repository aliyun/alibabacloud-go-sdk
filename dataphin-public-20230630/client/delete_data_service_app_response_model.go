// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataServiceAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataServiceAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataServiceAppResponseBody) *DeleteDataServiceAppResponse
	GetBody() *DeleteDataServiceAppResponseBody
}

type DeleteDataServiceAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataServiceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataServiceAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataServiceAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataServiceAppResponse) GetBody() *DeleteDataServiceAppResponseBody {
	return s.Body
}

func (s *DeleteDataServiceAppResponse) SetHeaders(v map[string]*string) *DeleteDataServiceAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataServiceAppResponse) SetStatusCode(v int32) *DeleteDataServiceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataServiceAppResponse) SetBody(v *DeleteDataServiceAppResponseBody) *DeleteDataServiceAppResponse {
	s.Body = v
	return s
}

func (s *DeleteDataServiceAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
