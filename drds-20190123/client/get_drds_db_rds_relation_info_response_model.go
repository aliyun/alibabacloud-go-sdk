// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDrdsDbRdsRelationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDrdsDbRdsRelationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDrdsDbRdsRelationInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDrdsDbRdsRelationInfoResponseBody) *GetDrdsDbRdsRelationInfoResponse
	GetBody() *GetDrdsDbRdsRelationInfoResponseBody
}

type GetDrdsDbRdsRelationInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDrdsDbRdsRelationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDrdsDbRdsRelationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDrdsDbRdsRelationInfoResponse) GetBody() *GetDrdsDbRdsRelationInfoResponseBody {
	return s.Body
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetHeaders(v map[string]*string) *GetDrdsDbRdsRelationInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetStatusCode(v int32) *GetDrdsDbRdsRelationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponse) SetBody(v *GetDrdsDbRdsRelationInfoResponseBody) *GetDrdsDbRdsRelationInfoResponse {
	s.Body = v
	return s
}

func (s *GetDrdsDbRdsRelationInfoResponse) Validate() error {
	return dara.Validate(s)
}
