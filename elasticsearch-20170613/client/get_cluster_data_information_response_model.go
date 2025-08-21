// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDataInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterDataInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterDataInformationResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterDataInformationResponseBody) *GetClusterDataInformationResponse
	GetBody() *GetClusterDataInformationResponseBody
}

type GetClusterDataInformationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterDataInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterDataInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDataInformationResponse) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterDataInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterDataInformationResponse) GetBody() *GetClusterDataInformationResponseBody {
	return s.Body
}

func (s *GetClusterDataInformationResponse) SetHeaders(v map[string]*string) *GetClusterDataInformationResponse {
	s.Headers = v
	return s
}

func (s *GetClusterDataInformationResponse) SetStatusCode(v int32) *GetClusterDataInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterDataInformationResponse) SetBody(v *GetClusterDataInformationResponseBody) *GetClusterDataInformationResponse {
	s.Body = v
	return s
}

func (s *GetClusterDataInformationResponse) Validate() error {
	return dara.Validate(s)
}
