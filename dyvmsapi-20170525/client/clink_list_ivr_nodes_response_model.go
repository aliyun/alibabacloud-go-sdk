// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListIvrNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListIvrNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListIvrNodesResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListIvrNodesResponseBody) *ClinkListIvrNodesResponse
	GetBody() *ClinkListIvrNodesResponseBody
}

type ClinkListIvrNodesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListIvrNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListIvrNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrNodesResponse) GoString() string {
	return s.String()
}

func (s *ClinkListIvrNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListIvrNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListIvrNodesResponse) GetBody() *ClinkListIvrNodesResponseBody {
	return s.Body
}

func (s *ClinkListIvrNodesResponse) SetHeaders(v map[string]*string) *ClinkListIvrNodesResponse {
	s.Headers = v
	return s
}

func (s *ClinkListIvrNodesResponse) SetStatusCode(v int32) *ClinkListIvrNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListIvrNodesResponse) SetBody(v *ClinkListIvrNodesResponseBody) *ClinkListIvrNodesResponse {
	s.Body = v
	return s
}

func (s *ClinkListIvrNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
