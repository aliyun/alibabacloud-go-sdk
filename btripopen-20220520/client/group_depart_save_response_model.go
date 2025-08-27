// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupDepartSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GroupDepartSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GroupDepartSaveResponse
	GetStatusCode() *int32
	SetBody(v *GroupDepartSaveResponseBody) *GroupDepartSaveResponse
	GetBody() *GroupDepartSaveResponseBody
}

type GroupDepartSaveResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupDepartSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupDepartSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s GroupDepartSaveResponse) GoString() string {
	return s.String()
}

func (s *GroupDepartSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GroupDepartSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GroupDepartSaveResponse) GetBody() *GroupDepartSaveResponseBody {
	return s.Body
}

func (s *GroupDepartSaveResponse) SetHeaders(v map[string]*string) *GroupDepartSaveResponse {
	s.Headers = v
	return s
}

func (s *GroupDepartSaveResponse) SetStatusCode(v int32) *GroupDepartSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupDepartSaveResponse) SetBody(v *GroupDepartSaveResponseBody) *GroupDepartSaveResponse {
	s.Body = v
	return s
}

func (s *GroupDepartSaveResponse) Validate() error {
	return dara.Validate(s)
}
