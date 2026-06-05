// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRequirementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppRequirementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppRequirementResponse
	GetStatusCode() *int32
	SetBody(v *GetAppRequirementResponseBody) *GetAppRequirementResponse
	GetBody() *GetAppRequirementResponseBody
}

type GetAppRequirementResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppRequirementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppRequirementResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppRequirementResponse) GoString() string {
	return s.String()
}

func (s *GetAppRequirementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppRequirementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppRequirementResponse) GetBody() *GetAppRequirementResponseBody {
	return s.Body
}

func (s *GetAppRequirementResponse) SetHeaders(v map[string]*string) *GetAppRequirementResponse {
	s.Headers = v
	return s
}

func (s *GetAppRequirementResponse) SetStatusCode(v int32) *GetAppRequirementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppRequirementResponse) SetBody(v *GetAppRequirementResponseBody) *GetAppRequirementResponse {
	s.Body = v
	return s
}

func (s *GetAppRequirementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
