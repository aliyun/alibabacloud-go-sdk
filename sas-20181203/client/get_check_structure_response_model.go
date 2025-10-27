// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckStructureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckStructureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckStructureResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckStructureResponseBody) *GetCheckStructureResponse
	GetBody() *GetCheckStructureResponseBody
}

type GetCheckStructureResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckStructureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckStructureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckStructureResponse) GoString() string {
	return s.String()
}

func (s *GetCheckStructureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckStructureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckStructureResponse) GetBody() *GetCheckStructureResponseBody {
	return s.Body
}

func (s *GetCheckStructureResponse) SetHeaders(v map[string]*string) *GetCheckStructureResponse {
	s.Headers = v
	return s
}

func (s *GetCheckStructureResponse) SetStatusCode(v int32) *GetCheckStructureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckStructureResponse) SetBody(v *GetCheckStructureResponseBody) *GetCheckStructureResponse {
	s.Body = v
	return s
}

func (s *GetCheckStructureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
