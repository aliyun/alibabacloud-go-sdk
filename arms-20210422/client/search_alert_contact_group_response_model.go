// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAlertContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAlertContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *SearchAlertContactGroupResponseBody) *SearchAlertContactGroupResponse
	GetBody() *SearchAlertContactGroupResponseBody
}

type SearchAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAlertContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAlertContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAlertContactGroupResponse) GetBody() *SearchAlertContactGroupResponseBody {
	return s.Body
}

func (s *SearchAlertContactGroupResponse) SetHeaders(v map[string]*string) *SearchAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactGroupResponse) SetStatusCode(v int32) *SearchAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactGroupResponse) SetBody(v *SearchAlertContactGroupResponseBody) *SearchAlertContactGroupResponse {
	s.Body = v
	return s
}

func (s *SearchAlertContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
