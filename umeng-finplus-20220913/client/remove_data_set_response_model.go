// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDataSetResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDataSetResponseBody) *RemoveDataSetResponse
	GetBody() *RemoveDataSetResponseBody
}

type RemoveDataSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSetResponse) GoString() string {
	return s.String()
}

func (s *RemoveDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDataSetResponse) GetBody() *RemoveDataSetResponseBody {
	return s.Body
}

func (s *RemoveDataSetResponse) SetHeaders(v map[string]*string) *RemoveDataSetResponse {
	s.Headers = v
	return s
}

func (s *RemoveDataSetResponse) SetStatusCode(v int32) *RemoveDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDataSetResponse) SetBody(v *RemoveDataSetResponseBody) *RemoveDataSetResponse {
	s.Body = v
	return s
}

func (s *RemoveDataSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
