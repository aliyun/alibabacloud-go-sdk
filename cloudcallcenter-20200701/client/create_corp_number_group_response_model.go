// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCorpNumberGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCorpNumberGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCorpNumberGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateCorpNumberGroupResponseBody) *CreateCorpNumberGroupResponse
	GetBody() *CreateCorpNumberGroupResponseBody
}

type CreateCorpNumberGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCorpNumberGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCorpNumberGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCorpNumberGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCorpNumberGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCorpNumberGroupResponse) GetBody() *CreateCorpNumberGroupResponseBody {
	return s.Body
}

func (s *CreateCorpNumberGroupResponse) SetHeaders(v map[string]*string) *CreateCorpNumberGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateCorpNumberGroupResponse) SetStatusCode(v int32) *CreateCorpNumberGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCorpNumberGroupResponse) SetBody(v *CreateCorpNumberGroupResponseBody) *CreateCorpNumberGroupResponse {
	s.Body = v
	return s
}

func (s *CreateCorpNumberGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
