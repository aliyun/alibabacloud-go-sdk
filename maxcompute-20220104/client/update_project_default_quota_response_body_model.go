// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectDefaultQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateProjectDefaultQuotaResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateProjectDefaultQuotaResponseBody
	GetRequestId() *string
}

type UpdateProjectDefaultQuotaResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0a06dfe716674588654372173ec0da
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectDefaultQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectDefaultQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectDefaultQuotaResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateProjectDefaultQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectDefaultQuotaResponseBody) SetData(v string) *UpdateProjectDefaultQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectDefaultQuotaResponseBody) SetRequestId(v string) *UpdateProjectDefaultQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectDefaultQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
