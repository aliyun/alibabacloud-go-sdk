// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManualDagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateManualDagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateManualDagResponse
	GetStatusCode() *int32
	SetBody(v *CreateManualDagResponseBody) *CreateManualDagResponse
	GetBody() *CreateManualDagResponseBody
}

type CreateManualDagResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateManualDagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateManualDagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateManualDagResponse) GoString() string {
	return s.String()
}

func (s *CreateManualDagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateManualDagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateManualDagResponse) GetBody() *CreateManualDagResponseBody {
	return s.Body
}

func (s *CreateManualDagResponse) SetHeaders(v map[string]*string) *CreateManualDagResponse {
	s.Headers = v
	return s
}

func (s *CreateManualDagResponse) SetStatusCode(v int32) *CreateManualDagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateManualDagResponse) SetBody(v *CreateManualDagResponseBody) *CreateManualDagResponse {
	s.Body = v
	return s
}

func (s *CreateManualDagResponse) Validate() error {
	return dara.Validate(s)
}
