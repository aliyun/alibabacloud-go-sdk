// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackedAssetDealResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackedAssetDealResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackedAssetDealResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackedAssetDealResponseBody) *GetAttackedAssetDealResponse
	GetBody() *GetAttackedAssetDealResponseBody
}

type GetAttackedAssetDealResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackedAssetDealResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackedAssetDealResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackedAssetDealResponse) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackedAssetDealResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackedAssetDealResponse) GetBody() *GetAttackedAssetDealResponseBody {
	return s.Body
}

func (s *GetAttackedAssetDealResponse) SetHeaders(v map[string]*string) *GetAttackedAssetDealResponse {
	s.Headers = v
	return s
}

func (s *GetAttackedAssetDealResponse) SetStatusCode(v int32) *GetAttackedAssetDealResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackedAssetDealResponse) SetBody(v *GetAttackedAssetDealResponseBody) *GetAttackedAssetDealResponse {
	s.Body = v
	return s
}

func (s *GetAttackedAssetDealResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
