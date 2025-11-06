// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocalEnsAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLocalEnsAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLocalEnsAssociationResponse
	GetStatusCode() *int32
	SetBody(v *QueryLocalEnsAssociationResponseBody) *QueryLocalEnsAssociationResponse
	GetBody() *QueryLocalEnsAssociationResponseBody
}

type QueryLocalEnsAssociationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLocalEnsAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLocalEnsAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalEnsAssociationResponse) GoString() string {
	return s.String()
}

func (s *QueryLocalEnsAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLocalEnsAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLocalEnsAssociationResponse) GetBody() *QueryLocalEnsAssociationResponseBody {
	return s.Body
}

func (s *QueryLocalEnsAssociationResponse) SetHeaders(v map[string]*string) *QueryLocalEnsAssociationResponse {
	s.Headers = v
	return s
}

func (s *QueryLocalEnsAssociationResponse) SetStatusCode(v int32) *QueryLocalEnsAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLocalEnsAssociationResponse) SetBody(v *QueryLocalEnsAssociationResponseBody) *QueryLocalEnsAssociationResponse {
	s.Body = v
	return s
}

func (s *QueryLocalEnsAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
