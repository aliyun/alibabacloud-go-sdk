// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenApiListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOpenApiListRequest
	GetRequestId() *string
}

type GetOpenApiListRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpenApiListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpenApiListRequest) GoString() string {
	return s.String()
}

func (s *GetOpenApiListRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpenApiListRequest) SetRequestId(v string) *GetOpenApiListRequest {
	s.RequestId = &v
	return s
}

func (s *GetOpenApiListRequest) Validate() error {
	return dara.Validate(s)
}
