// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataServiceAppMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDataServiceAppMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDataServiceAppMemberResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDataServiceAppMemberResponseBody) *RemoveDataServiceAppMemberResponse
	GetBody() *RemoveDataServiceAppMemberResponseBody
}

type RemoveDataServiceAppMemberResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDataServiceAppMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDataServiceAppMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataServiceAppMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveDataServiceAppMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDataServiceAppMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDataServiceAppMemberResponse) GetBody() *RemoveDataServiceAppMemberResponseBody {
	return s.Body
}

func (s *RemoveDataServiceAppMemberResponse) SetHeaders(v map[string]*string) *RemoveDataServiceAppMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveDataServiceAppMemberResponse) SetStatusCode(v int32) *RemoveDataServiceAppMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDataServiceAppMemberResponse) SetBody(v *RemoveDataServiceAppMemberResponseBody) *RemoveDataServiceAppMemberResponse {
	s.Body = v
	return s
}

func (s *RemoveDataServiceAppMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
