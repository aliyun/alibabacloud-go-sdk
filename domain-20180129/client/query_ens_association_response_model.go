// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnsAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEnsAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEnsAssociationResponse
	GetStatusCode() *int32
	SetBody(v *QueryEnsAssociationResponseBody) *QueryEnsAssociationResponse
	GetBody() *QueryEnsAssociationResponseBody
}

type QueryEnsAssociationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEnsAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEnsAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEnsAssociationResponse) GoString() string {
	return s.String()
}

func (s *QueryEnsAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEnsAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEnsAssociationResponse) GetBody() *QueryEnsAssociationResponseBody {
	return s.Body
}

func (s *QueryEnsAssociationResponse) SetHeaders(v map[string]*string) *QueryEnsAssociationResponse {
	s.Headers = v
	return s
}

func (s *QueryEnsAssociationResponse) SetStatusCode(v int32) *QueryEnsAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnsAssociationResponse) SetBody(v *QueryEnsAssociationResponseBody) *QueryEnsAssociationResponse {
	s.Body = v
	return s
}

func (s *QueryEnsAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
