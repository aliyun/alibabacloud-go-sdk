// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataServiceAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataServiceAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataServiceAppGroupResponseBody) *DeleteDataServiceAppGroupResponse
	GetBody() *DeleteDataServiceAppGroupResponseBody
}

type DeleteDataServiceAppGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataServiceAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataServiceAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceAppGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataServiceAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataServiceAppGroupResponse) GetBody() *DeleteDataServiceAppGroupResponseBody {
	return s.Body
}

func (s *DeleteDataServiceAppGroupResponse) SetHeaders(v map[string]*string) *DeleteDataServiceAppGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataServiceAppGroupResponse) SetStatusCode(v int32) *DeleteDataServiceAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataServiceAppGroupResponse) SetBody(v *DeleteDataServiceAppGroupResponseBody) *DeleteDataServiceAppGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDataServiceAppGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
