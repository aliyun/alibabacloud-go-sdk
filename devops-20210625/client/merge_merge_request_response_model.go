// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeMergeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MergeMergeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MergeMergeRequestResponse
	GetStatusCode() *int32
	SetBody(v *MergeMergeRequestResponseBody) *MergeMergeRequestResponse
	GetBody() *MergeMergeRequestResponseBody
}

type MergeMergeRequestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MergeMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MergeMergeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s MergeMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MergeMergeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MergeMergeRequestResponse) GetBody() *MergeMergeRequestResponseBody {
	return s.Body
}

func (s *MergeMergeRequestResponse) SetHeaders(v map[string]*string) *MergeMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *MergeMergeRequestResponse) SetStatusCode(v int32) *MergeMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeMergeRequestResponse) SetBody(v *MergeMergeRequestResponseBody) *MergeMergeRequestResponse {
	s.Body = v
	return s
}

func (s *MergeMergeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
