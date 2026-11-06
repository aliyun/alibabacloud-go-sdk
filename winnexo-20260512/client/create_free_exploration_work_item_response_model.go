// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeExplorationWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFreeExplorationWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFreeExplorationWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *CreateFreeExplorationWorkItemResponseBody) *CreateFreeExplorationWorkItemResponse
	GetBody() *CreateFreeExplorationWorkItemResponseBody
}

type CreateFreeExplorationWorkItemResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFreeExplorationWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFreeExplorationWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeExplorationWorkItemResponse) GoString() string {
	return s.String()
}

func (s *CreateFreeExplorationWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFreeExplorationWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFreeExplorationWorkItemResponse) GetBody() *CreateFreeExplorationWorkItemResponseBody {
	return s.Body
}

func (s *CreateFreeExplorationWorkItemResponse) SetHeaders(v map[string]*string) *CreateFreeExplorationWorkItemResponse {
	s.Headers = v
	return s
}

func (s *CreateFreeExplorationWorkItemResponse) SetStatusCode(v int32) *CreateFreeExplorationWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponse) SetBody(v *CreateFreeExplorationWorkItemResponseBody) *CreateFreeExplorationWorkItemResponse {
	s.Body = v
	return s
}

func (s *CreateFreeExplorationWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
