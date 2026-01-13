// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecallManagementTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecallManagementTableResponse
	GetStatusCode() *int32
	SetBody(v *GetRecallManagementTableResponseBody) *GetRecallManagementTableResponse
	GetBody() *GetRecallManagementTableResponseBody
}

type GetRecallManagementTableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecallManagementTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecallManagementTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementTableResponse) GoString() string {
	return s.String()
}

func (s *GetRecallManagementTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecallManagementTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecallManagementTableResponse) GetBody() *GetRecallManagementTableResponseBody {
	return s.Body
}

func (s *GetRecallManagementTableResponse) SetHeaders(v map[string]*string) *GetRecallManagementTableResponse {
	s.Headers = v
	return s
}

func (s *GetRecallManagementTableResponse) SetStatusCode(v int32) *GetRecallManagementTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecallManagementTableResponse) SetBody(v *GetRecallManagementTableResponseBody) *GetRecallManagementTableResponse {
	s.Body = v
	return s
}

func (s *GetRecallManagementTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
