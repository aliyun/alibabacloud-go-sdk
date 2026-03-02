// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLanesForServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpLanesForServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpLanesForServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *PdpLanesPageResult) *ListPdpLanesForServiceGroupResponse
	GetBody() *PdpLanesPageResult
}

type ListPdpLanesForServiceGroupResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpLanesPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpLanesForServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLanesForServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPdpLanesForServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpLanesForServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpLanesForServiceGroupResponse) GetBody() *PdpLanesPageResult {
	return s.Body
}

func (s *ListPdpLanesForServiceGroupResponse) SetHeaders(v map[string]*string) *ListPdpLanesForServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPdpLanesForServiceGroupResponse) SetStatusCode(v int32) *ListPdpLanesForServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpLanesForServiceGroupResponse) SetBody(v *PdpLanesPageResult) *ListPdpLanesForServiceGroupResponse {
	s.Body = v
	return s
}

func (s *ListPdpLanesForServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
