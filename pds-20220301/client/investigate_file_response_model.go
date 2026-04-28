// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvestigateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvestigateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvestigateFileResponse
	GetStatusCode() *int32
}

type InvestigateFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s InvestigateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s InvestigateFileResponse) GoString() string {
	return s.String()
}

func (s *InvestigateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvestigateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvestigateFileResponse) SetHeaders(v map[string]*string) *InvestigateFileResponse {
	s.Headers = v
	return s
}

func (s *InvestigateFileResponse) SetStatusCode(v int32) *InvestigateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *InvestigateFileResponse) Validate() error {
	return dara.Validate(s)
}
