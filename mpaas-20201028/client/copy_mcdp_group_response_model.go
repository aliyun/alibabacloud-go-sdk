// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyMcdpGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyMcdpGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyMcdpGroupResponse
	GetStatusCode() *int32
	SetBody(v *CopyMcdpGroupResponseBody) *CopyMcdpGroupResponse
	GetBody() *CopyMcdpGroupResponseBody
}

type CopyMcdpGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyMcdpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyMcdpGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyMcdpGroupResponse) GoString() string {
	return s.String()
}

func (s *CopyMcdpGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyMcdpGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyMcdpGroupResponse) GetBody() *CopyMcdpGroupResponseBody {
	return s.Body
}

func (s *CopyMcdpGroupResponse) SetHeaders(v map[string]*string) *CopyMcdpGroupResponse {
	s.Headers = v
	return s
}

func (s *CopyMcdpGroupResponse) SetStatusCode(v int32) *CopyMcdpGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyMcdpGroupResponse) SetBody(v *CopyMcdpGroupResponseBody) *CopyMcdpGroupResponse {
	s.Body = v
	return s
}

func (s *CopyMcdpGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
