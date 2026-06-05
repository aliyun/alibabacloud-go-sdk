// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAppRequirementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveAppRequirementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveAppRequirementResponse
	GetStatusCode() *int32
	SetBody(v *SaveAppRequirementResponseBody) *SaveAppRequirementResponse
	GetBody() *SaveAppRequirementResponseBody
}

type SaveAppRequirementResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAppRequirementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAppRequirementResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveAppRequirementResponse) GoString() string {
	return s.String()
}

func (s *SaveAppRequirementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveAppRequirementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveAppRequirementResponse) GetBody() *SaveAppRequirementResponseBody {
	return s.Body
}

func (s *SaveAppRequirementResponse) SetHeaders(v map[string]*string) *SaveAppRequirementResponse {
	s.Headers = v
	return s
}

func (s *SaveAppRequirementResponse) SetStatusCode(v int32) *SaveAppRequirementResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAppRequirementResponse) SetBody(v *SaveAppRequirementResponseBody) *SaveAppRequirementResponse {
	s.Body = v
	return s
}

func (s *SaveAppRequirementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
