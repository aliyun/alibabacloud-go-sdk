// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSearchAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageSearchAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageSearchAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *OperationResult) *ManageSearchAccountInfoResponse
	GetBody() *OperationResult
}

type ManageSearchAccountInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageSearchAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageSearchAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *ManageSearchAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageSearchAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageSearchAccountInfoResponse) GetBody() *OperationResult {
	return s.Body
}

func (s *ManageSearchAccountInfoResponse) SetHeaders(v map[string]*string) *ManageSearchAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *ManageSearchAccountInfoResponse) SetStatusCode(v int32) *ManageSearchAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageSearchAccountInfoResponse) SetBody(v *OperationResult) *ManageSearchAccountInfoResponse {
	s.Body = v
	return s
}

func (s *ManageSearchAccountInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
